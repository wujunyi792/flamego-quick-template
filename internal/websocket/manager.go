package websocket

import (
	"github.com/wujunyi792/flamego-quick-template/internal/core/logx"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var managers = make(map[string]*SocketManager)

const (
	writeWait = 10 * time.Second
)

var upGrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	// 跨域
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type SocketManager struct {
	clients    map[*socketClient]bool
	register   chan *socketClient
	unregister chan *socketClient
	receive    chan map[*socketClient][]byte
	broadcast  chan []byte
}

type socketClient struct {
	name    string
	manager *SocketManager
	conn    *websocket.Conn
	send    chan []byte
}

/**
 * @description: 创建Socket管理器
 */
func newManager() *SocketManager {
	return &SocketManager{
		clients:    make(map[*socketClient]bool),
		register:   make(chan *socketClient),
		unregister: make(chan *socketClient),
		receive:    make(chan map[*socketClient][]byte),
		broadcast:  make(chan []byte), // 广播
	}
}

/**
 * @description: 接收Socket连接
 */
func (m *SocketManager) run() {
	for {
		select {
		// 连接加入
		case client := <-m.register:
			// 设置clinet为true
			m.clients[client] = true
		// 连接关闭
		case client := <-m.unregister:
			if _, ok := m.clients[client]; ok {
				delete(m.clients, client)
				close(client.send)
			}
		// 发送message
		case message := <-m.broadcast:
			for client := range m.clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(m.clients, client)
				}
			}
		}
	}
}

/**
 * @description: 接收Socket信息
 */
func (c *socketClient) readPump() {
	defer func() {
		c.manager.unregister <- c
		err := c.conn.Close()
		if err != nil {
			logx.NameSpace("ws").Errorln(err)
		}
	}()
	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			logx.NameSpace("ws").Errorln("error: %v", err)
			break
		}
		c.manager.receive <- map[*socketClient][]byte{
			c: message,
		}
	}
}

/**
 * @description: 发送Socket信息
 */
func (c *socketClient) writePump() {
	defer func() {
		err := c.conn.Close()
		if err != nil {
			logx.NameSpace("ws").Errorln(err)
		}
	}()
	for {
		message, ok := <-c.send
		c.conn.SetWriteDeadline(time.Now().Add(writeWait))
		if !ok {
			c.conn.WriteMessage(websocket.CloseMessage, []byte{})
			return
		}

		w, err := c.conn.NextWriter(websocket.TextMessage)
		if err != nil {
			return
		}
		w.Write(message)

		// Add queued chat messages to the current websocket message.
		n := len(c.send)
		for i := 0; i < n; i++ {
			w.Write(<-c.send)
		}

		if err := w.Close(); err != nil {
			return
		}
	}
}

func (m *SocketManager) SendClientSocket(name string, message string) {
	for k := range m.clients {
		if k.name == name {
			k.send <- []byte(message)
		}
	}
}

func (m *SocketManager) SendAllSocket(message string) {
	m.broadcast <- []byte(message)
}

func (m *SocketManager) ServeSocket(w http.ResponseWriter, r *http.Request, n string) {
	conn, err := upGrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	client := &socketClient{name: n, manager: m, conn: conn, send: make(chan []byte, 256)}
	client.manager.register <- client
	go client.writePump()
	go client.readPump()
}

func InitSocketManager(key string) {
	if key == "" {
		key = "*"
	}
	if _, ok := managers[key]; ok {
		logx.NameSpace("ws").Fatalln("socket manager key duplication")
	}
	m := newManager()
	go m.run()
	managers[key] = m
}

func GetSocketManager(key string) *SocketManager {
	if m, ok := managers[key]; ok {
		return m
	}
	logx.NameSpace("ws").Errorln("socket client ", key, " not found")
	return nil
}

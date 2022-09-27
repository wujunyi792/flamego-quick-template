package handle

import (
	"github.com/wujunyi792/flamego-quick-template/internal/websocket"
	"github.com/wujunyi792/flamego-quick-template/pkg/utils/gen/xrandom"
	"net/http"
)

func HandleWebsocketEcho(w http.ResponseWriter, r *http.Request, socket *websocket.SocketManager) {
	socket.ServeSocket(w, r, xrandom.GetRandom(7, xrandom.RAND_ALL))
}

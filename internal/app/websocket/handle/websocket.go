package handle

import (
	"github.com/wujunyi792/gin-template-new/internal/app/websocket/service"
	"github.com/wujunyi792/gin-template-new/pkg/utils/gen/xrandom"
	"net/http"
)

func HandelWebsocketEcho(w http.ResponseWriter, r *http.Request) {
	service.SocketServer(w, r, xrandom.GetRandom(7, xrandom.RAND_ALL))
}

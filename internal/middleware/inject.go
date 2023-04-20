package middleware

import (
	"github.com/flamego/flamego"
	"github.com/wujunyi792/flamego-quick-template/internal/core/cache"
	"github.com/wujunyi792/flamego-quick-template/internal/core/database"
	"github.com/wujunyi792/flamego-quick-template/internal/models/jwtModel"
	"github.com/wujunyi792/flamego-quick-template/internal/websocket"
)

func InjectDB(key string) flamego.Handler {
	return func(c flamego.Context) {
		c.Map(database.GetDb(key))
	}
}

func InjectWebsocket(key string) flamego.Handler {
	return func(c flamego.Context) {
		c.Map(websocket.GetSocketManager(key))
	}
}

func InjectCache(key string) flamego.Handler {
	return func(c flamego.Context) {
		c.Map(cache.GetCache(key))
	}
}

func InjectUserInfo(info jwtModel.UserInfo) flamego.Handler {
	return func(c flamego.Context) {
		c.Map(info)
	}
}

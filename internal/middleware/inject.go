package middleware

import (
	"github.com/flamego/flamego"
	"github.com/wujunyi792/gin-template-new/internal/cache"
	"github.com/wujunyi792/gin-template-new/internal/database"
	"github.com/wujunyi792/gin-template-new/internal/models/jwtModel"
	"github.com/wujunyi792/gin-template-new/internal/websocket"
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

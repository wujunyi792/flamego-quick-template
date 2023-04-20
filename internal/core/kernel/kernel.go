package kernel

import (
	"context"
	"github.com/flamego/flamego"
)

type Engine struct {
	// 不由 Engine 统一管理
	//Mysql      *gorm.DB
	//Cache      *redis.Client
	Fg *flamego.Flame

	Ctx    context.Context
	Cancel context.CancelFunc
}

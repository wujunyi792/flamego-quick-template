package ping

import (
	"context"
	"github.com/wujunyi792/flamego-quick-template/internal/app"
	"github.com/wujunyi792/flamego-quick-template/internal/app/ping/dao"
	"github.com/wujunyi792/flamego-quick-template/internal/app/ping/router"
	"github.com/wujunyi792/flamego-quick-template/internal/core/database"
	"github.com/wujunyi792/flamego-quick-template/internal/core/kernel"

	"sync"
)

type Ping struct {
	Name string

	app.UnimplementedModule
}

func (p *Ping) Info() string {
	return p.Name
}

func (p *Ping) PreInit(engine *kernel.Engine) error {
	db := database.GetDb("*")
	if db == nil {
		return nil
	}
	err := dao.PingDto.Init(db)
	if err != nil {
		return err
	}
	dao.Ping = db
	return nil
}

func (p *Ping) Init(*kernel.Engine) error {
	return nil
}

func (p *Ping) Load(engine *kernel.Engine) error {
	// 加载flamego api
	router.AppPingInit(engine.Fg)
	return nil
}

func (p *Ping) Stop(wg *sync.WaitGroup, ctx context.Context) error {
	defer wg.Done()
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		return nil
	}
}

func (p *Ping) OnConfigChange() func(*kernel.Engine) error {
	return func(engine *kernel.Engine) error {
		return nil
	}
}

package app

import (
	"github.com/wujunyi792/flamego-quick-template/internal/core/kernel"
)

type Module interface {
	Info() string
	PreInit(*kernel.Engine) error
	Init(*kernel.Engine) error
	PostInit(*kernel.Engine) error
	Load(*kernel.Engine) error
	Start(*kernel.Engine) error

	OnConfigChange() func(*kernel.Engine) error

	mustEmbedUnimplementedModule()
}

type UnimplementedModule struct{}

func (*UnimplementedModule) Info() string {
	return "unimplementedModule"
}

func (*UnimplementedModule) PreInit(*kernel.Engine) error {
	return nil
}

func (*UnimplementedModule) Init(*kernel.Engine) error {
	return nil
}

func (*UnimplementedModule) PostInit(*kernel.Engine) error {
	return nil
}

func (*UnimplementedModule) Load(*kernel.Engine) error {
	return nil
}

func (*UnimplementedModule) Start(*kernel.Engine) error {
	return nil
}

func (*UnimplementedModule) OnConfigChange() func(*kernel.Engine) error {
	return func(engine *kernel.Engine) error {
		return nil
	}
}

func (*UnimplementedModule) mustEmbedUnimplementedModule() {}

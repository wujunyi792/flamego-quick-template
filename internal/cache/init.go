package cache

import (
	"github.com/wujunyi792/gin-template-new/config"
	"github.com/wujunyi792/gin-template-new/internal/cache/types"
	"github.com/wujunyi792/gin-template-new/pkg/logx"
	"sync"
)

var (
	dbs = make(map[string]types.Cache)
	mux sync.RWMutex
)

func InitCache() {
	sources := config.GetConfig().Caches
	for _, source := range sources {
		setCacheByKey(source.Key, mustCreateCache(source))
		if source.Key == "" {
			source.Key = "*"
		}
		logx.Info.Printf("create cache %s => %s:%s", source.Key, source.IP, source.PORT)
	}
}

func GetCache(key string) types.Cache {
	mux.Lock()
	defer mux.Unlock()
	return dbs[key]
}

func setCacheByKey(key string, cache types.Cache) {
	if key == "" {
		key = "*"
	}
	if GetCache(key) != nil {
		logx.Error.Fatalln("duplicate db key: " + key)
	}
	mux.Lock()
	defer mux.Unlock()
	dbs[key] = cache
}

func mustCreateCache(conf config.Cache) types.Cache {
	var creator = getCreatorByType(conf.Type)
	if creator == nil {
		logx.Error.Fatalf("fail to find creator for cache types:%s", conf.Type)
		return nil
	}
	cache, err := creator.Create(conf)
	if err != nil {
		logx.Error.Fatalln(err)
		return nil
	}
	return cache
}

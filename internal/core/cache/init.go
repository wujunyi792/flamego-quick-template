package cache

import (
	"github.com/wujunyi792/flamego-quick-template/config"
	"github.com/wujunyi792/flamego-quick-template/internal/core/cache/types"
	"github.com/wujunyi792/flamego-quick-template/internal/core/logx"
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
		logx.NameSpace("cache").Infoln("create cache %s => %s:%s", source.Key, source.IP, source.PORT)
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
		logx.NameSpace("cache").Fatalln("duplicate db key: " + key)
	}
	mux.Lock()
	defer mux.Unlock()
	dbs[key] = cache
}

func mustCreateCache(conf config.Cache) types.Cache {
	var creator = getCreatorByType(conf.Type)
	if creator == nil {
		logx.NameSpace("cache").Fatalln("fail to find creator for cache types:%s", conf.Type)
		return nil
	}
	cache, err := creator.Create(conf)
	if err != nil {
		logx.NameSpace("cache").Fatalf("create err: %s", err)
		return nil
	}
	return cache
}

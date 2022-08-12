package cache

import (
	"github.com/wujunyi792/gin-template-new/config"
	"github.com/wujunyi792/gin-template-new/internal/cache/typeCache"
	"github.com/wujunyi792/gin-template-new/internal/loging"
	"sync"
)

var (
	dbs = make(map[string]typeCache.Cache)
	mux sync.RWMutex
)

func InitCache() {
	sources := config.GetConfig().Caches
	for _, source := range sources {
		setCacheByKey(source.Key, mustCreateCache(source))
		if source.Key == "" {
			source.Key = "*"
		}
		loging.Info.Printf("create cache %s => %s:%s", source.Key, source.IP, source.PORT)
	}
}

func GetCache(key string) typeCache.Cache {
	mux.Lock()
	defer mux.Unlock()
	return dbs[key]
}

func setCacheByKey(key string, cache typeCache.Cache) {
	if key == "" {
		key = "*"
	}
	if GetCache(key) != nil {
		loging.Error.Fatalln("duplicate db key: " + key)
	}
	mux.Lock()
	defer mux.Unlock()
	dbs[key] = cache
}

func mustCreateCache(conf config.Cache) typeCache.Cache {
	var creator = getCreatorByType(conf.Type)
	if creator == nil {
		loging.Error.Fatalf("fail to find creator for cache typeCache:%s", conf.Type)
		return nil
	}
	cache, err := creator.Create(conf)
	if err != nil {
		loging.Error.Fatalln(err)
		return nil
	}
	return cache
}

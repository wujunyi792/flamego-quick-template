package cache

import (
	"github.com/wujunyi792/gin-template-new/config"
	"github.com/wujunyi792/gin-template-new/internal/cache/driver"
	"github.com/wujunyi792/gin-template-new/internal/cache/typeCache"
)

type Creator interface {
	Create(conf config.Cache) (typeCache.Cache, error)
}

func init() {
	typeMap["redis"] = driver.RedisCreator{}
}

var typeMap = make(map[string]Creator)

func getCreatorByType(cacheType string) Creator {
	return typeMap[cacheType]
}

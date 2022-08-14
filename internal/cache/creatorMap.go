package cache

import (
	"github.com/wujunyi792/flamego-quick-template/config"
	"github.com/wujunyi792/flamego-quick-template/internal/cache/driver"
	"github.com/wujunyi792/flamego-quick-template/internal/cache/types"
)

type Creator interface {
	Create(conf config.Cache) (types.Cache, error)
}

func init() {
	typeMap["redis"] = driver.RedisCreator{}
}

var typeMap = make(map[string]Creator)

func getCreatorByType(cacheType string) Creator {
	return typeMap[cacheType]
}

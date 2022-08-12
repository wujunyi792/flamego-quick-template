package database

import (
	"github.com/wujunyi792/gin-template-new/config"
	"github.com/wujunyi792/gin-template-new/internal/logx"
	"gorm.io/gorm"
	"sync"
)

var (
	dbs = make(map[string]*gorm.DB)
	mux sync.RWMutex
)

func InitDB() {
	sources := config.GetConfig().Databases
	for _, source := range sources {
		setDbByKey(source.Key, mustCreateGorm(source))
		if source.Key == "" {
			source.Key = "*"
		}
		logx.Info.Printf("create datasource %s => %s:%s", source.Key, source.IP, source.PORT)
	}
}

func GetDb(key string) *gorm.DB {
	mux.Lock()
	defer mux.Unlock()
	return dbs[key]
}

func setDbByKey(key string, db *gorm.DB) {
	if key == "" {
		key = "*"
	}
	if GetDb(key) != nil {
		logx.Error.Fatalln("duplicate db key: " + key)
	}
	mux.Lock()
	defer mux.Unlock()
	dbs[key] = db
}

func mustCreateGorm(database config.Datasource) *gorm.DB {
	var creator = getCreatorByType(database.Type)
	if creator == nil {
		logx.Error.Fatalf("fail to find creator for types:%s", database.Type)
		return nil
	}
	db, err := creator.Create(database.IP, database.PORT, database.USER, database.PASSWORD, database.DATABASE)
	if err != nil {
		logx.Error.Fatalln(err)
		return nil
	}

	return db
}

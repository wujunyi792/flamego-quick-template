package database

import (
	"github.com/wujunyi792/gin-template-new/config"
	"github.com/wujunyi792/gin-template-new/internal/loging"
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
		loging.Info.Printf("create datasource %s => %s:%s", source.Key, source.IP, source.PORT)
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
		loging.Error.Fatalln("duplicate db key: " + key)
	}
	mux.Lock()
	defer mux.Unlock()
	dbs[key] = db
}

func mustCreateGorm(database config.Datasource) *gorm.DB {
	var creator = getCreatorByType(database.Type)
	if creator == nil {
		loging.Error.Fatalf("fail to find creator for type:%s", database.Type)
		return nil
	}
	db, err := creator.Create(database.IP, database.PORT, database.USER, database.PASSWORD, database.DATABASE)
	if err != nil {
		loging.Error.Fatalln(err)
		return nil
	}

	return db
}

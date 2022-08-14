package example

import (
	"github.com/wujunyi792/flamego-quick-template/internal/database"
	"github.com/wujunyi792/flamego-quick-template/internal/models/databasesModel"
	"github.com/wujunyi792/flamego-quick-template/pkg/logx"
	"gorm.io/gorm"
	"sync"
)

func init() {
	logx.Info.Println("start routerInitialize Table ...")
	dbManage = GetManage()
}

type DBManage struct {
	mDB     *gorm.DB
	sDBLock sync.RWMutex
}

var dbManage *DBManage = nil

func (manager *DBManage) getGOrmDB() *gorm.DB {
	return manager.mDB
}

func (manager *DBManage) atomicDBOperation(op func()) {
	manager.sDBLock.Lock()
	op()
	manager.sDBLock.Unlock()
}

func GetManage() *DBManage {
	if dbManage == nil {
		var userDb = database.GetDb("*")
		err := userDb.AutoMigrate(&databasesModel.Example{})
		if err != nil {
			logx.Error.Fatalln(err)
			return nil
		}
		dbManage = &DBManage{mDB: userDb}
	}
	return dbManage
}

package example

import (
	"github.com/wujunyi792/gin-template-new/internal/database"
	"github.com/wujunyi792/gin-template-new/internal/loging"
	"github.com/wujunyi792/gin-template-new/internal/model/Mysql"
	"gorm.io/gorm"
	"sync"
)

func init() {
	loging.Info.Println("start init Table ...")
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
		err := userDb.AutoMigrate(&Mysql.Example{})
		if err != nil {
			loging.Error.Fatalln(err)
			return nil
		}
		dbManage = &DBManage{mDB: userDb}
	}
	return dbManage
}

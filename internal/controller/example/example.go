package example

import (
	"github.com/wujunyi792/gin-template-new/internal/db"
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
	mDB     *db.MainGORM
	sDBLock sync.RWMutex
}

var dbManage *DBManage = nil

func (manager *DBManage) getGOrmDB() *gorm.DB {
	return manager.mDB.GetDB()
}

func (manager *DBManage) atomicDBOperation(op func()) {
	manager.sDBLock.Lock()
	op()
	manager.sDBLock.Unlock()
}

func GetManage() *DBManage {
	if dbManage == nil {
		var userDb = db.MustCreateGorm()
		err := userDb.GetDB().AutoMigrate(&Mysql.Example{})
		if err != nil {
			loging.Error.Fatalln(err)
			return nil
		}
		dbManage = &DBManage{mDB: userDb}
	}
	return dbManage
}

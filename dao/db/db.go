package db

import (
	"log"
	"os"
	"sync"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var engineMap map[string]*xorm.Engine
var defalultEngine string
var useOnce sync.Once
var beanList []interface{}

func init() {
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
		}
	}()
	useOnce.Do(func() {
		var err error
		engineMap = make(map[string]*xorm.Engine)
		engineMap[defalultEngine], err = xorm.NewEngine("mysql", "root:123456@tcp(host.docker.internal)/sql_test?charset=utf8mb4")
		if err != nil {
			log.Println(err)
		}
	})
}
func GetDbengine() *xorm.Engine {
	return engineMap[defalultEngine]
}
func SyncAppend(bean interface{}) {
	beanList = append(beanList, bean)
}
func Sync2Init() {
	err := GetDbengine().Sync2(beanList...)
	if err != nil {
		log.Println(err)
		os.Exit(-1)
	}
}

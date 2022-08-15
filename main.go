package main

import (
	"log"

	_ "github.com/tanweiquan/webmake/controllers" // main包要导入controllers
	"github.com/tanweiquan/webmake/dao/db"
	"github.com/tanweiquan/webmake/router"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			log.Fatalln(err)
		}
	}()
	db.Sync2Init()
	service := router.GETRouter()
	log.Println("运行开始")
	// go cron.StartTaskJops()
	_ = service.Run(":8080")

}

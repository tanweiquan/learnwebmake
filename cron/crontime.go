package cron

import (
	"fmt"
	"sync"
	"time"

	"github.com/robfig/cron"
)

// 每隔5秒执行一次：*/5 * * * * *
// 每隔1分钟执行一次：0 */1 * * * *
// 每天23点执行一次：0 0 23 * * *
// 每天凌晨1点执行一次：0 0 1 * * *
// 每月1号凌晨1点执行一次：0 0 1 1 * *
// 在26分、29分、33分执行一次：0 26,29,33 * * * *
// 每天的0点、13点、18点、21点都执行一次：0 0 0,13,18,21 * * *

var cronInstance *cron.Cron
var oncedo sync.Once

// 懒汉模式，多线程同时调用，使用once，可以避免创建多个对象，相当于nil判断+锁
func getCronInstance() *cron.Cron {
	oncedo.Do(func() {
		cronInstance = cron.New()
	})
	return cronInstance
}

func StartTaskJops() {
	// 每5秒执行一次
	getCronInstance().AddJob("*/5 * * * * *", Job{})
	//启动计划任务
	getCronInstance().Start()
	//关闭着计划任务, 但是不能关闭已经在执行中的任务.
	defer getCronInstance().Stop()
	select {}
}

type Job struct{}

func (t Job) Run() {
	fmt.Println(time.Now())
	fmt.Println("定时任务开始...")
}

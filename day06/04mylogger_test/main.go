package main

import (
	"github.com/Altair-cyx/day06/mylogger"
)

var log mylogger.Logger // 声明一个全局的接口变量

// 测试我们自己写的日志库

func main() {
	f1()

}

func f1() {
	log = mylogger.NewConsoleLog("Info")
	log = mylogger.NewFileLogger("Info", "./", "zhoulin.log", 10*1024*1024)
	for {

		log.Debug("这是一条Debug日志")
		log.Info("这是一条Info日志")
		log.Warning("这是一条Warning日志")
		id := 10010
		name := "理想"
		log.Error("这是一条Error日志,id:%d, name:%s", id, name)
		log.Trace("这是一条Trace日志")
		log.Fatal("这是一条Fatal日志")
		// time.Sleep(time.Second)
	}
}

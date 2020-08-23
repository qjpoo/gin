package main

import (
	"fmt"
	log "github.com/sirupsen/logrus" //别名
	"os"
	"time"
)

// 创建一个新的logger实例。可以创建任意多个。
var log1 = log.New()

func main() {
	// logrus示例

	// time="2020-08-22T21:28:22+08:00" level=info msg="hello, world ..." name=chiling
	log.WithFields(log.Fields{
		"name": "chiling",
		"age":  99,
	}).Info("hello, world ...")

	fmt.Println("----------------")
	// 设置日志输出到os.Stdout
	log1.Out = os.Stdout
	log1.WithField("xxoo", "ok?").Info("yes | no ")

	// 设置写到文件里面
	log2 := log.New()
	file, err := os.OpenFile("./09-logrus/logrus.log", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0777)
	if err != nil {
		//fmt.Println(err)
		log2.Info(err.Error())
	}
	log2.Out = file
	// time="2020-08-22T21:44:36+08:00" level=info msg=1111111111111111 fields.msg=ok fields.time="2020-08-22 21:44:36.9828365 +0800 CST m=+0.009000501" meta="hello, world"
	log2.WithFields(log.Fields{
		"time": time.Now(),
		"msg":  "ok",
		"meta": "hello, world",
	}).Info("1111111111111111")


	/*
	log.Trace("Something very low level.")
	log.Debug("Useful debugging information.")
	log.Info("Something noteworthy happened!")
	log.Warn("You should probably take a look at this.")
	log.Error("Something failed but I'm not quitting.")
	// 记完日志后会调用os.Exit(1)
	log.Fatal("Bye.")
	// 记完日志后会调用 panic()
	log.Panic("I'm bailing.")
	 */

	log.Info("这样调用也是可以的...")



	// 输出一个Json格式的数据
	// {"age":99,"level":"info","msg":"hello, world ...","name":"chiling","time":"2020-08-22T21:54:13+08:00"}
	log.SetFormatter(&log.JSONFormatter{})  // 设置成json
	// 如果你希望将调用的函数名添加为字段
	log.SetReportCaller(true)
	// {"age":99,"file":"C:/Users/瞿健/go/src/gin/09-logrus/main.go:67","func":"main.main","level":"info","msg":"hello, world ...","name":"chiling","time":"2020-08-22
	//T21:57:25+08:00"}
	log.WithFields(log.Fields{
		"name": "chiling",
		"age":  99,
	}).Info("hello, world ...")
}

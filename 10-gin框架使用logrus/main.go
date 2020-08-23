package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"os"
)

var log = logrus.New()

func init()  {
	//设置日志为json格式
	log.Formatter = &logrus.JSONFormatter{}
	
//	f, _ := os.Create("./10-gin框架使用logrus/gin.log")
	f, _ := os.OpenFile("./10-gin框架使用logrus/gin.log", os.O_APPEND|os.O_RDWR, 777)
	// log日志输出到文件里面
	log.Out = f
	
	// 生产模式,有关debug的信息不会打印出来
	gin.SetMode(gin.ReleaseMode)
	// gin的日志也要打印到文件中
	gin.DefaultWriter  = log.Out
	// gin.DefaultWriter = io.MultiWriter(f, os.Stdout, 192.168.11.120:5001) 即向终端输出,又向文件里面写
	// 日志的级别
	log.Level = logrus.InfoLevel

}

func main()  {

	r := gin.Default()

	r.Any("/any", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"msg": "ok",
		})
	})

	r.GET("hello", func(c *gin.Context) {
		log.WithFields(logrus.Fields{
			"name": "chiling",
			"age": 48,
			"address": "HongKong",
		}).Info("骚裱子")
	})
	r.Run(":8000")

	/*
	Logrus配有内置钩子。在init中添加这些内置钩子或你自定义的钩子：
	向syslog里面发日志
	import (
	  log "github.com/sirupsen/logrus"
	  "gopkg.in/gemnasium/logrus-airbrake-hook.v2" // the package is named "airbrake"
	  logrus_syslog "github.com/sirupsen/logrus/hooks/syslog"
	  "log/syslog"
	)

	func init() {

	  // Use the Airbrake hook to report errors that have Error severity or above to
	  // an exception tracker. You can create custom hooks, see the Hooks section.
	  log.AddHook(airbrake.NewHook(123, "xyz", "production"))

	  hook, err := logrus_syslog.NewSyslogHook("udp", "localhost:514", syslog.LOG_INFO, "")
	  if err != nil {
	    log.Error("Unable to connect to local syslog daemon")
	  } else {
	    log.AddHook(hook)
	  }
	}
	 */
}

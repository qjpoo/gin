package main

import (
	"fmt"
	"github.com/gamelife1314/logging"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"os"
)

func main() {
	/*
		一直都是用结构体接收参数，假如事先不清楚参数名，或者参数是不固定的，就要动态获取。
	*/

	r := gin.Default()

	r.POST("/a", func(c *gin.Context) {
		c.Request.ParseForm()
		for k, v := range c.Request.PostForm {
			fmt.Println("k: ", k)
			fmt.Println("v: ", v)
		}
	})

	r.POST("/j", func(c *gin.Context) {
		data, _ := ioutil.ReadAll(c.Request.Body)
		fmt.Println("c.request.body: ", string(data))
		//logging.Debugf("c.Request.Method: %v", c.Request.Method)
		//logging.Debugf("c.Request.ContentType: %v", c.ContentType())
		//logger := logging.GetDefaultLogger()
		//logger.Debug("hello world, %s", "logging")
		//logger.Info("hello world, %s", "logging")
		//logger.Warning("hello world, %s", "logging")
		//logger.Error("hello world, %s", "logging")
		//logger.Critical("hello world, %s", "logging")

		logFile, _ := os.OpenFile("./03-获取querystring参数/gin框架获取post请求的所有参数/log.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		defer logFile.Close()
		logger2 := &logging.Logger{
			Level: logging.DEBUG,
			StreamHandler: &logging.StreamMessageHandler{
				Level: logging.DEBUG,
				Formatter: &logging.MessageFormatter{
					Format:     `{{.Color}}[{{.Time}}] {{.LevelString | printf "%8s"}}  {{.FuncName}} {{.ShortFileName}} {{.Line}} {{.ColorClear}} {{.Message}}`,
					TimeFormat: "2006-01-02 15:04:05",
				},
				Destination: os.Stdout,
			},
			FileHandler: &logging.FileMessageHandler{
				Level: logging.ERROR,
				Formatter: &logging.MessageFormatter{
					Format: "[{{.Time}}] {{.LevelString}}  {{.Pid}} {{.Program}} {{.FuncName}} {{.LongFileName}} {{.Line}}{{.Message}}\n",
				},
				Destination: logFile,
			},
		}
		logger2.Debug("c.Request.Method: %v", c.Request.Method)
		logger2.Debug("c.Request.ContentType: %v", c.ContentType())

	})

	r.Run(":8000")

}

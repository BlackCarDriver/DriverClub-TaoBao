package main

import (
	md "TaobaoServer/models"
	_ "TaobaoServer/routers"
	"os"
	"os/signal"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/plugins/cors"
)

func main() {
	//设置跨域请求
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET, POST, PUT, DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "X-Requested-With"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true}))
	go destructor()
	beego.Run()
}

func destructor() {
	c := make(chan os.Signal)
	signal.Notify(c)
	sin := <-c
	//save static data to database
	logs.Info("Interupt single:%v", sin)
	md.UpdateStatic()
	logs.Warn("update static data success!")
	os.Exit(0)
}

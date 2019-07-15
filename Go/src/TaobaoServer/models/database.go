package models

import (
	"fmt"
	"time"

	"github.com/astaxie/beego/config"
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
)

var (
	host     = ""
	port     = 0
	userName = ""
	database = ""
	password = ""
)

func init() {
	//获取连接数据库配置
	iniconf, err := config.NewConfig("ini", "./conf/database.conf")
	if err != nil {
		panic(err)
	}
	userName = iniconf.String("userName")
	database = iniconf.String("database")
	password = iniconf.String("password")
	host = iniconf.String("host")
	port, err = iniconf.Int("port")
	if err != nil {
		panic(err)
	}

	//连接数据库
	dataSource := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, userName, password, database)
	err = orm.RegisterDataBase("default", "postgres", dataSource)
	if err != nil {
		fmt.Println("Can't not connect to database! : ", err)
		return
	} else {
		fmt.Println("DataBase connect scuess!!")
	}
	//设置最大空闲连接
	orm.SetMaxIdleConns("default", 30)
	// 设置为 UTC 时间
	orm.DefaultTimeLoc = time.UTC

	//#################################test
}

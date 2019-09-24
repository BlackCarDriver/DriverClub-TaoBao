package models

import (
	"fmt"
	"time"

	"github.com/astaxie/beego/logs"

	"github.com/astaxie/beego/config"
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
)

//database connection config
var (
	host     = ""
	port     = 0
	userName = ""
	database = ""
	password = ""
)

//other config
var (
	maxCreditBuff = 0 //the maximun of ActiveNess map
)

//logger specially used by models package function
var mlog *logs.BeeLogger

func init() {
	var err error
	//setting up logger
	mlog = logs.NewLogger()
	mlog.SetLogger("file", `{"filename":"logs/models.log"}`)
	mlog.EnableFuncCallDepth(true)
	mlog.Info("Router logs init success!")
	//get database config
	if iniconf, err := config.NewConfig("ini", "./conf/database.conf"); err != nil {
		mlog.Critical("%v", err)
		panic(err)
	} else {
		userName = iniconf.String("userName")
		database = iniconf.String("database")
		password = iniconf.String("password")
		host = iniconf.String("host")
		if port, err = iniconf.Int("port"); err != nil {
			mlog.Critical("%v", err)
			panic(err)
		}
	}

	//get models config
	if iniconf, err := config.NewConfig("ini", "./conf/driver.conf"); err != nil {
		mlog.Critical("%v", err)
		panic(err)
	} else {
		if maxCreditBuff, err = iniconf.Int("maxCreditBuff"); err != nil {
			mlog.Critical("%v", err)
			panic(err)
		}
	}

	//connect to database
	dataSource := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, userName, password, database)
	err = orm.RegisterDataBase("default", "postgres", dataSource)
	if err != nil {
		err = fmt.Errorf("Can't not connect to database! : %v", err)
		mlog.Critical("%v", err)
		panic(err)
	} else {
		mlog.Info("DataBase connect scuess!!")
	}
	//max unmbers of connection
	orm.SetMaxIdleConns("default", 30)
	orm.DefaultTimeLoc = time.UTC
	initTempData()
}

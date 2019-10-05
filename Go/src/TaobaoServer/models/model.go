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
	mlog.SetLogger("file", `{"filename":"logs/models.log","daily":false,"maxsize":512000}`)

	mlog.EnableFuncCallDepth(true)
	mlog.SetLevel(2)
	mlog.Info("Router logs init success!")

	//get database config
	iniconf, err := config.NewConfig("ini", "./conf/database.conf")
	if err != nil {
		mlog.Critical("%v", err)
		panic(err)
	}
	userName = iniconf.String("userName")
	database = iniconf.String("database")
	password = iniconf.String("password")
	host = iniconf.String("host")
	if port, err = iniconf.Int("port"); err != nil {
		mlog.Critical("%v", err)
		panic(err)
	}

	//get cache config
	rdadress = iniconf.String("rdadress")
	rdpassword = iniconf.String("rdpassword")
	if rdisuse, err = iniconf.Bool("rdisuse"); err != nil {
		mlog.Critical("%v", err)
		panic(err)
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

	//connect to redis database
	if rdisuse {
		UseCache = true
		if err = initReids(); err != nil {
			mlog.Critical("%v", err)
			panic(err)
		} else {
			mlog.Info("Redis connect success!!")
		}
	} else {
		mlog.Info("Resis is close...")
		UseCache = false
	}

	//set up connection of database connection
	orm.SetMaxIdleConns("default", 30)
	orm.DefaultTimeLoc = time.UTC
	initTempData()
}

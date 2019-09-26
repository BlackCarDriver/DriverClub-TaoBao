package models

import (
	"errors"
	"strings"

	"github.com/astaxie/beego/logs"

	"github.com/go-redis/redis"
)

var RedisC *redis.Client

//database connection config
var (
	rdadress   = ""
	rdpassword = ""
	rdisuse    = false //whether use cache function
)

//global value
var (
	UseCache  bool  = false
	NotInUsed error = errors.New("Cache function is not using")
	IsExist   error = errors.New("The key is exist")
)

//init redis
func initReids() error {
	RedisC = redis.NewClient(&redis.Options{
		Addr:     rdadress,
		Password: rdpassword,
		DB:       0,
	})
	if _, err := RedisC.Ping().Result(); err != nil {
		mlog.Critical("connect to redis server failed %v", err)
		return err
	}
	return nil
}

//note that it is forbiden to set a new value before time-out or delete the value
func SetCache(value string, second int, tag ...string) error {
	if !UseCache {
		return NotInUsed
	}
	if second <= 0 {
		return errors.New("Invalid value of second.")
	}
	if len(tag) == 0 {
		return errors.New("No tag was given")
	}
	err, key := combineKey(tag...)
	if err != nil {
		return nil
	}
	isExist, _ := RedisC.Do("exists", key).Int()
	if isExist == 1 {
		return IsExist
	}
	err = RedisC.Do("set", key, value, "EX", second).Err()
	if err != nil {
		mlog.Error("%v", err)
		return err
	}
	return nil
}

func GetCache(tag ...string) (error, string) {
	if !UseCache {
		return NotInUsed, ""
	}
	err, key := combineKey(tag...)
	if err != nil {
		mlog.Error("%v", err)
		return err, ""
	}
	if isExist, err := RedisC.Do("exists", key).Bool(); err != nil {
		return err, ""
	} else if !isExist {
		return errors.New("value not exist"), ""
	}
	result, err := RedisC.Do("get", key).String()
	if err != nil {
		mlog.Error("%v", err)
		return err, ""
	}
	return nil, result
}

//note that it delete a value that not exist will return no error, because the value might be clear because of time-out
func DelCache(tag ...string) error {
	if !UseCache {
		return NotInUsed
	}
	err, key := combineKey(tag...)
	if err != nil {
		mlog.Error("%v", err)
		return err
	}
	if num, err := RedisC.Do("del", key).Int(); err != nil {
		mlog.Error("%v", err)
		return err
	} else if num > 1 {
		logs.Warn("delete '%s' affect more than one!", key)
	}
	return nil
}

//===================== tool function ================================

func combineKey(tag ...string) (error, string) {
	if len(tag) == 0 {
		return errors.New("no argument was given"), ""
	}
	return nil, strings.Join(tag, "-")
}

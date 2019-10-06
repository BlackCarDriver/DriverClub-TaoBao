package models

import (
	"encoding/json"
	"errors"
	"fmt"

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
	UseCache    bool  = false
	NotInUsed   error = errors.New("Cache function is not using")
	IsExist     error = errors.New("The key is exist")
	NotExist    error = errors.New("Cache not exist")
	NoCacheTime error = errors.New("CacheTime is invalid")
	KeyNotFound error = errors.New("CacheKey not found")
	WornResult  error = errors.New("Cache is null")
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
func SetCache(req *RequestProto, any interface{}) error {
	if !UseCache {
		return NotInUsed
	}
	if req.CacheTime <= 0 {
		return NoCacheTime
	}
	if req.CacheKey == "" {
		return KeyNotFound
	}
	var err error
	//parse interface into string
	value := parseToString(any)
	if value == "" {
		err = fmt.Errorf("Can't not parse interface %s to cache", req.CacheKey)
		mlog.Error("%v", err)
		return err
	}
	//check value if exist
	isExist, _ := RedisC.Do("exists", req.CacheKey).Int()
	if isExist == 1 {
		return IsExist
	}
	//save to redis database
	err = RedisC.Do("set", req.CacheKey, value, "EX", req.CacheTime).Err()
	if err != nil {
		mlog.Error("set cache fail: %v", err)
		return err
	}
	logs.Warn("Set cache success! %s", req.CacheKey)
	return nil
}

//get cache from redis
func GetCache(req *RequestProto) (cache string, err error) {
	if !UseCache {
		return "", NotInUsed
	}
	if req.CacheTime <= 0 {
		return "", NoCacheTime
	}
	if req.CacheKey == "" {
		return "", KeyNotFound
	}
	//check if exist
	if isExist, err := RedisC.Do("exists", req.CacheKey).Bool(); err != nil {
		mlog.Error("check exist fail: %v", err)
		return "", err
	} else if !isExist {
		return "", NotExist
	}
	//get value if is exist
	result, err := RedisC.Do("get", req.CacheKey).String()
	if err != nil {
		mlog.Error("get cache fail: %v", err)
		return "", err
	}
	if result == "" {
		return "", WornResult
	}
	return result, nil
}

//note that it delete a value that not exist will return no error,
//because the value might be clear because of time-out
func DelCache(req *RequestProto) error {
	if !UseCache {
		return NotInUsed
	}
	if req.CacheTime <= 0 {
		return NoCacheTime
	}
	if req.CacheKey == "" {
		return KeyNotFound
	}
	if num, err := RedisC.Do("del", req.CacheKey).Int(); err != nil {
		mlog.Error("delete cache fial:%v", err)
		return err
	} else if num > 1 {
		logs.Warn("delete '%s' affect more than one rows!", req.CacheKey)
	}
	return nil
}

//check if a operation in request is too frequent 🍜
//return true mean the operation is execute in near time
func CheckFrequent(req *RequestProto) bool {
	if req.CacheKey == "" || req.CacheTime <= 0 {
		mlog.Error("CheckFrequent function receive a null cache_key")
		return true
	}
	if _, err := GetCache(req); err == nil {
		return true
	}
	SetCache(req, "haha")
	return false
}

//===================== tool function ================================

//parse a obeject into json encoding string, return null string if error happend
func parseToString(any interface{}) string {
	bs, err := json.Marshal(any)
	if err != nil {
		mlog.Error("ParseToString fail: %v", err)
		return ""
	}
	return string(bs)
}

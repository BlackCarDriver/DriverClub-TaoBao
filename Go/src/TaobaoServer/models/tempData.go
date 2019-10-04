package models

import (
	"errors"
	"fmt"
	"time"

	"github.com/astaxie/beego/logs"
)

/*
tempDate.go save some galbol data to avoid excessive database too frequently
*/
//public varlue
var (
	RunHour = 0 //how many hours the backend already run
)

//Refresh all temp data when start the backend 🌰
func initTempData() {
	//Uas1 and Uas2 used to save the credits of users for a while
	Uas1 = NewActiveNess()
	Uas2 = NewActiveNess()
	//the value in ComfirmCode will be save for half hour
	ComfirmCode = NewTimeMap(60 * 30)

	RefreshTypeTagData()
	RefreshUserRank()

	go RunPreHour()
}

func RunPreHour() {
	for _ = range time.NewTicker(1 * time.Hour).C {
		nowHour := time.Now().Hour()
		RunHour++
		mlog.Info("==========The % bus is going to start!=======", RunHour)
		//refresh credits and clear activeness record each hour
		MainTainCredits()
		Uas1.ReBuild()
		Uas2.ReBuild()
		ComfirmCode.Clear()

		RefreshTypeTagData()

		MainTainGoodLike()
		MainTainGoodCollect()
		MainTainGoodTalk()

		//refresh user rank one times each day
		if nowHour == 0 {
			mlog.Info("Begin to refresh user rank...")
			RefreshUserRank()
		}
		//maintain level data and rank data three times each day
		if nowHour%6 == 0 {
			MainTainLevel()
			MainTainRank()
		}
	}
}

//============ Goods Type ===============
//save the goods type and tag data
var GoodsTypeTempDate = []GoodsType{
	{"学习用品", []GoodsSubType{}},
	{"体育用品", []GoodsSubType{}},
	{"生活用品", []GoodsSubType{}},
	{"电子产品", []GoodsSubType{}},
	{"手工diy", []GoodsSubType{}},
	{"虚拟商品", []GoodsSubType{}},
	{"其他", []GoodsSubType{}},
}

func RefreshTypeTagData() {
	for i := 0; i < len(GoodsTypeTempDate); i++ {
		GetTagsData(GoodsTypeTempDate[i].Type, &GoodsTypeTempDate[i].List)
	}
}

//============ User Rank ==================
//save top 10 user rank
var UserRank []Rank

func RefreshUserRank() {
	if err := GetRankList(&UserRank); err != nil {
		mlog.Error("%v", err)
	}
}

//===================== User active static ===========
//user active statistics
var Uas1, Uas2 *ActiveNess

//used to count the degree of user visit in a short time, It is what credits statics basis for 🌰
type ActiveNess struct {
	active map[string]int
	max    int
}

func NewActiveNess() *ActiveNess {
	var t *ActiveNess
	t = new(ActiveNess)
	t.Init(maxCreditBuff)
	return t
}
func (a *ActiveNess) Init(max int) {
	a.active = make(map[string]int)
	a.max = max
}
func (a *ActiveNess) Add(uid string) {
	if uid == "" || a.active[uid] >= a.max {
		return
	}
	a.active[uid]++
}
func (a *ActiveNess) ReBuild() {
	a.active = make(map[string]int)
}
func (a *ActiveNess) GetMap() map[string]int {
	return a.active
}

//============== Save some message for a limited time ============ 🍖
//save the comfirm code
var ComfirmCode *TimeMap

//timeMap is a countainer that only save the data for a duration
type TimeMap struct {
	Map  map[string]time.Time
	Life int
}

//create a timeMap
func NewTimeMap(second int) *TimeMap {
	var t *TimeMap = new(TimeMap)
	t.Map = make(map[string]time.Time)
	t.Life = second
	return t
}

//save a key in the map for a while 🍖
func (t *TimeMap) Add(val string) error {
	if _, ok := t.Map[val]; ok {
		err := fmt.Errorf("The value still exist: %s", val)
		mlog.Error("%v", err)
		return err
	}
	t.Map[val] = time.Now()
	logs.Info("Save comfirm code: %s", val)
	return nil
}

//judge if the key have out of data 🍖
//return nil mean the key is found and not out of data
func (t *TimeMap) Get(key string) error {
	var err error
	val, ok := t.Map[key]
	if ok == false {
		err = errors.New("验证码错误")
		mlog.Info("%v", err)
		return err
	}
	afterSecond := int(time.Since(val).Seconds())
	if afterSecond > t.Life {
		err = errors.New("验证码过期")
		mlog.Info("%v", err)
		return err
	}
	logs.Info("Get comfirm %s  afterSecond: %v", key, afterSecond)
	return nil
}

//clear all key that already out of date 🍖
func (t *TimeMap) Clear() {
	for k, v := range t.Map {
		duration := time.Since(v)
		logs.Warn("Comfirm code: %s \t\t\t %d", k, int(duration.Minutes()))
		if int(duration.Seconds()) > t.Life {
			mlog.Warn("timer key %s have been delete", k)
			delete(t.Map, k)
		}
	}
}

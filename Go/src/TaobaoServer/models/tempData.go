package models

import (
	"errors"
	"fmt"
	"time"

	"github.com/astaxie/beego/logs"

	"github.com/astaxie/beego/orm"
)

/*
tempDate.go save some galbol data to avoid excessive database too frequently
*/
//public varlue
var (
	RunHour = 0
)

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

//save top 10 user rank
var UserRank []Rank

//Refresh all temp data when start the process,🌰
func initTempData() {
	//Uas1 and Uas2 used to save the credits of users for a while
	Uas1 = NewActiveNess()
	Uas2 = NewActiveNess()
	//the value in ComfirmCode will be save for half hour
	ComfirmCode = NewTimeMap(60 * 30)

	RefreshTypeTagDate()
	UpdateUserRank()
	go RunPreHour()
}

func RefreshTypeTagDate() {
	for i := 0; i < len(GoodsTypeTempDate); i++ {
		GetTagsData(GoodsTypeTempDate[i].Type, &GoodsTypeTempDate[i].List)
	}
}

func UpdateUserRank() {
	if err := GetRankList(&UserRank); err != nil {
		mlog.Error("%v", err)
	}
}

//==================== active defintion =======================
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

//========= datastruct for save the comfirm code for a while ============ 🍖
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

//==================== Timmer Bus ===================================
//to execute those timed job
func RunPreHour() {
	for _ = range time.NewTicker(1 * time.Hour).C {
		RunHour++
		mlog.Info("==========The % bus is going to start!=======", RunHour)
		//refresh credits and clear activeness record each hour
		MainTainCredits()
		Uas1.ReBuild()
		Uas2.ReBuild()
		ComfirmCode.Clear()
		time.Sleep(10 * time.Second)
		RefreshTypeTagDate()
		time.Sleep(10 * time.Second)
		MainTainLevel()
		time.Sleep(10 * time.Second)
		MainTainRank()
		time.Sleep(10 * time.Second)
		MainTainGoodLike()
		time.Sleep(10 * time.Second)
		MainTainGoodCollect()
		time.Sleep(10 * time.Second)
		MainTainGoodTalk()
		time.Sleep(10 * time.Second)
		UpdateUserRank()

	}
}

//==================== database select function ======================

//get the list of top 10 user's rank, data include id, name and credits
func GetRankList(c *[]Rank) error {
	o := orm.NewOrm()
	if num, err := o.Raw(`select * from v_rank`).QueryRows(c); err != nil {
		mlog.Error("%v", err)
		return fmt.Errorf("Get data error: %v", err)
	} else if num == 0 {
		err := fmt.Errorf("the result is empty!")
		mlog.Error("%v", err)
		return err
	}
	return nil
}

//get all tag name and tag number of a type
func GetTagsData(gtype string, tag *[]GoodsSubType) error {
	if gtype == "" {
		return errors.New("Receive a null gtype")
	}
	o := orm.NewOrm()
	var tSubType []GoodsSubType
	num, err := o.Raw(`select tag, count(*) as number from t_goods where type = $1 group by tag`, gtype).QueryRows(&tSubType)
	if err != nil {
		mlog.Error("%v", err)
		return err
	}
	if num == 0 {
		err = fmt.Errorf("the result is empty!")
		mlog.Warn("%v", err)
		return err
	}
	var sum int64 = 0
	for i := 0; i < len(tSubType); i++ {
		sum += tSubType[i].Number
	}
	slice := make([]GoodsSubType, len(tSubType)+1)
	copy(slice, []GoodsSubType{{"全部", sum}})
	copy(slice[1:], tSubType)
	*tag = make([]GoodsSubType, len(tSubType)+1)
	copy(*tag, slice)
	return nil
}

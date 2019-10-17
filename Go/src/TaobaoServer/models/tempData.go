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
the value with flag '👀' need to add some motier code around the functions
*/
//============== static data ======================== 🍙
var (
	RunHour           = 0   //how many hours the backend already run
	TotalGoodsNum     = 0   //total goods number (state not small than 0)
	TotalTagNum       = 0   //total goods tag numbers
	TotalUserNum      = 0   //user number
	TotalGoodsPrice   = 0.0 //sum of all goods price（state not small than 0)
	TotalCommendNum   = 0   //total times of goods's commend 👀
	TotalPVMsgNum     = 0   //private message send times 👀
	TotalDealNumber   = 0   //number os removed goods	👀
	TotalDealPrice    = 0   //total price of removed goods	👀
	TotalFBTimes      = 0   //total feedback times	👀
	TodayNewUser      = 0   //how many user sign in today 👀
	TodayNewGoods     = 0   //how many goods uploaded today 👀
	TodayVStimes      = 0   //homepage visit times in last hour 👀
	TodayRequestTimes = 0   //how many request today have response 👀
	TotalVisitTimes   = 0   //total homepage visit times 👀
	TotalRequestTimes = 0   //total request times 👀
	LastUpdateTime    = ""  //last time of the static data refreshtion
)

//struct of static data fontend interface need
type Static struct {
	Key   string      `json:"key"`
	Value interface{} `json:"value"`
}

//the temp data return to fontend directly
var StaticData []Static

//update StaticData
func RefreshStaticData() {
	TotalGoodsNum = CountOnlineGoods()
	TotalGoodsPrice = CountTotalPrice()
	TotalTagNum = CountGoodsTag()
	TotalUserNum = CountTotalUser()

	TotalCommendNum = GetIntStaticData("TotalCommendNum")
	TotalPVMsgNum = GetIntStaticData("TotalPVMsgNum")
	TotalFBTimes = GetIntStaticData("TotalFBTimes")
	TotalDealNumber = GetIntStaticData("TotalDealNumber")
	TotalDealPrice = GetIntStaticData("TotalDealPrice")

	TotalVisitTimes = GetIntStaticData("TotalVisitTimes")
	TotalRequestTimes = GetIntStaticData("TotalRequestTimes")

	LastUpdateTime = time.Now().Format("01-02 15:04")
	StaticData = GetStaticData()
}

//manually update before close the process
func UpdateStatic() {
	UpdateStaticIntData("TotalVisitTimes", TodayVStimes)        //👀
	UpdateStaticIntData("TotalRequestTimes", TodayRequestTimes) //👀
}

//update some static data when a new day start
func UpdateStaticPreDay() {
	UpdateStaticIntData("TotalVisitTimes", TodayVStimes)        //👀
	UpdateStaticIntData("TotalRequestTimes", TodayRequestTimes) //👀
	his := GetStaticData()
	InsertToStatic("static_history", fmt.Sprintf("%#v", his))
	TodayRequestTimes = 0
	TodayVStimes = 0
	TodayNewUser = 0
	TodayNewGoods = 0
}

//Used to updata StaticData 🍙
func GetStaticData() []Static {
	var data = []Static{
		{"在线商品数", TotalGoodsNum},
		{"在线商品总价值", TotalGoodsPrice},
		{"标签总数", TotalTagNum},
		{"注册人数", TotalUserNum},
		{"评论总次数", TotalCommendNum},
		{"反馈总次数", TotalFBTimes},
		{"成功交易次数", TotalDealNumber},
		{"成功交易面额", TotalDealPrice},
		{"今日新增用户", TodayNewUser},
		{"今日新增商品", TodayNewGoods},
		{"主页总访问量", TotalVisitTimes},
		{"今日主页访问量", TodayVStimes},
		{"总处理请求总数", TotalRequestTimes},
		{"今日处理请求总数", TodayRequestTimes},
		{"本版本后端运行时长(h)", RunHour},
		{"本统计信息更新时间", LastUpdateTime},
	}
	return data
}

//================================= time bus ================================

//Refresh all temp data when start the backend 🌰
func initTempData() {
	//Uas1 and Uas2 used to save the credits of users for a while
	Uas1 = NewActiveNess()
	Uas2 = NewActiveNess()
	//the value in ComfirmCode will be save for half hour
	ComfirmCode = NewTimeMap(60 * 30)
	//init static data
	RefreshStaticData()
	StaticData = GetStaticData()
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
		MainTainRank()
		MainTainGoodLike()
		MainTainGoodCollect()
		MainTainGoodTalk()

		RefreshStaticData()

		//execute add 0:00 to 1:00
		if nowHour == 0 {
			mlog.Info("Begin to start mid night clock...")
			//refresh user rank one times each day
			RefreshUserRank()
			//reset some static data
			UpdateStaticPreDay()
			//add 50 credits to driver account each day
			AwardDriver()
		}
		//update goods rank at 3:00
		if nowHour == 3 {
			MaintainGoodsState()
			MainTainLevel()

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

//============================= User active static =============================
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

//======================== Save some message for a limited time ===================== 🍖
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

//################### function relate to get and update static ####################🍙

//add the init static data to database when first times run backend🍙
//Usually it function only called one times!!!
func InitStaticTable() {
	o := orm.NewOrm()
	number := 0
	o.Raw("select count(*) from t_static where keyname='TotalCommendNum'").QueryRow(&number)
	if number != 0 {
		return
	}
	var queryRows = []string{
		"INSERT INTO public.t_static(keyname, numberval) VALUES ('TotalCommendNum', 0)",
		"INSERT INTO public.t_static(keyname, numberval) VALUES ('TotalPVMsgNum', 0)",
		"INSERT INTO public.t_static(keyname, numberval) VALUES ('TotalDealNumber', 0)",
		"INSERT INTO public.t_static(keyname, numberval) VALUES ('TotalFBTimes', 0)",
		"INSERT INTO public.t_static(keyname, numberval) VALUES ('TotalDealPrice', 0)",
		"INSERT INTO public.t_static(keyname, numberval) VALUES ('TotalVisitTimes', 0)",
		"INSERT INTO public.t_static(keyname, numberval) VALUES ('TotalRequestTimes', 0)",
	}
	for _, row := range queryRows {
		_, err := o.Raw(row).Exec()
		if err != nil {
			mlog.Error("Init static data fail: %s  ===> %v", row, err)
		} else {
			mlog.Info("Init static data success: %s", row)
		}
	}
}

//get interge tpye data from t_static table
func GetIntStaticData(key string) int {
	o := orm.NewOrm()
	number := 0
	err := o.Raw("select numberval from t_static where keyname=?", key).QueryRow(&number)
	if err != nil {
		mlog.Critical("Get Int StaticData %s fail: %v", key, err)
		return -1
	}
	return number
}

//get string type data from t_static table
func GetStrStaticData(key string) string {
	o := orm.NewOrm()
	str := ""
	err := o.Raw("select stringval from t_static where keyname=?", key).QueryRow(&str)
	if err != nil {
		mlog.Critical("Get string StaticData fail: %v", err)
		return ""
	}
	return str
}

//add change to value if fund the key
func UpdateStaticIntData(key string, change int) {
	o := orm.NewOrm()
	_, err := o.Raw("update t_static set numberval= numberval+? where keyname=?", change, key).Exec()
	if err != nil {
		mlog.Error("try to update static fail: key=%s  change=%d", key, change)
	}
}

//replay value with change if find the key
func UpdateStaticStrData(key string, newStr string) {
	o := orm.NewOrm()
	_, err := o.Raw("update t_static set stringval=? where keyname=?", newStr, key).Exec()
	if err != nil {
		mlog.Error("try to update static fail: key=%s  newVal=%d", key, newStr)
	}
}

//add a record to t_static
func InsertToStatic(key string, value string) {
	o := orm.NewOrm()
	if _, err := o.Raw("INSERT INTO public.t_static(keyname, stringval) VALUES (?, ?)", key, value).Exec(); err != nil {
		mlog.Error("insert into t_static fail: %v", err)
	}

}

//=============== funciton about notification email receive =============  🍣
//only those user with change more than 1 will reveive email notification when his reveive a private message.

//check the setting of user whether receive email notification, return the times of recevie chance
//return 0 if any error happen
func GetRecevieChange(uid string) int {
	if uid == "" {
		return 0
	}
	change, err := GetIntCache("notifiChange_" + uid)
	if err != nil {
		return 0
	}
	return change
}

//decrease the times of receive email change
func SubReceiveChance(uid string) {
	if uid == "" {
		return
	}
	key := "notifiChange_" + uid
	if chance := GetRecevieChange(uid); chance <= 0 {
		return
	} else {
		SetIntCache(key, chance-1, 6000000)
	}
}

//add the times of reveive email chance after all change have been used
func ResetReceiveChange(uid string) {
	if uid == "" {
		return
	}
	key := "notifiChange_" + uid
	if chance := GetRecevieChange(key); chance > 0 {
		return
	} else {
		SetIntCache(key, 3, 6000000)
	}
}

//delete the cache of user receive email setting
func DelReveiveChange(uid string) {
	if uid == "" {
		return
	}
	key := "notifiChange_" + uid
	DelCacheByKey(key)
}

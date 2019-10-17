package models

import (
	"math"
	"time"
	"github.com/astaxie/beego/orm"
)

/*
maintain.go statistics and update some value in database
*/

//refresh the number of user like for each goods 🍄
func MainTainGoodLike() {
	type like struct {
		Id  string `json:"id"`
		Num int    `json:"num"`
	}
	var rows = make([]like, 0)
	o := orm.NewOrm()
	if _, err := o.Raw(`select goodsid as id, count(*) as num from t_goods_like group by goodsid`).QueryRows(&rows); err != nil {
		mlog.Error("%v",err)
	}
	for _, row := range rows {
		if _, err := o.Raw(`update t_goods set "like"=? where id=?`, row.Num, row.Id).Exec(); err != nil {
			mlog.Error("%v",err)
			continue
		}
	}
	mlog.Info("Goods like data maintain success!")
}

//refresh the number of collect for all goods🍄
func MainTainGoodCollect() {
	type like struct {
		Id  string `json:"id"`
		Num int    `json:"num"`
	}
	var rows = make([]like, 0)
	o := orm.NewOrm()
	if _, err := o.Raw(`select count(*) as num, goodsid as id from t_collect group by goodsid`).QueryRows(&rows); err != nil {
		mlog.Error("%v",err)
	}
	for _, row := range rows {
		if _, err := o.Raw(`update t_goods set collect=? where id=?`, row.Num, row.Id).Exec(); err != nil {
			mlog.Error("%v",err)
			continue
		}
	}
	mlog.Info("Goods collect data maintain success!")
}

//refresh the number of comment for each goods 🍄
func MainTainGoodTalk() {
	type like struct {
		Id  string `json:"id"`
		Num int    `json:"num"`
	}
	var rows = make([]like, 0)
	o := orm.NewOrm()
	if _, err := o.Raw(`select count(*) as num, goodsid as id from t_comment group by goodsid`).QueryRows(&rows); err != nil {
		mlog.Error("%v",err)
	}
	for _, row := range rows {
		if _, err := o.Raw(`update t_goods set talk=? where id=?`, row.Num, row.Id).Exec(); err != nil {
			mlog.Error("%v",err)
			continue
		}
	}
	mlog.Info("Goods talk maintin success!")
}

//refresh the leave of each user 🌰
func MainTainLevel() {
	type crdit struct {
		Id      string `json:"id"`
		Credits int    `json:"credits"`
	}
	var rows = make([]crdit, 0)
	o := orm.NewOrm()
	if _, err := o.Raw(`select id, credits from t_user`).QueryRows(&rows); err != nil {
		mlog.Error("%v",err)
	}
	for _, row := range rows {
		level := countLevel(row.Credits)
		if _, err := o.Raw(`update t_user set leave=? where id=?`, level, row.Id).Exec(); err != nil {
			mlog.Error("%v",err)
			continue
		}
	}
	mlog.Info("User level maintail success!")
}

//refresh the rank of all users 🌰
func MainTainRank() {
	type rank struct {
		Id   string `json:"id"`
		Rank int    `json:"rank"`
	}
	var rows = make([]rank, 0)
	o := orm.NewOrm()
	if _, err := o.Raw(`SELECT row_number() OVER (ORDER BY credits DESC) as rank, id from t_user`).QueryRows(&rows); err != nil {
		mlog.Error("%v",err)
	}
	for _, row := range rows {
		if _, err := o.Raw(`update t_user set rank=? where id=?`, row.Rank, row.Id).Exec(); err != nil {
			mlog.Error("%v",err)
			continue
		}
	}
	mlog.Info("User rank maintain success!")
}

//update user credits according to ActiveNess 🌰
func MainTainCredits() {
	o := orm.NewOrm()
	for id,val := range Uas1.GetMap() {
		if val == 0{
			continue
		}
		mlog.Info("update user credits %s + %d", id, val)
		if _, err := o.Raw(`update t_user set credits=credits+? where id=?`, val, id).Exec(); err != nil {
			mlog.Error("%v",err)
			continue
		}
	}
	for id,val := range Uas2.GetMap() {
		if val == 0{
			continue
		}
		mlog.Info("update user credits %s + %d", id, val)
		if _, err := o.Raw(`update t_user set credits=credits+? where id=?`, val, id).Exec(); err != nil {
			mlog.Error("%v",err)
			continue
		}
	}
	mlog.Info("User credits maintain success!")
}

//permanent delete those removed goods data from database
func DeleteRMGoods() {
	//TODO: design
}

//add credits to those account which have a driver name 🍛
func AwardDriver(){
	o := orm.NewOrm()
	var driverId []string 
	_, err := o.Raw("select id from t_user where name ~ '.*Driver.*' or name ~ '.*司机.*'").QueryRows(&driverId)
	if err != nil {
		mlog.Error("Find Driver fail: %v", err)
	}
	for _, id := range driverId {
		if _, err := o.Raw("update t_user set credits = credits + 50 where id =?", id).Exec();err!=nil {
			mlog.Error("Add credit to user %s fail: %v", id, err)
		}
	}
}

//update goods state which is used to rank the goods🍛
func MaintainGoodsState(){
	type tmpStruct struct{
		Goodsid string `json:"goodsid"`
		Rank int	`json:"rank"`
		Time time.Time `json:"time"`
		Like int 	`json:"like"`
	}
	var gd []tmpStruct 
	o := orm.NewOrm()
	_, err := o.Raw(`select g.id as goodsid, u.rank, ul.time, g.like from t_user as u,
	 t_upload as ul, t_goods as g where u.id=ul.userid and ul.goodsid=g.id and g.state>=0 `).QueryRows(&gd)
	if err != nil {
		mlog.Error("Maintain goods state fail: %v", err)
		return
	}
	for i:=0; i<len(gd); i++ {
		newstate := countGoodsRank(gd[i].Rank, gd[i].Like, gd[i].Time)
		_, err := o.Raw("Update t_goods set state=? where id=?", newstate, gd[i].Goodsid).Exec()
		if err!=nil {
			mlog.Error("Update %d state fail: %v", gd[i].Goodsid, err)
		}
	}
}


//=============== tool function ================
//It is the medhod how counting the level of user🌰
func countLevel(greditInt int) int {
	if greditInt <= 0 {
		return 0
	}
	gredits := float64(greditInt)
	var level float64
	if gredits <= 1024.0 {
		level = math.Floor(math.Log2(gredits))
	} else {
		level = 10.0 + math.Floor((gredits-1024.0)/1000)
	}
	return int(level)
}

//count how many hour have been go since 2019-10-01 🍛
func countSinceHour(t time.Time) int {
	startTime, _ := time.Parse("2006-01-02 15:04:05", "2019-1-01 22:22:04")
	duration := t.Sub(startTime)
	return int(duration.Hours())
}

//count the weight of a goods according to user rank, upload time and goods like times🍛
func countGoodsRank(rank int, like int, uploadTime time.Time) int {
	afterHour := countSinceHour(uploadTime)
	return afterHour + rank*20 + like
}

//get what is the hour of the day 🍛
func getNowHour() int {
	h, _, _ := time.Now().Clock()
	return h
}
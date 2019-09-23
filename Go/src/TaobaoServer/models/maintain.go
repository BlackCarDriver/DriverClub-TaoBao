package models

import (
	"github.com/astaxie/beego/logs"

	"github.com/astaxie/beego/orm"
)

//refresh the number of user like for each goods 🍄
func MainTainGoodLike() {
	type like struct {
		Id  string `json:"id"`
		Num int    `json:"num"`
	}
	var rows = make([]like, 0)
	o := orm.NewOrm()
	if _, err := o.Raw(`select goodsid as id, count(*) as num from t_goods_like group by goodsid`).QueryRows(&rows); err != nil {
		logs.Error(err)
	}
	for _, row := range rows {
		if _, err := o.Raw(`update t_goods set "like"=? where id=?`, row.Num, row.Id).Exec(); err != nil {
			logs.Error(err)
			continue
		}
	}
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
		logs.Error(err)
	}
	for _, row := range rows {
		if _, err := o.Raw(`update t_goods set collect=? where id=?`, row.Num, row.Id).Exec(); err != nil {
			logs.Error(err)
			continue
		}
	}
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
		logs.Error(err)
	}
	for _, row := range rows {
		if _, err := o.Raw(`update t_goods set talk=? where id=?`, row.Num, row.Id).Exec(); err != nil {
			logs.Error(err)
			continue
		}
	}
}

package models

import (
	"fmt"

	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

//update user base message 🍊
func UpdateUserBaseMsg(d UpdeteMsg) error {
	o := orm.NewOrm()
	rawSeter := o.Raw("update t_user set name=?,sex=?,sign=?,dorm=?,major=?,grade=? where id=?",
		d.Name, d.Sex, d.Sign, d.Dorm, d.Major, d.Grade, d.Id)
	result, err := rawSeter.Exec()
	if err != nil {
		logs.Error("user: %s, error:%v", d.Id, err)
		return err
	}
	effect, _ := result.RowsAffected()
	if effect == 0 {
		err = fmt.Errorf("No Roow Affected !")
		logs.Error("user: %s, error:%v", d.Id, err)
		return err
	}
	return nil
}

//update user's connection message 🍊
func UpdateUserConnectMsg(d UpdeteMsg) error {
	o := orm.NewOrm()
	rawSeter := o.Raw("update t_user set emails=?, phone=?, qq=? where id=?;",
		d.Emails, d.Phone, d.Qq, d.Id)
	result, err := rawSeter.Exec()
	if err != nil {
		logs.Error("user:%s, error:%v", d.Id, err)
		return err
	}
	effect, _ := result.RowsAffected()
	if effect == 0 {
		err := fmt.Errorf("No Roow Affected !")
		logs.Error("user: %s, error:%v", d.Id, err)
		return err
	}
	return nil
}

//更新用户头像
func UpdateUserHeadIMg(imgurl, userid string) error {
	o := orm.NewOrm()
	rawSeter := o.Raw("update t_user set headimg=? where id=?;", imgurl, userid)
	result, err := rawSeter.Exec()
	if err != nil {
		logs.Error(err)
		return err
	}
	effect, _ := result.RowsAffected()
	if effect == 0 {
		err = fmt.Errorf("No Roow Affected !")
		logs.Error(err)
		return err
	}
	return nil
}

//主页被浏览，更新浏览量
func UpdateUserVisit(uid string) error {
	o := orm.NewOrm()
	rawSeter := o.Raw(`UPDATE t_user SET visit=visit+1 WHERE id = ?`, uid)
	result, err := rawSeter.Exec()
	if err != nil {
		logs.Error(err)
		return err
	}
	effect, _ := result.RowsAffected()
	if effect == 0 {
		err = fmt.Errorf("No Roow Affected !")
		logs.Error(err)
		return err
	}
	return nil
}

//商品被浏览，更新浏览量
func UpdateGoodsVisit(gid string) error {
	o := orm.NewOrm()
	rawSeter := o.Raw(`UPDATE t_goods SET visit=visit+1 WHERE id = ?`, gid)
	result, err := rawSeter.Exec()
	if err != nil {
		logs.Error(err)
		return err
	}
	effect, _ := result.RowsAffected()
	if effect == 0 {
		err := fmt.Errorf("No Roow Affected !")
		logs.Error(err)
		return err
	}
	return nil
}

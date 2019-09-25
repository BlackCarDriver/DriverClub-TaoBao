package models

import (
	"github.com/astaxie/beego/logs"
	"errors"
	"fmt"

	"github.com/astaxie/beego/orm"
)

//update user base message 🍊
func UpdateUserBaseMsg(d UpdeteMsg) error {
	o := orm.NewOrm()
	rawSeter := o.Raw("update t_user set name=?,sex=?,sign=?,dorm=?,major=?,grade=? where id=?",
		d.Name, d.Sex, d.Sign, d.Dorm, d.Major, d.Grade, d.Id)
	result, err := rawSeter.Exec()
	if err != nil {
		mlog.Error("user: %s, error:%v", d.Id, err)
		return err
	}
	effect, _ := result.RowsAffected()
	if effect == 0 {
		err = fmt.Errorf("No Roow Affected !")
		mlog.Error("user: %s, error:%v", d.Id, err)
		return err
	}
	Uas2.Add(d.Id) 	//user change hiself profile, credits+1
	return nil
}

//update user's connection message 🍊
func UpdateUserConnectMsg(d UpdeteMsg) error {
	o := orm.NewOrm()
	rawSeter := o.Raw("update t_user set emails=?, phone=?, qq=? where id=?;",
		d.Emails, d.Phone, d.Qq, d.Id)
	result, err := rawSeter.Exec()
	if err != nil {
		mlog.Error("user:%s, error:%v", d.Id, err)
		return err
	}
	effect, _ := result.RowsAffected()
	if effect == 0 {
		err := fmt.Errorf("No Roow Affected !")
		mlog.Error("user: %s, error:%v", d.Id, err)
		return err
	}
	Uas2.Add(d.Id) 	//user change hiself profile, credits+1
	return nil
}

//upadate the profile image url of user 
func UpdateUserHeadIMg(userid,   imgurl string) error {
	o := orm.NewOrm()
	rawSeter := o.Raw("update t_user set headimg=? where id=?;", imgurl, userid)
	result, err := rawSeter.Exec()
	if err != nil {
		mlog.Error("%v", err)
		return err
	}
	effect, _ := result.RowsAffected()
	if effect == 0 {
		err = fmt.Errorf("No Row Affected !")
		mlog.Error("%v", err)
		return err
	}
	return nil
}

//update the visited times of a user profile page
func UpdateUserVisit(uid string) error {
	o := orm.NewOrm()
	rawSeter := o.Raw(`UPDATE t_user SET visit=visit+1 WHERE id = ?`, uid)
	result, err := rawSeter.Exec()
	if err != nil {
		mlog.Error("%v", err)
		return err
	}
	effect, _ := result.RowsAffected()
	if effect == 0 {
		err = fmt.Errorf("No Roow Affected !")
		mlog.Error("%v", err)
		return err
	}
	Uas2.Add(uid)	//homepage have been visited, credits+1
	return nil
}

//update the visited times of a goods
func UpdateGoodsVisit(gid string) error {
	o := orm.NewOrm()
	rawSeter := o.Raw(`UPDATE t_goods SET visit=visit+1 WHERE id = ?`, gid)
	result, err := rawSeter.Exec()
	if err != nil {
		mlog.Error("%v", err)
		return err
	}
	effect, _ := result.RowsAffected()
	if effect == 0 {
		err := fmt.Errorf("No Roow Affected !")
		mlog.Error("%v", err)
		return err
	}
	return nil
}

//set a goods state as -1 which mean it goods have been delete 🍑
func UpdateMyGoodsState(uid, gid string) error {
	if uid == "" || gid == "" {
		err := errors.New("Got a null userid or goodsid")
		mlog.Error("%v", err)
		return err
	}
	o := orm.NewOrm()
	//check whether this goods exit and the user is right
	count := 0
	if err := o.Raw(`select count(*) from v_mygoods where uid =? and id=?`, uid, gid).QueryRow(&count); err != nil {
		mlog.Error("Count row in v_mygoods fail: %v", err)
		return err
	} else if count == 0 {
		err = fmt.Errorf("No row found in v_mygoods when want to delete: uid:%s gid:%s", uid, gid)
		mlog.Error("%v", err)
		return err
	}
	if res, err := o.Raw("update t_goods set state = -1 where id = ?", gid).Exec(); err != nil {
		mlog.Error("%v", err)
		return err
	} else if af, err := res.RowsAffected(); err != nil {
		mlog.Warn("%v", err)
	} else if af == 0 {
		err = fmt.Errorf("No rows affacted when user %s update goods %s state", uid, gid)
		mlog.Error("%v", err)
		return err
	}
	return nil
}

//change the state of user message which mean already read  🍞
func UpdateMessageState(mid string) error {
	if mid=="" {
		return errors.New("Receive a empty message")
	}
	logs.Warn(mid)
	o := orm.NewOrm()
	if res, err := o.Raw("update t_message set state=1 WHERE id = ?", mid).Exec(); err != nil {
		mlog.Error("%v", err)
		return err
	}else if af, err := res.RowsAffected(); err != nil {
		mlog.Warn("%v", err)
	} else if af == 0 {
		err = fmt.Errorf("No row affected when update state of message %s", mid)
		mlog.Error("%v",err)
		return err
	}
	return nil
}
package models

import (
	"github.com/astaxie/beego/logs"
	"errors"
	"fmt"

	"github.com/astaxie/beego/orm"
)
/*
NOTE about goods state:
1 is the default value
-1 mean user remove the goods by hisself
*/

//update user base message 🍊 🍆🍖🍙
func UpdateUserBaseMsg(d UpdeteMsg, tid string) error {
	o := orm.NewOrm()
	//check user name, can't same with other people if name is changed
	if nameNum := CountOtherUserName(d.Name, tid); nameNum!=0 {
		logs.Error(nameNum)
		err := fmt.Errorf("User name %s already have been used! Please change to another", d.Name)
		mlog.Info("%v",err)
		return err
	}
	rawSeter := o.Raw(`update t_user set name=?,sex=?,sign=?,
	grade=?,major=?,colleage=?,dorm=? where id=?`,
		d.Name, d.Sex, d.Sign, d.Grade, d.Major, d.Colleage, d.Dorm, d.Id)
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
	//update goods state to -1
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
	goodsPrice := 0
	o.Raw("select price from t_goods where id=?", gid).QueryRow(&goodsPrice)
	//update static data 👀
	UpdateStaticIntData("TotalDealNumber",1)
	UpdateStaticIntData("TotalDealPrice",goodsPrice) 
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

//update the state of feedback record 🍗
func UpdateFeedbackState(fbid int) error {
	if fbid < 0 {
		err := errors.New("Receive id smaller than 0")
		mlog.Error("%v", err)
		return err
	}
	o := orm.NewOrm()
	updateTP := `update t_feedback set fb_status=1 where id=?`
	if res, err := o.Raw(updateTP, fbid).Exec(); err != nil {
		mlog.Error("%v", err)
		return err
	}else if af, err := res.RowsAffected(); err != nil {
		mlog.Warn("%v", err)
	} else if af == 0 {
		err = errors.New("No row affected when update state of message")
		mlog.Error("%v",err)
		return err
	}
	return nil
}

//update user's last login time 🍛
func UpdateLoginTime(uid string) error {
	o := orm.NewOrm()
	if res, err := o.Raw(`update t_user set lasttime = now() where id = ?`, uid).Exec(); err != nil {
		mlog.Error("%v", err)
		return err
	}else if af, err := res.RowsAffected(); err != nil {
		mlog.Warn("%v", err)
	} else if af == 0 {
		err = errors.New("No row affected when update logintime")
		mlog.Error("%v",err)
		return err
	}
	return nil
}
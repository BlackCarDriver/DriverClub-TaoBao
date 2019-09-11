package models

import (
	"errors"
	"fmt"

	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

//delete a message that user receive 🍑
func DeleteMyMessage(uid, mid string) error {
	if uid == "" || mid == "" {
		err := errors.New("Got a null userid or messageid")
		logs.Error(err)
		return err
	}
	o := orm.NewOrm()
	//delete message from database
	if res, err := o.Raw("delete from t_message where receiverid = ? and id = ?", uid, mid).Exec(); err != nil {
		logs.Error(err)
		return err
	} else if af, err := res.RowsAffected(); err != nil {
		logs.Warn(err)
	} else if af == 0 {
		err = fmt.Errorf("No rows affacted when user %s delete message %s", uid, mid)
		logs.Error(err)
		return err
	}
	return nil
}

//cancel a goods collect 🍑
func DeleteMyCollect(uid, gid string) error {
	if uid == "" || gid == "" {
		err := errors.New("Got a null userid or goodsid")
		logs.Error(err)
		return err
	}
	o := orm.NewOrm()
	//check whether this goods exit and the user is right
	count := 0
	if err := o.Raw(`select count(*) from v_mycollect where uid =? and id=?`, uid, gid).QueryRow(&count); err != nil {
		logs.Error("Count row in v_mycollect fail: %v", err)
		return err
	} else if count == 0 {
		err = fmt.Errorf("No row found in v_mycollect when want to cancel collect: uid:%s, gid:%s", uid, gid)
		logs.Error(err)
		return err
	}
	//delete the collect record
	if res, err := o.Raw("delete from t_collect where userid = ? and goodsid = ?", uid, gid).Exec(); err != nil {
		logs.Error(err)
		return err
	} else if af, err := res.RowsAffected(); err != nil {
		logs.Warn(err)
	} else if af == 0 {
		err = fmt.Errorf("No rows affacted when user %s cancel collect goods %s", uid, gid)
		logs.Error(err)
		return err
	}
	return nil
}

//calcel a user collect 🍑
func DeleteMyConcern(uid1, uid2 string) error {
	if uid1 == "" || uid2 == "" {
		err := errors.New("Got a null userid")
		logs.Error(err)
		return err
	}
	o := orm.NewOrm()
	//check whether record is exist
	count := 0
	if err := o.Raw(`select count(*) from t_concern where id1 =? and id2=?`, uid1, uid2).QueryRow(&count); err != nil {
		logs.Error("Count row in t_concern fail: %v", err)
		return err
	} else if count == 0 {
		err = fmt.Errorf("No row found in t_concern when want to cancel concern: uid1:%s, uid2:%s", uid1, uid2)
		logs.Error(err)
		return err
	}
	//delete the collect record
	if res, err := o.Raw("delete from t_concern where id1 = ? and id2 = ?", uid1, uid2).Exec(); err != nil {
		logs.Error(err)
		return err
	} else if af, err := res.RowsAffected(); err != nil {
		logs.Warn(err)
	} else if af == 0 {
		err = fmt.Errorf("No rows affacted when user %s cancel concern user %s", uid1, uid2)
		logs.Error(err)
		return err
	}
	return nil
}

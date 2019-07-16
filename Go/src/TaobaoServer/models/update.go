package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
)

//更新用户基本信息
//需要确保ID是用户自己的ID
func UpdateUserBaseMsg(d UpdeteMsg) error {
	o := orm.NewOrm()
	rawSeter := o.Raw("update t_user set,name=?,sex=?,sign=?,dorm=?,major=? where id=?",
		d.Name, d.Sex, d.Sign, d.Dorm, d.Major, d.Id)
	result, err := rawSeter.Exec()
	if err != nil {
		return err
	}
	effect, _ := result.RowsAffected()
	if effect == 0 {
		return fmt.Errorf("No Roow Affected !")
	}
	return nil
}

//更新联系方式
func UpdateUserConnectMsg(d UpdeteMsg) error {
	o := orm.NewOrm()
	rawSeter := o.Raw("update t_user set emails=?, phone=?, qq=? where id=?;",
		d.Emails, d.Phone, d.Qq, d.Id)
	result, err := rawSeter.Exec()
	if err != nil {
		return err
	}
	effect, _ := result.RowsAffected()
	if effect == 0 {
		return fmt.Errorf("No Roow Affected !")
	}
	return nil
}

//更新用户头像
func UpdateUserHeadIMg(imgurl, userid string) error {
	o := orm.NewOrm()
	rawSeter := o.Raw("update t_user set headimg=? where id=?;", imgurl, userid)
	result, err := rawSeter.Exec()
	if err != nil {
		return err
	}
	effect, _ := result.RowsAffected()
	if effect == 0 {
		return fmt.Errorf("No Roow Affected !")
	}
	return nil
}

package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
)

//更新用户基本信息
//需要确保ID是用户自己的ID
func UpdateUserBaseMsg(d UpdeteMsg) error {
	o := orm.NewOrm()
	rawSeter := o.Raw("update t_user set name=?,sex=?,sign=?,dorm=?,major=?,grade=? where id=?",
		d.Name, d.Sex, d.Sign, d.Dorm, d.Major, d.Grade, d.Id)
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

//某商品被点赞，点赞数加1
func UpdateGoodsLike(gid string) error {
	o := orm.NewOrm()
	rawSeter := o.Raw(`update t_goods set "like" = "like" + 1 where id = ?`, gid)
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

//某人被点赞，点赞数加1
func UpdateUserLike(uid string) error {
	o := orm.NewOrm()
	rawSeter := o.Raw(`UPDATE t_user SET likes= likes+1 WHERE id = ?`, uid)
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

//主页被浏览，更新浏览量
func UpdateUserVisit(uid string) error {
	o := orm.NewOrm()
	rawSeter := o.Raw(`UPDATE t_user SET visit=visit+1 WHERE id = ?`, uid)
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

//商品被浏览，更新浏览量
func UpdateGoodsVisit(gid string) error {
	o := orm.NewOrm()
	rawSeter := o.Raw(`UPDATE t_goods SET visit=visit+1 WHERE id = ?`, gid)
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

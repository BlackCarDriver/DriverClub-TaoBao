package models

import (
	"fmt"
	"time"

	"github.com/astaxie/beego/logs"

	"database/sql"

	"github.com/astaxie/beego/orm"
)

//一些数据库默认值
const (
	dfUserHeadimg = `https://gss0.bdstatic.com/6LZ1dD3d1sgCo2Kml5_Y_D3/sys/portrait/item/2f6de585abe7baa7e5a4a7e78b82e9a38e5a`
	dfUserName    = `尊贵的用户`
	dfGoodHeadimg = `https://gss0.bdstatic.com/6LZ1dD3d1sgCo2Kml5_Y_D3/sys/portrait/item/c62bcfccb5b0b9b7c8cbc132?t=1526199816`
)

//创建用户账号
//id自动生成,注意在调用此函数前需要确保name,password,email非空
func CreateAccount(user RegisterData) error {
	o := orm.NewOrm()
	userNumber := CountUser() + 1
	t := time.Now()
	userid := fmt.Sprintf("%02d%02d%04d", t.Year()%100, t.Month(), userNumber)
	rawSeter := o.Raw("insert into t_user(id, email, password, name, headimg) values(?,?,?,?,?)",
		userid, user.Email, user.Password, user.Name, dfUserHeadimg)
	_, err := rawSeter.Exec()
	if err != nil {
		logs.Error(err)
	}
	return err
}

//insert a good to database 🍋
func CreateGoods(goods UploadGoodsData) error {
	o := orm.NewOrm()
	var err error
	goodsNumber := CountGoods() + 1
	t := time.Now()
	goodsid := fmt.Sprintf("%02d%02d%02d%04d", t.Year()%100, t.Month(), t.Day(), goodsNumber)
	logs.Info(goods.UserId, goodsid)
	insertGoods := o.Raw("insert into t_goods(id, name, title, type, tag, price, file, headimg)values(?,?,?,?,?,?,?,?)",
		goodsid, goods.Name, goods.Title, goods.Type, goods.Tag, goods.Price, goods.Text, goods.Imgurl)
	insertUpload := o.Raw("insert into t_upload(userid, goodsid) values(?, ?)", goods.UserId, goodsid)

	err = o.Begin()
	if err != nil {
		logs.Error(err)
		return err
	}
	_, err1 := insertGoods.Exec()
	_, err2 := insertUpload.Exec()

	if err1 != nil || err2 != nil {
		logs.Warn("Need to Rollback!! t_goods: %v ,  t_upload: %v ", err1, err2)
		//rollback
		if err3 := o.Rollback(); err3 != nil {
			logs.Error("Rollback fail: ", err3)
		} else {
			logs.Info("RollBack success!")
		}
		if err1 != nil {
			err = fmt.Errorf("updata t_goods fail: %v", err1)
		} else {
			err = fmt.Errorf("update t_upload fail: %v", err2)
		}
	} else {
		logs.Info("Create Goods Scuueed!!")
		err = o.Commit()
	}
	return err
}

//某商品被收藏，记录收藏信息
func AddCollectRecord(uid, gid string) error {
	o := orm.NewOrm()
	rawSeter := o.Raw(`INSERT INTO t_collect(userid, goodsid) VALUES (?, ?)`, uid, gid)
	_, err := rawSeter.Exec()
	if err != nil {
		logs.Error(err)
	}
	return err
}

//发出私信，更新消息表
func AddUserMessage(uid, targetid, message string) error {
	o := orm.NewOrm()
	rawSeter := o.Raw(`INSERT INTO public.t_message(senderid, receiverid, content) VALUES (?, ?, ?)`, uid, targetid, message)
	_, err := rawSeter.Exec()
	if err != nil {
		logs.Error(err)
	}
	return err
}

//某商品被收藏，更新收藏表
func AddGoodsCollect(uid, gid string) error {
	o := orm.NewOrm()
	var err error
	var result sql.Result
	count := 0
	fmt.Println(uid, gid)
	err = o.Raw(`SELECT count(*) from t_collect where userid=? and goodsid=?`, uid, gid).QueryRow(&count)
	if err != nil {
		err := fmt.Errorf("Error when select: %s", err)
		logs.Error(err)
		return err
	}
	if count > 0 {
		err := fmt.Errorf("You are already Collect it goods!")
		logs.Error(err)
		return err
	}
	result, err = o.Raw(`INSERT INTO t_collect(userid, goodsid)VALUES (?, ?)`, uid, gid).Exec()
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

//某人被关注，更新关注表
func AddUserConcern(id1, id2 string) error {
	o := orm.NewOrm()
	var err error
	var result sql.Result
	count := 0
	err = o.Raw(`SELECT count(*) FROM t_concern where id1=? and id2=?`, id1, id2).QueryRow(&count)
	if err != nil {
		err := fmt.Errorf("Error when select from t_concern: %s", err)
		logs.Error(err)
		return err
	}
	if count > 0 {
		err := fmt.Errorf("You are already concern it user!")
		logs.Error(err)
		return err
	}
	result, err = o.Raw(`INSERT INTO t_concern(id1, id2)VALUES (?, ?)`, id1, id2).Exec()
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

//insert a goods_like record
func AddGoodsLike(uid, gid string) error {
	o := orm.NewOrm()
	var err error
	var result sql.Result
	result, err = o.Raw(`INSERT INTO public.t_goods_like(userid, goodsid)VALUES (?, ?)`, uid, gid).Exec()
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

//save a user_like record
func AddUserLike(uid1, uid2 string) error {
	o := orm.NewOrm()
	var err error
	var result sql.Result
	result, err = o.Raw(`INSERT INTO public.t_user_like(userid1, userid2)VALUES (?, ?)`, uid1, uid2).Exec()
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

//save a goods comment
func AddGoodsComment(uid, gid, conetnt string) error {
	o := orm.NewOrm()
	var err error
	var result sql.Result
	result, err = o.Raw(`INSERT INTO public.t_comment(userid, goodsid, content)VALUES (?, ?, ?)`, uid, gid, conetnt).Exec()
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

package models

import (
	"errors"
	"fmt"
	"time"

	"database/sql"

	"github.com/astaxie/beego/orm"
)

//一些数据库默认值
const (
	dfUserHeadimg = `https://gss0.bdstatic.com/6LZ1dD3d1sgCo2Kml5_Y_D3/sys/portrait/item/2f6de585abe7baa7e5a4a7e78b82e9a38e5a`
	dfUserName    = `尊贵的用户`
	dfGoodHeadimg = `https://gss0.bdstatic.com/6LZ1dD3d1sgCo2Kml5_Y_D3/sys/portrait/item/c62bcfccb5b0b9b7c8cbc132?t=1526199816`
	masterId = "19070010"
)

//Create a account autoly by provided name, password and email 🍖
func CreateAccount(user RegisterData) error {
	o := orm.NewOrm()
	//check the username again
	if nameNumber := CountUserName(user.Name); nameNumber!= 0 {
		err := fmt.Errorf("User name %s already have been used!", user.Name)
		mlog.Error("%v",err)
		return err
	}
	//make a userid by the following regular
	userNumber := CountUser() + 1
	t := time.Now()
	userid := fmt.Sprintf("%02d%02d%04d", t.Year()%100, t.Month(), userNumber)
	rawSeter := o.Raw("insert into t_user(id, email, password, name, headimg) values(?,?,?,?,?)",
		userid, user.Email, user.Password, user.Name, dfUserHeadimg)
	_, err := rawSeter.Exec()
	if err != nil {
		mlog.Error("%v",err)
		return err
	}
	//send user a private message
	if err := SendSystemMsg(userid, HelloMsgToNewUser); err!= nil {
		mlog.Error("%v",err)
	}
	return nil
}

//insert a good to database 🍋
func CreateGoods(goods UploadGoodsData) error {
	o := orm.NewOrm()
	var err error
	goodsNumber := CountGoods() + 1
	t := time.Now()
	goodsid := fmt.Sprintf("%02d%02d%02d%04d", t.Year()%100, t.Month(), t.Day(), goodsNumber)
	mlog.Info(goods.UserId, goodsid)
	insertGoods := o.Raw("insert into t_goods(id, name, title, type, tag, price, file, headimg)values(?,?,?,?,?,?,?,?)",
		goodsid, goods.Name, goods.Title, goods.Type, goods.Tag, goods.Price, goods.Text, goods.Imgurl)
	insertUpload := o.Raw("insert into t_upload(userid, goodsid) values(?, ?)", goods.UserId, goodsid)

	err = o.Begin()
	if err != nil {
		mlog.Error("%v",err)
		return err
	}
	_, err1 := insertGoods.Exec()
	_, err2 := insertUpload.Exec()

	if err1 != nil || err2 != nil {
		mlog.Warn("Need to Rollback!! t_goods: %v ,  t_upload: %v ", err1, err2)
		//rollback
		if err3 := o.Rollback(); err3 != nil {
			mlog.Error("Rollback fail: ", err3)
		} else {
			mlog.Info("RollBack success!")
		}
		if err1 != nil {
			err = fmt.Errorf("updata t_goods fail: %v", err1)
		} else {
			err = fmt.Errorf("update t_upload fail: %v", err2)
		}
	} else {
		mlog.Info("Create Goods Scuueed!!")
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
		mlog.Error("%v",err)
	}
	Uas2.Add(uid) 	//user collect a goods, credits+1
	return err
}

//private message sending
func AddUserMessage(uid, targetid, message string) error {
	o := orm.NewOrm()
	rawSeter := o.Raw(`INSERT INTO public.t_message(senderid, receiverid, content) VALUES (?, ?, ?)`, uid, targetid, message)
	_, err := rawSeter.Exec()
	if err != nil {
		mlog.Error("%v",err)
	}
	Uas2.Add(uid) //user send message, credits+1
	return err
}

//user collect a goods, update t_user_collect
func AddGoodsCollect(uid, gid string) error {
	o := orm.NewOrm()
	var err error
	var result sql.Result
	count := 0
	err = o.Raw(`SELECT count(*) from t_collect where userid=? and goodsid=?`, uid, gid).QueryRow(&count)
	if err != nil {
		err := fmt.Errorf("Error when select: %s", err)
		mlog.Error("%v",err)
		return err
	}
	if count > 0 {
		err := fmt.Errorf("You are already Collect it goods!")
		mlog.Error("%v",err)
		return err
	}
	result, err = o.Raw(`INSERT INTO t_collect(userid, goodsid)VALUES (?, ?)`, uid, gid).Exec()
	if err != nil {
		mlog.Error("%v",err)
		return err
	}
	effect, _ := result.RowsAffected()
	if effect == 0 {
		err := fmt.Errorf("No Roow Affected !")
		mlog.Error("%v",err)
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
		mlog.Error("%v",err)
		return err
	}
	if count > 0 {
		err := fmt.Errorf("You are already concern it user!")
		mlog.Error("%v",err)
		return err
	}
	result, err = o.Raw(`INSERT INTO t_concern(id1, id2)VALUES (?, ?)`, id1, id2).Exec()
	if err != nil {
		mlog.Error("%v",err)
		return err
	}
	effect, _ := result.RowsAffected()
	if effect == 0 {
		err := fmt.Errorf("No Roow Affected !")
		mlog.Error("%v",err)
		return err
	}
	Uas2.Add(id1) //two user credits +1
	Uas2.Add(id2) 
	return nil
}

//insert a goods_like record
func AddGoodsLike(uid, gid string) error {
	o := orm.NewOrm()
	var err error
	var result sql.Result
	result, err = o.Raw(`INSERT INTO public.t_goods_like(userid, goodsid)VALUES (?, ?)`, uid, gid).Exec()
	if err != nil {
		mlog.Error("Insert t_goods_like fail: %v, user:%s, goods:%s", err, uid, gid)
		return err
	}
	effect, _ := result.RowsAffected()
	if effect == 0 {
		err := fmt.Errorf("No Roow Affected !")
		mlog.Error("%v",err)
		return err
	}
	Uas2.Add(uid)
	return nil
}

//save a user_like record
func AddUserLike(uid1, uid2 string) error {
	o := orm.NewOrm()
	var err error
	var result sql.Result
	result, err = o.Raw(`INSERT INTO public.t_user_like(userid1, userid2)VALUES (?, ?)`, uid1, uid2).Exec()
	if err != nil {
		mlog.Error("uid1: %s, uid2:%s, error:%v", uid1, uid2, err)
		return err
	}
	effect, _ := result.RowsAffected()
	if effect == 0 {
		err := fmt.Errorf("No Roow Affected !")
		mlog.Error("%v",err)
		return err
	}
	Uas2.Add(uid1)
	return nil
}

//save a goods comment 🍉
func AddGoodsComment(uid, gid, conetnt string) error {
	if uid == "" || gid == "" || conetnt == "" {
		return errors.New("Argument not right, get a empty id or comment content")
	}
	o := orm.NewOrm()
	var err error
	var result sql.Result
	result, err = o.Raw(`INSERT INTO public.t_comment(userid, goodsid, content)VALUES (?, ?, ?)`, uid, gid, conetnt).Exec()
	if err != nil {
		mlog.Error("%v",err)
		return err
	}
	effect, _ := result.RowsAffected()
	if effect == 0 {
		err := fmt.Errorf("No Roow Affected !")
		mlog.Error("%v",err)
		return err
	}
	Uas2.Add(uid)
	return nil
}

//send a system message to user 🍖
func SendSystemMsg(uid, msg string) error {
	o := orm.NewOrm()
	rawSeter := o.Raw(`INSERT INTO public.t_message(senderid, receiverid, content) VALUES (?, ?, ?)`, masterId, uid, msg)
	_, err := rawSeter.Exec()
	if err != nil {
		mlog.Error("%v",err)
		return err
	}
	return nil
}

//save a feedback record 🍗
func AddFeedback(d *FeedBackData) error {
	var err error 
	if d == nil {
		err = errors.New("Receive a nil pointer")
		mlog.Error("%v",err)
		return nil
	}
	insertTP := `INSERT INTO t_feedback(user_id, fb_location, 
		fb_type, imgurl, describes, email)VALUES (?,?,?,?,?,?)`
	o := orm.NewOrm()
	if _, err = o.Raw(insertTP,d.UserId, d.Location, d.Type, d.Imgurl, d.Describes, d.Email).Exec();err != nil {
		mlog.Error("%v",err)
		return err
	}
	return nil
}
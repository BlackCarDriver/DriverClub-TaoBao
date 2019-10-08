package models

import (
	"errors"
	"fmt"
	"time"

	"database/sql"

	"github.com/astaxie/beego/orm"
)

//some default value of database
const (
	dfUserHeadimg = `https://img-blog.csdnimg.cn/20191003114954113.jpg`
	dfGoodHeadimg = `https://img-blog.csdnimg.cn/20191003114954113.jpg`
	masterId      = "19100001"
)

// some pulic message template
const (
	HelloMsgToNewUser = ` [系统消息] 欢迎并感谢你成为本站的会员！本站仍然在开发之中，很多地方有待完善，欢迎到反馈页面反馈问题以及向我发送私聊，
我会认对待每一条建议和反馈，谢谢！ 让我们共同努力，将本站打造成一个实用和有趣的社区！`
	GoodsHanveBeenCollectTp = `[系统消息] 你的商品 %s 刚刚被用户 %s 收藏了哦！`
	GoodsHanveBeenLikeTP    = `[系统消息] 你的商品 %s 刚刚被用户 %s 点赞了哦！`
	GoodsHanveBeenTalkTP    = `[系统消息] 你的商品 %s 刚刚收到了来自用户 %s 的评论哦！`
	UserHaveBeenConcernTP   = `[系统消息] 刚才 %s 在你的主页关注了你~`
	UserHaveBeenLikeTP      = `[系统消息] 刚才 %s 在你的主页点赞了~`
)

//Create a account autoly by provided name, password and email 🍖🍚🍙🍜
//note that the password  should be md5 encoded
func CreateAccount(user RegisterData) error {
	o := orm.NewOrm()
	//check the username and email again
	if nameNumber := CountUserName(user.Name); nameNumber != 0 {
		err := fmt.Errorf("User name %s already have been used", user.Name)
		mlog.Error("%v", err)
		return err
	}
	if emailNumber := CountRegistEmail(user.Email); emailNumber != 0 {
		err := fmt.Errorf("Email %s already have been used", user.Name)
		mlog.Error("%v", err)
		return err
	}
	if CountUserId(user.Name) != 0 {
		err := fmt.Errorf("User name %s is same as a exist id", user.Name)
		mlog.Error("%v", err)
		return err
	}
	//make a userid by the following regular
	userNumber := CountTotalUser() + 1
	t := time.Now()
	userid := fmt.Sprintf("%02d%02d%04d", t.Year()%100, t.Month(), userNumber)
	rawSeter := o.Raw("insert into t_user(id, email, password, name, headimg, rank) values(?,?,?,?,?,?)",
		userid, user.Email, user.Password, user.Name, dfUserHeadimg, userNumber)
	_, err := rawSeter.Exec()
	if err != nil {
		mlog.Error("create new account fail: %v", err)
		return err
	}
	//send user a private message to user account
	if err := SendSystemMsg(userid, HelloMsgToNewUser); err != nil {
		mlog.Error("%v", err)
	}
	//update static data 👀
	TodayNewUser++
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
		mlog.Error("%v", err)
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
	//update static data 👀
	TodayNewGoods++
	return err
}

//private message sending 🍚
func AddUserMessage(uid, targetid, message string) error {
	if uid == "" || targetid == "" || message == "" {
		err := errors.New("get empty userid or message content")
		mlog.Error("%v", err)
		return err
	}
	o := orm.NewOrm()
	rawSeter := o.Raw(`INSERT INTO public.t_message(senderid, receiverid, content) VALUES (?, ?, ?)`, uid, targetid, message)
	_, err := rawSeter.Exec()
	if err != nil {
		mlog.Error("%v", err)
	}
	Uas2.Add(uid)                           //user send message, credits+1
	UpdateStaticIntData("TotalPVMsgNum", 1) //👀
	return err
}

//user collect a goods, update t_user_collect 🍠
func AddGoodsCollect(uid, gid string) error {
	o := orm.NewOrm()
	var err error
	var result sql.Result
	count := 0
	//check if collect before
	err = o.Raw(`SELECT count(*) from t_collect where userid=? and goodsid=?`, uid, gid).QueryRow(&count)
	if err != nil {
		err := fmt.Errorf("Error when select: %s", err)
		mlog.Error("%v", err)
		return err
	}
	if count > 0 {
		err := fmt.Errorf("You are already Collect it goods!")
		mlog.Error("%v", err)
		return err
	}
	result, err = o.Raw(`INSERT INTO t_collect(userid, goodsid)VALUES (?, ?)`, uid, gid).Exec()
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
	//send a message to owner of goods
	if oid, err := GetOwnerId(gid); err != nil {
		msg := fmt.Sprintf(GoodsHanveBeenCollectTp, GetGNameById(gid), GetUNameById(uid))
		err = SendSystemMsg(oid, msg)
		if err != nil {
			mlog.Error("send goods collect msg to user %s fail : %v", uid, err)
		}
	}
	Uas2.Add(uid) //user collect a goods, credits+1
	return nil
}

//user concern by others, id1 concern id2 🍠
func AddUserConcern(id1, id2 string) error {
	o := orm.NewOrm()
	var err error
	var result sql.Result
	count := 0
	err = o.Raw(`SELECT count(*) FROM t_concern where id1=? and id2=?`, id1, id2).QueryRow(&count)
	if err != nil {
		err := fmt.Errorf("Error when select from t_concern: %s", err)
		mlog.Error("%v", err)
		return err
	}
	if count > 0 {
		err := fmt.Errorf("You are already concern it user!")
		mlog.Error("%v", err)
		return err
	}
	result, err = o.Raw(`INSERT INTO t_concern(id1, id2)VALUES (?, ?)`, id1, id2).Exec()
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
	//tell user he or she have been concern by other
	msg := fmt.Sprintf(UserHaveBeenConcernTP, GetUNameById(id1))
	if SendSystemMsg(id2, msg); err != nil {
		mlog.Error("send concern message fail: %v", err)
	}
	Uas2.Add(id1) //two user credits +1
	Uas2.Add(id2)
	return nil
}

//insert a goods_like record 🍠
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
		mlog.Error("%v", err)
		return err
	}
	//send a message to owner of goods
	if oid, err := GetOwnerId(gid); err != nil {
		msg := fmt.Sprintf(GoodsHanveBeenLikeTP, GetGNameById(gid), GetUNameById(uid))
		err = SendSystemMsg(oid, msg)
		if err != nil {
			mlog.Error("send goods like msg to user: %v", err)
		}
	}
	Uas2.Add(uid)
	return nil
}

//save a user_like record,uid1 like uid2 🍠
func AddUserLike(uid1, uid2 string) error {
	if uid1 == "" || uid2 == "" {
		err := errors.New("uid or uid is null")
		mlog.Error("%v", err)
		return err
	}
	o := orm.NewOrm()
	result, err := o.Raw(`INSERT INTO public.t_user_like(userid1, userid2)VALUES (?, ?)`, uid1, uid2).Exec()
	if err != nil {
		mlog.Error("uid1: %s, uid2:%s, error:%v", uid1, uid2, err)
		return err
	}
	effect, _ := result.RowsAffected()
	if effect == 0 {
		err := fmt.Errorf("No Roow Affected !")
		mlog.Error("%v", err)
		return err
	}
	//send a message to owner of goods
	msg := fmt.Sprintf(UserHaveBeenLikeTP, GetUNameById(uid1))
	if err := SendSystemMsg(uid2, msg); err != nil {
		mlog.Error("send user like msg to fail: %v", err)
	}
	Uas2.Add(uid1)
	return nil
}

//save a goods comment 🍉🍠
func AddGoodsComment(uid, gid, conetnt string) error {
	if uid == "" || gid == "" || conetnt == "" {
		return errors.New("Argument not right, get a empty id or comment content")
	}
	o := orm.NewOrm()
	result, err := o.Raw(`INSERT INTO public.t_comment(userid, goodsid, content)VALUES (?, ?, ?)`, uid, gid, conetnt).Exec()
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
	//send a message to owner of goods
	if oid, err := GetOwnerId(gid); err != nil {
		msg := fmt.Sprintf(GoodsHanveBeenTalkTP, GetGNameById(gid), GetUNameById(uid))
		err = SendSystemMsg(oid, msg)
		if err != nil {
			mlog.Error("send goods talk msg to user: %v", err)
		}
	}
	Uas2.Add(uid)
	UpdateStaticIntData("TotalCommendNum", 1) //👀
	return nil
}

//send a system message to user 🍖
func SendSystemMsg(uid, msg string) error {
	o := orm.NewOrm()
	rawSeter := o.Raw(`INSERT INTO public.t_message(senderid, receiverid, content) VALUES (?, ?, ?)`, masterId, uid, msg)
	_, err := rawSeter.Exec()
	if err != nil {
		mlog.Error("%v", err)
		return err
	}
	return nil
}

//save a feedback record 🍗
func AddFeedback(d *FeedBackData) error {
	var err error
	if d == nil {
		err = errors.New("Receive a nil pointer")
		mlog.Error("%v", err)
		return nil
	}
	insertTP := `insert into t_feedback(user_id, fb_location, 
		fb_type, imgurl, describes, email)VALUES (?,?,?,?,?,?)`
	o := orm.NewOrm()
	if _, err = o.Raw(insertTP, d.Userid, d.Location, d.Type, d.Imgurl, d.Describes, d.Email).Exec(); err != nil {
		mlog.Error("%v", err)
		return err
	}
	UpdateStaticIntData("TotalFBTimes", 1) //👀
	return nil
}

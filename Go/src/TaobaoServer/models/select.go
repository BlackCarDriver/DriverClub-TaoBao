package models

import (
	"errors"
	"fmt"

	"github.com/astaxie/beego/logs"

	"github.com/astaxie/beego/orm"
)

var (
	NoResultErr = errors.New("No Result")
	MutiRowsErr = errors.New("More than noe rows were found")
)

//get the goods list that need to show in homepage, return the total rows databse have 🍇
//note: v_goodslist only select goods with state = 1;
func SelectHomePageGoods(gstype string, tag string, offset int, limit int, g *[]Goods1) (int, error) {
	var err error
	totalrows := 0
	o := orm.NewOrm()
	if gstype == "all" { // search all goods
		if _, err = o.Raw(`select * from v_hpgoodslist offset ? limit ?`, offset, limit).QueryRows(g); err != nil {
			mlog.Error("%v", err)
			return 0, err
		} else if err = o.Raw(`select count(*) from v_hpgoodslist`).QueryRow(&totalrows); err != nil {
			mlog.Error("Can't not count total rows number: %v", err)
		}
		return totalrows, err
	}
	if gstype == "like" { //search by input keyword
		tag = fmt.Sprintf("%%%s%%", tag)
		if _, err = o.Raw(`select distinct * from v_hpgoodslist where tag like ? or name like ? or title like ? offset ? limit ?`, tag, tag, tag, offset, limit).QueryRows(g); err != nil {
			mlog.Error("%v", err)
			return 0, err
		} else if err = o.Raw(`select distinct count(*) from v_hpgoodslist where tag like ? or name like ? or title like ?`, tag, tag, tag).QueryRow(&totalrows); err != nil {
			mlog.Error("Can't not count total rows number: %v", err)
		}
		return totalrows, err
	}
	if tag == "全部" { //search by given type
		if _, err = o.Raw(`select * from v_hpgoodslist where type=? offset ? limit ?`, gstype, offset, limit).QueryRows(g); err != nil {
			mlog.Error("%v", err)
			return 0, err
		} else if err = o.Raw(`select count(*) from v_hpgoodslist where type=?`, gstype).QueryRow(&totalrows); err != nil {
			mlog.Error("Can't not count total rows number: %v", err)
		}
		return totalrows, err
	}
	//search by given type and tag
	if _, err = o.Raw(`select * from v_hpgoodslist where type=? and tag=? offset ? limit ?`, gstype, tag, offset, limit).QueryRows(g); err != nil {
		mlog.Error("%v", err)
		return 0, err
	} else if o.Raw(`select count(*) from v_hpgoodslist where type=? and tag=?`, gstype, tag).QueryRow(&totalrows); err != nil {
		mlog.Error("Can't not count total rows number: %v", err)
	}
	return totalrows, err
}

//get a goods detail message
func GetGoodsById(gid string, c *GoodsDetail) error {
	if gid == "" {
		return errors.New("Receive a null gid")
	}
	o := orm.NewOrm()
	err := o.Raw(`select * from v_goods_detail where goodsid=$1`, gid).QueryRow(c)
	if err != nil {
		mlog.Error("%v", err)
		return err
	}
	return nil
}

//get the data need to show in personal page 🍊
func GetUserData(uid string, u *UserMessage) error {
	if uid == "" {
		return errors.New("Receive a null uid")
	}
	o := orm.NewOrm()
	err := o.Raw(`select * from v_mydata where id = ?`, uid).QueryRow(&u)
	if err != nil {
		mlog.Error("%v", err)
		return err
	}
	if u.Care, err = CountIcare(uid); err != nil {
		mlog.Error("%v", err)
		return err
	}
	if u.Becare, err = CountCareMe(uid); err != nil {
		mlog.Error("%v", err)
		return err
	}
	return nil
}

//get my messages 🍊 🍉🍏 🍑
func GetMyMessage(uid string, c *[]MyMessage, offset, limit int) error {
	if uid == "" {
		return errors.New("Receive a null uid")
	}
	if offset < 0 || limit <= 0 {
		mlog.Error("Unsuppose offset or limit argument, offset=%d, limit=%d", offset, limit)
		return errors.New("Unsuppose offset or limit argument")
	}
	o := orm.NewOrm()
	_, err := o.Raw(`select * from v_mymessage where uid = ? offset ? limit ?`, uid, offset, limit).QueryRows(c)
	if err != nil {
		mlog.Error("%v", err)
		return err
	}
	return nil
}

//get the goods list that i have collect 🍊 🍉
func GetMyCollectGoods(uid string, c *[]GoodsShort, offset, limit int) error {
	if uid == "" {
		return errors.New("Receive a null uid")
	}
	if offset < 0 || limit <= 0 {
		return errors.New("Unsuppose offset or limit argument")
	}
	o := orm.NewOrm()
	_, err := o.Raw(`select * from v_mycollect where uid = ? offset ? limit ?`, uid, offset, limit).QueryRows(c)
	if err != nil {
		mlog.Error("%v", err)
		return err
	}
	return nil
}

//get the goods list that i upload 🍉
//note: v_mygoods only return the goods with sate = 1;
func GetMyGoods(uid string, c *[]GoodsShort, offset, limit int) error {
	if uid == "" {
		return errors.New("Receive a null uid")
	}
	if offset < 0 || limit <= 0 {
		return errors.New("Unsuppose offset or limit argument")
	}
	o := orm.NewOrm()
	_, err := o.Raw(`select * from v_mygoods where uid = ? offset ? limit ?`, uid, offset, limit).QueryRows(c)
	if err != nil {
		mlog.Error("%v", err)
		return err
	}
	return nil
}

//get the list of user i care and which acre me 🍊
func GetCareMeData(uid string, c *[2][]UserShort) error {
	o := orm.NewOrm()
	//the list of user that care abtout me
	_, err := o.Raw(`select headimg, name, id from v_concern where id2=?`, uid).QueryRows(&c[0])
	if err != nil {
		mlog.Error("%v", err)
		return err
	}
	//get the list of user that i care
	_, err = o.Raw(`select headimg2 as headimg, id2 as  id, name2 as name  from v_concern where id = ?`, uid).QueryRows(&c[1])
	if err != nil {
		return err
	}
	return nil
}

//get the data that need to show in naving componment 🍓 🍞
func GetNavingMsg(uid string, c *MyStatus) error {
	o := orm.NewOrm()
	if err := o.Raw(`select * from v_navingmsg where id =?`, uid).QueryRow(&c); err != nil {
		return err
	}
	logs.Info(c)
	return nil
}

//get comment data of a goods
func GetGoodsComment(goodsid string, c *[]GoodsComment) error {
	o := orm.NewOrm()
	_, err := o.Raw(`select u.name as "username", c.time as "time", c.content as "comment" from t_user as u, t_comment as c where u.id=c.userid and c.goodsid=?`, goodsid).QueryRows(c)
	if err != nil {
		mlog.Error("%v", err)
		return err
	}
	return nil
}

//get statement of user-goods from database  🍌
//return result = like*1 + collect*2
func GetGoodsStatement(userid, goodid string) (int, error) {
	var result = 0
	var tmp = 0
	o := orm.NewOrm()
	//check if have collect
	if err := o.Raw(`SELECT count(*) FROM public.t_collect where userid=? and goodsid=?`, userid, goodid).QueryRow(&tmp); err != nil {
		mlog.Error("%v", err)
		return 0, err
	} else {
		result += tmp * 2
	}
	//check if have like
	if err := o.Raw(`SELECT count(*) FROM public.t_goods_like where userid=? and goodsid=?`, userid, goodid).QueryRow(&tmp); err != nil {
		mlog.Error("%v", err)
		return result, err
	} else {
		result += tmp
	}
	return result, nil
}

//get statement of user-user from database 🍉
//return result = like*1 + concern*2
func GetUserStatement(uid1, uid2 string) (int, error) {
	var result = 0
	var tmp = 0
	o := orm.NewOrm()
	//check if have concern
	if err := o.Raw(`SELECT count(*) FROM public.t_concern where id1=? and id2=?`, uid1, uid2).QueryRow(&tmp); err != nil {
		mlog.Error("%v", err)
		return 0, err
	} else {
		result += tmp * 2
	}
	//check if have like
	if err := o.Raw(`SELECT count(*) FROM public.t_user_like where userid1=? and userid2=?`, uid1, uid2).QueryRow(&tmp); err != nil {
		mlog.Error("%v", err)
		return result, err
	} else {
		result += tmp
	}
	return result, nil
}

//get user's message to that needed to display in changmsg page 🍏
func GetSettingMsg(uid string, c *UserSetData) error {
	var err error
	if uid == "" {
		err = errors.New("Receive a null uid")
		mlog.Error("%v", err)
	}
	o := orm.NewOrm()
	err = o.Raw(`select * from v_mydata where id = ?`, uid).QueryRow(&c)
	if err != nil {
		mlog.Error("%v", err)
		return err
	}
	return nil
}

//comfirm the login request and return the true id 🍓🍄
//note that the password is md5 encoding
func ComfirmLogin(identifi, password string) (id string, err error) {
	if identifi == "" || password == "" {
		err = errors.New("Receive null id or password")
		mlog.Error("%v", err)
		return "", err
	}
	count := 0
	o := orm.NewOrm()
	//use identifi as id firstly
	if err := o.Raw("select count(*) from t_user where id=? and password=?", identifi, password).QueryRow(&count); err != nil {
		mlog.Error("%v", err)
		return "", err
	} else {
		if count == 1 {
			return identifi, nil
		} else if count > 1 {
			err = errors.New("Find two account with same id!")
			mlog.Error("%v", err)
			return "", err
		}
	}
	//if no result with finding id, then use identfi as name and search again
	if err := o.Raw("select count(*) from t_user where name=? and password=?", identifi, password).QueryRow(&count); err != nil {
		mlog.Error("%v", err)
		return "", err
	} else {
		if count == 1 { //find true user id
			err = o.Raw("select id from t_user where name=? and password=?", identifi, password).QueryRow(&id)
			if err != nil {
				mlog.Error("%v", err)
				return "", err
			}
			return id, nil
		} else if count > 1 {
			err = errors.New("Find two account with same name!")
			mlog.Error("%v", err)
			return "", err
		} else if count == 0 {
			return "", NoResultErr
		}
	}
	return id, err
}

//#################### count ###########################

//get the user's number who car me
func CountCareMe(myid string) (int, error) {
	if myid == "" {
		return 0, errors.New("Receive a null myid")
	}
	o := orm.NewOrm()
	userNumber := 0
	err := o.Raw("select count(*) from t_concern where id2 = ?", myid).QueryRow(&userNumber)
	if err != nil {
		mlog.Error("%v", err)
		return 0, err
	}
	return userNumber, nil
}

//get the user's number i am cared 🍏
func CountIcare(myid string) (int, error) {
	o := orm.NewOrm()
	userNumber := 0
	err := o.Raw("select count(*) from t_concern where id1 = ?", myid).QueryRow(&userNumber)
	if err != nil {
		mlog.Error("%v", err)
		return 0, err
	}
	return userNumber, nil
}

//get the total number of user 🍏
func CountUser() int {
	o := orm.NewOrm()
	userNumber := 0
	err := o.Raw("select count(*) from t_user").QueryRow(&userNumber)
	if err != nil {
		mlog.Error("%v", err)
		return 0
	}
	return userNumber
}

//get the total number of total upload goods
func CountGoods() int {
	o := orm.NewOrm()
	goodsNumber := 0
	err := o.Raw("select count(*) from t_goods").QueryRow(&goodsNumber)
	if err != nil {
		mlog.Error("%v", err)
		return 0
	}
	return goodsNumber
}

//count how many goods a user have upload 🍉🍆
func CountMyCoods(uid string) int {
	o := orm.NewOrm()
	userNumber := 0
	err := o.Raw("select count(*) from v_mygoods where uid = ?", uid).QueryRow(&userNumber)
	if err != nil {
		mlog.Error("%v", err)
		return 0
	}
	return userNumber
}

//count how many goods a user have collect 🍉
func CountMyCollect(uid string) int {
	o := orm.NewOrm()
	userNumber := 0
	err := o.Raw("select count(*) from t_collect where userid = ?", uid).QueryRow(&userNumber)
	if err != nil {
		mlog.Error("%v", err)
		return 0
	}
	return userNumber
}

//count how many message a user have receive 🍉
func CountMyAllMsg(uid string) int {
	o := orm.NewOrm()
	userNumber := 0
	err := o.Raw("select count(*) from v_mymessage where uid = ?", uid).QueryRow(&userNumber)
	if err != nil {
		mlog.Error("%v", err)
		return 0
	}
	return userNumber
}

//count the unread number of user 🍞
func CountUnreadMsg(uid string) int {
	o := orm.NewOrm()
	userNumber := 0
	err := o.Raw("select count(*) from t_message where receiverid=? and state=0", uid).QueryRow(&userNumber)
	if err != nil {
		mlog.Error("CountUnreadMsg fail: %v", err)
		return 0
	}
	return userNumber
}

//count the specificed name numbers of other user  🍆
func CountUserName(name, id string) int {
	o := orm.NewOrm()
	userNumber := 0
	err := o.Raw("select count(*) from t_user where name=? and id != ?", name, id).QueryRow(&userNumber)
	if err != nil {
		mlog.Error("CountUserName fail: %v", err)
		return 0
	}
	return userNumber
}

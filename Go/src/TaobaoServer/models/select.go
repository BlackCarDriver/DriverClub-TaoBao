package models

import (
	tb "TaobaoServer/toolsbox"
	"errors"
	"fmt"

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

//get user name by userid 🍠
func GetUNameById(uid string) string {
	if uid == "" {
		return "unknow"
	}
	o := orm.NewOrm()
	username := ""
	err := o.Raw(`select name from t_user where id = ?`, uid).QueryRow(&username)
	if err != nil {
		mlog.Error("select name from t_user fail: %v", err)
		return "unknow"
	}
	return username
}

//find owner id by goods id 🍠
func GetOwnerId(gid string) (string, error) {
	if gid == "" {
		err := errors.New("Receive a null goods id")
		mlog.Error("%v", err)
		return "", err
	}
	o := orm.NewOrm()
	userid := ""
	err := o.Raw(`select userid from t_upload where goodsid = ?`, gid).QueryRow(&userid)
	if err != nil {
		mlog.Error("select userid from t_upload fail: %v", err)
		return "", err
	}
	mlog.Info("Userid: %s", userid)
	return userid, nil
}

//get goods name by goods id 🍠
func GetGNameById(gid string) string {
	if gid == "" {
		mlog.Error("Receive a null goods id")
		return ""
	}
	o := orm.NewOrm()
	name := ""
	err := o.Raw(`select name from t_goods where id = ?`, gid).QueryRow(&name)
	if err != nil {
		mlog.Error("select name from t_goods fail: %v", err)
		return "unknow"
	}
	return name
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
	return nil
}

//get comment data of a goods🍚
func GetGoodsComment(goodsid string, c *[]GoodsComment) error {
	o := orm.NewOrm()
	if goodsid == "" {
		return errors.New("Receive a null goodsid")
	}
	_, err := o.Raw(`select u.name as "username", 
	c.time as "time", c.content as "comment" from t_user as u,
	 t_comment as c where u.id=c.userid and c.goodsid=? order by time`, goodsid).QueryRows(c)
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

//comfirm the login request and return the true id 🍓🍄🍚
//identifi can be: id, username, email
//note that the password is md5 encoding
func ComfirmLogin(identifi, password string) (id string, err error) {
	if identifi == "" || password == "" {
		err = errors.New("Receive null id or password")
		mlog.Error("%v", err)
		return "", err
	}
	count := 0
	o := orm.NewOrm()
	//use use identifi as id firstly
	if tb.CheckUserID(identifi) {
		if err := o.Raw("select count(*) from t_user where id=? and password=?", identifi, password).QueryRow(&count); err != nil {
			mlog.Error("%v", err)
			return "", err
		} else {
			if count == 1 {
				return identifi, nil
			} else if count > 1 {
				err = errors.New("Find two user with same id!")
				mlog.Critical("%v", err)
				return "", err
			}
		}
	}
	//use user identifi as email
	if tb.CheckEmail(identifi) {
		if err := o.Raw("select count(*) from t_user where email=? and password=?", identifi, password).QueryRow(&count); err != nil {
			mlog.Error("%v", err)
			return "", err
		}
		if count == 1 { //find true user id
			err = o.Raw("select id from t_user where email=? and password=?", identifi, password).QueryRow(&id)
			if err != nil {
				mlog.Error("%v", err)
				return "", err
			}
			return id, nil
		} else if count > 1 {
			err = errors.New("Find two user with same email and passwod!")
			mlog.Critical("%v", err)
			return "", err
		}
	}
	//use identifi as user name in the last
	if tb.CheckUserName(identifi) {
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
				err = errors.New("Find two user with same name and password!")
				mlog.Critical("%v", err)
				return "", err
			}
			return "", NoResultErr
		}

	}
	return "", NoResultErr
}

//get feedback data, read at most 12 rows of feedback record from database  🍗
//note that the limit of rows number is setting in the sql command
func GetFeedBack(data *[]FeedBackData, offset int) error {
	if offset < 0 {
		err := fmt.Errorf("Offset smaller than 0!")
		mlog.Error("%v", err)
		return err
	}
	o := orm.NewOrm()
	selectTP := `SELECT id, 
	 user_id as "userid",
	 fb_location as location,
	 fb_type as "type", 
	 imgurl, describes, 
	 fb_time as "time", 
	 fb_status as status, 
	 email  FROM t_feedback order by fb_time desc limit 20 offset $1;`
	if _, err := o.Raw(selectTP, offset).QueryRows(data); err != nil {
		mlog.Error("%v", err)
		return err
	}
	return nil
}

//get the state of a goods and return some statement if the goods can't be read 🍜
func GetGoodsStat(gid string) string {
	o := orm.NewOrm()
	state := 0
	if err := o.Raw("select state from t_goods where id = ?", gid).QueryRow(&state); err != nil {
		mlog.Error("Search goods fail: %v", err)
		return "找不到此商品"
	}
	switch {
	case state > 0:
		return ""
	case state == -1:
		return "该商品已被用户下架"
	case state == -99:
		return "该商品已被管理员删除"
	case state < 0:
		return "该商品已被删除"
	}
	return ""
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
		mlog.Critical("%v", err)
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
		mlog.Critical("%v", err)
		return 0, err
	}
	return userNumber, nil
}

//get the total number of user 🍏
func CountTotalUser() int {
	o := orm.NewOrm()
	userNumber := 0
	err := o.Raw("select count(*) from t_user").QueryRow(&userNumber)
	if err != nil {
		mlog.Critical("%v", err)
		return 0
	}
	return userNumber
}

//get the total number of total upload goods (all state)
func CountGoods() int {
	o := orm.NewOrm()
	goodsNumber := 0
	err := o.Raw("select count(*) from t_goods").QueryRow(&goodsNumber)
	if err != nil {
		mlog.Critical("%v", err)
		return 0
	}
	return goodsNumber
}

//count numbers of goods which can be showed to user🍙
func CountOnlineGoods() int {
	o := orm.NewOrm()
	goodsNumber := 0
	err := o.Raw("select count(*) from t_goods where state >= 0").QueryRow(&goodsNumber)
	if err != nil {
		mlog.Critical("%v", err)
		return 0
	}
	return goodsNumber
}

//count numbers of goods tag (state >=0)🍙
func CountGoodsTag() int {
	o := orm.NewOrm()
	goodsNumber := 0
	err := o.Raw("select count(tag) from t_goods where state >=0").QueryRow(&goodsNumber)
	if err != nil {
		mlog.Critical("%v", err)
		return 0
	}
	return goodsNumber
}

//count total price of goods that are selling🍙
func CountTotalPrice() float64 {
	o := orm.NewOrm()
	var totalPrice float64 = 0.0
	err := o.Raw("select sum(price) from t_goods where state >= 0").QueryRow(&totalPrice)
	if err != nil {
		mlog.Critical("%v", err)
		return 0.0
	}
	return totalPrice
}

//count how many goods a user have upload 🍉🍆
func CountMyCoods(uid string) int {
	o := orm.NewOrm()
	userNumber := 0
	err := o.Raw("select count(*) from v_mygoods where uid = ?", uid).QueryRow(&userNumber)
	if err != nil {
		mlog.Critical("%v", err)
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
		mlog.Critical("%v", err)
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
		mlog.Critical("%v", err)
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
		mlog.Critical("CountUnreadMsg fail: %v", err)
		return 0
	}
	return userNumber
}

//count the specificed name numbers of all user  🍖
//return the numbers of user name or -1 which mean unexpect error happen
func CountUserName(name string) int {
	o := orm.NewOrm()
	userNumber := 0
	err := o.Raw("select count(*) from t_user where name=?", name).QueryRow(&userNumber)
	if err != nil {
		mlog.Critical("CountUserName fail: %v", err)
		return -1
	}
	return userNumber
}

//count the numbers of a name of ohter user name🍙
func CountOtherUserName(name, myid string) int {
	o := orm.NewOrm()
	if name == "" || myid == "" {
		return -1
	}
	userNumber := 0
	err := o.Raw("select count(*) from t_user where name=? and id !=?", name, myid).QueryRow(&userNumber)
	if err != nil {
		mlog.Critical("CountOtherUserName fail: %v", err)
		return -1
	}
	return userNumber
}

//count the numbers of using email that have been registe  🍚
func CountRegistEmail(email string) int {
	o := orm.NewOrm()
	number := 0
	err := o.Raw("select count(*) from t_user where email=?", email).QueryRow(&number)
	if err != nil {
		mlog.Critical("CountRegistEmail fail: %v", err)
		return -1
	}
	return number
}

//count id number to judge wheteher it id is exist 🍙
func CountUserId(id string) int {
	o := orm.NewOrm()
	number := 0
	err := o.Raw("select count(*) from t_user where id=?", id).QueryRow(&number)
	if err != nil {
		mlog.Critical("CountUserId fail: %v", err)
		return -1
	}
	return number
}

//#################### function relate to tempdata and maintain ####################

//get the list of top 10 user's rank, data include id, name and credits
func GetRankList(c *[]Rank) error {
	o := orm.NewOrm()
	if num, err := o.Raw(`select * from v_rank`).QueryRows(c); err != nil {
		mlog.Error("%v", err)
		return fmt.Errorf("Get user rank data error: %v", err)
	} else if num == 0 {
		err := fmt.Errorf("User rank result empty!")
		mlog.Error("%v", err)
		return err
	}
	return nil
}

//get all tag name and tag number of a type
func GetTagsData(gtype string, tag *[]GoodsSubType) error {
	if gtype == "" {
		return errors.New("Receive a null gtype")
	}
	o := orm.NewOrm()
	var tSubType []GoodsSubType
	num, err := o.Raw(`select tag, count(*) as number from t_goods where type = $1 and state >=0 group by tag`, gtype).QueryRows(&tSubType)
	if err != nil {
		mlog.Error("get type and tag data fail: %v", err)
		return err
	}
	if num == 0 {
		err = fmt.Errorf("the result of seaching type %s empty!", gtype)
		mlog.Error("%v", err)
		return err
	}
	var sum int64 = 0
	//add a new tag as "all"
	for i := 0; i < len(tSubType); i++ {
		sum += tSubType[i].Number
	}
	slice := make([]GoodsSubType, len(tSubType)+1)
	copy(slice, []GoodsSubType{{"全部", sum}})
	copy(slice[1:], tSubType)
	*tag = make([]GoodsSubType, len(tSubType)+1)
	copy(*tag, slice)
	return nil
}

package models

import (
	"errors"
	"fmt"

	"github.com/astaxie/beego/logs"

	"github.com/astaxie/beego/orm"
)

//get the goods list that need to show in homepage
func SelectHomePageGoods(gstype string, tag string, skip int, g *[]Goods1) error {
	var err error
	var num int64
	skip = (skip - 1) * 40
	o := orm.NewOrm()
	if gstype == "all" { //不筛选
		num, err = o.Raw(`select * from v_hpgoodslist offset ?`, skip).QueryRows(g)
		goto tail
	}
	if gstype == "like" { //模糊搜索
		tag = fmt.Sprintf("%%%s%%", tag)
		num, err = o.Raw(`select distinct * from v_hpgoodslist where tag like ? or name like ? or title like ?`, tag, tag, tag).QueryRows(g)
		goto tail
	}
	if tag == "全部" { //筛选类型
		num, err = o.Raw(`select * from v_hpgoodslist where type=? offset ?`, gstype, skip).QueryRows(g)
		goto tail
	}
	//筛选标签
	num, err = o.Raw(`select * from v_hpgoodslist where type=? and tag=? offset ?`, gstype, tag, skip).QueryRows(g)
tail:
	if err != nil {
		logs.Error(err)
		return err
	}
	if num == 0 {
		logs.Warn("Search a goods and no result, type: %s, tag: %s", gstype, tag)
	}
	return nil
}

//get all type name and tag
func GetTagsData(gtype string, tag *[]GoodsSubType) error {
	if gtype == "" {
		return errors.New("Receive a null gtype")
	}
	o := orm.NewOrm()
	var tSubType []GoodsSubType
	num, err := o.Raw(`select tag, count(*) as number from t_goods where type = $1 group by tag`, gtype).QueryRows(&tSubType)
	if err != nil {
		logs.Error(err)
		return err
	}
	if num == 0 {
		err = fmt.Errorf("the result is empty!")
		logs.Warn(err)
		return err
	}
	var sum int64 = 0
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

//get a goods detail message
func GetGoodsById(gid string, c *GoodsDetail) error {
	if gid == "" {
		return errors.New("Receive a null gid")
	}
	o := orm.NewOrm()
	err := o.Raw(`select * from v_goods_detail where goodsid=$1`, gid).QueryRow(c)
	if err != nil {
		logs.Error(err)
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
		logs.Error(err)
		return err
	}
	if u.Care, err = CountIcare(uid); err != nil {
		logs.Error(err)
		return err
	}
	if u.Becare, err = CountCareMe(uid); err != nil {
		logs.Error(err)
		return err
	}
	return nil
}

//get my messages 🍊 🍉
func GetMyMessage(uid string, c *[]MyMessage, offset, limit int) error {
	if uid == "" {
		return errors.New("Receive a null uid")
	}
	if offset < 0 || limit <= 0 {
		return errors.New("Unsuppose offset or limit argument")
	}
	o := orm.NewOrm()
	_, err := o.Raw(`select * from v_mymessage where id = ? offset ? limit ?`, uid, offset, limit).QueryRows(c)
	if err != nil {
		logs.Error(err)
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
		logs.Error(err)
		return err
	}
	return nil
}

//get the goods list that i upload 🍉
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
		logs.Error(err)
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
		logs.Error(err)
		return err
	}
	//get the list of user that i care
	_, err = o.Raw(`select headimg2 as headimg, id2 as  id, name2 as name  from v_concern where id = ?`, uid).QueryRows(&c[1])
	if err != nil {
		return err
	}
	return nil
}

//get the list of user's rank
func GetRankList(c *[]Rank) error {
	o := orm.NewOrm()
	num, err := o.Raw(`select * from v_rank`).QueryRows(c)
	if err != nil {
		logs.Error(err)
		return fmt.Errorf("GetMessage error: %v", err)
	}
	if num == 0 {
		err := fmt.Errorf("the result is empty!")
		logs.Error(err)
		return err
	}
	return nil
}

//get the data that need to show in naving componment
func GetNavingMsg(uid string, c *MyStatus) error {
	o := orm.NewOrm()
	err := o.Raw(`select * from v_navingmsg where id =?`, uid).QueryRow(&c)
	if err != nil {
		logs.Error(err)
		return err
	}
	return nil
}

//get comment data of a goods
func GetGoodsComment(goodsid string, c *[]GoodsComment) error {
	o := orm.NewOrm()
	_, err := o.Raw(`select u.name as "username", c.time as "time", c.content as "comment" from t_user as u, t_comment as c where u.id=c.userid and c.goodsid=?`, goodsid).QueryRows(c)
	if err != nil {
		logs.Error(err)
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
		logs.Error(err)
		return 0, err
	} else {
		result += tmp * 2
	}
	//check if have like
	if err := o.Raw(`SELECT count(*) FROM public.t_goods_like where userid=? and goodsid=?`, userid, goodid).QueryRow(&tmp); err != nil {
		logs.Error(err)
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
		logs.Error(err)
		return 0, err
	} else {
		result += tmp * 2
	}
	//check if have like
	if err := o.Raw(`SELECT count(*) FROM public.t_user_like where userid1=? and userid2=?`, uid1, uid2).QueryRow(&tmp); err != nil {
		logs.Error(err)
		return result, err
	} else {
		result += tmp
	}
	return result, nil
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
		logs.Error(err)
		return 0, err
	}
	return userNumber, nil
}

//get the user's number i am cared 🍊
func CountIcare(myid string) (int, error) {
	o := orm.NewOrm()
	userNumber := 0
	err := o.Raw("select count(*) from t_concern where id1 = ?", myid).QueryRow(&userNumber)
	if err != nil {
		logs.Error(err)
		return 0, err
	}
	return userNumber, nil
}

//get the total number of user 🍊
func CountUser() int {
	o := orm.NewOrm()
	userNumber := 0
	err := o.Raw("select count(*) from t_user").QueryRow(&userNumber)
	if err != nil {
		logs.Error(err)
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
		logs.Error(err)
		return 0
	}
	return goodsNumber
}

//count how many goods a user have upload 🍉
func CountMyCoods(uid string) int {
	o := orm.NewOrm()
	userNumber := 0
	err := o.Raw("select count(*) from t_upload where userid = ?", uid).QueryRow(&userNumber)
	if err != nil {
		logs.Error(err)
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
		logs.Error(err)
		return 0
	}
	return userNumber
}

//count how many message a user have receive 🍉
func CountMyAllMsg(uid string) int {
	o := orm.NewOrm()
	userNumber := 0
	err := o.Raw("select count(*) from t_message where senderid = ?", uid).QueryRow(&userNumber)
	if err != nil {
		logs.Error(err)
		return 0
	}
	return userNumber
}

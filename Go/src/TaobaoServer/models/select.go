package models

import (
	"fmt"

	"github.com/astaxie/beego/logs"

	"github.com/astaxie/beego/orm"
)

//获取主页商品的商品列表数据
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
		err = fmt.Errorf("the result is empty!")
		logs.Error(err)
		return err
	}
	return nil
}

//获取已有用户人数
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

//获取已有商品数
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

//获取某个类型的所有标签以及对应商品的数量
func GetTagsData(gtype string, tag *[]GoodsSubType) error {
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

//根据商品的id获得这个商品的属性信息
func GetGoodsById(gid string, c *GoodsDetail) error {
	o := orm.NewOrm()
	err := o.Raw(`select * from v_goods_detail where goodsid=$1`, gid).QueryRow(c)
	if err != nil {
		logs.Error(err)
		return err
	}
	fmt.Println(c.Time)
	return nil
}

//获取某个用户的展示数据
func GetUserData(uid string, u *UserMessage) error {
	o := orm.NewOrm()
	err := o.Raw(`select * from v_mydata where id = ?`, uid).QueryRow(&u)
	if err != nil {
		logs.Error(err)
		return err
	}
	return nil
}

//获取我的消息
func GetMyMessage(uid string, c *[]MyMessage) error {
	o := orm.NewOrm()
	num, err := o.Raw(`select * from v_mymessage where id = ?`, uid).QueryRows(c)
	if err != nil {
		logs.Error(err)
		return err
	}
	if num == 0 {
		err = fmt.Errorf("the result is empty!")
		logs.Error(err)
		return err
	}
	return nil
}

//获取我的收藏商品
func GetMyCollectGoods(uid string, c *[]GoodsShort) error {
	o := orm.NewOrm()
	num, err := o.Raw(`select * from v_mycollect where uid = ?`, uid).QueryRows(c)
	if err != nil {
		logs.Error(err)
		return err
	}
	if num == 0 {
		err := fmt.Errorf("the result is empty!")
		logs.Error(err)
		return err
	}
	return nil
}

//获取我上传的商品数据
func GetMyGoods(uid string, c *[]GoodsShort) error {
	o := orm.NewOrm()
	num, err := o.Raw(`select * from v_mygoods where uid = ?`, uid).QueryRows(c)
	if err != nil {
		logs.Error(err)
		return err
	}
	if num == 0 {
		err := fmt.Errorf("the result is empty!")
		logs.Error(err)
		return err
	}
	return nil
}

//获取我关注的和关注我的用户数据
func GetCareMeData(uid string, c *[2][]UserShort) error {
	o := orm.NewOrm()
	num, err := o.Raw(`select * from v_concern where myid=?`, uid).QueryRows(c[0])
	if err != nil {
		logs.Error(err)
		return err
	}
	if num == 0 {
		err := fmt.Errorf("the result of c[0] is empty!")
		logs.Error(err)
		return err
	}
	num, err = o.Raw(`select myid as id, name, headimg from v_iconcern where id = ?`, uid).QueryRows(c[1])
	if err != nil {
		return err
	}
	if num == 0 {
		err := fmt.Errorf("the result of c[1] is empty!")
		logs.Error(err)
		return err
	}
	return nil
}

//获取排名信息
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

//获取导航栏我的信息框数据
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
//like:1 collect:2
func GetStatement(userid, goodid string) (int, error) {
	var result = 0
	var tmp = 0
	o := orm.NewOrm()
	//check if have collect
	if _, err := o.Raw(`SELECT count(*) FROM public.t_collect where userid=? and goodsid=?`, userid, goodid).QueryRows(&tmp); err != nil {
		logs.Error(err)
		return 0, err
	} else {
		result += tmp * 2
	}
	//check if have like
	if _, err := o.Raw(`SELECT count(*) FROM public.t_goods_like where userid=? and goodsid=?`, userid, goodid).QueryRows(&tmp); err != nil {
		logs.Error(err)
		return result, err
	} else {
		result += tmp
	}
	return result, nil
}

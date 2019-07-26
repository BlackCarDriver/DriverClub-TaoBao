package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
)

//获取主页商品的商品列表数据(不筛选)
func SelectHomePageGoods(gstype string, tag string, skip int, g *[]Goods1) error {
	o := orm.NewOrm()
	num, err := o.Raw(`select uid as userid, uname as username, gid as Id, gname as Name, title,
	 	price, time, headimg from goods_list where state = 1`).QueryRows(g)
	if err != nil {
		return err
	}
	if num == 0 {
		return fmt.Errorf("the result is empty!")
	}
	return nil
}

//获取已有用户人数
func CountUser() int {
	o := orm.NewOrm()
	userNumber := 0
	err := o.Raw("select count(*) from t_user").QueryRow(&userNumber)
	if err != nil {
		fmt.Println("CountUser fall ! ", err)
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
		fmt.Println("CountGoods fall ! ", err)
		return 0
	}
	return goodsNumber
}

//获取某个类型的所有标签以及对应商品的数量
func GetTagsData(gtype string, tag *[]GoodsSubType) error {
	o := orm.NewOrm()
	num, err := o.Raw(`select tag, count(*) as number from t_goods where type = $1 group by tag`, gtype).QueryRows(tag)
	if err != nil {
		return err
	}
	if num == 0 {
		return fmt.Errorf("the result is empty!")
	}
	return nil
}

//根据商品的id获得这个商品的属性信息
func GetGoodsById(gid string, c *GoodsDetail) error {
	o := orm.NewOrm()
	err := o.Raw(`select * from v_goods_detail where goodsid=$1`, gid).QueryRow(c)
	if err != nil {
		return err
	}
	//还需要加入收藏数量和评论数量信息
	//还需要转换时间格式
	return nil
}

//获取某个用户的展示数据
func GetUserData(uid string, u *UserMessage) error {
	o := orm.NewOrm()
	err := o.Raw(`select * from t_user where id = $1`, uid).QueryRow(&u)
	if err != nil {
		fmt.Println("GetOtherUserData error: ", err)
		return err
	}
	fmt.Println(u)
	return nil
}

//获取我的消息
func GetMyMessage(uid string, c *[]MyMessage) error {
	o := orm.NewOrm()
	num, err := o.Raw(`select * from v_mymessage where id = ?`, uid).QueryRows(c)
	if err != nil {
		return fmt.Errorf("GetMessage error: %v", err)
	}
	if num == 0 {
		return fmt.Errorf("the result is empty!")
	}
	return nil
}

//获取我的收藏商品
func GetMyCollectGoods(uid string, c *[]GoodsShort) error {
	o := orm.NewOrm()
	num, err := o.Raw(`select * from v_mycollect where uid = ?`, uid).QueryRows(c)
	if err != nil {
		return err
	}
	if num == 0 {
		return fmt.Errorf("the result is empty!")
	}
	return nil
}

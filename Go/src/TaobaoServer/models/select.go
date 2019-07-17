package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
)

//获取主页商品的商品列表数据(不筛选)
func SelectHomePageGoods(gstype string, tag string, skip int, g *[]Goods1) error {
	o := orm.NewOrm()
	num, err := o.Raw(`select uname as Userid, gid as Id, gname as Name, title,
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

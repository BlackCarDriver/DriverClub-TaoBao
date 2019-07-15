package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
)

//获取主页商品的商品列表数据
func SelectHomePageGoods(gstype string, tag string, skip int, container []Goods1) error {
	fmt.Println(gstype, "  ", tag, "  ", skip, "  ", container)
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

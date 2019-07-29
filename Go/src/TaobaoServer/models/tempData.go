package models

import "fmt"

/*
tempDate.go保存一些全局的数据,避免每次前端请求都查询一次数据库
*/

//更新所有tempDate,
func initAllTempData() {
	UpdateGoodsTypeTempDate()
	UpdateUserRank()
}

//商品类型及标签的数据
var GoodsTypeTempDate = []GoodsType{
	{"学习用品", []GoodsSubType{}},
	{"体育用品", []GoodsSubType{}},
	{"生活用品", []GoodsSubType{}},
	{"电子产品", []GoodsSubType{}},
	{"手工diy", []GoodsSubType{}},
	{"虚拟商品", []GoodsSubType{}},
	{"其他", []GoodsSubType{}},
}

var UserRank []Rank

func UpdateGoodsTypeTempDate() {
	for i := 0; i < len(GoodsTypeTempDate); i++ {
		GetTagsData(GoodsTypeTempDate[i].Type, &GoodsTypeTempDate[i].List)
	}
}

func UpdateUserRank() {
	err := GetRankList(&UserRank)
	if err != nil {
		fmt.Println(err)
	}
}

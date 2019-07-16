package controllers

import (
	md "TaobaoServer/models"
	"encoding/json"
	"fmt"
)

//主页商品封面数据
func (this *HPGoodsController) Post() {
	PostBody := md.PostBody1{}
	var err error
	if err = json.Unmarshal(this.Ctx.Input.RequestBody, &PostBody); err != nil {
		return
	}
	//需要根据不同的类型，标签和页数从数据库中获取真实数据
	fmt.Println(PostBody.GoodsTag, "------------------", PostBody.GoodsIndex)
	this.Data["json"] = &md.MockGoodsData
	this.ServeJSON()
}

//商品详情获取数据接口
func (this *GoodsDetailController) Post() {
	postBody := md.GoodsPostBody{}
	var err error
	if err = json.Unmarshal(this.Ctx.Input.RequestBody, &postBody); err != nil {
		return
	}
	goodId := postBody.GoodId
	datatype := postBody.DataType
	if goodId == 0 || datatype == "" {
		return
	}
	fmt.Println("GoodsDetail postBody :", postBody)
	switch datatype {
	case "message":
		this.Data["json"] = &md.MockGoodsMessage
	case "detail":
		this.Data["json"] = &md.MockGoodsDetail
	}

	this.ServeJSON()
}

//上传商品
func (this *UploadGoodsController) Post() {
	var goodsdata md.UploadGoodsData
	var returnRes md.UpLoadResult
	var err error
	if err = json.Unmarshal(this.Ctx.Input.RequestBody, &goodsdata); err != nil {
		returnRes = md.CreateUploadRes(-100, err, "")
		goto tail
	}
	if err = md.CreateGoods(goodsdata); err != nil {
		returnRes = md.CreateUploadRes(-200, err, "")
		goto tail
	}
	returnRes = md.CreateUploadRes(0, nil, "")
tail:
	this.Data["json"] = &returnRes
	this.ServeJSON()
}

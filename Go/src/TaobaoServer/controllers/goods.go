package controllers

import (
	md "TaobaoServer/models"
	"encoding/json"
	"fmt"
)

//主页商品列表数据
func (this *HPGoodsController) Post() {
	PostBody := md.PostBody1{}
	var err error
	if err = json.Unmarshal(this.Ctx.Input.RequestBody, &PostBody); err != nil {
		return
	}
	var goodslist []md.Goods1
	fmt.Println(PostBody)
	err = md.SelectHomePageGoods(PostBody.GoodsType, PostBody.GoodsTag, PostBody.GoodsIndex, &goodslist)
	if err != nil {
		fmt.Println(err)
	}
	this.Data["json"] = goodslist

	this.ServeJSON()
}

//返回商品分类和标签列表数据
func (this *GoodsTypeController) Get() {
	this.Data["json"] = &md.GoodsTypeTempDate
	this.ServeJSON()
}

//商品详情获取数据接口
func (this *GoodsDetailController) Post() {
	postBody := md.GoodsPostBody{}
	var err error
	if err = json.Unmarshal(this.Ctx.Input.RequestBody, &postBody); err != nil {
		fmt.Println("error : ", err)
		return
	}
	goodId := postBody.GoodId
	datatype := postBody.DataType
	if goodId == "" || datatype == "" {
		//请求不规范，这里应该返回一个错误页面或重定向
		this.Data["json"] = &md.MockGoodsMessage
		goto tail
	}
	err = md.UpdateGoodsVisit(goodId)
	if err != nil {
		fmt.Println("update visit of goods fall!! ", err)
	}
	switch datatype {
	case "goodsmessage": //base message
		var gooddata md.GoodsDetail
		err = md.GetGoodsById(goodId, &gooddata)
		if err == nil {
			this.Data["json"] = &gooddata
			goto tail
		}
	case "goodscomment": //goods comment
		var comment []md.GoodsComment
		err = md.GetGoodsComment(goodId, &comment)
		fmt.Println(comment)
		if err == nil {
			this.Data["json"] = comment
			goto tail
		}
	}

	//请求不规范或id找不到，应该或返回错误重定向
	this.Data["json"] = "empty"
tail:
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

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
	fmt.Println(postBody)
	if goodId == "" || datatype == "" {
		//请求不规范，这里应该返回一个错误页面或重定向
		this.Data["json"] = &md.MockGoodsMessage
		goto tail
	}
	switch datatype {
	case "goodsmessage":
		var gooddata md.GoodsDetail
		err = md.GetGoodsById(goodId, &gooddata)
		if err == nil {
			this.Data["json"] = &gooddata
			fmt.Println(gooddata.Userid)
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

//更新各种信息的接口,如商品点赞数，私信，商品收藏
func (this *UpdateController) Post() {
	postBody := md.UpdatePostBody{}
	result := md.UpdateResult{}
	var err error
	if err = json.Unmarshal(this.Ctx.Input.RequestBody, &postBody); err != nil {
		fmt.Println("error : ", err)
		return
	}
	tag := postBody.Tag
	switch tag {
	case "likegoods":
		err = md.UpdateGoodsLike(postBody.TargetId)
		if err == nil {
			result.Status = 0
			goto tail
		}
		result.Status = -1
		result.Describe = fmt.Sprintf("点赞失败： %s ", err)
	case "sendmessage":
		err = md.AddUserMessage(postBody.UserId, postBody.TargetId, postBody.StrData)
		if err == nil {
			result.Status = 0
			goto tail
		}
		result.Status = -1
		result.Describe = fmt.Sprintf("发送失败: %s", err)
	case "addcollect":
		err = md.AddGoodsCollect(postBody.UserId, postBody.TargetId)
		if err == nil {
			result.Status = 0
			goto tail
		}
		result.Status = -1
		result.Describe = fmt.Sprintf("收藏失败: %s", err)
	default:
		fmt.Println(tag)
	}
tail:
	this.Data["json"] = result
	this.ServeJSON()
}

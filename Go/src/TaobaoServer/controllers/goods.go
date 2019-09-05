package controllers

import (
	md "TaobaoServer/models"
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego/logs"
)

//主页商品列表数据
func (this *HPGoodsController) Post() {
	PostBody := md.PostBody1{}
	var err error
	if err = json.Unmarshal(this.Ctx.Input.RequestBody, &PostBody); err != nil {
		logs.Error(err)
		return
	}
	var goodslist []md.Goods1
	err = md.SelectHomePageGoods(PostBody.GoodsType, PostBody.GoodsTag, PostBody.GoodsIndex, &goodslist)
	if err != nil {
		logs.Error(err)
	}
	this.Data["json"] = goodslist

	this.ServeJSON()
}

//返回商品分类和标签列表数据
func (this *GoodsTypeController) Get() {
	this.Data["json"] = &md.GoodsTypeTempDate
	this.ServeJSON()
}

//get all kind of data in goodspage  🍌🔥
//response for GetGoodsDeta() in fontend
func (this *GoodsDetailController) Post() {
	postBody := md.RequestProto{}
	response := md.ReplyProto{}
	response.StatusCode = 0
	var err error
	var api, goodId, userid string
	//parse request protocol
	if err = json.Unmarshal(this.Ctx.Input.RequestBody, &postBody); err != nil {
		response.StatusCode = -1
		response.Msg = fmt.Sprintf("Can not parse postbody: %v", err)
		logs.Error(response.Msg)
		goto tail
	}
	api = postBody.Api
	goodId = postBody.TargetId
	userid = postBody.UserId
	//check that the data is complete
	if api == "" || goodId == "" {
		response.StatusCode = -2
		response.Msg = fmt.Sprintf("Can't get api or goodsid from request data")
		logs.Error(response.Msg)
		goto tail
	}
	//update some statistical
	if err = md.UpdateGoodsVisit(goodId); err != nil {
		logs.Error(err)
	}
	//handle the request
	switch api {
	case "goodsmessage": // base message in goodsdetail page
		var gooddata md.GoodsDetail
		if err := md.GetGoodsById(goodId, &gooddata); err != nil {
			response.StatusCode = -3
			response.Msg = fmt.Sprintf("Get goodsmessage fail: %v", err)
			logs.Error(response.Msg)
		} else {
			response.Data = &gooddata
		}
		goto tail

	case "goodscomment": //comment or discuss date in goodsdetail page
		var comment []md.GoodsComment
		if err := md.GetGoodsComment(goodId, &comment); err != nil {
			response.StatusCode = -4
			response.Msg = fmt.Sprintf("Get goods %s 's comment fail: %v", goodId, err)
			logs.Error(response.Msg)
		} else {
			response.Data = comment
		}
		goto tail

	case "usergoodsstate": //user state for specified goods in goodsdetail page
		tmp := md.UserGoodsState{Like: false, Collect: false}
		if userid == "" { // if user havn't login then return default date
			response.Data = tmp
		} else if res, err := md.GetStatement(userid, goodId); err != nil {
			response.StatusCode = -5
			response.Msg = fmt.Sprintf("Getstatement fail: %v", err)
			logs.Error(response.Msg)
		} else {
			if res&1 != 0 {
				tmp.Like = true
			}
			if res&2 != 0 {
				tmp.Collect = true
			}
			response.Data = tmp
		}
		goto tail

	default:
		response.StatusCode = -10
		response.Msg = "No such method"
		logs.Warn(response.Msg)
		goto tail
	}
tail:
	this.Data["json"] = response
	this.ServeJSON()
}

//上传商品
func (this *UploadGoodsController) Post() {
	var goodsdata md.UploadGoodsData
	var returnRes md.UpLoadResult
	var err error
	if err = json.Unmarshal(this.Ctx.Input.RequestBody, &goodsdata); err != nil {
		logs.Error(err)
		returnRes = md.CreateUploadRes(-100, err, "")
		goto tail
	}
	if err = md.CreateGoods(goodsdata); err != nil {
		returnRes = md.CreateUploadRes(-200, err, "")
		logs.Error(err)
		goto tail
	}
	returnRes = md.CreateUploadRes(0, nil, "")
tail:
	this.Data["json"] = &returnRes
	this.ServeJSON()
}

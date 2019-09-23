package controllers

import (
	md "TaobaoServer/models"
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego/logs"
)

//return homepage goods list data 🍋🔥🍇
func (this *HPGoodsController) Post() {
	postBody := md.RequestProto{}
	response := md.ReplyProto{}
	response.StatusCode = 0
	var goodslist []md.Goods1
	var err error
	var goodstype, goodstag string
	var appendData map[string]interface{}
	//parse postbody
	if err = json.Unmarshal(this.Ctx.Input.RequestBody, &postBody); err != nil {
		response.StatusCode = -1
		response.Msg = fmt.Sprintf("Can not parse postbody: %v", err)
		logs.Error(response.Msg)
		goto tail
	}
	//check argument
	if postBody.Limit <= 0 || postBody.Offset < 0 {
		response.StatusCode = -2
		response.Msg = fmt.Sprintf("Unexpect limited or offset")
		logs.Error(response.Msg)
		goto tail
	}
	//get and chekc additional argument
	appendData = postBody.Data.(map[string]interface{})
	goodstype = appendData["goodstype"].(string)
	goodstag = appendData["goodstag"].(string)
	if goodstype == "" || goodstag == "" {
		response.StatusCode = -2
		response.Msg = fmt.Sprintf("Unexpect argument")
		logs.Error(response.Msg)
		goto tail
	}
	//get data from database
	if sum, err := md.SelectHomePageGoods(goodstype, goodstag, postBody.Offset, postBody.Limit, &goodslist); err != nil {
		response.StatusCode = -1
		response.Msg = fmt.Sprintf("Get goods list data fail: %v", err)
		logs.Error(response.Msg)
		goto tail
	} else {
		response.Data = goodslist
		response.Rows = len(goodslist)
		response.Sum = sum
	}
tail:
	this.Data["json"] = response
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
	//catch the unexpect panic
	defer func() {
		if err, ok := recover().(error); ok {
			response.StatusCode = -99
			response.Msg = fmt.Sprintf("Unexpect error happen, api: %s , error: %v", api, err)
			logs.Error(response.Msg)
			this.Data["json"] = response
			this.ServeJSON()
		}
	}()
	//handle the request
	switch api {
	case "goodsmessage": // base message in goodsdetail page 🍄
		var gooddata md.GoodsDetail
		if err := md.GetGoodsById(goodId, &gooddata); err != nil {
			response.StatusCode = -3
			response.Msg = fmt.Sprintf("Get goodsmessage fail: %v", err)
			logs.Error(response.Msg)
		} else {
			response.Data = &gooddata
		}
		//update some statistical
		if err = md.UpdateGoodsVisit(goodId); err != nil {
			logs.Error(err)
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
		} else if res, err := md.GetGoodsStatement(userid, goodId); err != nil {
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

//user upload a goods 🍋🔥
func (this *UploadGoodsController) Post() {
	postBody := md.RequestProto{}
	response := md.ReplyProto{}
	var goodsdata md.UploadGoodsData
	var err error
	//parse request protocol
	if err = json.Unmarshal(this.Ctx.Input.RequestBody, &postBody); err != nil {
		response.StatusCode = -1
		response.Msg = fmt.Sprintf("Can not parse postbody: %v", err)
		logs.Error(response.Msg)
		goto tail
	}
	//catch the unexpect panic
	defer func() {
		if err, ok := recover().(error); ok {
			response.StatusCode = -99
			response.Msg = fmt.Sprintf("Unexpect error happen, error: %v", err)
			logs.Error(response.Msg)
			this.Data["json"] = response
			this.ServeJSON()
		}
	}()
	//parse postBody.Data
	if err := json.Unmarshal([]byte(postBody.Data.(string)), &goodsdata); err != nil {
		response.StatusCode = -2
		response.Msg = fmt.Sprintf("Marshal fail: %v", err)
		logs.Error(response.Msg)
		goto tail
	}
	//save to database
	if err = md.CreateGoods(goodsdata); err != nil {
		response.StatusCode = -3
		response.Msg = fmt.Sprintf("Save goods fail: %v", err)
		logs.Error(response.Msg)
		goto tail
	}
	response.StatusCode = 0
	response.Msg = "Success!"
tail:
	this.Data["json"] = response
	this.ServeJSON()
}

//note that it is GET method
//return goods type list and tag list 🍋🔥
func (this *GoodsTypeController) Get() {
	this.Data["json"] = &md.GoodsTypeTempDate
	this.ServeJSON()
}

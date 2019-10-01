package controllers

import (
	md "TaobaoServer/models"
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego/logs"
)

//return homepage goods list data 🍋🍇🌽
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
		rlog.Error(response.Msg)
		goto tail
	}
	//check argument
	if postBody.Limit <= 0 || postBody.Offset < 0 {
		response.StatusCode = -2
		response.Msg = fmt.Sprintf("Unexpect limited or offset")
		rlog.Error(response.Msg)
		goto tail
	}
	//get and chekc additional argument
	appendData = postBody.Data.(map[string]interface{})
	goodstype = appendData["goodstype"].(string)
	goodstag = appendData["goodstag"].(string)
	if goodstype == "" || goodstag == "" {
		response.StatusCode = -2
		response.Msg = fmt.Sprintf("Unexpect argument")
		rlog.Error(response.Msg)
		goto tail
	}
	//get data from cache
	if cache, err := md.GetCache(&postBody); err == nil {
		if err = json.Unmarshal([]byte(cache), &response); err != nil {
			rlog.Error("Unmarshal cache %s fail:%v", postBody.CacheKey, err)
		} else {
			logs.Info("Get cache %s success! ", postBody.CacheKey)
			goto tail
		}
	} else {
		logs.Error(err)
	}
	//get data from database
	if sum, err := md.SelectHomePageGoods(goodstype, goodstag, postBody.Offset, postBody.Limit, &goodslist); err != nil {
		response.StatusCode = -1
		response.Msg = fmt.Sprintf("Get goods list data fail: %v", err)
		rlog.Error(response.Msg)
	} else {
		response.Data = goodslist
		response.Rows = len(goodslist)
		response.Sum = sum
	}
	//save response to cache
	md.SetCache(&postBody, &response)

tail:
	this.Data["json"] = response
	this.ServeJSON()
}

//get all kind of data in goodspage  🍌🌽
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
		rlog.Error(response.Msg)
		goto tail
	}
	api = postBody.Api
	goodId = postBody.TargetId
	userid = postBody.UserId
	//check that the data is complete
	if api == "" || goodId == "" {
		response.StatusCode = -2
		response.Msg = fmt.Sprintf("Can't get api or goodsid from request data")
		rlog.Error(response.Msg)
		goto tail
	}
	//get data from cache
	if cache, err := md.GetCache(&postBody); err == nil {
		if err = json.Unmarshal([]byte(cache), &response); err != nil {
			rlog.Error("Unmarshal cache %s fail:%v", postBody.CacheKey, err)
		} else {
			logs.Info("Get cache %s success! ", postBody.CacheKey)
			goto tail
		}
	} else {
		logs.Error(err)
	}
	//catch the unexpect panic
	defer func() {
		if err, ok := recover().(error); ok {
			response.StatusCode = -99
			response.Msg = fmt.Sprintf("Unexpect error happen, api: %s , error: %v", api, err)
			rlog.Error(response.Msg)
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
			rlog.Error(response.Msg)
		} else {
			response.Data = &gooddata
		}
		//update some statistical
		if err = md.UpdateGoodsVisit(goodId); err != nil {
			rlog.Error("%v", err)
		}
		md.Uas1.Add(userid) //user see other goods

	case "goodscomment": //comment or discuss date in goodsdetail page
		var comment []md.GoodsComment
		if err := md.GetGoodsComment(goodId, &comment); err != nil {
			response.StatusCode = -4
			response.Msg = fmt.Sprintf("Get goods %s 's comment fail: %v", goodId, err)
			rlog.Error(response.Msg)
		} else {
			response.Data = comment
		}

	case "usergoodsstate": //user state for specified goods in goodsdetail page
		tmp := md.UserGoodsState{Like: false, Collect: false}
		if userid == "" { // if user havn't login then return default date
			response.Data = tmp
		} else if res, err := md.GetGoodsStatement(userid, goodId); err != nil {
			response.StatusCode = -5
			response.Msg = fmt.Sprintf("Getstatement fail: %v", err)
			rlog.Error(response.Msg)
		} else {
			if res&1 != 0 {
				tmp.Like = true
			}
			if res&2 != 0 {
				tmp.Collect = true
			}
			response.Data = tmp
		}

	default:
		response.StatusCode = -10
		response.Msg = "No such method"
		rlog.Warn(response.Msg)
		goto tail
	}
	//save response to cache
	md.SetCache(&postBody, &response)

tail:
	this.Data["json"] = response
	this.ServeJSON()
}

//user upload a goods 🍋🍔
func (this *UploadGoodsController) Post() {
	postBody := md.RequestProto{}
	response := md.ReplyProto{}
	var goodsdata md.UploadGoodsData
	var err error
	var token string
	//parse request protocol
	if err = json.Unmarshal(this.Ctx.Input.RequestBody, &postBody); err != nil {
		response.StatusCode = -1
		response.Msg = fmt.Sprintf("Can not parse postbody: %v", err)
		rlog.Error(response.Msg)
		goto tail
	}
	//catch the unexpect panic
	defer func() {
		if err, ok := recover().(error); ok {
			response.StatusCode = -99
			response.Msg = fmt.Sprintf("Unexpect error happen, error: %v", err)
			rlog.Error(response.Msg)
			this.Data["json"] = response
			this.ServeJSON()
		}
	}()
	//parse postBody.Data
	if err := json.Unmarshal([]byte(postBody.Data.(string)), &goodsdata); err != nil {
		response.StatusCode = -2
		response.Msg = fmt.Sprintf("Marshal fail: %v", err)
		rlog.Error(response.Msg)
		goto tail
	}
	//check token
	if token == "" || !CheckToken(goodsdata.UserId, token) {
		rlog.Warn("User %s request upload goods with worng token", goodsdata.UserId)
		response.StatusCode = -1000
		response.Msg = "Token错误或过期,请重新登录！"
		goto tail
	}
	//save to database
	if err = md.CreateGoods(goodsdata); err != nil {
		response.StatusCode = -3
		response.Msg = fmt.Sprintf("Save goods fail: %v", err)
		rlog.Error(response.Msg)
		goto tail
	}
	response.StatusCode = 0
	response.Msg = "Success!"
	md.Uas2.Add(goodsdata.UserId) //upload goods scuess, credits+1
tail:
	this.Data["json"] = response
	this.ServeJSON()
}

//note that it is GET method
//return goods type list and tag list 🍋
func (this *GoodsTypeController) Get() {
	this.Data["json"] = &md.GoodsTypeTempDate
	this.ServeJSON()
}

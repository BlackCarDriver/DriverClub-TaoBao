package controllers

import (
	md "TaobaoServer/models"
	tb "TaobaoServer/toolsbox"
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego/logs"
)

//return homepage goods list data 🍋🍇🌽🍙
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
	//get and check additional argument
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
			goto tail
		}
	}
	//get data from database
	if sum, err := md.SelectHomePageGoods(goodstype, goodstag, postBody.Offset, postBody.Limit, &goodslist); err != nil {
		response.StatusCode = -1
		response.Msg = fmt.Sprintf("Get goods list data fail: %v", err)
		rlog.Error(response.Msg)
	} else {
		response.Data = goodslist
		response.Rows = len(goodslist)
		if response.Rows == 0 {
			logs.Info("No goods result found: type:%s  tag:%s", goodstype, goodstag)
		}
		response.Sum = sum
	}
	//save response to cache
	md.SetCache(&postBody, &response)
tail:
	//update static data 👀
	md.TodayRequestTimes++
	md.TodayVStimes++

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
			goto tail
		}
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
		//check the statement firstly🍜
		statement := md.GetGoodsStat(goodId)
		if statement != "" {
			response.StatusCode = -3
			response.Msg = statement
			goto tail
		}
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

	case "goodscomment": //comment or discuss date in goodsdetail page 🍚
		if goodId == "" {
			response.StatusCode = -4
			response.Msg = "获取商品ID失败"
			rlog.Error(response.Msg)
			goto tail
		}
		var comment []md.GoodsComment
		if err := md.GetGoodsComment(goodId, &comment); err != nil {
			response.StatusCode = -4
			response.Msg = fmt.Sprintf("获取商品数据失败：id '%s'  原因：' %v'", goodId, err)
			rlog.Error(response.Msg)
			goto tail
		}
		response.Data = comment
		goto tail

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
	//update static data 👀
	md.TodayRequestTimes++
	this.Data["json"] = response
	this.ServeJSON()
}

//user upload a goods in upload page 🍋🍔🍚
func (this *UploadGoodsController) Post() {
	postBody := md.RequestProto{}
	response := md.ReplyProto{}
	var goodsdata md.UploadGoodsData
	var err error
	var token, reason string
	//parse request protocol
	if err = json.Unmarshal(this.Ctx.Input.RequestBody, &postBody); err != nil {
		response.StatusCode = -1
		response.Msg = fmt.Sprintf("解析请求主体失败: %v", err)
		rlog.Error(response.Msg)
		goto tail
	}
	//catch the unexpect panic
	defer func() {
		if err, ok := recover().(error); ok {
			response.StatusCode = -99
			response.Msg = fmt.Sprintf("Unexpect error happen, error: %v", err)
			rlog.Critical("%s", response.Msg)
			this.Data["json"] = response
			this.ServeJSON()
		}
	}()
	//parse postBody.Data
	if err := json.Unmarshal([]byte(postBody.Data.(string)), &goodsdata); err != nil {
		response.StatusCode = -2
		response.Msg = fmt.Sprintf("解析主体数据失败: %v", err)
		rlog.Error(response.Msg)
		goto tail
	}
	//check token
	token = postBody.Token
	if token == "" || !CheckToken(goodsdata.UserId, token) {
		rlog.Warn("User %s request upload goods with worng token", goodsdata.UserId)
		response.StatusCode = -1000
		response.Msg = "Token 错误或过期,请重新登录！"
		goto tail
	}
	//check goodsdata format
	switch {
	case goodsdata.UserId == "":
		reason = "无效的用户ID"
	case goodsdata.Imgurl == "":
		reason = "无效的图片连接"
	case !tb.CheckGoodsName(goodsdata.Name):
		reason = "商品名称不通过"
	case len(goodsdata.Text) < 5, len(goodsdata.Text) >= 45:
		reason = "商品标题不通过"
	case goodsdata.Price < 0 || goodsdata.Price > 10000:
		reason = "商品专让价不通过"
	case goodsdata.Type == "", goodsdata.Tag == "":
		reason = "无法获取分类或标签数据"
	case len(goodsdata.Text) > 500<<10:
		reason = "商品描述的长度超过了500kb"
	}
	if reason != "" {
		response.StatusCode = -3
		response.Msg = reason
		rlog.Error(response.Msg)
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
	md.Uas2.Add(goodsdata.UserId)
tail:
	//update static data 👀
	md.TodayRequestTimes++
	this.Data["json"] = response
	this.ServeJSON()
}

//note that it is GET method
//return goods type list and tag list 🍋
func (this *GoodsTypeController) Get() {
	//update static data 👀
	md.TodayRequestTimes++
	this.Data["json"] = &md.GoodsTypeTempDate
	this.ServeJSON()
}

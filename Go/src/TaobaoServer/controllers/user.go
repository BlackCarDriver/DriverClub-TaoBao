package controllers

import (
	md "TaobaoServer/models"
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego/logs"
)

//个人详情页面或其他用户主页信息获取接口
func (this *PersonalDataController) Post() {
	postBody := md.PersonalPostBody{}
	var err error
	if err = json.Unmarshal(this.Ctx.Input.RequestBody, &postBody); err != nil {
		logs.Error(err)
		return
	}
	userName := postBody.Name
	dataTag := postBody.Tag
	if userName == "" || dataTag == "" { //userName 就是 id
		return
	}
	switch dataTag {
	case "mymsg": //我的数据
		var data md.UserMessage
		err = md.GetUserData(userName, &data)
		if err != nil {
			logs.Error(err)
			this.Data["json"] = ""
			goto tail
		}
		this.Data["json"] = data

	case "mygoods": //我的商品
		this.Data["json"] = &md.MockGoodsShort
		var data []md.GoodsShort
		err = md.GetMyGoods(userName, &data)
		if err != nil {
			logs.Error(err)
			this.Data["json"] = ""
			goto tail
		}
		this.Data["json"] = data

	case "mycollect": //我的收藏
		var data []md.GoodsShort
		err = md.GetMyCollectGoods(userName, &data)
		if err != nil {
			logs.Error(err)
			this.Data["json"] = "do something..."
			goto tail
		}
		this.Data["json"] = &data

	case "message": //我的消息
		var data []md.MyMessage
		err = md.GetMyMessage(userName, &data)
		if err != nil {
			//do something...
			this.Data["json"] = ""
			goto tail
		}
		this.Data["json"] = data

	case "rank": //用户排名数据
		this.Data["json"] = &md.UserRank

	case "mycare": //关注我的和我关注的用户数据
		var data [2][]md.UserShort
		err = md.GetCareMeData(userName, &data)
		if err != nil {
			logs.Error(err)
			this.Data["json"] = ""
			goto tail
		}
		this.Data["json"] = data

	case "naving": //导航栏我的数据
		var data md.MyStatus
		err = md.GetNavingMsg(userName, &data)
		if err != nil {
			logs.Error(err)
			this.Data["json"] = ""
			goto tail
		}
		this.Data["json"] = data

	case "othermsg": //看其他人的数据
		var data md.UserMessage
		err = md.GetUserData(userName, &data)
		if err != nil {
			goto tail
		}
		err = md.UpdateUserVisit(userName)
		if err != nil {
			fmt.Println("update user visit fall, ", err)
		} else {
			fmt.Println(userName)
		}
		this.Data["json"] = data
		goto tail

	case "setdata":
		this.Data["json"] = &md.MockUserSetData

	}
tail:
	this.ServeJSON()
}

//update personal message, such as base information, connect wags. 🍍
//server for UpdateMessage() in frontend
func (this *UpdataMsgController) Post() {
	postBody := md.RequestProto{}
	response := md.ReplyProto{}
	response.StatusCode = 0
	var api, userid string
	var err error
	//parse request protocol
	if err = json.Unmarshal(this.Ctx.Input.RequestBody, &postBody); err != nil {
		response.StatusCode = -1
		response.Msg = fmt.Sprintf("Can not parse postbody: %v", err)
		logs.Error(response.Msg)
		goto tail
	}
	api = postBody.Api
	userid = postBody.UserId
	//check that the data is complete
	if api == "" || userid == "" {
		response.StatusCode = -2
		response.Msg = fmt.Sprintf("Can't get api or userid from request data")
		logs.Error(response.Msg)
		goto tail
	}
	logs.Info(userid)
	//handle the request
	switch api {
	case "MyBaseMessage": //base information of users
		postData := md.UpdeteMsg{}
		if err = Parse(postBody.Data, &postData); err != nil {
			response.StatusCode = -3
			response.Msg = fmt.Sprintf("Can't parse postbody data: %v", err)
			logs.Error(response.Msg)
			goto tail
		}
		if err = md.UpdateUserBaseMsg(postData); err != nil {
			response.StatusCode = -4
			response.Msg = fmt.Sprintf("Update message fail: %v", err)
			logs.Error(response.Msg)
		}
		goto tail
	case "MyConnectMessage": //connect information
		postData := md.UpdeteMsg{}
		if err = Parse(postBody.Data, &postData); err != nil {
			response.StatusCode = -5
			response.Msg = fmt.Sprintf("Can't parse postbody data: %v", err)
			logs.Error(response.Msg)
			goto tail
		}
		if err = md.UpdateUserConnectMsg(postData); err != nil {
			response.StatusCode = -6
			response.Msg = fmt.Sprintf("Update connection message fail: %v", err)
			logs.Error(response.Msg)
		}
		goto tail
	case "MyHeadImage": //update profile picture
		if postBody.Data.(string) == "" { //here imgurl save in data directrly
			response.StatusCode = -9
			response.Msg = fmt.Sprintf("Can't get imagurl from postbody")
			logs.Error(response.Msg)
			goto tail
		}
		if err = md.UpdateUserHeadIMg(userid, postBody.Data.(string)); err != nil {
			response.StatusCode = -8
			response.Msg = fmt.Sprintf("Update profile iamge fail: %v", err)
			logs.Error(response.Msg)
		}
		goto tail
	default:
		response.StatusCode = -100
		response.Msg = fmt.Sprintf("No such method: %s", api)
		logs.Error(response.Msg)
	}
tail:
	this.Data["json"] = response
	this.ServeJSON()
}

//登录，注册，更换验证码， 获取验证码
func (this *EntranceController) Post() {
	postBody := md.EntranceBody{}
	var err error
	if err = json.Unmarshal(this.Ctx.Input.RequestBody, &postBody); err != nil {
		return
	}
	updateTag := postBody.Tag
	switch updateTag {
	case "login":
		fmt.Println("user login ...")
	case "CheckRegister":
		tregister := md.RegisterData{}
		Parse(postBody.Data, &tregister)
		err := md.CreateAccount(tregister)
		if err != nil {
			logs.Error(err)
		}
	case "confirmcode":
		fmt.Println("confirmcode...")
	}
	this.Data["json"] = &md.MockRequireResult
	this.ServeJSON()
}

package controllers

import (
	md "TaobaoServer/models"
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego/logs"
)

//get myprofile data or other user profile data 🍋 🔥
//server for GetMyMsg() from frontend
func (this *PersonalDataController) Post() {
	postBody := md.RequestProto{}
	response := md.ReplyProto{}
	response.StatusCode = 0
	var err error
	var api, targetid string
	//parse request protocol
	if err = json.Unmarshal(this.Ctx.Input.RequestBody, &postBody); err != nil {
		response.StatusCode = -1
		response.Msg = fmt.Sprintf("Can not parse postbody: %v", err)
		logs.Error(response.Msg)
		goto tail
	}
	api = postBody.Api
	targetid = postBody.TargetId
	//check that the data is complete
	if api == "" || targetid == "" {
		response.StatusCode = -2
		response.Msg = fmt.Sprintf("Can't get api or targetid from request data")
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
	// logs.Info(api, "\t\t", targetid)
	//handle the request
	switch api {
	case "mymsg": //my profile data
		var data md.UserMessage
		if err = md.GetUserData(targetid, &data); err != nil {
			response.StatusCode = -3
			response.Msg = fmt.Sprintf("Get user data fail: %v", err)
			logs.Error(response.Msg)
		} else {
			response.Data = data
		}
		goto tail

	case "mygoods": //my goods data 🍉
		var data []md.GoodsShort
		if err = md.GetMyGoods(targetid, &data, postBody.Offset, postBody.Limit); err != nil {
			response.StatusCode = -4
			response.Msg = fmt.Sprintf("Can't get goods data: %v ", err)
			logs.Error(response.Msg)
		} else {
			response.Data = data
			response.Rows = len(data)
			response.Sum = md.CountMyCoods(targetid)
		}
		goto tail

	case "mycollect": //my collect goods data 🍉
		var data []md.GoodsShort
		if err = md.GetMyCollectGoods(targetid, &data, postBody.Offset, postBody.Limit); err != nil {
			response.StatusCode = -5
			response.Msg = fmt.Sprintf("Can't get collect data: %v ", err)
			logs.Error(response.Msg)
		} else {
			response.Data = data
			response.Rows = len(data)
			response.Sum = md.CountMyCollect(targetid)
		}
		goto tail

	case "message": //my receive messages 🍉🍏
		var data []md.MyMessage
		if err = md.GetMyMessage(targetid, &data, postBody.Offset, postBody.Limit); err != nil {
			response.StatusCode = -9
			response.Msg = fmt.Sprintf("Can't get message data: %v ", err)
			logs.Error(response.Msg)
		} else {
			response.Data = data
			response.Rows = len(data)
			response.Sum = md.CountMyAllMsg(targetid)
		}
		goto tail

	case "mycare": //get my favorite and who care me
		var data [2][]md.UserShort
		if err = md.GetCareMeData(targetid, &data); err != nil {
			response.StatusCode = -6
			response.Msg = fmt.Sprintf("Can't get care data: %v ", err)
			logs.Error(response.Msg)
		} else {
			response.Data = data
		}
		goto tail

	case "naving": //get naving data
		var data md.MyStatus
		if err = md.GetNavingMsg(targetid, &data); err != nil {
			response.StatusCode = -7
			response.Msg = fmt.Sprintf("Can't get naving data: %v ", err)
			logs.Error(response.Msg)
		} else {
			response.Data = data
		}
		goto tail

	case "othermsg": //other people profile data
		var data md.UserMessage
		if err = md.GetUserData(targetid, &data); err != nil {
			response.StatusCode = -8
			response.Msg = fmt.Sprintf("Can't get user data: %v ", err)
			logs.Error(response.Msg)
			goto tail
		} else {
			response.Data = data
		}
		if err = md.UpdateUserVisit(targetid); err != nil {
			logs.Error("Update visit number fail: %v", err)
		}
		goto tail

	case "getuserstatement": //the statement of user to user 🍉
		tmp := md.UserState{Like: false, Concern: false}
		if postBody.UserId == "" { // if user havn't login then return default date
			response.Data = tmp
		} else if res, err := md.GetUserStatement(postBody.UserId, targetid); err != nil {
			response.StatusCode = -9
			response.Msg = fmt.Sprintf("UserGoodsState fail: %v", err)
			logs.Error(response.Msg)
		} else {
			if res&1 != 0 {
				tmp.Like = true
			}
			if res&2 != 0 {
				tmp.Concern = true
			}
			response.Data = tmp
		}
		goto tail

	case "rank": //user rank
		response.Data = md.UserRank
		//TODO: make a function 🍉
		goto tail

	case "settingdata": //user message in the changemsg page 🍏
		data := md.UserSetData{}
		if err = md.GetSettingMsg(targetid, &data); err != nil {
			response.StatusCode = -10
			response.Msg = fmt.Sprintf("获取数据失败：%v", err)
			logs.Error(response.Msg)
		} else {
			response.Data = data
		}
		goto tail

	default:
		response.StatusCode = -100
		response.Msg = fmt.Sprintf("Unsupose metho: %s", api)
		logs.Error(response.Msg)
		goto tail
	}

tail:
	this.Data["json"] = response
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
	//catch the unexpect panic
	defer func() {
		if err, ok := recover().(error); ok {
			response.StatusCode = -99
			response.Msg = fmt.Sprintf("Unexpect error, api: %s , des: %v", api, err)
			logs.Error(response.Msg)
			this.Data["json"] = response
			this.ServeJSON()
		}
	}()
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
		logs.Info(postData)
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

//login, regeist, comfirm code, change password... 🍏🍓🍄
func (this *EntranceController) Post() {
	postBody := md.RequestProto{}
	response := md.ReplyProto{}
	var api, targetid string
	var err error
	//parse request protocol
	if err = json.Unmarshal(this.Ctx.Input.RequestBody, &postBody); err != nil {
		response.StatusCode = -1
		response.Msg = fmt.Sprintf("Can not parse postbody: %v", err)
		logs.Error(response.Msg)
		goto tail
	}
	//check that the data is complete
	api = postBody.Api
	targetid = postBody.TargetId
	if api == "" || targetid == "" {
		response.StatusCode = -2
		response.Msg = fmt.Sprintf("Can't get api or id from request body")
		logs.Error(response.Msg)
		goto tail
	}
	switch api {
	case "login": //login, note that the target id can be true id or name 🍓🍄
		//TODO: check to format of password
		password := MD5Parse(postBody.Data.(string))
		logs.Info(password)
		//TODO: check if the identifi is student number and vertify it

		//check the account from database and get true id
		tid, err := md.ComfirmLogin(targetid, password)
		if err != nil {
			logs.Error(err)
			response.StatusCode = -3
			if err == md.NoResultErr {
				response.Msg = "没有此账号或密码错误"
			} else {
				response.Msg = fmt.Sprint(err)
				logs.Error(response.Msg)
			}
			goto tail
		}
		//if the password is passed than return user data
		var data md.MyStatus
		if err = md.GetNavingMsg(tid, &data); err != nil {
			response.StatusCode = -4
			response.Msg = fmt.Sprintf("Can't get usre message: %v ", err)
			logs.Error(response.Msg)
		} else {
			data.ID = tid
			response.Data = data
		}
		goto tail

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
	response.StatusCode = 0
	response.Msg = "Success"
tail:
	this.Data["json"] = response
	this.ServeJSON()
}

package controllers

import (
	md "TaobaoServer/models"
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego/logs"
)

//get myprofile data or other user profile data 🍋🔥🌽
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
		rlog.Error(response.Msg)
		goto tail
	}
	api = postBody.Api
	targetid = postBody.TargetId
	//check that the data is complete
	if api == "" || targetid == "" {
		response.StatusCode = -2
		response.Msg = fmt.Sprintf("Can't get api or targetid from request data")
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
		logs.Info(err)
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
	// rlog.Info(api, "\t\t", targetid)
	//handle the request
	switch api {
	case "mymsg": //my profile data
		var data md.UserMessage
		if err = md.GetUserData(targetid, &data); err != nil {
			response.StatusCode = -3
			response.Msg = fmt.Sprintf("Get user data fail: %v", err)
			rlog.Error(response.Msg)
		} else {
			response.Data = data
		}

	case "mygoods": //my goods data 🍉
		var data []md.GoodsShort
		if err = md.GetMyGoods(targetid, &data, postBody.Offset, postBody.Limit); err != nil {
			response.StatusCode = -4
			response.Msg = fmt.Sprintf("Can't get goods data: %v ", err)
			rlog.Error(response.Msg)
		} else {
			response.Data = data
			response.Rows = len(data)
			response.Sum = md.CountMyCoods(targetid)
		}
		md.Uas1.Add(targetid) //user see himself personal page, active+1

	case "mycollect": //my collect goods data 🍉
		var data []md.GoodsShort
		if err = md.GetMyCollectGoods(targetid, &data, postBody.Offset, postBody.Limit); err != nil {
			response.StatusCode = -5
			response.Msg = fmt.Sprintf("Can't get collect data: %v ", err)
			rlog.Error(response.Msg)
		} else {
			response.Data = data
			response.Rows = len(data)
			response.Sum = md.CountMyCollect(targetid)
		}

	case "message": //my receive messages 🍉🍏 🍞
		var data []md.MyMessage
		if err = md.GetMyMessage(targetid, &data, postBody.Offset, postBody.Limit); err != nil {
			response.StatusCode = -9
			response.Msg = fmt.Sprintf("Can't get message data: %v ", err)
			rlog.Error(response.Msg)
		} else {
			response.Data = data
			response.Rows = len(data)
			response.Sum = md.CountMyAllMsg(targetid)
		}

	case "mycare": //get my favorite and who care me
		var data [2][]md.UserShort
		if err = md.GetCareMeData(targetid, &data); err != nil {
			response.StatusCode = -6
			response.Msg = fmt.Sprintf("Can't get care data: %v ", err)
			rlog.Error(response.Msg)
		} else {
			response.Data = data
		}

	case "naving": //get naving data 🍔
		var data md.MyStatus
		userid := targetid
		//check user token
		if postBody.Token == "" {
			rlog.Error("User %s request naving with null token", userid)
			response.StatusCode = -1000
			response.Msg = "获取Token失败,请重新登录"
			goto tail
		} else if !CheckToken(userid, postBody.Token) {
			rlog.Warn("User %s request naving with worng token", userid)
			response.StatusCode = -1000
			response.Msg = "Token错误或过期,请重新登录！"
			goto tail
		}
		//get and return naving data
		if err = md.GetNavingMsg(userid, &data); err != nil {
			response.StatusCode = -7
			response.Msg = fmt.Sprintf("Can't get naving data: %v ", err)
			rlog.Error(response.Msg)
		} else {
			response.Data = data
		}

	case "othermsg": //other people profile data
		var data md.UserMessage
		if err = md.GetUserData(targetid, &data); err != nil {
			response.StatusCode = -8
			response.Msg = fmt.Sprintf("Can't get user data: %v ", err)
			rlog.Error(response.Msg)
			goto tail
		} else {
			response.Data = data
		}
		if err = md.UpdateUserVisit(targetid); err != nil {
			rlog.Error("Update visit number fail: %v", err)
		}

	case "getuserstatement": //the statement of user to user 🍉
		tmp := md.UserState{Like: false, Concern: false}
		if postBody.UserId == "" { // if user havn't login then return default date
			response.Data = tmp
		} else if res, err := md.GetUserStatement(postBody.UserId, targetid); err != nil {
			response.StatusCode = -9
			response.Msg = fmt.Sprintf("UserGoodsState fail: %v", err)
			rlog.Error(response.Msg)
		} else {
			if res&1 != 0 {
				tmp.Like = true
			}
			if res&2 != 0 {
				tmp.Concern = true
			}
			response.Data = tmp
		}

	case "rank": //user rank
		response.Data = md.UserRank
		//TODO: make a function 🍉

	case "settingdata": //user message in the changemsg page 🍏
		data := md.UserSetData{}
		if err = md.GetSettingMsg(targetid, &data); err != nil {
			response.StatusCode = -10
			response.Msg = fmt.Sprintf("获取数据失败：%v", err)
			rlog.Error(response.Msg)
		} else {
			response.Data = data
		}
	default:
		response.StatusCode = -100
		response.Msg = fmt.Sprintf("Unsupose metho: %s", api)
		rlog.Error(response.Msg)
		goto tail
	}
	//save response to cache
	md.SetCache(&postBody, &response)
tail:
	this.Data["json"] = response
	this.ServeJSON()
}

//update personal message, such as base information, connect wags. 🍍🍔
//server for UpdateMessage() in frontend
//all function here need to vertiry with token
func (this *UpdataMsgController) Post() {
	postBody := md.RequestProto{}
	response := md.ReplyProto{}
	response.StatusCode = 0
	var api, userid, token string
	var err error
	//parse request protocol
	if err = json.Unmarshal(this.Ctx.Input.RequestBody, &postBody); err != nil {
		response.StatusCode = -1
		response.Msg = fmt.Sprintf("Can not parse postbody: %v", err)
		rlog.Error(response.Msg)
		goto tail
	}
	api = postBody.Api
	userid = postBody.UserId
	token = postBody.Token
	//check that the data is complete
	if api == "" || userid == "" {
		response.StatusCode = -2
		response.Msg = fmt.Sprintf("Can't get api or userid from request data")
		rlog.Error(response.Msg)
		goto tail
	}
	//check token
	if token == "" || !CheckToken(userid, token) {
		rlog.Warn("User %s request update %s with worng token", userid, api)
		response.StatusCode = -1000
		response.Msg = "Token错误或过期,请重新登录！"
		goto tail
	}
	//catch the unexpect panic
	defer func() {
		if err, ok := recover().(error); ok {
			response.StatusCode = -99
			response.Msg = fmt.Sprintf("Unexpect error, api: %s , des: %v", api, err)
			rlog.Error(response.Msg)
			this.Data["json"] = response
			this.ServeJSON()
		}
	}()
	//handle the request 🍆
	switch api {
	case "changemybasemsg": //base information of users
		postData := md.UpdeteMsg{}
		if err = Parse(postBody.Data, &postData); err != nil {
			response.StatusCode = -3
			response.Msg = fmt.Sprintf("Can't parse postbody data: %v", err)
			rlog.Error(response.Msg)
			goto tail
		}
		if err = md.UpdateUserBaseMsg(postData); err != nil {
			response.StatusCode = -4
			response.Msg = fmt.Sprintf("Update message fail: %v", err)
			rlog.Error(response.Msg)
		}
		goto tail
	case "MyConnectMessage": //connect information
		postData := md.UpdeteMsg{}
		if err = Parse(postBody.Data, &postData); err != nil {
			response.StatusCode = -5
			response.Msg = fmt.Sprintf("Can't parse postbody data: %v", err)
			rlog.Error(response.Msg)
			goto tail
		}
		if err = md.UpdateUserConnectMsg(postData); err != nil {
			response.StatusCode = -6
			response.Msg = fmt.Sprintf("Update connection message fail: %v", err)
			rlog.Error(response.Msg)
		}
		goto tail
	case "MyHeadImage": //update profile picture
		if postBody.Data.(string) == "" { //here imgurl save in data directrly
			response.StatusCode = -9
			response.Msg = fmt.Sprintf("Can't get imagurl from postbody")
			rlog.Error(response.Msg)
			goto tail
		}
		if err = md.UpdateUserHeadIMg(userid, postBody.Data.(string)); err != nil {
			response.StatusCode = -8
			response.Msg = fmt.Sprintf("Update profile iamge fail, error:%v, uid:%s", err, userid)
			rlog.Error(response.Msg)
		}
		goto tail
	default:
		response.StatusCode = -100
		response.Msg = fmt.Sprintf("No such method: %s", api)
		rlog.Error(response.Msg)
	}
tail:
	this.Data["json"] = response
	this.ServeJSON()
}

//login, regeist, comfirm code, change password... 🍏🍓🍄🍖
func (this *EntranceController) Post() {
	postBody := md.RequestProto{}
	response := md.ReplyProto{}
	var api, targetid string
	var err error
	//parse request protocol
	if err = json.Unmarshal(this.Ctx.Input.RequestBody, &postBody); err != nil {
		response.StatusCode = -1
		response.Msg = fmt.Sprintf("Can not parse postbody: %v", err)
		rlog.Error(response.Msg)
		goto tail
	}
	//check that the data is complete
	api = postBody.Api
	targetid = postBody.TargetId
	if api == "" || targetid == "" {
		response.StatusCode = -2
		response.Msg = fmt.Sprintf("Can't get api or id from request body")
		rlog.Error(response.Msg)
		goto tail
	}
	switch api {
	case "login": //login, note that the target id can be true id or name 🍓🍄🍔
		//TODO: check to format of password
		password := MD5Parse(postBody.Data.(string))
		//TODO: check if the identifi is student number and vertify it

		//check the account from database and get true id
		tid, err := md.ComfirmLogin(targetid, password)
		if err != nil {
			rlog.Warn("%v", err)
			response.StatusCode = -3
			if err == md.NoResultErr {
				response.Msg = "没有此账号或密码错误"
			} else {
				response.Msg = fmt.Sprint(err)
				rlog.Error(response.Msg)
			}
			goto tail
		}
		//if the password is passed than return user data
		var data md.MyStatus
		if err = md.GetNavingMsg(tid, &data); err != nil {
			response.StatusCode = -4
			response.Msg = fmt.Sprintf("Can't get usre message: %v ", err)
			rlog.Error(response.Msg)
		} else {
			data.ID = tid
			response.Data = data
			if token := CreateToken(tid); token == "" {
				response.StatusCode = -5
				response.Msg = "Sorry, Create token fail!"
				rlog.Error(response.Msg)
			} else {
				response.Msg = token
			}
		}
		goto tail

	case "getcomfirmcode": //comfrim the signup message form user and return a comfrim code 🍖
		register := md.RegisterData{}
		if err = Parse(postBody.Data, &register); err != nil {
			response.StatusCode = -5
			response.Msg = fmt.Sprintf("Parse postbody fail: %v", err)
			goto tail
		}
		logs.Info(register)
		//TODO:check the postdata message
		if md.CountUserName(register.Name) != 0 {
			response.StatusCode = -6
			response.Msg = "It name have been used, please try another one"
			goto tail
		}
		code := GetRandomCode()
		logs.Info(code)
		register.Code = code
		if err = SendConfrimEmail(register, md.CountUser()); err != nil {
			response.StatusCode = -7
			response.Msg = fmt.Sprintf("Send Email fail: %v", err)
			logs.Error(response.Msg)
			goto tail
		}
		//save the code into timer map
		keyData := fmt.Sprintf("%v", register)
		logs.Info(keyData)
		if err = md.ComfirmCode.Add(keyData); err != nil {
			response.StatusCode = -8
			response.Msg = fmt.Sprintf("Save comfirm code fail: %v", err)
			logs.Error(response.Msg)
			goto tail
		}

	case "comfirmAndRegisit": //vertify the comfirm code and create a new account if pass  🍖
		register := md.RegisterData{}
		if err = Parse(postBody.Data, &register); err != nil {
			response.StatusCode = -9
			response.Msg = fmt.Sprintf("Parse postbody fail: %v", err)
			goto tail
		}
		//TODO:check the postdata message
		//check the comfirm code
		keyData := fmt.Sprintf("%v", register)
		logs.Info(keyData)
		if err = md.ComfirmCode.Get(keyData); err != nil {
			rlog.Warn("%v", err)
			response.StatusCode = -10
			response.Msg = fmt.Sprintf("验证失败：%v", err)
			goto tail
		}
		//create a new account
		register.Password = MD5Parse(register.Password)
		if err = md.CreateAccount(register); err != nil {
			rlog.Error("%v", err)
			response.StatusCode = -11
			response.Msg = fmt.Sprintf("：( 创建账号失败：%v, 请稍后重试", err)
			goto tail
		}
		rlog.Info("New account have been create! %s", register.Name)
	}
	response.StatusCode = 0
	response.Msg = "Success"
tail:
	this.Data["json"] = response
	this.ServeJSON()
}

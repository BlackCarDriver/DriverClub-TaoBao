package controllers

import (
	md "TaobaoServer/models"
	tb "TaobaoServer/toolsbox"
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
	case "mymsg": //my user profile data, targetid mean userid 🍚
		var data md.UserMessage
		if err = md.GetUserData(targetid, &data); err != nil {
			response.StatusCode = -3
			response.Msg = fmt.Sprintf("获取数据失败: %v", err)
			rlog.Error(response.Msg)
			goto tail
		} else {
			response.Data = data
		}

	case "mygoods": //my user's goods data 🍉🍚
		var data []md.GoodsShort
		if postBody.Offset < 0 || postBody.Limit <= 0 {
			response.StatusCode = -4
			response.Msg = "非法的 offset 值或 limit 值"
			rlog.Error(response.Msg)
			goto tail
		}
		if err = md.GetMyGoods(targetid, &data, postBody.Offset, postBody.Limit); err != nil {
			response.StatusCode = -4
			response.Msg = fmt.Sprintf("获取商品列表失败: %v ", err)
			rlog.Error(response.Msg)
			goto tail
		}
		response.Data = data
		response.Rows = len(data)
		response.Sum = md.CountMyCoods(targetid)

		md.Uas1.Add(targetid)

	case "mycollect": //my collect goods data 🍉🍚
		if postBody.Offset < 0 || postBody.Limit <= 0 {
			response.StatusCode = -5
			response.Msg = "非法的 offset 值或 limit 值"
			rlog.Error(response.Msg)
			goto tail
		}
		var data []md.GoodsShort
		if err = md.GetMyCollectGoods(targetid, &data, postBody.Offset, postBody.Limit); err != nil {
			response.StatusCode = -5
			response.Msg = fmt.Sprintf("获取收藏列表失败: %v ", err)
			rlog.Error(response.Msg)
			goto tail
		}
		response.Data = data
		response.Rows = len(data)
		response.Sum = md.CountMyCollect(targetid)

	case "message": //my receive messages 🍉🍏 🍞🍚
		if postBody.Offset < 0 || postBody.Limit <= 0 {
			response.StatusCode = -9
			response.Msg = "非法的 offset 值或 limit 值"
			rlog.Error(response.Msg)
			goto tail
		}
		var data []md.MyMessage
		if err = md.GetMyMessage(targetid, &data, postBody.Offset, postBody.Limit); err != nil {
			response.StatusCode = -9
			response.Msg = fmt.Sprintf("获取消息列表失败: %v ", err)
			rlog.Error(response.Msg)
			goto tail
		}
		response.Data = data
		response.Rows = len(data)
		response.Sum = md.CountMyAllMsg(targetid)

	case "mycare": //get my care list and who care me
		var data [2][]md.UserShort
		if err = md.GetCareMeData(targetid, &data); err != nil {
			response.StatusCode = -6
			response.Msg = fmt.Sprintf("获取关注列表失败: %v ", err)
			rlog.Error(response.Msg)
			goto tail
		}
		response.Data = data

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

	case "othermsg": //other people profile data  🍙
		var data md.UserMessage
		if err = md.GetUserData(targetid, &data); err != nil {
			response.StatusCode = -8
			response.Msg = fmt.Sprintf("获取用户数据失败: %v ", err)
			rlog.Error(response.Msg)
			goto tail
		}
		if err = md.UpdateUserVisit(targetid); err != nil {
			rlog.Error("Update visit number fail: %v", err)
		}
		response.Data = data

	case "getuserstatement": //the statement of user to user 🍉🍙
		tmp := md.UserState{Like: false, Concern: false}
		if postBody.UserId == "" { // if user havn't login then return default date
			response.Data = tmp
			goto tail
		}
		if res, err := md.GetUserStatement(postBody.UserId, targetid); err != nil {
			response.StatusCode = -9
			response.Msg = fmt.Sprintf("获取用户状态失败: %v", err)
			rlog.Error(response.Msg)
			goto tail
		} else {
			if res&1 != 0 {
				tmp.Like = true
			}
			if res&2 != 0 {
				tmp.Concern = true
			}
		}
		response.Data = tmp

	case "rank": //user rank message
		response.Data = md.UserRank

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
	//update static data 👀
	md.TodayRequestTimes++
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
	case "changemybasemsg": //base information of users 🍙
		postData := md.UpdeteMsg{}
		var reason string
		if err = Parse(postBody.Data, &postData); err != nil {
			response.StatusCode = -3
			response.Msg = fmt.Sprintf("解析请求体失败: %v", err)
			rlog.Error(response.Msg)
			goto tail
		}
		reason = ""
		switch {
		case !tb.CheckUserName(postData.Name):
			reason = "用户名称不合规则"
		case postData.Sex != "GIRL" && postData.Sex != "BOY":
			reason = "性别信息不合规则"
		case len(postData.Sign) > 50:
			reason = "前面长度超出限制"
		case !tb.CheckGrade(postData.Grade):
			reason = "年级信息不合规则"
		case len(postData.Colleage) > 50:
			reason = "学院信息不合规则"
		case len(postData.Major) > 50:
			reason = "专业信息不合规则"
		case len(postData.Dorm) > 50:
			reason = "宿舍楼栋信息不合规则"
		}
		if reason != "" {
			response.StatusCode = -4
			response.Msg = reason
			goto tail
		}
		if err = md.UpdateUserBaseMsg(postData, userid); err != nil {
			response.StatusCode = -4
			response.Msg = fmt.Sprintf("Update message fail: %v", err)
			rlog.Error(response.Msg)
		}
		goto tail
	case "MyConnectMessage": //connect information update🍙
		reason := ""
		postData := md.UpdeteMsg{}
		if err = Parse(postBody.Data, &postData); err != nil {
			response.StatusCode = -5
			response.Msg = fmt.Sprintf("无法解析请求体数据: %v", err)
			rlog.Error(response.Msg)
			goto tail
		}
		switch {
		case !tb.CheckEmail(postData.Emails):
			reason = "邮箱名称不合规则"
		case len(postData.Qq) > 20:
			reason = "qq信息不合规则"
		case len(postData.Phone) > 20:
			reason = "电话号码信息不合规则"
		}
		if reason != "" {
			response.StatusCode = -6
			response.Msg = reason
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
	//update static data 👀
	md.TodayRequestTimes++
	this.Data["json"] = response
	this.ServeJSON()
}

//login, regeist, comfirm code, change password... 🍏🍓🍄🍖🍚
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
	switch api {
	case "login": //note that the target id can be true id or name or email🍓🍄🍔🍚
		//user name format checking
		if !tb.CheckUserName(targetid) && !tb.CheckEmail(targetid) {
			rlog.Warn("User name '%s' unpass checking in login", targetid)
			response.StatusCode = -3
			response.Msg = "用户名、ID、或邮箱地址格式不对"
			goto tail
		}
		//password format checking
		password := postBody.Data.(string)
		if !tb.CheckPassword(password) {
			rlog.Warn("User %s password '%s' unpass checking in login", targetid, password)
			response.StatusCode = -3
			response.Msg = "密码格式错误"
			goto tail
		}
		//check the account from database and get true id
		password = MD5Parse(postBody.Data.(string))
		tid, err := md.ComfirmLogin(targetid, password)
		if err != nil {
			rlog.Warn("ComfirmLogin fail: %v", err)
			response.StatusCode = -3
			if err == md.NoResultErr {
				response.Msg = "没有此账号或密码错误"
			} else {
				response.Msg = fmt.Sprint(err)
				rlog.Error(response.Msg)
			}
			goto tail
		}
		//if the password is passed than response data is user naving data
		var data md.MyStatus
		if err = md.GetNavingMsg(tid, &data); err != nil {
			response.StatusCode = -4
			response.Msg = fmt.Sprintf("Can't get usre naving data: %v ", err)
			rlog.Error(response.Msg)
		} else {
			//create a token for user next times identify, token will place at response msg
			data.ID = tid
			response.Data = data
			if token := CreateToken(tid); token == "" {
				response.StatusCode = -5
				response.Msg = "Sorry, Create token fail!"
				rlog.Critical(response.Msg)
			} else {
				response.Msg = token
			}
		}
		goto tail

	case "getcomfirmcode": //comfrim sign up data and return a comfrim code when user regiest  🍖🍚
		register := md.RegisterData{}
		if err = Parse(postBody.Data, &register); err != nil {
			response.StatusCode = -5
			response.Msg = fmt.Sprintf("解析请求主体失败: %v", err)
			goto tail
		}
		//check username, password and email format
		if !tb.CheckUserName(register.Name) {
			response.StatusCode = -5
			response.Msg = fmt.Sprintf("用户名格式错误")
			goto tail
		}
		if !tb.CheckEmail(register.Email) {
			response.StatusCode = -5
			response.Msg = fmt.Sprintf("邮箱地址格式错误")
			goto tail
		}
		if !tb.CheckPassword(register.Password) {
			response.StatusCode = -5
			response.Msg = fmt.Sprintf("密码格式错误")
			goto tail
		}
		//check the if the user name and email have been used
		if md.CountUserName(register.Name) != 0 {
			response.StatusCode = -6
			response.Msg = "这个用户名正在被使用，请更换一个哦"
			goto tail
		}
		if md.CountRegistEmail(register.Email) != 0 {
			response.StatusCode = -6
			response.Msg = "这个邮箱地址已经被注册，请更换一个或重置密码"
			goto tail
		}
		//send a comfirm code to user's email
		code := GetRandomCode()
		logs.Debug(code)
		register.Code = code
		if err = SendConfrimEmail(register, md.CountTotalUser()); err != nil {
			response.StatusCode = -7
			response.Msg = fmt.Sprintf("发送邮件失败:'%v' ,请稍后重试", err)
			logs.Critical(response.Msg)
			goto tail
		}
		//save the comfirm code into timer map
		keyData := fmt.Sprintf("%v", register)
		if err = md.ComfirmCode.Add(keyData); err != nil {
			response.StatusCode = -8
			response.Msg = fmt.Sprintf("保存验证码失败, '%v' ,请稍后重试", err)
			logs.Critical(response.Msg)
			goto tail
		}

	case "comfirmAndRegisit": //vertify the comfirm code and create a new account if pass  🍖🍚
		register := md.RegisterData{}
		if err = Parse(postBody.Data, &register); err != nil {
			response.StatusCode = -9
			response.Msg = fmt.Sprintf("解析请求体数据失败:' %v', 请稍后重试", err)
			goto tail
		}
		//check user name, password and email format
		if !tb.CheckUserName(register.Name) {
			response.StatusCode = -5
			response.Msg = fmt.Sprintf("用户名格式错误")
			goto tail
		}
		if !tb.CheckEmail(register.Email) {
			response.StatusCode = -5
			response.Msg = fmt.Sprintf("邮箱地址格式错误")
			goto tail
		}
		if !tb.CheckPassword(register.Password) {
			response.StatusCode = -5
			response.Msg = fmt.Sprintf("密码格式错误")
			goto tail
		}
		//check the comfirm code
		keyData := fmt.Sprintf("%v", register)
		if err = md.ComfirmCode.Get(keyData); err != nil {
			rlog.Warn("%v", err)
			response.StatusCode = -10
			response.Msg = fmt.Sprintf("验证失败：%v", err)
			goto tail
		}
		//comfirm success, create a new account for user
		register.Password = MD5Parse(register.Password)
		if err = md.CreateAccount(register); err != nil {
			rlog.Error("%v", err)
			response.StatusCode = -11
			response.Msg = fmt.Sprintf("💣 创建账号失败：%v, 请稍后重试", err)
			goto tail
		}
		rlog.Info("New account have been create! %s", register.Name)
	case "staticdata": //get static data from about-page 🍙
		response.Data = md.GetStaticData()
	}

	response.StatusCode = 0
	response.Msg = "Success"
tail:
	//update static data 👀
	md.TodayRequestTimes++
	this.Data["json"] = response
	this.ServeJSON()
}

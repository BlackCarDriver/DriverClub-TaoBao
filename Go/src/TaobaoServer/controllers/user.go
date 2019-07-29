package controllers

import (
	md "TaobaoServer/models"
	"encoding/json"
	"fmt"
)

//个人详情页面或其他用户主页信息获取接口
func (this *PersonalDataController) Post() {
	postBody := md.PersonalPostBody{}
	var err error
	if err = json.Unmarshal(this.Ctx.Input.RequestBody, &postBody); err != nil {
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
			fmt.Println(err)
			this.Data["json"] = ""
			goto tail
		}
		this.Data["json"] = data

	case "mygoods": //我的商品
		this.Data["json"] = &md.MockGoodsShort
		var data []md.GoodsShort
		err = md.GetMyGoods(userName, &data)
		if err != nil {
			fmt.Println(err)
			this.Data["json"] = ""
			goto tail
		}
		this.Data["json"] = data

	case "mycollect": //我的收藏
		var data []md.GoodsShort
		err = md.GetMyCollectGoods(userName, &data)
		if err != nil {
			fmt.Println(err)
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
			fmt.Println(err)
			this.Data["json"] = ""
			goto tail
		}
		this.Data["json"] = data

	case "naving": //导航栏我的数据
		var data md.MyStatus
		err = md.GetNavingMsg(userName, &data)
		if err != nil {
			fmt.Println(err)
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

//个人主页更新信息接口
func (this *UpdataMsgController) Post() {
	postBody := md.UpdateBody{}
	var updateTag string
	var err error
	if err = json.Unmarshal(this.Ctx.Input.RequestBody, &postBody); err != nil {
		this.Data["json"] = md.GetUpdateResult(-100, err)
		goto tail
	}
	updateTag = postBody.Tag
	switch updateTag {
	//更新基本信息
	case "MyBaseMessage":
		fmt.Println("Updata base message ...")
		postData := md.UpdeteMsg{}
		Parse(postBody.Data, &postData)
		err = md.UpdateUserBaseMsg(postData)
		if err != nil {
			this.Data["json"] = md.GetUpdateResult(-200, err)
			goto tail
		}
		//更新联系方式
	case "MyConnectMessage":
		fmt.Println("Updata connect ways...")
		postData := md.UpdeteMsg{}
		Parse(postBody.Data, &postData)
		err = md.UpdateUserConnectMsg(postData)
		if err != nil {
			this.Data["json"] = md.GetUpdateResult(-300, err)
			goto tail
		}
		//修改头像
	case "MyHeadImage":
		fmt.Println("Updata User headimg...")
		postData := md.UpdeteMsg{}
		Parse(postBody.Data, &postData)
		err = md.UpdateUserHeadIMg(postBody.UserId, postData.Headimg)
		if err != nil {
			this.Data["json"] = md.GetUpdateResult(-300, err)
			goto tail
		}
	default:
		this.Data["json"] = md.GetUpdateResult(-600, fmt.Errorf("No such tag!"))
	}
	this.Data["json"] = md.GetUpdateResult(0, nil)
tail:
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
			fmt.Println(err)
		}
	case "confirmcode":
		fmt.Println("confirmcode...")
	}
	this.Data["json"] = &md.MockRequireResult
	this.ServeJSON()
}

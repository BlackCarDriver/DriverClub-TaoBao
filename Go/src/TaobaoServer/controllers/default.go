package controllers

import (
	"TaobaoServer/models"
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego"
)

const (
	imgPath = `E:\tempfile\taobaosource\`
)

type HPGoodsController struct {
	beego.Controller
}
type GoodsTypeController struct {
	beego.Controller
}
type GoodsDetailController struct {
	beego.Controller
}
type PersonalDataController struct {
	beego.Controller
}

type UpdataMsgController struct {
	beego.Controller
}

type UploadGoodsController struct {
	beego.Controller
}

type UploadImagesController struct {
	beego.Controller
}

type EntranceController struct {
	beego.Controller
}

//主页商品封面数据
func (this *HPGoodsController) Post() {
	PostBody := models.PostBody1{}
	var err error
	if err = json.Unmarshal(this.Ctx.Input.RequestBody, &PostBody); err != nil {
		return
	}
	//需要根据不同的类型，标签和页数从数据库中获取真实数据
	fmt.Println(PostBody.GoodsTag, "------------------", PostBody.GoodsIndex)
	this.Data["json"] = &models.MockGoodsData
	this.ServeJSON()
}

//返回商品分类和标签列表数据
func (this *GoodsTypeController) Get() {
	//需要从数据库获取真实数据返回
	this.Data["json"] = &models.MockTypeData
	this.ServeJSON()
}

//商品详情获取数据接口
func (this *GoodsDetailController) Post() {
	postBody := models.GoodsPostBody{}
	var err error
	if err = json.Unmarshal(this.Ctx.Input.RequestBody, &postBody); err != nil {
		return
	}
	goodId := postBody.GoodId
	datatype := postBody.DataType
	if goodId == 0 || datatype == "" {
		return
	}
	fmt.Println("GoodsDetail postBody :", postBody)
	switch datatype {
	case "message":
		this.Data["json"] = &models.MockGoodsMessage
	case "detail":
		this.Data["json"] = &models.MockGoodsDetail
	}

	this.ServeJSON()
}

//个人详情页面信息获取接口
func (this *PersonalDataController) Post() {
	postBody := models.PersonalPostBody{}
	var err error
	if err = json.Unmarshal(this.Ctx.Input.RequestBody, &postBody); err != nil {
		return
	}
	userName := postBody.Name
	dataTag := postBody.Tag
	if userName == "" || dataTag == "" {
		return
	}
	fmt.Println(userName, " -------------- ", dataTag)
	switch dataTag {
	case "mymsg":
		this.Data["json"] = &models.MockUserMessage
	case "mygoods":
		this.Data["json"] = &models.MockGoodsShort
	case "mycollect":
		this.Data["json"] = &models.MockGoodsShort
	case "message":
		this.Data["json"] = &models.MockMyMessage
	case "rank":
		this.Data["json"] = &models.MockRank
	case "mycare":
		this.Data["json"] = &models.MockCare
	case "naving":
		this.Data["json"] = &models.MockMystatus
	case "othermsg":
		this.Data["json"] = &models.MockUserMessage
	case "setdata":
		this.Data["json"] = &models.MockUserSetData

	}
	this.ServeJSON()
}

//更新信息接口
func (this *UpdataMsgController) Post() {
	postBody := models.UpdateBody{}
	var err error
	if err = json.Unmarshal(this.Ctx.Input.RequestBody, &postBody); err != nil {
		return
	}
	updateTag := postBody.Tag
	switch updateTag {
	case "MyBaseMessage":
		fmt.Println("Updata base message ...")
		this.Data["json"] = &models.MockUpdateResult
	case "MyConnectMessage":
		fmt.Println("Updata connect ways...")
		this.Data["json"] = &models.MockUpdateResult
	case "MyHeadImage":
		fmt.Println("Updata head img ...")
		this.Data["json"] = &models.MockUpdateResult
	}
	this.ServeJSON()
}

//上传商品
func (this *UploadGoodsController) Post() {
	postBody := models.UploadGoodsData{}
	var err error
	if err = json.Unmarshal(this.Ctx.Input.RequestBody, &postBody); err != nil {
		return
	}
	fmt.Println(postBody)
	this.Data["json"] = &models.MockUpLoadResult
	this.ServeJSON()
}

//上传图片
func (this *UploadImagesController) Post() {
	f, h, err := this.GetFile("file")
	if err != nil {
		return
	}
	defer f.Close()
	err = this.SaveToFile("file", imgPath+h.Filename)
	if err != nil {
		fmt.Println("savetofile error : ", err)
		return
	}
	this.Data["json"] = &models.MockUpLoadResult
	this.ServeJSON()
}

//登录，注册， 更换验证码， 获取验证码
func (this *EntranceController) Post() {
	postBody := models.EntranceBody{}
	var err error
	if err = json.Unmarshal(this.Ctx.Input.RequestBody, &postBody); err != nil {
		return
	}
	updateTag := postBody.Tag
	switch updateTag {
	case "login":
		fmt.Println("user login ...")
	case "CheckRegister":
		fmt.Println("CheckRegister...")
	case "confirmcode":
		fmt.Println("confirmcode...")
	}
	this.Data["json"] = &models.MockRequireResult
	this.ServeJSON()
}

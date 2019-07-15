package controllers

import (
	md "TaobaoServer/models"
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego"
)

const (
	imgPath = "E:\\tempfile\\taobaosource\\"
)

var (
	tmpImgurl = "https://gss0.bdstatic.com/6LZ1dD3d1sgCo2Kml5_Y_D3/sys/portrait/item/2f6de585abe7baa7e5a4a7e78b82e9a38e5a"
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
	PostBody := md.PostBody1{}
	var err error
	if err = json.Unmarshal(this.Ctx.Input.RequestBody, &PostBody); err != nil {
		return
	}
	//需要根据不同的类型，标签和页数从数据库中获取真实数据
	fmt.Println(PostBody.GoodsTag, "------------------", PostBody.GoodsIndex)
	this.Data["json"] = &md.MockGoodsData
	this.ServeJSON()
}

//返回商品分类和标签列表数据
func (this *GoodsTypeController) Get() {
	//需要从数据库获取真实数据返回
	this.Data["json"] = &md.MockTypeData
	this.ServeJSON()
}

//商品详情获取数据接口
func (this *GoodsDetailController) Post() {
	postBody := md.GoodsPostBody{}
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
		this.Data["json"] = &md.MockGoodsMessage
	case "detail":
		this.Data["json"] = &md.MockGoodsDetail
	}

	this.ServeJSON()
}

//个人详情页面信息获取接口
func (this *PersonalDataController) Post() {
	postBody := md.PersonalPostBody{}
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
		this.Data["json"] = &md.MockUserMessage
	case "mygoods":
		this.Data["json"] = &md.MockGoodsShort
	case "mycollect":
		this.Data["json"] = &md.MockGoodsShort
	case "message":
		this.Data["json"] = &md.MockMyMessage
	case "rank":
		this.Data["json"] = &md.MockRank
	case "mycare":
		this.Data["json"] = &md.MockCare
	case "naving":
		this.Data["json"] = &md.MockMystatus
	case "othermsg":
		this.Data["json"] = &md.MockUserMessage
	case "setdata":
		this.Data["json"] = &md.MockUserSetData

	}
	this.ServeJSON()
}

//更新信息接口
func (this *UpdataMsgController) Post() {
	postBody := md.UpdateBody{}
	var err error
	if err = json.Unmarshal(this.Ctx.Input.RequestBody, &postBody); err != nil {
		return
	}
	updateTag := postBody.Tag
	switch updateTag {
	case "MyBaseMessage":
		fmt.Println("Updata base message ...")
		this.Data["json"] = &md.MockUpdateResult
	case "MyConnectMessage":
		fmt.Println("Updata connect ways...")
		this.Data["json"] = &md.MockUpdateResult
	case "MyHeadImage":
		fmt.Println("Updata head img ...")
		this.Data["json"] = &md.MockUpdateResult
	}
	this.ServeJSON()
}

//上传商品
func (this *UploadGoodsController) Post() {
	var goodsdata md.UploadGoodsData
	var returnRes md.UpLoadResult
	var err error
	if err = json.Unmarshal(this.Ctx.Input.RequestBody, &goodsdata); err != nil {
		returnRes = md.CreateUploadRes(-100, err, "")
		goto tail
	}
	if err = md.CreateGoods(goodsdata); err != nil {
		returnRes = md.CreateUploadRes(-200, err, "")
		goto tail
	}
	returnRes = md.CreateUploadRes(0, nil, "")
tail:
	this.Data["json"] = &returnRes
	this.ServeJSON()
}

//保存用户上传的图片，返回访问这个图片的url
func (this *UploadImagesController) Post() {
	var uploadRes md.UpLoadResult
	f, h, err := this.GetFile("file")
	if err != nil {
		uploadRes = md.CreateUploadRes(-100, err, "")
		goto tail
	}
	defer f.Close()
	err = this.SaveToFile("file", imgPath+h.Filename)
	if err != nil {
		uploadRes = md.CreateUploadRes(-200, err, "")
		goto tail
	}
	uploadRes = md.CreateUploadRes(0, err, tmpImgurl)
tail:
	fmt.Println(uploadRes)
	this.Data["json"] = &uploadRes
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

//将map[string]interface{} 转换成相应结构体
func Parse(data interface{}, container interface{}) error {
	tdata, err := json.Marshal(data)
	if err != nil {
		return err
	}
	err = json.Unmarshal(tdata, container)
	return err
}

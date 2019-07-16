package controllers

import (
	md "TaobaoServer/models"
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego"
)

const (
	//上传的图片保存到的位置
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

//将map[string]interface{} 转换成相应结构体
func Parse(data interface{}, container interface{}) error {
	tdata, err := json.Marshal(data)
	if err != nil {
		return err
	}
	err = json.Unmarshal(tdata, container)
	return err
}

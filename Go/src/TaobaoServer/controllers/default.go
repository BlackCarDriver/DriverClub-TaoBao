package controllers

import (
	md "TaobaoServer/models"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/astaxie/beego/logs"

	"github.com/astaxie/beego"
)

const (
	//上传的图片保存到的位置
	imgPath    = "E:\\tempfile\\taobaosource\\"
	imgUrlRoot = "https://blackcardriver.com/images/"
)

var (
	tmpImgurl = "https://gss0.bdstatic.com/6LZ1dD3d1sgCo2Kml5_Y_D3/sys/portrait/item/2f6de585abe7baa7e5a4a7e78b82e9a38e5a"
	random    *rand.Rand
)

func init() {
	beego.SetLogFuncCall(true)
	beego.SetLogger("file", `{"filename":"logs/test.log"}`)
	logs.SetLogFuncCallDepth(3)
	random = rand.New(rand.NewSource(time.Now().UnixNano()))
}

type TestController struct {
	beego.Controller
}
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

type UpdateController struct {
	beego.Controller
}

type DeleteController struct {
	beego.Controller
}

//test interface 🍍
func (this *TestController) Get() {
	logs.Info("Soneont test the program by it iterface")
	this.Data["Ip"] = this.Ctx.Input.IP()
	this.Data["Host"] = this.Ctx.Input.Host()
	this.Data["Domain"] = this.Ctx.Input.Domain()
	this.Data["SubDomains"] = this.Ctx.Input.SubDomains()
	this.Data["Scheme"] = this.Ctx.Input.Scheme()
	this.Data["Protocaol"] = this.Ctx.Input.Protocol()
	this.Data["Refer"] = this.Ctx.Input.Refer()
	this.Data["Referer"] = this.Ctx.Input.Referer()
	this.Data["Site"] = this.Ctx.Input.Site()
	this.Data["UserAgent"] = this.Ctx.Input.UserAgent()
	this.TplName = "test.tpl"
}

//saved user's upload images into dist and return a url that get it images 🍍
//response to UploadImg() in fontend
func (this *UploadImagesController) Post() {
	response := md.ReplyProto{}
	response.StatusCode = 0
	f, h, err := this.GetFile("file")
	if err != nil {
		response.StatusCode = -1
		response.Msg = fmt.Sprintf("Can not get file from request: %v", err)
		logs.Error(response.Msg)
		goto tail
	}
	h.Filename, err = CheckImgName(h.Filename)
	if err != nil {
		response.StatusCode = -2
		response.Msg = fmt.Sprintf("Can not get images name from request: %v", err)
		logs.Error(response.Msg)
		goto tail
	}
	defer f.Close()
	err = this.SaveToFile("file", imgPath+h.Filename)
	if err != nil {
		response.StatusCode = -3
		response.Msg = fmt.Sprintf("Can not save file: %v", err)
		logs.Error(response.Msg)
		goto tail
	}
	//DOTO: use true image url
	// response.Data = imgUrlRoot + h.Filename
	// response.Data = "https://blackcardriver.cn/images/huawei.jpg"
	response.Data = "https://blackcardriver.cn/images/huaji.png"
tail:
	this.Data["json"] = response
	this.ServeJSON()
}

//small update all kind of record such as like numbers, collect numbers 🍍
//response for SmallUpdate() in fontend
func (this *UpdateController) Post() {
	postBody := md.RequestProto{}
	response := md.ReplyProto{}
	response.StatusCode = 0
	var err error
	var api, targetid, userid string
	//parse request protocol
	if err = json.Unmarshal(this.Ctx.Input.RequestBody, &postBody); err != nil {
		response.StatusCode = -1
		response.Msg = fmt.Sprintf("Can not parse postbody: %v", err)
		logs.Error(response.Msg)
		goto tail
	}
	api = postBody.Api
	userid = postBody.UserId
	targetid = postBody.TargetId
	//check whether the data is complete
	if api == "" || targetid == "" {
		response.StatusCode = -2
		response.Msg = fmt.Sprintf("Can't get api or targetid from request data")
		logs.Error(response.Msg)
		goto tail
	}
	switch api {
	case "likegoods": //add like to a goods 🔥
		err = md.AddGoodsLike(userid, targetid)
		if err != nil {
			response.StatusCode = -3
			response.Msg = fmt.Sprintf("AddGoodsLike fail: %v", err)
			logs.Error(response.Msg)
		}
		goto tail

	case "sendmessage": //send a private message to goods owner 🔥
		logs.Info(postBody.Data)
		appendData := postBody.Data.(map[string]interface{})
		message := ""
		if message = appendData["message"].(string); message == "" {
			response.StatusCode = -4
			response.Msg = "Can't get message on postbody"
			logs.Error(response.Msg)
			goto tail
		}
		if err = md.AddUserMessage(userid, targetid, message); err != nil {
			response.StatusCode = -5
			response.Msg = fmt.Sprintf("AddUserMssage() fail: %v", err)
			logs.Error(response.Msg)
		}
		goto tail

	case "addcollect": //add a goods to favorite	🔥
		if err = md.AddGoodsCollect(userid, targetid); err != nil {
			response.StatusCode = -6
			response.Msg = fmt.Sprintf("AddGoodsCollect() fail: %v", err)
			logs.Error(response.Msg)
		}
		goto tail

	case "addcomment": // reviews a goods 🔥
		appendData := postBody.Data.(map[string]interface{})
		comment := ""
		if comment = appendData["comment"].(string); comment == "" {
			response.StatusCode = -10
			response.Msg = "Can't get comment on postbody"
			logs.Error(response.Msg)
			goto tail
		}
		if err = md.AddGoodsComment(userid, targetid, comment); err != nil {
			response.StatusCode = -11
			response.Msg = fmt.Sprintf("AddGoodsComment() fail %v", err)
			logs.Error(response.Msg)
		}
		goto tail

	case "likeuser": //add a like to a user profile
		if err = md.AddUserLike(userid, targetid); err != nil {
			response.StatusCode = -7
			response.Msg = fmt.Sprintf("AddUserLike() fail: %v", err)
			logs.Error(response.Msg)
		}
		goto tail

	case "addconcern": // add someone to favorite
		if err = md.AddUserConcern(userid, targetid); err != nil {
			response.StatusCode = -8
			response.Msg = fmt.Sprintf("AddUserConcern() fail: %v", err)
			logs.Error(response.Msg)
			goto tail
		}
		goto tail

	default:
		response.StatusCode = -100
		response.Msg = fmt.Sprintf("No such api %s", api)
		logs.Error(response.Msg)
	}
tail:
	this.Data["json"] = response
	this.ServeJSON()
}

//delete data such as collect's goods and user and receive message 🍑
//DeleteMyData()
func (this *DeleteController) Post() {
	postBody := md.RequestProto{}
	response := md.ReplyProto{}
	response.StatusCode = 0
	var err error
	var api, targetid, userid string
	//parse request protocol
	if err = json.Unmarshal(this.Ctx.Input.RequestBody, &postBody); err != nil {
		response.StatusCode = -1
		response.Msg = fmt.Sprintf("Can not parse postbody: %v", err)
		logs.Error(response.Msg)
		goto tail
	}
	api = postBody.Api
	userid = postBody.UserId
	targetid = postBody.TargetId
	//check whether the data is complete
	if api == "" || targetid == "" || userid == "" {
		response.StatusCode = -2
		response.Msg = fmt.Sprintf("Can't get targetid, userid, or api from request data")
		logs.Error(response.Msg)
		goto tail
	}
	switch api {
	case "deletemygoods": //user delete his/her goods, temply change the goods's state
		if err = md.UpdateMyGoodsState(userid, targetid); err != nil {
			response.StatusCode = -3
			response.Msg = fmt.Sprintf("删除商品失败： %v", err)
			logs.Error(response.Msg)
		}
		goto tail

	case "deletemymessage": //delete user's message from database
		if err = md.DeleteMyMessage(userid, targetid); err != nil {
			response.StatusCode = -4
			response.Msg = fmt.Sprintf("删除消息失败： %v", err)
			logs.Error(response.Msg)
		}
		goto tail

	case "uncollectgoods": //delete user collect's goods from database
		if err = md.DeleteMyCollect(userid, targetid); err != nil {
			response.StatusCode = -5
			response.Msg = fmt.Sprintf("取消收藏失败： %v", err)
			logs.Error(response.Msg)
		}
		goto tail

	case "uncollectuser": //delete the record of user concern another user
		if err = md.DeleteMyConcern(userid, targetid); err != nil {
			response.StatusCode = -3
			response.Msg = fmt.Sprintf("取消关注失败： %v", err)
			logs.Error(response.Msg)
		}
		goto tail

	default:
		response.StatusCode = -99
		response.Msg = fmt.Sprintf("No such api: %s", api)
		logs.Error(response.Msg)
	}
tail:
	this.Data["json"] = response
	this.ServeJSON()
}

//################## tool function #######################

//create a random string with length l
func GetRandomString(l int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyz_"
	bytes := []byte(str)
	result := []byte{}
	for i := 0; i < l; i++ {
		result = append(result, bytes[random.Intn(len(bytes))])
	}
	return string(result)
}

//check whether a iamge file name is ok and chenged it name to a random string
func CheckImgName(filename string) (newName string, err error) {
	c := strings.Count(filename, ".")
	if c > 1 {
		err := fmt.Errorf("Comma numbers in image name more than one!")
		logs.Error(err)
		return "", err
	}
	l := strings.LastIndex(filename, ".")
	if l < 0 {
		err := fmt.Errorf("No comma in the image name!")
		logs.Error(err)
		return "", err
	}
	suffix := strings.ToLower(filename[l+1:])
	if suffix != "png" && suffix != "jpg" {
		err := fmt.Errorf("not an png or jpg type images!")
		logs.Error(err)
		return "", err
	}
	return GetRandomString(10) + "." + suffix, nil
}

//parse a interface map into specified struct
func Parse(data interface{}, container interface{}) error {
	tdata, err := json.Marshal(data)
	if err != nil {
		logs.Error(err)
		return err
	}
	err = json.Unmarshal(tdata, container)
	logs.Error(err)
	return err
}

//md5 encrypt
func MD5Parse(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

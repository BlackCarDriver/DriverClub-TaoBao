package controllers

import (
	md "TaobaoServer/models"
	"encoding/json"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/astaxie/beego"
)

const (
	//上传的图片保存到的位置
	imgPath    = "E:\\tempfile\\taobaosource\\"
	imgUrlRoot = "https://blackcardriver.com/taobao/images/"
)

var (
	tmpImgurl = "https://gss0.bdstatic.com/6LZ1dD3d1sgCo2Kml5_Y_D3/sys/portrait/item/2f6de585abe7baa7e5a4a7e78b82e9a38e5a"
	random    *rand.Rand
)

func init() {
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

//测试接口
func (this *TestController) Get() {
	fmt.Println("##### test #####")

	this.Data["json"] = "test"
	this.ServeJSON()
}

//保存用户上传的图片，返回访问这个图片的url
func (this *UploadImagesController) Post() {
	f, h, err := this.GetFile("file")
	if err != nil {
		this.Data["json"] = md.CreateUploadRes(-100, err, "")
		goto tail
	}
	h.Filename, err = CheckImgName(h.Filename)
	if err != nil {
		this.Data["json"] = md.CreateUploadRes(-200, err, "")
		goto tail
	}
	defer f.Close()
	err = this.SaveToFile("file", imgPath+h.Filename)
	if err != nil {
		this.Data["json"] = md.CreateUploadRes(-300, err, "")
		goto tail
	}
	// this.Data["json"] = md.CreateUploadRes(0, err, imgUrlRoot+h.Filename)
	this.Data["json"] = md.CreateUploadRes(0, err, tmpImgurl)
tail:
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

//得到一个长度为l的随机字符串
func GetRandomString(l int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyz_"
	bytes := []byte(str)
	result := []byte{}
	for i := 0; i < l; i++ {
		result = append(result, bytes[random.Intn(len(bytes))])
	}
	return string(result)
}

//检查一个图片名字是否合法，合法则改成另外一个随机字符串
func CheckImgName(filename string) (newName string, err error) {
	c := strings.Count(filename, ".")
	if c > 1 {
		return "", fmt.Errorf("Comma numbers in image name more than one!")
	}
	l := strings.LastIndex(filename, ".")
	if l < 0 {
		return "", fmt.Errorf("No comma in the image name!")
	}
	suffix := strings.ToLower(filename[l+1:])
	if suffix != "png" && suffix != "jpg" {
		return "", fmt.Errorf("not an png or jpg type images!")
	}
	return GetRandomString(10) + "." + suffix, nil
}

//更新各种信息的接口,如商品点赞数，私信，商品收藏
func (this *UpdateController) Post() {
	postBody := md.UpdatePostBody{}
	result := md.UpdateResult{}
	var err error
	if err = json.Unmarshal(this.Ctx.Input.RequestBody, &postBody); err != nil {
		fmt.Println("error : ", err)
		return
	}
	tag := postBody.Tag
	switch tag {
	case "likegoods": //商品点赞
		err = md.AddGoodsLike(postBody.UserId, postBody.TargetId)
		if err == nil {
			result.Status = 0
			goto tail
		}
		result.Status = -1
		result.Describe = fmt.Sprintf("点赞失败： %s ", err)
	case "sendmessage": //商品页面向卖家发私信
		err = md.AddUserMessage(postBody.UserId, postBody.TargetId, postBody.StrData)
		if err == nil {
			result.Status = 0
			goto tail
		}
		result.Status = -1
		result.Describe = fmt.Sprintf("发送失败: %s", err)
	case "addcollect": //收藏商品
		err = md.AddGoodsCollect(postBody.UserId, postBody.TargetId)
		if err == nil {
			result.Status = 0
			goto tail
		}
		result.Status = -1
		result.Describe = fmt.Sprintf("收藏失败: %s", err)
	case "likeuser": //个人信息页面点赞
		err = md.AddUserLike(postBody.UserId, postBody.TargetId)
		if err == nil {
			result.Status = 0
			goto tail
		}
		result.Status = -1
		result.Describe = fmt.Sprintf("点赞失败： %s", err)
	case "addconcern": //个人信息页面关注
		err = md.AddUserConcern(postBody.UserId, postBody.TargetId)
		if err == nil {
			result.Status = 0
			goto tail
		}
		result.Status = -1
		result.Describe = fmt.Sprintf("关注失败： %s", err)
	case "addcomment": // reviews a goods
		err = md.AddGoodsComment(postBody.UserId, postBody.TargetId, postBody.StrData)
		if err == nil {
			result.Status = 0
		} else {
			result.Status = -1
			result.Describe = fmt.Sprintf("评论失败: %v", err)
		}
		goto tail
	default:
		fmt.Println(tag)
	}
tail:
	this.Data["json"] = result
	this.ServeJSON()
}

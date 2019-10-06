package controllers

import (
	md "TaobaoServer/models"
	tb "TaobaoServer/toolsbox"
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"math/rand"
	"mime/multipart"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/config"
	"github.com/astaxie/beego/logs"
	"github.com/nfnt/resize"
)

//config varlue
//https://blackcardriver.cn/mkimg/%s
var (
	imgPath               = ""
	imgUrlTP              = ""
	maxGoodsHeadImgSizekb = int64(0)
	maxFeedbackImgSizekb  = int64(200)
)

//random seed
var random *rand.Rand
var rlog *logs.BeeLogger

func init() {
	logs.SetLogFuncCallDepth(3)
	//make a logger specially used by router
	rlog = logs.NewLogger()
	rlog.SetLogger("file", `{"filename":"logs/router.log", "daily":false, "maxsize":512000}`)
	rlog.EnableFuncCallDepth(true)
	rlog.Info("Router logs init success!")
	rlog.SetLogFuncCallDepth(2)
	//read config from conf file
	if iniconf, err := config.NewConfig("ini", "./conf/driver.conf"); err != nil {
		rlog.Error("%v", err)
		panic(err)
	} else {
		//images upload and saving config
		imgPath = iniconf.String("imgPath")
		imgUrlTP = iniconf.String("imgUrlTP")
		if maxGoodsHeadImgSizekb, err = iniconf.Int64("maxGoodsHeadImgSizekb"); err != nil {
			logs.Error(err)
			maxGoodsHeadImgSizekb = 40
		}
		//get toeken key
		tokenkey = iniconf.String("tokenkey")
		if tokenkey == "" {
			err = errors.New("Can't get tokenkey from config file")
			logs.Error(err)
			panic(err)
		} else {
			secretKey = []byte(tokenkey)
		}

		//email sending server config 🍖
		stmpHost = iniconf.String("stmpHost")
		myemail = iniconf.String("myemail")
		mypassword = iniconf.String("mypassword")
		stmpPort, _ = iniconf.Int("stmpPort")
		sendEmail, _ = iniconf.Bool("sendEmail")
	}
	imgPath = strings.TrimRight(imgPath, `\`)
	random = rand.New(rand.NewSource(time.Now().UnixNano()))
}

type TestController struct {
	beego.Controller
}
type PublicController struct {
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
type PostFormController struct {
	beego.Controller
}

//test interface 🍍🌰
func (this *TestController) Get() {
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
	this.Data["Runhour"] = md.RunHour
	tlog, _ := ParseFile("./logs/router.log")
	this.Data["RouterLog"] = strings.ReplaceAll(tlog, "[E]", "🍉")
	tlog, _ = ParseFile("./logs/models.log")
	this.Data["ModelsLog"] = strings.ReplaceAll(tlog, "[E]", "🍉")
	this.TplName = "test.tpl"
	//update static data 👀
	md.TodayRequestTimes++
}

//saved user's upload images into dist and return a url that get it images 🍍🌰🍚
//response to UploadImg() in fontend
///upload/images
func (this *UploadImagesController) Post() {
	response := md.ReplyProto{}
	var h *multipart.FileHeader
	var f multipart.File
	var err error
	var savePath = ""
	//only let sign in user upload images
	userid := this.GetString("userid")
	token := this.GetString("token")
	if userid == "" || token == "" {
		response.StatusCode = -1
		response.Msg = "无法获取 userid 或 token"
		goto tail
	}
	if !CheckToken(userid, token) {
		response.StatusCode = -1
		response.Msg = "Token 错误或已过期: %v"
		rlog.Error("User '%s' upload image fail because token unpass ", userid)
		goto tail
	}
	//get file from from data
	f, h, err = this.GetFile("file")
	if err != nil {
		response.StatusCode = -1
		response.Msg = fmt.Sprintf("Can not get file from request: %v", err)
		rlog.Error("%v", response.Msg)
		goto tail
	}
	//check the size of upload images (300kb)
	rlog.Info("Upload images name:%s,  size:%d", h.Filename, h.Size)
	if h.Size > maxGoodsHeadImgSizekb<<10 {
		response.StatusCode = -2
		response.Msg = fmt.Sprintf("images size is too larger")
		rlog.Error("%v", response.Msg)
		goto tail
	}
	//check and change the images name
	h.Filename, err = CheckImgName(h.Filename)
	if err != nil {
		response.StatusCode = -3
		response.Msg = fmt.Sprintf("File type not right: %v", err)
		rlog.Error("%v", response.Msg)
		goto tail
	}
	defer f.Close()
	//save the file into dish
	savePath = fmt.Sprintf("%s/%s", imgPath, h.Filename)
	if err = this.SaveToFile("file", savePath); err != nil {
		response.StatusCode = -4
		response.Msg = fmt.Sprintf("Can not save file: %v", err)
		rlog.Error("%v", response.Msg)
		goto tail
	}
	//compress the image,
	if err = CompressImg(savePath, 100); err != nil {
		rlog.Critical("images compress function fail:%v", err)
	}
	response.StatusCode = 0
	response.Data = fmt.Sprintf(imgUrlTP, h.Filename)
tail:
	//update static data 👀
	md.TodayRequestTimes++
	this.Data["json"] = response
	this.ServeJSON()
}

//small update all kind of record such as like numbers, collect numbers 🍍🍔
//response for SmallUpdate() in fontend
//all request need token here
func (this *UpdateController) Post() {
	postBody := md.RequestProto{}
	response := md.ReplyProto{}
	response.StatusCode = 0
	var err error
	var api, targetid, userid, token string
	//parse request protocol
	if err = json.Unmarshal(this.Ctx.Input.RequestBody, &postBody); err != nil {
		response.StatusCode = -1
		response.Msg = fmt.Sprintf("Can not parse postbody: %v", err)
		rlog.Error("%v", response.Msg)
		goto tail
	}
	api = postBody.Api
	userid = postBody.UserId
	targetid = postBody.TargetId
	token = postBody.Token
	//check whether the data is complete
	if api == "" || targetid == "" {
		response.StatusCode = -2
		response.Msg = fmt.Sprintf("Can't get api or targetid from request data")
		rlog.Error("%v", response.Msg)
		goto tail
	}
	//check token
	if token == "" || !CheckToken(userid, token) {
		rlog.Warn("User %s request update %s with worng token", userid, api)
		response.StatusCode = -1000
		response.Msg = "Token错误或过期,请重新登录！"
		goto tail
	}
	switch api {
	case "likegoods": //add like to a goods
		err = md.AddGoodsLike(userid, targetid)
		if err != nil {
			response.StatusCode = -3
			response.Msg = fmt.Sprintf("AddGoodsLike fail: %v", err)
			rlog.Error("%v", response.Msg)
		}
		goto tail

	case "sendmessage": //send a private message to goods owner🍚🍜
		if md.CheckFrequent(&postBody) {
			response.StatusCode = -4
			response.Msg = "操作太频繁,请稍后再试哦 :)"
			goto tail
		}
		appendData := postBody.Data.(map[string]interface{})
		if userid == targetid {
			response.StatusCode = -4
			response.Msg = "不能发消息给自己哦 :)"
			rlog.Error("%v", response.Msg)
			goto tail
		}
		message := ""
		if message = appendData["message"].(string); !tb.CheckMessage(message) {
			response.StatusCode = -4
			response.Msg = "消息格式不通过"
			rlog.Error("%v", response.Msg)
			goto tail
		}
		if err = md.AddUserMessage(userid, targetid, message); err != nil {
			response.StatusCode = -5
			response.Msg = fmt.Sprintf("保存消息数据失败： %v", err)
			rlog.Error("%v", response.Msg)
			goto tail
		}
		goto tail

	case "addcollect": //add a goods to favorite🍚
		if userid == "" || targetid == "" {
			response.StatusCode = -6
			response.Msg = fmt.Sprintf("读取用户id或目标id失败")
			rlog.Error("%v", response.Msg)
		}
		if err = md.AddGoodsCollect(userid, targetid); err != nil {
			response.StatusCode = -6
			response.Msg = fmt.Sprintf("更新数据库失败: %v", err)
			rlog.Error("%v", response.Msg)
		}
		md.Uas2.Add(userid)
		goto tail

	case "addcomment": //user comment at a goods 🍚🍜
		if md.CheckFrequent(&postBody) {
			response.StatusCode = -4
			response.Msg = "操作太频繁,请稍后再试哦 :)"
			goto tail
		}
		appendData := postBody.Data.(map[string]interface{})
		comment := ""
		//read and check text data from request struct
		if comment = appendData["comment"].(string); !tb.CheckComment(comment) {
			response.StatusCode = -10
			response.Msg = "评论格式不不通过"
			rlog.Error("%v", response.Msg)
			goto tail
		}
		if err = md.AddGoodsComment(userid, targetid, comment); err != nil {
			response.StatusCode = -11
			response.Msg = fmt.Sprintf("保存评论失败：%v", err)
			rlog.Error("%v", response.Msg)
			goto tail
		}
		goto tail

	case "likeuser": //add a like to a user profile
		if userid == "" || targetid == "" {
			response.StatusCode = -7
			response.Msg = fmt.Sprintf("无法读取用户id或目标id")
			rlog.Error("%v", response.Msg)
			goto tail
		}
		if err = md.AddUserLike(userid, targetid); err != nil {
			response.StatusCode = -7
			response.Msg = fmt.Sprintf("AddUserLike() fail: %v", err)
			rlog.Error("%v", response.Msg)
			goto tail
		}

	case "addconcern": // add someone to favorite
		if err = md.AddUserConcern(userid, targetid); err != nil {
			response.StatusCode = -8
			response.Msg = fmt.Sprintf("AddUserConcern() fail: %v", err)
			rlog.Error("%v", response.Msg)
			goto tail
		}
		goto tail
	case "msgisread": //change the state of user message, tag than as isread 🍞
		if err = md.UpdateMessageState(targetid); err != nil {
			response.StatusCode = -9
			response.Msg = fmt.Sprintf("Change message state fail:%v", err)
			rlog.Error(response.Msg)
			goto tail
		}
	default:
		response.StatusCode = -100
		response.Msg = fmt.Sprintf("No such api %s", api)
		rlog.Error("%v", response.Msg)
	}
tail:
	//update static data 👀
	md.TodayRequestTimes++
	this.Data["json"] = response
	this.ServeJSON()
}

//delete data such as collect's goods and user and receive message 🍑🍔
//all operation here neeed to vertify with token
//DeleteMyData()
func (this *DeleteController) Post() {
	postBody := md.RequestProto{}
	response := md.ReplyProto{}
	response.StatusCode = 0
	var err error
	var api, targetid, userid, token string
	//parse request protocol
	if err = json.Unmarshal(this.Ctx.Input.RequestBody, &postBody); err != nil {
		response.StatusCode = -1
		response.Msg = fmt.Sprintf("Can not parse postbody: %v", err)
		rlog.Error("%v", response.Msg)
		goto tail
	}
	api = postBody.Api
	userid = postBody.UserId
	targetid = postBody.TargetId
	token = postBody.Token
	//check whether the data is complete
	if api == "" || targetid == "" || userid == "" {
		response.StatusCode = -2
		response.Msg = fmt.Sprintf("Can't get targetid, userid, or api from request data")
		rlog.Error("%v", response.Msg)
		goto tail
	}
	//check token
	if token == "" || !CheckToken(userid, token) {
		rlog.Warn("User %s request delete %s with worng token", userid, api)
		response.StatusCode = -1000
		response.Msg = "Token错误或过期,请重新登录！"
		goto tail
	}
	switch api {
	case "deletemygoods": //user delete his/her goods, temply change the goods's state
		if err = md.UpdateMyGoodsState(userid, targetid); err != nil {
			response.StatusCode = -3
			response.Msg = fmt.Sprintf("删除商品失败： %v", err)
			rlog.Error("%v", response.Msg)
		}
		goto tail

	case "deletemymessage": //delete user's message from database
		if err = md.DeleteMyMessage(userid, targetid); err != nil {
			response.StatusCode = -4
			response.Msg = fmt.Sprintf("删除消息失败： %v", err)
			rlog.Error("%v", response.Msg)
		}
		goto tail

	case "uncollectgoods": //delete user collect's goods from database
		if err = md.DeleteMyCollect(userid, targetid); err != nil {
			response.StatusCode = -5
			response.Msg = fmt.Sprintf("取消收藏失败： %v", err)
			rlog.Error("%v", response.Msg)
		}
		goto tail

	case "uncollectuser": //delete the record of user concern another user
		if err = md.DeleteMyConcern(userid, targetid); err != nil {
			response.StatusCode = -3
			response.Msg = fmt.Sprintf("取消关注失败： %v", err)
			rlog.Error("%v", response.Msg)
		}
		goto tail

	default:
		response.StatusCode = -99
		response.Msg = fmt.Sprintf("No such api: %s", api)
		rlog.Error("%v", response.Msg)
	}
tail:
	//update static data 👀
	md.TodayRequestTimes++
	this.Data["json"] = response
	this.ServeJSON()
}

//get of upload static data such as feedback message  🍗
func (this *PublicController) Post() {
	var postBody struct {
		Api   string      `json:"api"`
		Token string      `json:"token"`
		Data  interface{} `json:"data"`
	}
	var response struct {
		Status int         `json:"status"`
		Msg    string      `json:"msg"`
		Data   interface{} `json:"data"`
	}
	var err error
	//parse request protocol
	if err = json.Unmarshal(this.Ctx.Input.RequestBody, &postBody); err != nil {
		response.Status = -1
		response.Msg = fmt.Sprintf("Can not parse postbody: %v", err)
		rlog.Error("%v", response.Msg)
		goto tail
	}
	//switch to different function
	switch postBody.Api {
	case "getfeedback": //get user feedback data
		var data []md.FeedBackData
		offset := 0
		if err = Parse(postBody.Data, &offset); err != nil {
			response.Status = -2
			response.Msg = fmt.Sprintf("Can not get offset: %v", err)
			rlog.Error("%v", response.Msg)
			goto tail
		}
		if err = md.GetFeedBack(&data, offset); err != nil {
			response.Status = -3
			response.Msg = fmt.Sprintf("get feedback data fail: %v", err)
			rlog.Error("%v", response.Msg)
			goto tail
		}
		response.Data = data
	case "setfbisread": //update the state of a feedback record as is read
		fbid := 0
		if err = Parse(postBody.Data, &fbid); err != nil {
			response.Status = -4
			response.Msg = fmt.Sprintf("Can not get feedback id: %v", err)
			rlog.Error("%v", response.Msg)
			goto tail
		}
		if err = md.UpdateFeedbackState(fbid); err != nil {
			response.Status = -5
			response.Msg = fmt.Sprintf("Update state fail: %v", err)
			rlog.Error("%v", response.Msg)
			goto tail
		}

	case "staticdata": //get the static data 🍙
		response.Data = md.StaticData

	default:
		response.Status = -99
		response.Msg = fmt.Sprintf("Unsuppose API: %s", postBody.Api)
		rlog.Error("%v", response.Msg)
		goto tail
	}
	response.Status = 0
tail:
	//update static data 👀
	md.TodayRequestTimes++
	this.Data["json"] = response
	this.ServeJSON()
}

//Receive and handle multipart from request 🍗
//add feedback record is here
func (this *PostFormController) Post() {
	postBody := md.RequestProto{}
	response := md.ReplyProto{}
	api := this.GetString("api")
	if api == "" {
		response.StatusCode = -1
		response.Msg = fmt.Sprintf("Can't get api from from data")
		rlog.Error(response.Msg)
		goto tail
	}
	switch api {
	case "feedback": //User feedback, save a record into database  🍙
		data := md.FeedBackData{}
		reason := ""
		//get string value from postfrom
		data.Type = this.GetString("fb_type")
		data.Location = this.GetString("fb_location")
		data.Email = this.GetString("email")
		data.Userid = this.GetString("userid")
		data.Describes = this.GetString("describes")
		if data.Type == "" || data.Location == "" || data.Describes == "" {
			response.StatusCode = -2
			response.Msg = fmt.Sprintf("Can't get type or describtion in feedback data")
			rlog.Warn(response.Msg)
			goto tail
		}
		switch {
		case len(data.Type) > 15:
			reason = "反馈类型不通过"
		case len(data.Location) > 200:
			reason = "报错位置超出长度"
		case len(data.Describes) > 480:
			reason = "反馈描述超出长度"
		case data.Email != "" && !tb.CheckEmail(data.Email):
			reason = "邮箱格式不通过"
		}
		if reason != "" {
			response.StatusCode = -3
			response.Msg = reason
			rlog.Error(response.Msg)
			goto tail
		}
		//get images from postform
		if f, h, err := this.GetFile("images"); err == nil {
			//check the size of upload images
			rlog.Info("Upload images name:%s,  size:%d", h.Filename, h.Size)
			if h.Size > maxFeedbackImgSizekb<<10 {
				response.StatusCode = -3
				response.Msg = fmt.Sprintf("Images size is too larger! feedback is cancel")
				rlog.Error(response.Msg)
				goto tail
			}
			//check and change the images name
			h.Filename, err = CheckImgName(h.Filename)
			if err != nil {
				response.StatusCode = -4
				response.Msg = fmt.Sprintf("Images type not right: %v", err)
				rlog.Error("%v", response.Msg)
				goto tail
			}
			defer f.Close()
			//save the images into dist
			savePath := fmt.Sprintf("%s/%s", imgPath, h.Filename)
			if err = this.SaveToFile("images", savePath); err != nil {
				response.StatusCode = -5
				response.Msg = fmt.Sprintf("Can not save upload iamges: %v", err)
				rlog.Error("%v", response.Msg)
				goto tail
			}
			data.Imgurl = fmt.Sprintf(imgUrlTP, h.Filename)
		} else {
			logs.Error(err)
		}
		//save the record into database
		if err := md.AddFeedback(&data); err != nil {
			response.StatusCode = -6
			response.Msg = fmt.Sprintf("Save to database fail: %v", err)
			rlog.Error("%v", response.Msg)
			goto tail
		}
	default:
		response.StatusCode = -99
		response.Msg = fmt.Sprintf("Unsuppose API: %s", postBody.Api)
		rlog.Error("%v", response.Msg)
		goto tail
	}
	response.StatusCode = 0
tail:
	//update static data 👀
	md.TodayRequestTimes++
	this.Data["json"] = response
	this.ServeJSON()
}

//################## tool function #######################

//create a random string with length l
func GetRandomString(l int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWSYZ"
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
		rlog.Error("%v", err)
		return "", err
	}
	l := strings.LastIndex(filename, ".")
	if l < 0 {
		err := fmt.Errorf("No comma in the image name!")
		rlog.Error("%v", err)
		return "", err
	}
	suffix := strings.ToLower(filename[l+1:])
	if suffix != "png" && suffix != "jpg" {
		err := fmt.Errorf("not an png or jpg type images!")
		rlog.Error("%v", err)
		return "", err
	}
	return GetRandomString(10) + "." + suffix, nil
}

//parse a interface map into specified struct
func Parse(data interface{}, container interface{}) error {
	tdata, err := json.Marshal(data)
	if err != nil {
		rlog.Error("%v", err)
		return err
	}
	if err = json.Unmarshal(tdata, container); err != nil {
		rlog.Error("%v", err)
		return err
	}
	return nil
}

//md5 encrypt
func MD5Parse(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

//read a file and parse it into a string 📂
func ParseFile(path string) (text string, err error) {
	file, err := os.Open(path)
	if err != nil {
		return "", fmt.Errorf("Open %s fall: %v", path, err)
	}
	defer file.Close()
	buf := bufio.NewReader(file)
	bytes, err := ioutil.ReadAll(buf)
	if err != nil {
		return "", fmt.Errorf("ioutil.ReadAll fall : %v", err)
	}
	return string(bytes), nil
}

//compress a jpg or png format image, the new images will be named autoly 🌆
func CompressImg(source string, hight uint) error {
	var err error
	var file *os.File
	reg, _ := regexp.Compile(`^.*\.((png)|(jpg))$`)
	if !reg.MatchString(source) {
		err = errors.New("%s is not a .png or .jpg file")
		logs.Error(err)
		return err
	}
	if file, err = os.Open(source); err != nil {
		logs.Error(err)
		return err
	}
	defer file.Close()
	name := file.Name()
	var img image.Image
	switch {
	case strings.HasSuffix(name, ".png"):
		if img, err = png.Decode(file); err != nil {
			logs.Error(err)
			return err
		}
	case strings.HasSuffix(name, ".jpg"):
		if img, err = jpeg.Decode(file); err != nil {
			logs.Error(err)
			return err
		}
	default:
		err = fmt.Errorf("Images %s name not right!", name)
		logs.Error(err)
		return err
	}
	resizeImg := resize.Resize(hight, 0, img, resize.Lanczos3)
	newName := newName(source)
	if outFile, err := os.Create(newName); err != nil {
		logs.Error(err)
		return err
	} else {
		defer outFile.Close()
		err = jpeg.Encode(outFile, resizeImg, nil)
		if err != nil {
			logs.Error(err)
			return err
		}
	}
	abspath, _ := filepath.Abs(newName)
	logs.Info("New imgs successfully save at: %s", abspath)
	return nil
}

//create a file name for the iamges that after resize 🌆
func newName(name string) string {
	dir, file := filepath.Split(name)
	return fmt.Sprintf("%s_%s", dir, file)
}

//create a comfirm code made by six number
func GetRandomCode() string {
	str := "0123456789"
	bytes := []byte(str)
	result := []byte{}
	for i := 0; i < 6; i++ {
		result = append(result, bytes[random.Intn(len(bytes))])
	}
	return string(result)
}

package models

import (
	"fmt"
	"time"
)

//######################################### 公用协议 ##########################################
//public struct that used to request 🍌 🍉
type RequestProto struct {
	Tag       string      `json:"tag"`
	Api       string      `json:"api"`
	UserId    string      `json:"userid"`
	TargetId  string      `json:"targetid"`
	CacheTime int         `json:"cachetime"`
	Data      interface{} `json:"data"`
	Offset    int         `json:"offset"`
	Limit     int         `json:"limit"`
}

//public struct that response by server 🍌 🍉
type ReplyProto struct {
	StatusCode int         `json:"statuscode"`
	Msg        string      `json:"msg"`
	Data       interface{} `json:"data"`
	Rows       int         `json:"rows"`
	Sum        int         `json:"sum"`
}

//########################################## 主页结构和模拟数据 ################################
type Goods1 struct {
	Userid   string    `json:"userid"`
	Username string    `json:"username"`
	Id       string    `json:"id"`
	Name     string    `json:"name"`
	Price    float64   `json:"price"`
	Time     time.Time `json:"time"`
	Headimg  string    `json:"headimg"`
	Title    string    `json:"title"`
	Type     string    `json:"type"`
	Tag      string    `json:"tag"`
}

//主页获取商品封面数据时提供的信息
type PostBody1 struct {
	GoodsType  string `json:"goodstype"`
	GoodsTag   string `json:"goodstag"`
	GoodsIndex int    `json:"goodsindex"`
}

//商品分类
type GoodsType struct {
	Type string         `json:"type"`
	List []GoodsSubType `json:"list"`
}

//分类中的标签
type GoodsSubType struct {
	Tag    string `json:"tag"`
	Number int64  `json:"number"`
}

//########################################## 商品详情页面结构体和模拟数据 #################################################

//goods data shown in goodsdetail page 🍉
type GoodsDetail struct {
	Headimg  string    `json:"headimg"`
	Userid   string    `json:"userid"`
	Username string    `json:"username"`
	Time     time.Time `json:"time"`
	Title    string    `json:"title"`
	Type     string    `json:"type"`
	Tag      string    `json:"tag"`
	Price    float64   `json:"price"`
	Id       string    `json:"id"`
	Name     string    `json:"name"`
	Visit    int       `json:"visit"`
	Like     int       `json:"like"`
	Talk     int       `json:"talk"`
	Collect  int       `json:"collect"`
	Detail   string    `json:"detail"`
}

type GoodsComment struct { //comment of goods
	Username string    `json:"username"`
	Time     time.Time `json:"time"`
	Comment  string    `json:"comment"`
}

type GoodsPostBody struct {
	GoodId   string `json:"goodid"`
	DataType string `json:"datatype"`
}

//可用于更新点赞数，收藏表，和私信表
type UpdatePostBody struct {
	Tag      string `json:"tag"`
	UserId   string `json:"userid"`
	TargetId string `json:"targetid"`
	StrData  string `json:"strdata"`
	IntData  int    `json:"intdata"`
}

//user state for goods
type UserGoodsState struct {
	Like    bool `json:"like"`
	Collect bool `json:"collect"`
}

//########################################## 个人详情页结构体和模拟数据 #################################################
type PersonalPostBody struct {
	Name string `json:"name"`
	Tag  string `json:"tag"`
}

type UserMessage struct {
	Headimg  string    `json:"headimg"`
	Name     string    `json:"name"`
	Id       string    `json:"id"`
	Sex      string    `json:"sex"`
	Sign     string    `json:"sign"`
	Grade    string    `json:"grade"`
	Colleage string    `json:"colleage"`
	Major    string    `json:"major"`
	Emails   string    `json:"emails"`
	Qq       string    `json:"qq"`
	Phone    string    `json:"phone"`
	Lasttime time.Time `json:"lasttime"`
	Dorm     string    `json:"dorm"`
	Leave    int       `json:"leave"`
	Credits  int       `json:"credits"`
	Rank     int       `json:"rank"`
	Becare   int       `json:"becare"`
	Likes    int       `json:"likes"`
	Visit    int       `json:"visit"`
	Goodsnum int       `json:"goodsnum"`
	Scuess   int       `json:"scuess"`
	Care     int       `json:"care"`
}

type GoodsShort struct {
	Headimg string  `json:"headimg"`
	Id      string  `json:"id"`
	Name    string  `json:"name"`
	Title   string  `json:"title"`
	Price   float64 `json:"price"`
}

type MyMessage struct {
	Uid     string    `json:"uid"` //senderid
	Mid     string    `json:"mid"`
	Time    time.Time `json:"time"`
	Name    string    `json:"name"`
	Title   string    `json:"title"`
	Content string    `json:"content"`
	Headimg string    `json:"headimg"`
}

type User struct {
	Id2     string `json:"id2"`
	Name    string `json:"name"`
	Headimg string `json:"headimg"`
}

type Rank struct {
	Rank   int64  `json:"rank"`
	Name   string `json:"name"`
	Userid string `json:"userid"`
}

type UserShort struct {
	Headimg string `json:"headimg"`
	Name    string `json:"name"`
	Id      string `json:"id"`
}

type UserState struct {
	Like    bool `json:"like"`
	Concern bool `json:"concern"`
}

//########################################## 修改信息页面数据结构和模拟数据 #################################################

//修改数据请求的主体结构
type UpdateBody struct {
	UserId string      `json:"userid"`
	Tag    string      `json:"tag"`
	Data   interface{} `json:"data"`
}

type UpdeteMsg struct {
	Id         string `json:"id"`
	Headimg    string `json:"headimg"`
	UpdataType string `json:"updatatype"`
	Name       string `json:"name"`
	Sex        string `json:"sex"`
	Sign       string `json:"sign"`
	Grade      string `json:"grade"`
	Colleage   string `json:"colleage"`
	Dorm       string `json:"dorm"`
	Major      string `json:"major"`
	Emails     string `json:"emails"`
	Qq         string `json:"qq"`
	Phone      string `json:"phone"`
}

type UpdateResult struct {
	Status   int    `json:"status"`
	Describe string `json:"describe"`
}

func GetUpdateResult(status int, err error) UpdateResult {
	return UpdateResult{Status: status, Describe: fmt.Sprint(err)}
}

//########################################## 上传商品页面数据结构和模拟数据 #################################################

type UpLoadResult struct {
	Status   int    `json:"status"`
	Describe string `json:"describe"`
	ImgUrl   string `json:"imgurl"`
}

type UploadGoodsData struct {
	UserId     string  `json:"userid"`
	Name       string  `json:"name"`
	Title      string  `json:"title"`
	Date       string  `json:"date"`
	Price      float64 `json:"price"`
	Imgurl     string  `json:"imgurl"`
	Type       string  `json:"type"`
	Tag        string  `json:"tag"`
	Usenewtag  bool    `json:"usenewtag"`
	Newtagname string  `json:"newtagname"`
	Text       string  `json:"text"`
}

func CreateUploadRes(status int, err error, imgurl string) UpLoadResult {
	return UpLoadResult{status, fmt.Sprint(err), imgurl}
}

//########################################## 更新个人信息页面数据结构和模拟数据 #################################################

type UserSetData struct {
	Headimg  string `json:"headimg"`
	Name     string `json:"name"`
	Id       string `json:"id"`
	Sex      string `json:"sex"`
	Sign     string `json:"sign"`
	Grade    string `json:"grade"`
	Colleage string `json:"colleage"`
	Major    string `json:"major"`
	Emails   string `json:"emails"`
	Qq       string `json:"qq"`
	Phone    string `json:"phone"`
}

//########################################## 导航栏页面数据结构和模拟数据 #################################################

type MyStatus struct {
	ID         string    `json:"id"`
	Name       string    `json:"name"`
	Headimg    string    `json:"headimg"`
	Leave      int       `json:"leave"`
	Credits    int       `json:"credits"`
	MessageNUm int       `json:"messagenum"`
	GoodsNum   int       `json:"goodsnum"`
	Lasttime   time.Time `json:"lasttime"`
}

type EntranceBody struct {
	UserId string      `json:"userid"`
	Tag    string      `json:"tag"`
	Data   interface{} `json:"data"`
}

type LoginData struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

//注册账号时发来的结构体
type RegisterData struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Code     string `json:"code"`
}

//请求登录，注册，更换验证码时返回的结构
type RequireResult struct {
	Status   int    `json:"status"`
	Describe string `json:"describe"`
}

//#####################

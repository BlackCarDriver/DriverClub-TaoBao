package models

import (
	"fmt"
	"time"
)

//######################################### public protocol ##########################################
//public struct that used to request 🍌 🍉🍔
type RequestProto struct {
	Tag       string      `json:"tag"`
	Api       string      `json:"api"`
	UserId    string      `json:"userid"`
	TargetId  string      `json:"targetid"`
	Token     string      `json:"token"`
	CacheTime int         `json:"cachetime"`
	CacheKey  string      `json:"cachekey"`
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

//########################################## homepage component ################################
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

//used when request a goods list
type PostBody1 struct {
	GoodsType  string `json:"goodstype"`
	GoodsTag   string `json:"goodstag"`
	GoodsIndex int    `json:"goodsindex"`
}

type GoodsType struct {
	Type string         `json:"type"`
	List []GoodsSubType `json:"list"`
}

type GoodsSubType struct {
	Tag    string `json:"tag"`
	Number int64  `json:"number"`
}

//########################################## goodspage component #################################################

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

//########################################## personal component #################################################
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
	Becare   int       `json:"becare"`
	Visit    int       `json:"visit"`
	Goodsnum int       `json:"goodsnum"`
	Scuess   int       `json:"scuess"`
	Care     int       `json:"care"`
	Likes    int       `json:"likes"`
	Rank     int       `json:"rank"`
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
	State   int       `json:"state"`
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

//########################################## chmymsg conpoment #################################################
type UpdateBody struct {
	UserId string      `json:"userid"`
	Tag    string      `json:"tag"`
	Data   interface{} `json:"data"`
}
type UserSetData struct {
	Headimg  string `json:"headimg"`
	Name     string `json:"name"`
	Id       string `json:"id"`
	Sex      string `json:"sex"`
	Sign     string `json:"sign"`
	Grade    string `json:"grade"`
	Colleage string `json:"colleage"`
	Major    string `json:"major"`
	Dorm     string `json:"dorm"`
	Emails   string `json:"emails"`
	Qq       string `json:"qq"`
	Phone    string `json:"phone"`
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

//########################################## upload component #################################################

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

//########################################## naving component #################################################

type MyStatus struct {
	ID         string    `json:"id"`
	Name       string    `json:"name"`
	Headimg    string    `json:"headimg"`
	Leave      int       `json:"leave"`
	Credits    int       `json:"credits"`
	Messagenum int       `json:"messagenum"`
	Goodsnum   int       `json:"goodsnum"`
	Lasttime   time.Time `json:"lasttime"`
}

type EntranceBody struct {
	UserId string      `json:"userid"`
	Tag    string      `json:"tag"`
	Data   interface{} `json:"data"`
}

//used when sign up
type RegisterData struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Code     string `json:"code"`
}

//##################### feedback component #############################
type FeedBackData struct {
	Id        int64     `json:"id"`
	Userid    string    `json:"userid"`
	Email     string    `json:"email"`
	Time      time.Time `json:"time"`
	Status    int64     `json:"status"`
	Type      string    `json:"fbtype"`
	Location  string    `json:"location"` //Where the problem occurred
	Describes string    `json:"describes"`
	Imgurl    string    `json:"imgurl"` //screenshot saving name
}

package models

import "time"

type T_upload struct {
	Time    time.Time `json:"time"`
	Goodsid string    `json:"goodsid"`
	Userid  string    `json:"userid"`
}

type T_user struct {
	Id       string    `json:"id" orm:"pk"`
	Emails   string    `json:"emails"`
	Lasttime time.Time `json:"lasttime"`
	Visit    int32     `json:"visit"`
	Credits  int32     `json:"credits"`
	Sign     string    `json:"sign"`
	Name     string    `json:"name"`
	Dorm     string    `json:"dorm"`
	Leave    int32     `json:"leave"`
	Rank     int32     `json:"rank"`
	Phone    string    `json:"phone"`
	Major    string    `json:"major"`
	Sex      string    `json:"sex"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
	Likes    int32     `json:"likes"`
	Qq       string    `json:"qq"`
	Headimg  string    `json:"headimg"`
}

type T_collect struct {
	Userid  string    `json:"userid"`
	Time    time.Time `json:"time"`
	Goodsid string    `json:"goodsid"`
}

type T_comment struct {
	Userid  string    `json:"userid"`
	Time    time.Time `json:"time"`
	Content string    `json:"content"`
	Goodsid string    `json:"goodsid"`
}

type T_concern struct {
	Id1  string    `json:"id1" orm:"pk"`
	Id2  string    `json:"id2"`
	Time time.Time `json:"time"`
}

type T_goods struct {
	File    string  `json:"file"`
	Tag     string  `json:"tag"`
	Title   string  `json:"title"`
	Name    string  `json:"name"`
	Like    int32   `json:"like"`
	Headimg string  `json:"headimg"`
	Price   float64 `json:"price"`
	Type    string  `json:"type"`
	Visit   int32   `json:"visit"`
	State   int32   `json:"state"`
	Id      string  `json:"id"`
}

type T_message struct {
	Content    string    `json:"content"`
	Senderid   string    `json:"senderid"`
	Receiverid string    `json:"receiverid"`
	Time       time.Time `json:"time"`
	State      int32     `json:"state"`
}

package models

//########################################## 主页结构和模拟数据 ################################
type Goods1 struct {
	Headimg string  `json:"headimg"`
	Userid  string  `json:"userid"`
	Time    string  `json:"time"`
	Title   string  `json:"title"`
	Price   float64 `json:"price"`
	Id      string  `json:"id"`
	Name    string  `json:"name"`
}

//主页获取商品封面数据时提供的信息
type PostBody1 struct {
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
	Tag string `json:"tag"`
	int int64  `json:"int"`
}

//首页封面 模拟数据
var MockGoodsData = []Goods1{
	{"http://www.mycodes.net/upload_files/article/162/1_20190319070316_nu1ok.jpg", "BlackCardriver", "2019-2-10", "adsfasdf阿斯顿发生大发大法第三方", 100.3, "123123", ""},
	{"http://www.mycodes.net/upload_files/article/162/1_20190319070316_nu1ok.jpg", "BlackCardriver", "2019-2-10", "adsfasdf阿斯顿发生大发大法第三方", 120.3, "1231", ""},
	{"http://www.mycodes.net/upload_files/article/162/1_20190319070316_nu1ok.jpg", "BlackCardriver", "2019-2-10", "adsfasdf阿斯顿发生大发大法第三方", 140, "120", ""},
	{"http://www.mycodes.net/upload_files/article/162/1_20190319070316_nu1ok.jpg", "BlackCardriver", "2019-2-10", "adsfasdf阿斯顿发生大发大法第三方", 100.3, "123123", ""},
	{"http://www.mycodes.net/upload_files/article/162/1_20190319070316_nu1ok.jpg", "BlackCardriver", "2019-2-10", "adsfasdf阿斯顿发生大发大法第三方", 120.3, "1231", ""},
	{"http://www.mycodes.net/upload_files/article/162/1_20190319070316_nu1ok.jpg", "BlackCardriver", "2019-2-10", "adsfasdf阿斯顿发生大发大法第三方", 140, "120", ""},
	{"http://www.mycodes.net/upload_files/article/162/1_20190319070316_nu1ok.jpg", "BlackCardriver", "2019-2-10", "adsfasdf阿斯顿发生大发大法第三方", 100.3, "123123", ""},
	{"http://www.mycodes.net/upload_files/article/162/1_20190319070316_nu1ok.jpg", "BlackCardriver", "2019-2-10", "adsfasdf阿斯顿发生大发大法第三方", 120.3, "1231", ""},
}

//首页分类 模拟数据
var MockTypeData = []GoodsType{
	{"stydy", []GoodsSubType{{"book", 111}, {"ruler", 123}, {"cap", 123}}},
	{"sport", []GoodsSubType{{"ball", 333}, {"ruler", 123}, {"cap", 123}}},
	{"live", []GoodsSubType{{"rice", 11}, {"ruler", 123}, {"cap", 123}}},
	{"electrit", []GoodsSubType{{"phone", 1}, {"ruler", 123}, {"cap", 123}}},
	{"handdiv", []GoodsSubType{{"wool", 1}, {"ruler", 123}, {"cap", 123}}},
	{"virutal", []GoodsSubType{{"link", 111}, {"ruler", 123}, {"cap", 123}}},
	{"other", []GoodsSubType{{"water", 111}, {"ruler", 123}, {"cap", 123}}},
}

//########################################## 商品详情页面结构体和模拟数据 #################################################

type GoodsDetail struct {
	Headimg string  `json:"headimg"`
	Userid  string  `json:"userid"`
	Time    string  `json:"time"`
	Title   string  `json:"title"`
	Price   float64 `json:"price"`
	Id      string  `json:"id"`
	Name    string  `json:"name"`
	Visit   int     `json:"visit"`
	Like    int     `json:"like"`
	Talk    int     `json:"talk"`
	Collect int     `json:"collect"`
}

var MockGoodsDetail = GoodsDetail{
	Headimg: "https://tb1.bdstatic.com/tb/r/image/2019-05-22/a5e3c00f38b64d9ff86b2015746e5584.jpg",
	Userid:  "123456",
	Time:    "2019-22-22",
	Price:   33.33,
	Id:      "000043",
	Name:    "九阳电饭煲",
	Visit:   112233,
	Like:    33322,
	Talk:    100,
	Collect: 40,
}

//########################################## 个人详情页结构体和模拟数据 #################################################

type UserMessage struct {
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
	Leave    int    `json:"leave"`
	Credits  int    `json:"credits"`
	Rank     int    `json:"rank"`
	Becare   int    `json:"becare"`
	Like     int    `json:"like"`
	Lasttime int    `json:"lasttime"`
	Visit    int    `json:"visit"`
	Goodsnum int    `json:"goodsnum"`
	Scuess   int    `json:"scuess"`
	Care     int    `json:"care"`
}

type GoodsShort struct {
	Id      string  `json:"id"`
	Headimg string  `json:"headimg"`
	Name    string  `json:"name"`
	Title   string  `json:"title"`
	Price   float64 `json:"price"`
}

type MyMessage struct {
	Senderid   string `json:"senderid"`
	Sendername string `json:"sendername"`
	Content    string `json:"content"`
	Time       string `json:"time"`
}

type User struct {
	Id2     string `json:"id2"`
	Name    string `json:"name"`
	Headimg string `json:"headimg"`
}

//########################################## 修改信息页面数据结构和模拟数据 #################################################

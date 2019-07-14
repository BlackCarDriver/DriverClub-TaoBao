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
	Tag    string `json:"tag"`
	Number int64  `json:"number"`
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

type GoodsPostBody struct {
	GoodId   int    `json:"goodid"`
	DataType string `json:"datatype"`
}

//########################################## 个人详情页结构体和模拟数据 #################################################
type PersonalPostBody struct {
	Name string `json:"name"`
	Tag  string `json:"tag"`
}

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
	Headimg string  `json:"headimg"`
	Id      string  `json:"id"`
	Name    string  `json:"name"`
	Title   string  `json:"title"`
	Price   float64 `json:"price"`
}

//数据库未有title
type MyMessage struct {
	Time    string `json:"time"`
	Name    string `json:"name"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Headimg string `json:"headimg"`
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

//########################################## 修改信息页面数据结构和模拟数据 #################################################

//修改数据请求的主体结构
type UpdateBody struct {
	UserId string      `json:"userid"`
	Tag    string      `json:"tag"`
	Data   interface{} `json:"data"`
}

type UpdeteMsg struct {
	UpdataType string `json:"updatatype"`
	Name       string `json:"name"`
	Sex        string `json:"sex"`
	Sign       string `json:"sign"`
	Grade      string `json:"grade"`
	Colleage   string `json:"colleage"`
	Major      string `json:"major"`
	Emails     string `json:"emails"`
	Qq         string `json:"qq"`
	Phone      string `json:"phone"`
}

type UpdateResult struct {
	Status   int    `json:"status"`
	Describe string `json:"describe"`
}

//########################################## 上传商品页面数据结构和模拟数据 #################################################

type UpLoadResult struct {
	Status   int    `json:"status"`
	Describe string `json:"describe"`
	ImgUrl   string `json:"imgurl"`
}

type UploadGoodsData struct {
	Username   string `json:"username"`
	Title      string `json:"title"`
	Date       string `json:"date"`
	Price      int    `json:"price"`
	Imgurl     string `json:"imgurl"`
	Type       string `json:"type"`
	Tag        string `json:"tag"`
	Usenewtag  bool   `json:"usenewtag"`
	Newtagname string `json:"newtagname"`
	Text       string `json:"text"`
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
	Headimg    string `json:"headimg"`
	Leave      int    `json:"leave"`
	Credits    int    `json:"credits"`
	MessageNUm int    `json:"messagenum"`
	GoodsNum   int    `json:"goodsnum"`
	Lasttime   string `json:"lasttime"`
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

type RegisterData struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Code     string `json:"code"`
}

type RequireResult struct {
	Status   int    `json:"status"`
	Describe string `json:"describe"`
}

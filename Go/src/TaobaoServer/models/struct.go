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

//########################################## 导航栏 #################################################

type MyStatus struct {
	Headimg    string `json:"headimg"`
	Leave      int    `json:"leave"`
	Credits    int    `json:"credits"`
	MessageNUm int    `json:"messagenum"`
	GoodsNum   int    `json:"goodsnum"`
	Lasttime   string `json:"lasttime"`
}

var MockMystatus = MyStatus{
	"https://tb1.bdstatic.com/tb/r/image/2019-05-22/a5e3c00f38b64d9ff86b2015746e5584.jpg",
	122, 123213, 11, 33, "2019-22-22",
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

var MockGoodsMessage = GoodsDetail{
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

var MockUserMessage = UserMessage{
	"https://tb1.bdstatic.com/tb/r/image/2019-05-22/a5e3c00f38b64d9ff86b2015746e5584.jpg",
	"BlackCarDriver", "id123345", "sexboy", "sign:it is sing",
	"grad2015", "collect:计算机学院", "major:计算机", "emils:123123123.com",
	"Qq123213213213", "phone2134213213", 123, 1234, 12, 321, 2134, 4545,
	3453, 56756, 6789, 55,
}

var MockGoodsShort = []GoodsShort{
	{"http://www.mycodes.net/upload_files/article/162/1_20190319070316_nu1ok.jpg", "1234567", "1234567", "hahahahahhahahahaha", 123.123},
	{"http://www.mycodes.net/upload_files/article/162/1_20190319070316_nu1ok.jpg", "1234567", "1234567", "hahahahahhahahahaha", 0.9},
	{"http://www.mycodes.net/upload_files/article/162/1_20190319070316_nu1ok.jpg", "1234567", "1234567", "hahahahahhahahahaha", 14},
	{"http://www.mycodes.net/upload_files/article/162/1_20190319070316_nu1ok.jpg", "1234567", "1234567", "hahahahahhahahahaha", 123},
	{"http://www.mycodes.net/upload_files/article/162/1_20190319070316_nu1ok.jpg", "1234567", "1234567", "hahahahahhahahahaha", 13},
}

var MockMyMessage = []MyMessage{
	{"2019-10-10", "BlackCarDriver", "Hello！", "I will give you ten yuan...", "https://avatar.csdn.net/0/E/6/3_blackcardriver.jpg"},
	{"2019-10-10", "BlackCarDriver", "Hello！", "I will give you ten yuan...", "https://avatar.csdn.net/0/E/6/3_blackcardriver.jpg"},
	{"2019-10-10", "BlackCarDriver", "Hello！", "I will give you ten yuan...", "https://avatar.csdn.net/0/E/6/3_blackcardriver.jpg"},
	{"2019-10-10", "BlackCarDriver", "Hello！", "I will give you ten yuan...", "https://avatar.csdn.net/0/E/6/3_blackcardriver.jpg"},
	{"2019-10-10", "BlackCarDriver", "Hello！", "I will give you ten yuan...", "https://avatar.csdn.net/0/E/6/3_blackcardriver.jpg"},
}

var MockRank = []Rank{
	{1, "Driver", "12322"}, {2, "DDridd", "123222"}, {3, "openid", "123421"},
}

var MockCare = [2][]UserShort{
	{{"https://avatar.csdn.net/0/E/6/3_blackcardriver.jpg", "BlackCarDriver", "123123"},
		{"https://avatar.csdn.net/0/E/6/3_blackcardriver.jpg", "BlackCarDriver", "123123"},
		{"https://avatar.csdn.net/0/E/6/3_blackcardriver.jpg", "BlackCarDriver", "123123"},
		{"https://avatar.csdn.net/0/E/6/3_blackcardriver.jpg", "BlackCarDriver", "123123"},
	},
	{
		{"https://avatar.csdn.net/0/E/6/3_blackcardriver.jpg", "BlackCarDriver", "123123"},
		{"https://avatar.csdn.net/0/E/6/3_blackcardriver.jpg", "BlackCarDriver", "123123"},
		{"https://avatar.csdn.net/0/E/6/3_blackcardriver.jpg", "BlackCarDriver", "123123"},
	},
}

//########################################## 修改信息页面数据结构和模拟数据 #################################################

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

var MockUpdateResult = UpdateResult{
	100, "Scuess!!!",
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

var MockUpLoadResult = UpLoadResult{
	101, "成功上传！", "https://tb1.bdstatic.com/tb/电视剧.jpg",
}

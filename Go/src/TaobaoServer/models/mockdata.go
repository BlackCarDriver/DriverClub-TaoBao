package models

//商品主页数据
var MockGoodsMessage = GoodsDetail{
	Headimg: "https://tb1.bdstatic.com/tb/r/image/2019-05-22/a5e3c00f38b64d9ff86b2015746e5584.jpg",
	Userid:  "4444444",
	Time:    "4444-44-44",
	Price:   44.44,
	Id:      "000043",
	Name:    "错误名字",
	Visit:   44444,
	Like:    44444,
	Tag:     "错误类型",
	Title:   "错误标题",
	Talk:    44444,
	Collect: 44444,
	Detail:  `<div style="color: rgb(212, 212, 212); background-color: rgb(30, 30, 30); font-family: Consolas, &quot;Courier New&quot;, monospace; font-size: 13px; line-height: 18px; white-space: pre;"><div><div style="color: rgb(212, 212, 212); line-height: 18px;"><br><br><br><br><br><br><div style="color: rgb(212, 212, 212); line-height: 18px;"><span style="color: #d4d4d4;">asdfasdf</span></div><br><br><br><br><br><br><br><br><br><br><br><br><br><br><br><br><br><br><br></div></div></div>`,
}

//首页封面 模拟数据
var MockGoodsData = []Goods1{
	{"http://www.mycodes.net/upload_files/article/162/1_20190319070316_nu1ok.jpg", "BlackCardriver", "2019-2-10", "adsfasdf阿斯顿发生大发大法第三方", 100.3, "123123", "", ""},
	{"http://www.mycodes.net/upload_files/article/162/1_20190319070316_nu1ok.jpg", "BlackCardriver", "2019-2-10", "adsfasdf阿斯顿发生大发大法第三方", 120.3, "1231", "", ""},
	{"http://www.mycodes.net/upload_files/article/162/1_20190319070316_nu1ok.jpg", "BlackCardriver", "2019-2-10", "adsfasdf阿斯顿发生大发大法第三方", 140, "120", "", ""},
	{"http://www.mycodes.net/upload_files/article/162/1_20190319070316_nu1ok.jpg", "BlackCardriver", "2019-2-10", "adsfasdf阿斯顿发生大发大法第三方", 100.3, "123123", "", ""},
	{"http://www.mycodes.net/upload_files/article/162/1_20190319070316_nu1ok.jpg", "BlackCardriver", "2019-2-10", "adsfasdf阿斯顿发生大发大法第三方", 120.3, "1231", "", ""},
	{"http://www.mycodes.net/upload_files/article/162/1_20190319070316_nu1ok.jpg", "BlackCardriver", "2019-2-10", "adsfasdf阿斯顿发生大发大法第三方", 140, "120", "", ""},
	{"http://www.mycodes.net/upload_files/article/162/1_20190319070316_nu1ok.jpg", "BlackCardriver", "2019-2-10", "adsfasdf阿斯顿发生大发大法第三方", 100.3, "123123", "", ""},
	{"http://www.mycodes.net/upload_files/article/162/1_20190319070316_nu1ok.jpg", "BlackCardriver", "2019-2-10", "adsfasdf阿斯顿发生大发大法第三方", 120.3, "1231", "", ""},
}

//首页分类 模拟数据
var MockTypeData = []GoodsType{
	{"学习用品", []GoodsSubType{{"book", 111}, {"ruler", 123}, {"cap", 123}}},
	{"体育用品", []GoodsSubType{{"ball", 333}, {"ruler", 123}, {"cap", 123}}},
	{"生活用品", []GoodsSubType{{"rice", 11}, {"ruler", 123}, {"cap", 123}}},
	{"电子产品", []GoodsSubType{{"phone", 1}, {"ruler", 123}, {"cap", 123}}},
	{"手工diy", []GoodsSubType{{"wool", 1}, {"ruler", 123}, {"cap", 123}}},
	{"虚拟商品", []GoodsSubType{{"link", 111}, {"ruler", 123}, {"cap", 123}}},
	{"其他", []GoodsSubType{{"water", 111}, {"ruler", 123}, {"cap", 123}}},
}

//导航栏我的信息下拉框
var MockMystatus = MyStatus{
	"https://tb1.bdstatic.com/tb/r/image/2019-05-22/a5e3c00f38b64d9ff86b2015746e5584.jpg",
	122, 123213, 11, 33, "2019-22-22",
}

//个人主页，用户数据
var MockUserMessage = UserMessage{
	"https://tb1.bdstatic.com/tb/r/image/2019-05-22/a5e3c00f38b64d9ff86b2015746e5584.jpg",
	"BlackCarDriver", "id123345", "sexboy", "sign:it is sing",
	"grad2015", "collect:计算机学院", "major:计算机", "emils:123123123.com",
	"Qq123213213213", "phone2134213213", "123", "dorm", 1234, 12, 321, 2134, 4545,
	3453, 56756, 6789, 55,
}

//个人主页，我的商品
var MockGoodsShort = []GoodsShort{
	{"http://www.mycodes.net/upload_files/article/162/1_20190319070316_nu1ok.jpg", "1234567", "1234567", "hahahahahhahahahaha", 123.123},
	{"http://www.mycodes.net/upload_files/article/162/1_20190319070316_nu1ok.jpg", "1234567", "1234567", "hahahahahhahahahaha", 0.9},
	{"http://www.mycodes.net/upload_files/article/162/1_20190319070316_nu1ok.jpg", "1234567", "1234567", "hahahahahhahahahaha", 14},
	{"http://www.mycodes.net/upload_files/article/162/1_20190319070316_nu1ok.jpg", "1234567", "1234567", "hahahahahhahahahaha", 123},
	{"http://www.mycodes.net/upload_files/article/162/1_20190319070316_nu1ok.jpg", "1234567", "1234567", "hahahahahhahahahaha", 13},
}

//个人主页，我的消息
var MockMyMessage = []MyMessage{
	{"2019-10-10", "BlackCarDriver", "Hello！", "I will give you ten yuan...", "https://avatar.csdn.net/0/E/6/3_blackcardriver.jpg"},
	{"2019-10-10", "BlackCarDriver", "Hello！", "I will give you ten yuan...", "https://avatar.csdn.net/0/E/6/3_blackcardriver.jpg"},
	{"2019-10-10", "BlackCarDriver", "Hello！", "I will give you ten yuan...", "https://avatar.csdn.net/0/E/6/3_blackcardriver.jpg"},
	{"2019-10-10", "BlackCarDriver", "Hello！", "I will give you ten yuan...", "https://avatar.csdn.net/0/E/6/3_blackcardriver.jpg"},
	{"2019-10-10", "BlackCarDriver", "Hello！", "I will give you ten yuan...", "https://avatar.csdn.net/0/E/6/3_blackcardriver.jpg"},
}

//个人主页，用户排名
var MockRank = []Rank{
	{1, "Driver", "12322"}, {2, "DDridd", "123222"}, {3, "openid", "123421"},
}

//个人主页，我关注的和关注我的数据
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

//上传图片返回结果
var MockUpLoadResult = UpLoadResult{
	101, "成功上传！", "https://tb1.bdstatic.com/tb/电视剧.jpg",
}

//更新数据返回结果
var MockUpdateResult = UpdateResult{
	100, "Scuess!!!",
}

//模拟更新个人信息页面获取的已有信息数据
var MockUserSetData = UserSetData{
	"https://tb1.bdstatic.com/tb/电视剧.jpg", "blackcardriver...", "123222", "boy",
	"it is sign used in update my message setting page....", "2019", "computer", "hahahahah",
	"123123123.@122.comn", "123123213", "12354555666",
}

var MockRequireResult = RequireResult{
	1, "成功！",
}

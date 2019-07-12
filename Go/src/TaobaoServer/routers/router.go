package routers

import (
	"TaobaoServer/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/homepage/goodsdata", &controllers.HPGoodsController{})
	beego.Router("/homepage/goodstypemsg", &controllers.GoodsTypeController{})
	beego.Router("/goodsdetail", &controllers.GoodsDetailController{})
	beego.Router("/personal/data", &controllers.PersonalDataController{})
}

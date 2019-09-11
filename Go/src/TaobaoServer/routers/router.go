package routers

import (
	"TaobaoServer/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/test", &controllers.TestController{})
	beego.Router("/homepage/goodsdata", &controllers.HPGoodsController{})
	beego.Router("/homepage/goodstypemsg", &controllers.GoodsTypeController{})
	beego.Router("/goodsdeta", &controllers.GoodsDetailController{})
	beego.Router("/personal/data", &controllers.PersonalDataController{})
	beego.Router("/update", &controllers.UpdataMsgController{})
	beego.Router("/upload/newgoods", &controllers.UploadGoodsController{})
	beego.Router("/upload/images", &controllers.UploadImagesController{})
	beego.Router("/entrance", &controllers.EntranceController{})
	beego.Router("/smallupdate", &controllers.UpdateController{})
	beego.Router("/deleteapi", &controllers.DeleteController{})
}

package routers

import (
	"github.com/astaxie/beego"
	"go_shop/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/test", &controllers.UserController{}, "get:ShowIndex")
}

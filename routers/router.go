package routers

import (
	"https://github.com/188660660/go_shop/tree/master/controllers"

	"github.com/TruthHun/DocHub/controllers/HomeControllers"
	"github.com/TruthHun/DocHub/helper"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"go_shop/controllers"
)
func init(){
	// 判断用户是否登录
	beego.InsertFilter("/goods/*", beego.BeforeExec, filterFunc)

	front()
	back()
}

// 后台路由
func back()  {
	// 后台-登录
	beego.Router("/admin/login", &controllers.AdminUserController{}, "get:ShowIndex")
}

// 前台路由
func front() {
	// 首页路由
	beego.Router("/", &controllers.IndexController{}, "get:ShowIndex")
	beego.Router("/index", &controllers.IndexController{}, "get:ShowIndex")
	beego.Router("/home/*", &controllers.IndexController{}, "get:ShowIndex")

	// 注册路由
	beego.Router("/register", &controllers.UserController{}, "get:ShowRegister;post:HandleRegister")
	// 登录路由
	beego.Router("/login", &controllers.UserController{}, "get:ShowLogin;post:HandleLogin")
	// 激活账号
	beego.Router("/active", &controllers.UserController{}, "get:HandleActivation")

	// 用户中心
	beego.Router("/goods/userCenterInfo", &controllers.UserController{}, "get:ShowUserCenterInfo;post:HandleUserCenterInfo")
	// 全部订单
	beego.Router("/goods/userCenterOrder", &controllers.UserController{}, "get:ShowUserCenterOrder;post:HandleUserCenterOrder")
	// 收货地址
	beego.Router("/goods/userCenterSite", &controllers.UserController{}, "get:ShowUserCenterSite;post:HandleUserCenterSite")

	// 退出登录
	beego.Router("/logout", &controllers.UserController{}, "get:HandleLogout")
}

func filterFunc(ctx *context.Context) {
	// 判断用户是否登录
	userName := ctx.Input.Session("userName")
	if userName == nil {
		ctx.Redirect(302, "/login")
		return
	}
}

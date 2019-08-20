package controller

import "github.com/astaxie/beego"

type AdminUserController struct {
	beego.Controller
}

func (this *AdminUserController) ShowIndex() {

	this.TplName = "admin/index.html"
}

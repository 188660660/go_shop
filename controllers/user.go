package controllers

import "github.com/astaxie/beego"

type UserController struct {
	beego.Controller
}

func (this *UserController) ShowIndex() {
	this.TplName = "index.html"
}

package controllers

import "github.com/astaxie/beego"

type IndexController struct {
	beego.Controller
}

func (this *IndexController) ShowIndex() {
	// 设置登录状态
	uName := this.GetSession("userName")
	if uName == nil {
		this.Data["uName"] = ""
	} else {
		this.Data["uName"] = uName
	}

	this.TplName = "home/index.html"
}

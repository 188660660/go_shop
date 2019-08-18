package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/utils"
	"go_shop/models"
	"regexp"
	"strconv"
)

type UserController struct {
	beego.Controller
}

// 注册-展示
func (this *UserController) ShowRegister() {
	this.TplName = "home/register.html"
}

// 注册-处理
func (this *UserController) HandleRegister() {
	// 1.获取数据
	userName := this.GetString("user_name")
	userPwd := this.GetString("pwd")
	userCpwd := this.GetString("pwd")
	userEmail := this.GetString("email")

	// 2.校验数据
	// a)非空校验
	if userName == "" || userPwd == "" || userCpwd == "" || userEmail == "" {
		this.Data["errmsg"] = "输入数据不完整,请您重新输入！"
		this.TplName = "home/login.html"
		return
	}
	// 正则创建
	reg, err := regexp.Compile(`^[a-zA-Z0-9_-]+@[a-zA-Z0-9_-]+(\.[a-zA-Z0-9_-]+)+$`)
	if err != nil {
		this.Data["errmsg"] = "正则表达式创建失败,请您稍后输入！"
		this.TplName = "home/login.html"
		return
	}
	// b)邮箱校验
	res := reg.MatchString(userEmail)
	if res == false {
		this.Data["errmsg"] = "正则表达式创建失败,请您稍后输入！"
		this.TplName = "home/login.html"
		return
	}
	// c)密码校验
	if userPwd != userCpwd {
		this.Data["errmsg"] = "两次密码输入不一致,请您重新输入！"
		this.TplName = "home/login.html"
		return
	}

	// 3.处理数据
	o := orm.NewOrm()
	var user models.User

	user.UserName = userName
	user.Pwd = userPwd
	user.Email = userEmail

	_, err = o.Insert(&user)
	if err != nil {
		this.Data["errmsg"] = "用户信息写入失败,请重新注册！"
		this.TplName = "home/login.html"
		return
	}
	beego.Info(user.Id)
	// 发送邮件
	config := `{"username":"188660660@qq.com","password":"exgksqkyniphbjaa","host":"smtp.qq.com","port":587}`
	email := utils.NewEMail(config)
	email.From = "188660660@qq.com"
	email.To = []string{userEmail}
	email.Subject = "天天生鲜"
	email.HTML = `<a href="http://192.168.1.108:8083/active?uid=` + strconv.Itoa(user.Id) + `">点击激活</a>`

	// email.Text = "请点击下方激活地址,完成账号激活" | 未生效
	// email.AttachFile("1.jpg") // 附件
	// email.AttachFile("1.jpg", "1") // 内嵌资源

	err = email.Send()
	if err != nil {
		this.Data["errmsg"] = "注册邮件发送失败,请您稍后尝试！"
		this.TplName = "home/login.html"
		return
	}

	// 4.返回数据
	// this.Redirect("home/index",302)
	this.Ctx.WriteString("注册成功,请您稍后查看注册邮箱激活账号！")
}

// 注册-激活
func (this *UserController) HandleActivation() {
	// 获取数据
	uid, err := this.GetInt("uid")
	if err != nil {
		this.Data["errmsg"] = "数据ID不存在！"
		this.TplName = "home/login.html"
		return
	}

	// 操作数据
	var user models.User
	o := orm.NewOrm()
	user.Id = uid
	// user.Active = 1
	err = o.Read(&user)
	if err != nil {
		this.Data["errmsg"] = "非法ID！"
		this.TplName = "home/login.html"
		return
	}
	user.Active = 1
	o.Update(&user)

	// 执行跳转
	this.Redirect("login", 302)
}

// 登录-展示
func (this *UserController) ShowLogin() {
	uName := this.Ctx.GetCookie("uName")
	if uName != "" {
		this.Data["userName"] = uName
		this.Data["remember"] = "checked"
	} else {
		this.Data["userName"] = ""
		this.Data["remember"] = ""
	}
	this.TplName = "home/login.html"
}

// 登录-处理
func (this *UserController) HandleLogin() {
	// 获取数据
	userName := this.GetString("username")
	userPwd := this.GetString("pwd")

	// 记住用户名
	remember := this.GetString("remember")
	// 判断cookie
	if remember == "on" {
		// param->3 指有效时常 默认3600s
		this.Ctx.SetCookie("uName", userName, 3600)
	} else {
		this.Ctx.SetCookie("uName", userName, -1)
	}

	// 校验数据
	if userName == "" || userPwd == "" {
		this.Data["errMsg"] = "用户名或密码不可以为空！"
		this.TplName = "home/login.html"
		return
	}

	// 执行操作
	var user models.User
	o := orm.NewOrm()
	user.UserName = userName
	user.Pwd = userPwd

	// 用户名检查
	err := o.Read(&user, "userName")
	if err != nil {
		this.Data["errMsg"] = "用户名不存在,请检查后输入！"
		this.TplName = "home/login.html"
		return
	}
	// 用户密码检查
	if user.Pwd != userPwd {
		this.Data["errMsg"] = "密码输入不正确,请重新输入！"
		this.TplName = "home/login.html"
		return
	}

	// 添加session | 使用session前 需开启相关配置 否则报错
	this.SetSession("userName", userName)

	// 返回数据
	this.Redirect("/goods/userCenterInfo", 302)
}

// 用户中心-展示
func (this *UserController) ShowUserCenterInfo() {
	// 获取用户名
	uName := this.GetSession("userName")
	this.Data["uName"] = uName

	this.Layout = "home/layout.html"
	this.TplName = "home/user_center_info.html"
}

// 用户中心-操作
func (this *UserController) HandleUserCenterInfo() {

}

// 用户中心-全部订单-展示
func (this *UserController) ShowUserCenterOrder() {
	this.Layout = "home/layout.html"
	this.TplName = "home/user_center_order.html"
}

// 用户中心-全部订单-操作
func (this *UserController) HandleUserCenterOrder() {

}

// 用户中心-收货地址-展示
func (this *UserController) ShowUserCenterSite() {
	this.Layout = "home/layout.html"
	this.TplName = "home/user_center_site.html"
}

// 用户中心-收货地址-操作
func (this *UserController) HandleUserCenterSite() {

}

// 用户-退出登录
func (this *UserController) HandleLogout() {
	// 删除session
	this.DelSession("uName")
	// 跳转页面
	this.Redirect("/", 302)
}

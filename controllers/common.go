package controllers

import (
	"regexp"
)

// 获取session用户名
func GetSessionUserName(this *UserController) string {
	// 设置登录状态
	uName := this.GetSession("userName")
	if uName == nil {
		this.Data["uName"] = ""
	} else {
		this.Data["uName"] = uName
	}
	return uName.(string)
}

// 获取正则校验的电话验证
func GetRegCheckPhone(this *UserController, Phone string) bool {
	// 1.创建正则
	regPhone, err := regexp.Compile(`^1([38][0-9]|4[579]|5[0-3,5-9]|6[6]|7[0135678]|9[89])\d{8}$`)
	if err != nil {
		return false
	}
	// 2.执行校验
	res := regPhone.MatchString(Phone)
	return res
}

// 获取正则校验的邮箱验证
func GetRegCheckEmail(this *UserController, Email string) bool {
	// 1.创建正则
	reg, err := regexp.Compile(`^[a-zA-Z0-9_-]+@[a-zA-Z0-9_-]+(\.[a-zA-Z0-9_-]+)+$`)
	if err != nil {
		return false
	}
	// 2.执行校验
	res := reg.MatchString(Email)
	return res
}

// 获取正则校验的邮政编码
func GetRegCheckZipCode(this *UserController, ZipCode string) bool {
	// 1.创建正则
	reg, err := regexp.Compile(`^[0-9]{6}$`)
	if err != nil {
		return false
	}
	// 2.执行校验
	res := reg.MatchString(ZipCode)
	return res
}

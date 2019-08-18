package main

import (
	"github.com/astaxie/beego"
	_ "go_shop/models" // 注册model层
	_ "go_shop/routers"
)

func main() {
	beego.Run()
}

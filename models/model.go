package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // ★ 手动导入MySQL驱动包
)

// 用户-信息表
type User struct {
	Id       int
	UserName string `orm:"unique;size(100)"` // 用户名
	Pwd      string `orm:"size(100)"`        // 密码
	Email    string // 邮箱
	Power    int    `orm:"default(0)"` // 0标识普通用户 1标识管理员 用户权限
	Active   int    `orm:"default(0)"` // 0标识未激活 1标识已激活 激活状态

	Receivers []*Receiver `orm:"reverse(many)"`
}

// 用户-地址表(关联)
type Receiver struct {
	Id        int
	Name      string // 收件人名字
	ZipCode   string // 收件人邮编
	Addr      string // 收件人地址
	Phone     string // 收件人电话号
	IsDefault bool   `orm:"default(false)"` // 是否默认收件地址

	User *User `orm:"rel(fk)"` // 添加外键约束
}

func init() {
	orm.RegisterDataBase("default", "mysql", "niequn:admin123@tcp(127.0.0.1:3306)/go_shop?charset=utf8") // 注册数据库 | 需提前创建指定数据库
	orm.RegisterModel(new(User), new(Receiver))                                                          // 注册数据表
	orm.RunSyncdb("default", false, true)                                                                // 执行上述操作 param 1.别名 2.强制更新 3.执行过程
}

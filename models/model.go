package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // ★ 手动导入MySQL驱动包
	"time"
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

// 主要放表的设计

type Goods struct { // 商品SPU表
	Id       int
	Name     string      `orm:"size(20)"`  // 商品名称
	Detail   string      `orm:"size(200)"` // 详细描述
	GoodsSKU []*GoodsSKU `orm:"reverse(many)"`
}

type GoodsType struct { // 商品类型表
	Id                   int
	Name                 string                  // 种类名称
	Logo                 string                  // logo
	Image                string                  // 图片
	GoodsSKU             []*GoodsSKU             `orm:"reverse(many)"`
	IndexTypeGoodsBanner []*IndexTypeGoodsBanner `orm:"reverse(many)"`
}

type GoodsSKU struct { // 商品SKU表
	Id                   int
	Goods                *Goods                  `orm:"rel(fk)"` // 商品SPU
	GoodsType            *GoodsType              `orm:"rel(fk)"` // 商品所属种类
	Name                 string                  // 商品名称
	Desc                 string                  // 商品简介
	Price                int                     // 商品价格
	Unite                string                  // 商品单位
	Image                string                  // 商品图片
	Stock                int                     `orm:"default(1)"`   // 商品库存
	Sales                int                     `orm:"default(0)"`   // 商品销量
	Status               int                     `orm:"default(1)"`   // 商品状态
	Time                 time.Time               `orm:"auto_now_add"` // 添加时间
	GoodsImage           []*GoodsImage           `orm:"reverse(many)"`
	IndexGoodsBanner     []*IndexGoodsBanner     `orm:"reverse(many)"`
	IndexTypeGoodsBanner []*IndexTypeGoodsBanner `orm:"reverse(many)"`
}

type GoodsImage struct { // 商品图片表
	Id       int
	Image    string    // 商品图片
	GoodsSKU *GoodsSKU `orm:"rel(fk)"` // 商品SKU
}
type IndexGoodsBanner struct { // 首页轮播商品展示表
	Id       int
	GoodsSKU *GoodsSKU `orm:"rel(fk)"` // 商品sku
	Image    string    // 商品图片
	Index    int       `orm:"default(0)"` // 展示顺序
}

type IndexTypeGoodsBanner struct { // 首页分类商品展示表
	Id          int
	GoodsType   *GoodsType `orm:"rel(fk)"`    // 商品类型
	GoodsSKU    *GoodsSKU  `orm:"rel(fk)"`    // 商品sku
	DisplayType int        `orm:"default(1)"` // 展示类型 0代表文字，1代表图片
	Index       int        `orm:"default(0)"` // 展示顺序
}

type IndexPromotionBanner struct { // 首页促销商品展示表
	Id    int
	Name  string `orm:"size(20)"` // 活动名称
	Url   string `orm:"size(50)"` // 活动链接
	Image string // 活动图片
	Index int    `orm:"default(0)"` // 展示顺序
}

func init() {
	orm.RegisterDataBase("default", "mysql", "niequn:admin123@tcp(127.0.0.1:3306)/go_shop?charset=utf8") // 注册数据库 | 需提前创建指定数据库
	// 注册数据表
	orm.RegisterModel(new(User), new(Receiver), new(GoodsSKU), new(Goods), new(IndexTypeGoodsBanner), new(IndexGoodsBanner), new(GoodsType), new(GoodsImage), new(IndexPromotionBanner))
	orm.RunSyncdb("default", false, true) // 执行上述操作 param 1.别名 2.强制更新 3.执行过程
}

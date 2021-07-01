package main

import (
	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
	_ "godbmonitor/routers"
	"godbmonitor/utils"
)

func AddOne(index int) (index1 int) {
	index1 = index + 1
	return
}

type JSONStruct struct {
	result  bool
	message string
}

func main() {
	beego.BConfig.WebConfig.Session.SessionOn = true

	beego.SetStaticPath("/static", "static")

	orm.RegisterDriver("mysql", orm.DRMySQL)

	orm.RegisterDataBase("default", "mysql", "godbmonitor:godbmonitor@/godbmonitor?charset=utf8")

	orm.Debug = true

	beego.AddFuncMap("AddOne", AddOne)
	beego.AddFuncMap("GetDBType", utils.GetDBType)

	beego.Run()
}

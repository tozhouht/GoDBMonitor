package routers

import (
	"godbmonitor/controllers"
	"strings"

	beego "github.com/beego/beego/v2/server/web"

	context "github.com/beego/beego/v2/server/web/context"
)

var FilterUser = func(ctx *context.Context) {
	ok := ctx.Input.Session("userId")
	if ok == nil && !strings.Contains(ctx.Request.RequestURI, "/login") && !strings.Contains(ctx.Request.RequestURI, "/error") && !strings.Contains(ctx.Request.RequestURI, "/static") {
		ctx.Redirect(302, "/login/index")
	}
}

func init() {
	beego.InsertFilter("/*", beego.BeforeRouter, FilterUser)

	beego.Router("/", &controllers.MainController{})
	beego.Router("/login", &controllers.LoginController{})

	beego.AutoRouter(&controllers.MainController{})
	beego.AutoRouter(&controllers.LoginController{})
}

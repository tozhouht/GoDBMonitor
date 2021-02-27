package controllers

import (
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/session"
	"godbmonitor/models"
)

var globalSessions *session.Manager

func init() {
	sessionConfig := &session.ManagerConfig{
		CookieName:"godbmonitor",
		EnableSetCookie: true,
		Gclifetime:3600,
		Maxlifetime: 3600,
		Secure: false,
		CookieLifeTime: 3600,
		ProviderConfig: "./tmp",
	}
	globalSessions, _ = session.NewManager("memory",sessionConfig)
	go globalSessions.GC()
}

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Index() {
	c.TplName = "login/index.html"
}

func (c *LoginController) Login() {
	userName := c.GetString("userName")
	password := c.GetString("password")

	fmt.Println("userName" + userName)
	fmt.Println("password" + password)

	o := orm.NewOrm()
	var user models.User
	err := o.QueryTable(user).Filter("UserName", userName).One(&user)

	if err == orm.ErrNoRows {
		fmt.Println("无法找到指定用户")
		c.Redirect("/error/404", 302)
	} else if user.Password != password {
		fmt.Println("用户名密码错误")
		c.Redirect("/error/500", 302)
	} else {
		c.SetSession("userId", user.Id)
		c.SetSession("userName", user.UserName)
		c.Redirect("/", 302)
	}
}

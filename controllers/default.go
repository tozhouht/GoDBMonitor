package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Layout = "layout.html"
	c.TplName = "index.html"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["ScriptContent"] = "index-script.html"
}

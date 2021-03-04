package controllers

import (
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
	"godbmonitor/models"
	"math"
)

type CommonController struct {
	beego.Controller
}

func (c *CommonController) Index(){
	c.Layout = "layout.html"
	c.TplName = "common/dataSourceList.html"
}

func (c *CommonController) DataSourceList(){
	count, error := models.DatabaseSourceCount()
	if error != nil{
		count = 1
	}

	pageSize := 15
	pageCount := math.Ceil((float64(count) / float64(pageSize)))

	pageNo, err := c.GetInt("pageNo")
	if err != nil {
		pageNo = 1
	}

	//获取分页数据
	dataSourceList, error2 := models.DatabaseSourceList(pageSize, pageNo)
	if error2 != nil{

	}

	c.Data["DataSourceList"] = dataSourceList
	c.Data["Count"] = count
	c.Data["PageCount"] = pageCount
	c.Data["PageNo"] = pageNo

	c.Layout = "layout.html"
	c.TplName = "common/dataSourceList.html"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["ScriptContent"] = "common/dataSourceList-script.html"
}

func (c *CommonController) AddDataSource(){
	c.Layout = "layout.html"
	c.TplName = "common/addDataSource.html"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["ScriptContent"] = "common/addDataSource-script.html"
}

func (c *CommonController) SaveDataSource(){
	name := c.GetString("name")
	dbType, err := c.GetInt("dbType")
	if err != nil {
		fmt.Print("数据库类型不能为空")
	}
	ip := c.GetString("ip")
	dbUserName := c.GetString("dbUserName")
	dbPassword := c.GetString("dbPassword")
	dbPort, err2 := c.GetInt("dbPort")
	if err2 != nil {
		fmt.Print("数据库端口不能为空")
	}

	o := orm.NewOrm()
	databaseSource := models.DatabaseSource{}
	databaseSource.DbType = dbType
	databaseSource.Ip = ip
	databaseSource.Name = name
	databaseSource.DbUser = dbUserName
	databaseSource.DbPassword = dbPassword
	databaseSource.DbPort = dbPort

	id, err3 := o.Insert(&databaseSource)
	if err3!= nil {
		fmt.Println(err3)
	}else {
		fmt.Println(id)
	}

	c.Redirect("/common/datasourcelist", 302)
}
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

type JSONStruct struct {
	result  bool
	message string
}

func (c *CommonController) Index() {
	c.Layout = "layout.html"
	c.TplName = "common/dataSourceList.html"
}

func (c *CommonController) DataSourceList() {
	count, error := models.DatabaseSourceCount()
	if error != nil {
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
	if error2 != nil {

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

func (c *CommonController) AddDataSource() {
	c.Layout = "layout.html"
	c.TplName = "common/addDataSource.html"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["ScriptContent"] = "common/addDataSource-script.html"
}

func (c *CommonController) SaveDataSource() {
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
	if err3 != nil {
		fmt.Println(err3)
	} else {
		fmt.Println(id)
	}

	c.Redirect("/common/datasourcelist", 302)
}

func (c *CommonController) EditDataSource() {
	id, err := c.GetInt("id")
	if err != nil {
		fmt.Println("获取参数失败")
		return
	}

	o := orm.NewOrm()
	databaseSource := models.DatabaseSource{Id: id}
	err2 := o.Read(&databaseSource)
	if err2 != nil {
		fmt.Println("获取数据源失败")
		return
	}

	c.Data["DatabaseSource"] = databaseSource

	c.Layout = "layout.html"
	c.TplName = "common/editDataSource.html"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["ScriptContent"] = "common/editDataSource-script.html"
}

func (c *CommonController) UpdateDataSource() {
	id, err := c.GetInt("id")
	if err != nil {
		fmt.Println("获取参数失败")
		return
	}
	name := c.GetString("name")
	dbType, err1 := c.GetInt("dbType")
	if err1 != nil {
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
	databaseSource := models.DatabaseSource{Id: id}
	err3 := o.Read(&databaseSource)
	if err3 != nil {
		fmt.Println("获取数据源失败")
		return
	}
	databaseSource.DbType = dbType
	databaseSource.Ip = ip
	databaseSource.Name = name
	databaseSource.DbUser = dbUserName
	databaseSource.DbPassword = dbPassword
	databaseSource.DbPort = dbPort

	_, err4 := o.Update(&databaseSource)

	if err4 != nil {
		fmt.Println(err3)
	}

	c.Redirect("/common/datasourcelist", 302)
}

func (c *CommonController) DeleteDataSource() {
	id, err := c.GetInt("id")
	if err != nil {
		fmt.Println("获取参数失败")
		return
	}

	o := orm.NewOrm()
	if _, err := o.Delete(&models.DatabaseSource{Id: id}); err == nil {
		c.Data["json"] = map[string]interface{}{"result": true, "message": ""}

		c.ServeJSON()
	} else {
		c.Data["json"] = map[string]interface{}{"result": false, "message": "删除错误，" + err.Error()}

		c.ServeJSON()
	}
}

func (c *CommonController) MonitorTargetList() {
	c.Layout = "layout.html"
	c.TplName = "common/monitorTargetList.html"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["ScriptContent"] = "common/monitorTargetList-script.html"
}

func (c *CommonController) AddMonitorTarget() {
	c.Layout = "layout.html"
	c.TplName = "common/addMonitorTarget.html"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["ScriptContent"] = "common/addMonitorTarget-script.html"
}

func (c *CommonController) SaveMonitorTarget(){
	
}

package models

import (
	"fmt"
	"github.com/beego/beego/v2/client/orm"
)

func init() {
	orm.RegisterModelWithPrefix("common_", new(DatabaseSource), new(MonitorTarget))
}

type DatabaseSource struct {
    Id         int
	Name       string `orm:"size(40)"`
	DbType     int    `orm:"type(smallint)"`
	Ip         string `orm:"size(20)"`
	DbUser     string `orm:"size(40)"`
	DbPassword string `orm:"size(40)"`
	DbPort     int    `orm:"type(smallint)"`
}

type MonitorTarget struct {
	Id             int
	Name           string          `orm:"size(100)"`
	DatabaseSource *DatabaseSource `orm:"rel(fk)"`
	MonitorDb      string          `orm:"size(100)"`
}

func DatabaseSourceCount() (int64, error) {
	o := orm.NewOrm()
	var databaseSource DatabaseSource
	count, error := o.QueryTable(databaseSource).Count()
	if error != nil {
		count = 0
	}

	return count, error
}

func DatabaseSourceList(pageSize int, pageNo int) (*[]DatabaseSource, error) {
	o := orm.NewOrm()
	var databaseSources []DatabaseSource
	num, err := o.Raw("SELECT * FROM common_database_source limit ?, ?", (pageNo-1)*pageSize, pageSize).QueryRows(&databaseSources)
	if err == nil {
		fmt.Println("databasesource nums: ", num)
	}
	return &databaseSources, err
}

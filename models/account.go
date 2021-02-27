package models

import (
	"github.com/beego/beego/v2/client/orm"
)

func init() {
	orm.RegisterModelWithPrefix("account_", new(User))
}

type User struct {
	Id       int
	UserName string `orm:"size(40)"`
	Password string `orm:"size(40)"`
	TrueName string `orm:"size(40)"`
	UserType int    `orm:"type(smallint)"`
}

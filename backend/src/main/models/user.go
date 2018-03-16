package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

func init() {
	orm.RegisterModel(new(User))
}

type User struct {
	ID      int `orm:"auto;pk;unique;index;column(id)"`
	Weight  int
	Height  float32
	Created time.Time `orm:"auto_now_add;type(datetime)"`
	Updated time.Time `orm:"auto_now;type(datetime)"`
}

func (u *User) TableName() string {
	return "users"
}

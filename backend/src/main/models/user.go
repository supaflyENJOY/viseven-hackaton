package models

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/astaxie/beego/orm"
)

func init() {
	orm.RegisterModel(new(User))
}

type User struct {
	ID               int `orm:"auto;pk;unique;index;column(id)"`
	Height           int
	Weight           float32
	Name             string
	Avatar           string
	Created          time.Time          `orm:"auto_now_add;type(datetime)"`
	Updated          time.Time          `orm:"auto_now;type(datetime)"`
	WorkoutTemplates []*WorkoutTemplate `orm:"reverse(many)"`
}

func (u *User) TableName() string {
	return "user"
}

func CheckUserRegistration(id int, token string) error {
	o := orm.NewOrm()
	user := User{ID: id}
	err := o.Read(&user)

	if err == orm.ErrNoRows {
		resp, err := http.Get("https://www.googleapis.com/oauth2/v1/userinfo?access_token=" + token)
		if err != nil {
			return err
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		data := make(map[string]interface{})
		err = json.Unmarshal(body, &data)
		if err != nil {
			return err
		}
		user.Name = data["name"].(string)
		user.Avatar = data["picture"].(string)
		_, err = o.Insert(&user)
		if err != nil {
			return err
		}
	}
	return nil
}

package controllers

import (
	"main/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"

	"encoding/json"
)

// Operations about Users
type UserController struct {
	beego.Controller
}

// @router /get [get]
func (u *UserController) Get() {
	uid, ok := IsUserLogin(u.Ctx)
	u.Data["json"] = make(map[string]interface{})
	if !ok {
		(u.Data["json"].(map[string]interface{}))["logged"] = false
	} else {
		(u.Data["json"].(map[string]interface{}))["logged"] = true
		(u.Data["json"].(map[string]interface{}))["id"] = uid
		o := orm.NewOrm()
		user := models.User{ID: uid}
		err := o.Read(&user)
		if err != nil {
			(u.Data["json"].(map[string]interface{}))["error"] = err
		} else {
			(u.Data["json"].(map[string]interface{}))["name"] = user.Name
			(u.Data["json"].(map[string]interface{}))["avatar"] = user.Avatar
			(u.Data["json"].(map[string]interface{}))["weight"] = user.Weight
			(u.Data["json"].(map[string]interface{}))["height"] = user.Height
		}

	}

	u.ServeJSON()
}

// @router /templates/workout [get]
func (u *UserController) WorkoutTemplates() {
	uid, ok := IsUserLogin(u.Ctx)
	u.Data["json"] = make(map[string]interface{})
	if !ok {
		(u.Data["json"].(map[string]interface{}))["error"] = "Not logged in"
	} else {
		o := orm.NewOrm()
		user := models.User{ID: uid}
		err := o.Read(&user)
		if err != nil {
			(u.Data["json"].(map[string]interface{}))["error"] = err
		} else {
			_, err = o.LoadRelated(&user, "WorkoutTemplates")
			if err != nil {
				(u.Data["json"].(map[string]interface{}))["error"] = err
			} else {
				for i := range user.WorkoutTemplates {
					o.LoadRelated(user.WorkoutTemplates[i], "WorkoutExercises")

					for j := range user.WorkoutTemplates[i].WorkoutExercises {
						o.LoadRelated(user.WorkoutTemplates[i].WorkoutExercises[j], "Muscles")

					}
				}
				(u.Data["json"].(map[string]interface{}))["templates"] = user.WorkoutTemplates
			}
		}

	}

	u.ServeJSON()
}

// @router /change [post]
func (u *UserController) Change() {
	//allowedFields := []string{"name", "weight", "height"}

	uid, ok := IsUserLogin(u.Ctx)
	u.Data["json"] = make(map[string]interface{})
	if !ok {
		(u.Data["json"].(map[string]interface{}))["error"] = "Not logged in"
	} else {
		o := orm.NewOrm()
		user := models.User{ID: uid}
		err := o.Read(&user)
		if err != nil {
			(u.Data["json"].(map[string]interface{}))["error"] = err
		} else {
			var ob map[string]interface{}
			json.Unmarshal(u.Ctx.Input.RequestBody, &ob)
			value, ok := ob["name"]
			if ok {
				if result, ok := value.(string); ok {
					user.Name = result
				}
			}
			value, ok = ob["weight"]
			if ok {
				if result, ok := value.(float32); ok {
					user.Weight = result
				}
			}
			value, ok = ob["height"]
			if ok {
				if result, ok := value.(int); ok {
					user.Height = result
				}
			}
			_, err := o.Update(&user)
			if err != nil {
				(u.Data["json"].(map[string]interface{}))["error"] = err
			} else {

				(u.Data["json"].(map[string]interface{}))["result"] = "ok"
			}
		}
	}
	u.ServeJSON()
}

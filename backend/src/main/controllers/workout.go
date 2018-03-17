package controllers

import (
	"main/models"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"

	"encoding/json"
)

// Operations about Users
type WorkoutController struct {
	beego.Controller
}

// @router /get/:id [get]
func (u *WorkoutController) Get() {
	userid, ok := IsUserLogin(u.Ctx)
	u.Data["json"] = make(map[string]interface{})
	if !ok {
		u.Data["json"] = "{\"error\": \"Not logged in\"}"
		u.ServeJSON()
		return
	}
	uidStr := u.Ctx.Input.Param(":id")
	uid64, err := strconv.ParseInt(uidStr, 10, 32)
	if err != nil {
		u.Data["json"] = "{\"error\": \"" + err.Error() + "\"}"
		u.ServeJSON()
		return
	}
	o := orm.NewOrm()
	template := models.WorkoutTemplate{}
	err = o.QueryTable(&template).Filter("id", uid64).RelatedSel("user").One(&template)
	o.LoadRelated(&template, "WorkoutExercises")
	o.LoadRelated(&template, "WorkoutPlans")
	if err != nil {
		u.Data["json"] = "{\"error\": \"" + err.Error() + "\"}"
	} else {
		if template.User.ID != userid {
			(u.Data["json"].(map[string]interface{}))["error"] = "Not allowed"
			u.ServeJSON()
			return
		}
		for i := range template.WorkoutExercises {
			o.LoadRelated(template.WorkoutExercises[i], "Muscles")
		}
		u.Data["json"] = template
	}

	u.ServeJSON()
}

// @router /:id/add/:exercise [get]
func (u *WorkoutController) Add() {
	userid, ok := IsUserLogin(u.Ctx)
	u.Data["json"] = make(map[string]interface{})
	if !ok {
		u.Data["json"] = "{\"error\": \"Not logged in\"}"
		u.ServeJSON()
		return
	}
	uidStr := u.Ctx.Input.Param(":id")
	uid64, err := strconv.ParseInt(uidStr, 10, 32)
	if err != nil {
		u.Data["json"] = "{\"error\": \"" + err.Error() + "\"}"
		u.ServeJSON()
		return
	}
	exidStr := u.Ctx.Input.Param(":exercise")
	exid64, err := strconv.ParseInt(exidStr, 10, 32)
	if err != nil {
		u.Data["json"] = "{\"error\": \"" + err.Error() + "\"}"
		u.ServeJSON()
		return
	}
	o := orm.NewOrm()
	template := models.WorkoutTemplate{}
	err = o.QueryTable(&template).Filter("id", uid64).RelatedSel("user").One(&template)
	if err != nil {
		u.Data["json"] = "{\"error\": \"" + err.Error() + "\"}"
	} else {
		if template.User.ID != userid {
			(u.Data["json"].(map[string]interface{}))["error"] = "Not allowed"
			u.ServeJSON()
			return
		}
		exercise := models.WorkoutExercise{ID: int(exid64)}
		err := o.Read(&exercise)
		if err != nil {
			u.Data["json"] = "{\"error\": \"" + err.Error() + "\"}"
		} else {
			m2m := o.QueryM2M(&template, "WorkoutExercises")
			m2m.Add(exercise)
			o.LoadRelated(&exercise, "Muscles")
			u.Data["json"] = exercise

		}
	}

	u.ServeJSON()
}

// @router /:id/remove/:exercise [get]
func (u *WorkoutController) Remove() {
	userid, ok := IsUserLogin(u.Ctx)
	u.Data["json"] = make(map[string]interface{})
	if !ok {
		u.Data["json"] = "{\"error\": \"Not logged in\"}"
		u.ServeJSON()
		return
	}
	uidStr := u.Ctx.Input.Param(":id")
	uid64, err := strconv.ParseInt(uidStr, 10, 32)
	if err != nil {
		u.Data["json"] = "{\"error\": \"" + err.Error() + "\"}"
		u.ServeJSON()
		return
	}
	exidStr := u.Ctx.Input.Param(":exercise")
	exid64, err := strconv.ParseInt(exidStr, 10, 32)
	if err != nil {
		u.Data["json"] = "{\"error\": \"" + err.Error() + "\"}"
		u.ServeJSON()
		return
	}
	o := orm.NewOrm()
	template := models.WorkoutTemplate{}
	err = o.QueryTable(&template).Filter("id", uid64).RelatedSel("user").One(&template)
	if err != nil {
		u.Data["json"] = "{\"error\": \"" + err.Error() + "\"}"
	} else {
		if template.User.ID != userid {
			(u.Data["json"].(map[string]interface{}))["error"] = "Not allowed"
			u.ServeJSON()
			return
		}
		exercise := models.WorkoutExercise{ID: int(exid64)}
		err := o.Read(&exercise)
		if err != nil {
			u.Data["json"] = "{\"error\": \"" + err.Error() + "\"}"
		} else {
			m2m := o.QueryM2M(&template, "WorkoutExercises")
			m2m.Remove(&exercise)
			u.Data["json"] = "{\"result\": \"ok\"}"

		}
	}

	u.ServeJSON()
}

// @router /create [get]
func (u *WorkoutController) Create() {
	userid, ok := IsUserLogin(u.Ctx)
	if !ok {
		u.Data["json"] = "{\"error\": \"Not logged in\"}"
		u.ServeJSON()
		return
	}

	o := orm.NewOrm()
	user := models.User{ID: userid}
	err := o.Read(&user)
	if err != nil {
		u.Data["json"] = "{\"error\": \"" + err.Error() + "\"}"
		u.ServeJSON()
		return
	}
	template := models.WorkoutTemplate{User: &user, Name: "One more perfect template"}
	_, err = o.Insert(&template)
	if err != nil {
		u.Data["json"] = "{\"error\": \"" + err.Error() + "\"}"
	} else {
		template.User = nil
		u.Data["json"] = template
	}

	u.ServeJSON()
}

// @router /change [post]
func (u *WorkoutController) Change() {
	uid, ok := IsUserLogin(u.Ctx)
	u.Data["json"] = make(map[string]interface{})
	if !ok {
		(u.Data["json"].(map[string]interface{}))["error"] = "Not logged in"
	} else {
		var ob map[string]interface{}
		json.Unmarshal(u.Ctx.Input.RequestBody, &ob)
		o := orm.NewOrm()
		templateId := 0
		value, ok := ob["id"]
		if ok {
			if result, ok := value.(int); ok {
				templateId = result
			} else {
				(u.Data["json"].(map[string]interface{}))["error"] = ":id should be int"
				u.ServeJSON()
				return
			}
		} else {
			(u.Data["json"].(map[string]interface{}))["error"] = "Invalid input :id"
			u.ServeJSON()
			return
		}
		template := models.WorkoutTemplate{}
		err := o.QueryTable(&template).Filter("id", templateId).RelatedSel("user").One(&template)
		if err != nil {
			u.Data["json"] = "{\"error\": \"" + err.Error() + "\"}"
		} else {
			if template.User.ID != uid {
				(u.Data["json"].(map[string]interface{}))["error"] = "Not allowed"
				u.ServeJSON()
				return
			}
			value, ok := ob["name"]
			if ok {
				if result, ok := value.(string); ok {
					template.Name = result
				}
			}
			_, err := o.Update(&template)
			if err != nil {
				u.Data["json"] = "{\"error\": \"" + err.Error() + "\"}"
			} else {

				(u.Data["json"].(map[string]interface{}))["result"] = "ok"
			}
		}
	}
	u.ServeJSON()
}

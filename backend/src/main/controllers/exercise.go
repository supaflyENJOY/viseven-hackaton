package controllers

import (
	"encoding/json"
	"main/models"
	"strings"

	"github.com/astaxie/beego/orm"

	"github.com/astaxie/beego"
)

// Operations about Users
type ExerciseController struct {
	beego.Controller
}

type GetInput struct {
	Input []int `json:"input"`
}

// @router /find [post]
func (u *ExerciseController) Get() {
	_, ok := IsUserLogin(u.Ctx)
	u.Data["json"] = make(map[string]interface{})
	if !ok {
		(u.Data["json"].(map[string]interface{}))["error"] = "You are not logged on!"
	} else {
		var ob GetInput
		err := json.Unmarshal(u.Ctx.Input.RequestBody, &ob)
		if err != nil {
			u.Data["json"] = "{\"error\": \"" + err.Error() + "\"}"
			u.ServeJSON()
			return
		}
		if len(ob.Input) == 0 {
			var exercises []*models.WorkoutExercise
			o := orm.NewOrm()
			_, err := o.QueryTable(new(models.WorkoutExercise)).All(&exercises)
			if err != nil {
				u.Data["json"] = "{\"error\": \"" + err.Error() + "\"}"
				u.ServeJSON()
				return
			}

			u.Data["json"] = exercises
		} else {

			qb, _ := orm.NewQueryBuilder("mysql")

			// Construct query object
			qb.Select("workout_exercise.*").
				From("workout_exercise_muscles").
				LeftJoin("workout_exercise").On("workout_exercise_muscles.workout_exercise_id = workout_exercise.id").
				Where("workout_exercise_muscles.muscle_id IN (" + strings.Repeat(", ?", len(ob.Input))[2:] + ")").GroupBy("workout_exercise_muscles.workout_exercise_id").Having("COUNT(workout_exercise_muscles.workout_exercise_id) = ?")

			// export raw query string from QueryBuilder object
			sql := qb.String()

			var exercises []*models.WorkoutExercise
			o := orm.NewOrm()
			o.Raw(sql, ob.Input, len(ob.Input)).QueryRows(&exercises)

			u.Data["json"] = exercises
		}
	}

	u.ServeJSON()
}

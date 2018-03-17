package models

import (
	"github.com/astaxie/beego/orm"
)

func init() {
	orm.RegisterModel(new(WorkoutTemplate))
}

type WorkoutTemplate struct {
	ID               int   `orm:"auto;pk;unique;index;column(id)"`
	User             *User `orm:"null;rel(fk)"`
	Name             string
	WorkoutExercises []*WorkoutExercise `orm:"rel(m2m)"`
	WorkoutPlans     []*WorkoutPlan     `orm:"reverse(many)"`
}

func (u *WorkoutTemplate) TableName() string {
	return "workout_template"
}

package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

func init() {
	orm.RegisterModel(new(WorkoutPlan))
}

type WorkoutPlan struct {
	ID               int `orm:"auto;pk;unique;index;column(id)"`
	Date             time.Time
	WorkoutTemplates *WorkoutTemplate `orm:"rel(fk)"`
}

func (u *WorkoutPlan) TableName() string {
	return "workout_plan"
}

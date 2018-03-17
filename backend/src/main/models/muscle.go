package models

import (
	"github.com/astaxie/beego/orm"
)

func init() {
	orm.RegisterModel(new(Muscle))
}

type Muscle struct {
	ID               int `orm:"auto;pk;unique;index;column(id)"`
	Name             string
	WorkoutExercises []*WorkoutExercise `orm:"reverse(many)"`
}

func (u *Muscle) TableName() string {
	return "muscle"
}

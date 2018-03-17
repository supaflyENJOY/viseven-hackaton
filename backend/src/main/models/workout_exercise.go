package models

import (
	"github.com/astaxie/beego/orm"
)

func init() {
	orm.RegisterModel(new(WorkoutExercise))
}

type WorkoutExercise struct {
	ID          int `orm:"auto;pk;unique;index;column(id)"`
	Title       string
	Description string
	Image       string
	Muscles     []*Muscle `orm:"rel(m2m)"`
	WorkoutTemplates []*WorkoutTemplate `orm:"reverse(many)"`
}

func (u *WorkoutExercise) TableName() string {
	return "workout_exercise"
}

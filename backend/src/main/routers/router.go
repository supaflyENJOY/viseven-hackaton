package routers

import (
	"main/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),
		beego.NSNamespace("/login",
			beego.NSInclude(
				&controllers.AuthController{},
			),
		),
		beego.NSNamespace("/workout",
			beego.NSInclude(
				&controllers.WorkoutController{},
			),
		),
		beego.NSNamespace("/exercise",
			beego.NSInclude(
				&controllers.ExerciseController{},
			),
		),
	)
	beego.AddNamespace(ns)
}

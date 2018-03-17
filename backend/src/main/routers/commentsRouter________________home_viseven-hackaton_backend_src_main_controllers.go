package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["main/controllers:AuthController"] = append(beego.GlobalControllerRouter["main/controllers:AuthController"],
		beego.ControllerComments{
			Method: "Login",
			Router: `/available`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["main/controllers:ExerciseController"] = append(beego.GlobalControllerRouter["main/controllers:ExerciseController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/find`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["main/controllers:UserController"] = append(beego.GlobalControllerRouter["main/controllers:UserController"],
		beego.ControllerComments{
			Method: "Change",
			Router: `/change`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["main/controllers:UserController"] = append(beego.GlobalControllerRouter["main/controllers:UserController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/get`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["main/controllers:UserController"] = append(beego.GlobalControllerRouter["main/controllers:UserController"],
		beego.ControllerComments{
			Method: "WorkoutTemplates",
			Router: `/templates/workout`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["main/controllers:WorkoutController"] = append(beego.GlobalControllerRouter["main/controllers:WorkoutController"],
		beego.ControllerComments{
			Method: "Add",
			Router: `/:id/add/:exercise`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["main/controllers:WorkoutController"] = append(beego.GlobalControllerRouter["main/controllers:WorkoutController"],
		beego.ControllerComments{
			Method: "Remove",
			Router: `/:id/remove/:exercise`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["main/controllers:WorkoutController"] = append(beego.GlobalControllerRouter["main/controllers:WorkoutController"],
		beego.ControllerComments{
			Method: "Change",
			Router: `/change`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["main/controllers:WorkoutController"] = append(beego.GlobalControllerRouter["main/controllers:WorkoutController"],
		beego.ControllerComments{
			Method: "Create",
			Router: `/create`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["main/controllers:WorkoutController"] = append(beego.GlobalControllerRouter["main/controllers:WorkoutController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/get/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

}

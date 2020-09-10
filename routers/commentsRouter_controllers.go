package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	// group
	beego.GlobalControllerRouter["dbmanager/controllers:GroupController"] = append(beego.GlobalControllerRouter["dbmanager/controllers:GroupController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["dbmanager/controllers:GroupController"] = append(beego.GlobalControllerRouter["dbmanager/controllers:GroupController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["dbmanager/controllers:GroupController"] = append(beego.GlobalControllerRouter["dbmanager/controllers:GroupController"],
		beego.ControllerComments{
			Method:           "Get",
			Router:           "/:name",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["dbmanager/controllers:GroupController"] = append(beego.GlobalControllerRouter["dbmanager/controllers:GroupController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           "/:name",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["dbmanager/controllers:GroupController"] = append(beego.GlobalControllerRouter["dbmanager/controllers:GroupController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           "/:name",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	// object
	beego.GlobalControllerRouter["dbmanager/controllers:ObjectController"] = append(beego.GlobalControllerRouter["dbmanager/controllers:ObjectController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["dbmanager/controllers:ObjectController"] = append(beego.GlobalControllerRouter["dbmanager/controllers:ObjectController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["dbmanager/controllers:ObjectController"] = append(beego.GlobalControllerRouter["dbmanager/controllers:ObjectController"],
		beego.ControllerComments{
			Method:           "Get",
			Router:           "/:objectId",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["dbmanager/controllers:ObjectController"] = append(beego.GlobalControllerRouter["dbmanager/controllers:ObjectController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           "/:objectId",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["dbmanager/controllers:ObjectController"] = append(beego.GlobalControllerRouter["dbmanager/controllers:ObjectController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           "/:objectId",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["dbmanager/controllers:UserController"] = append(beego.GlobalControllerRouter["dbmanager/controllers:UserController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["dbmanager/controllers:UserController"] = append(beego.GlobalControllerRouter["dbmanager/controllers:UserController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["dbmanager/controllers:UserController"] = append(beego.GlobalControllerRouter["dbmanager/controllers:UserController"],
		beego.ControllerComments{
			Method:           "Get",
			Router:           "/:uid",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["dbmanager/controllers:UserController"] = append(beego.GlobalControllerRouter["dbmanager/controllers:UserController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           "/:uid",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["dbmanager/controllers:UserController"] = append(beego.GlobalControllerRouter["dbmanager/controllers:UserController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           "/:uid",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["dbmanager/controllers:UserController"] = append(beego.GlobalControllerRouter["dbmanager/controllers:UserController"],
		beego.ControllerComments{
			Method:           "Login",
			Router:           "/login",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["dbmanager/controllers:UserController"] = append(beego.GlobalControllerRouter["dbmanager/controllers:UserController"],
		beego.ControllerComments{
			Method:           "Logout",
			Router:           "/logout",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

}

package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {
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
			Router:           "/:name",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["dbmanager/controllers:UserController"] = append(beego.GlobalControllerRouter["dbmanager/controllers:UserController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           "/:name",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["dbmanager/controllers:UserController"] = append(beego.GlobalControllerRouter["dbmanager/controllers:UserController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           "/:name",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})
}

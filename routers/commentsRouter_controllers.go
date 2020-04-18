package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["phonebook/controllers:PhonebookController"] = append(beego.GlobalControllerRouter["phonebook/controllers:PhonebookController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["phonebook/controllers:PhonebookController"] = append(beego.GlobalControllerRouter["phonebook/controllers:PhonebookController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["phonebook/controllers:PhonebookController"] = append(beego.GlobalControllerRouter["phonebook/controllers:PhonebookController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/:phonebookId`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["phonebook/controllers:PhonebookController"] = append(beego.GlobalControllerRouter["phonebook/controllers:PhonebookController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:phonebookId`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["phonebook/controllers:PhonebookController"] = append(beego.GlobalControllerRouter["phonebook/controllers:PhonebookController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:phonebookId`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["phonebook/controllers:UserController"] = append(beego.GlobalControllerRouter["phonebook/controllers:UserController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["phonebook/controllers:UserController"] = append(beego.GlobalControllerRouter["phonebook/controllers:UserController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["phonebook/controllers:UserController"] = append(beego.GlobalControllerRouter["phonebook/controllers:UserController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/:uid`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["phonebook/controllers:UserController"] = append(beego.GlobalControllerRouter["phonebook/controllers:UserController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:uid`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["phonebook/controllers:UserController"] = append(beego.GlobalControllerRouter["phonebook/controllers:UserController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:uid`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["phonebook/controllers:UserController"] = append(beego.GlobalControllerRouter["phonebook/controllers:UserController"],
        beego.ControllerComments{
            Method: "Login",
            Router: `/login`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["phonebook/controllers:UserController"] = append(beego.GlobalControllerRouter["phonebook/controllers:UserController"],
        beego.ControllerComments{
            Method: "Logout",
            Router: `/logout`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}

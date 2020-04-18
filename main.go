package main

import (
	_ "phonebook/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/auth"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	authPlugin := auth.NewBasicAuthenticator(SecretAuth, "Authorization Required")
	beego.InsertFilter("*", beego.BeforeRouter,authPlugin)
	beego.Run()
}

func SecretAuth(username, password string) bool {
	return username == "astaxie" && password == "helloBeego"
}

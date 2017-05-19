package main

import (
	_ "system_service/email_service/routers"

	"github.com/astaxie/beego"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	var cfg = beego.AppConfig
	systemEmailUser := cfg.String("system_email_user")
	systemEmailHost := cfg.String("system_email_host")
	beego.Debug("system email user:", systemEmailUser)
	beego.Debug("system email host:", systemEmailHost)

	beego.Run()
}

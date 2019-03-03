package main

import (
	"os"

	"github.com/louisevanderlith/mango"
	"github.com/louisevanderlith/mango/enums"
	"github.com/louisevanderlith/router/routers"

	"github.com/astaxie/beego"
)

func main() {
	mode := os.Getenv("RUNMODE")

	// Register with router
	appName := beego.BConfig.AppName
	srv := mango.NewService(mode, appName, enums.API)

	routers.Setup(srv)
	beego.Run()
}

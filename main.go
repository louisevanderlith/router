package main

import (
	"os"
	"path"

	"github.com/louisevanderlith/router/logic"

	"github.com/louisevanderlith/mango"
	"github.com/louisevanderlith/mango/enums"
	"github.com/louisevanderlith/router/routers"

	"github.com/astaxie/beego"
)

func main() {
	mode := os.Getenv("RUNMODE")
	keyPath := os.Getenv("KEYPATH")
	pubName := os.Getenv("PUBLICKEY")
	host := os.Getenv("HOST")
	pubPath := path.Join(keyPath, pubName)

	// Register with router
	appName := beego.BConfig.AppName
	srv := mango.NewService(mode, appName, pubPath, enums.API)

	//Doesn't have to make a request to register, but it still needs to Register
	_, err := logic.AddService(srv)

	if err != nil {
		panic(err)
	}

	routers.Setup(srv, host)
	beego.Run()
}

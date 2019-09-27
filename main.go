package main

import (
	"os"
	"path"
	"strconv"

	"github.com/louisevanderlith/droxolite"
	"github.com/louisevanderlith/droxolite/bodies"
	"github.com/louisevanderlith/droxolite/element"
	"github.com/louisevanderlith/droxolite/resins"
	"github.com/louisevanderlith/droxolite/servicetype"
	"github.com/louisevanderlith/router/logic"
	"github.com/louisevanderlith/router/routers"
)

func main() {
	keyPath := os.Getenv("KEYPATH")
	pubName := os.Getenv("PUBLICKEY")
	host := os.Getenv("HOST")
	profile := os.Getenv("PROFILE")
	httpport, _ := strconv.Atoi(os.Getenv("HTTPPORT"))
	appName := os.Getenv("APPNAME")
	pubPath := path.Join(keyPath, pubName)

	// Register with router
	srv := bodies.NewService(appName, profile, pubPath, host, httpport, servicetype.API)

	//Doesn't have to make a request to register, but it still needs to Register
	_, err := logic.AddService(srv)

	if err != nil {
		panic(err)
	}

	poxy := resins.NewMonoEpoxy(srv, element.GetNoTheme(host, srv.ID, profile))
	routers.Setup(poxy)
	poxy.EnableCORS(host)

	err = droxolite.Boot(poxy)

	if err != nil {
		panic(err)
	}
}

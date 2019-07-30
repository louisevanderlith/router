package main

import (
	"os"
	"path"

	"github.com/louisevanderlith/droxolite"
	"github.com/louisevanderlith/droxolite/servicetype"
	"github.com/louisevanderlith/router/logic"
	"github.com/louisevanderlith/router/routers"
)

func main() {
	keyPath := os.Getenv("KEYPATH")
	pubName := os.Getenv("PUBLICKEY")
	//host := os.Getenv("HOST")
	pubPath := path.Join(keyPath, pubName)

	conf, err := droxolite.LoadConfig()

	if err != nil {
		panic(err)
	}

	// Register with router
	srv := droxolite.NewService(conf.Appname, pubPath, conf.HTTPPort, servicetype.API)

	//Doesn't have to make a request to register, but it still needs to Register
	_, err = logic.AddService(srv)

	if err != nil {
		panic(err)
	}

	poxy := droxolite.NewEpoxy(srv)
	routers.Setup(poxy)

	err = poxy.Boot()

	if err != nil {
		panic(err)
	}
}

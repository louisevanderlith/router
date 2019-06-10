package routers

import (
	"fmt"
	"strings"

	"github.com/louisevanderlith/mango"
	"github.com/louisevanderlith/mango/control"
	"github.com/louisevanderlith/router/controllers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
	secure "github.com/louisevanderlith/secure/core"
	"github.com/louisevanderlith/secure/core/roletype"
)

//Setup creates the routing paths and attaches security filters
func Setup(s *mango.Service, host string) {
	ctrlmap := EnableFilter(s, host)

	discoCtrl := controllers.NewDiscoveryCtrl(ctrlmap)

	beego.Router("/v1/discovery", discoCtrl, "post:Post")
	beego.Router("/v1/discovery/:appID/:serviceName", discoCtrl, "get:GetDirty")
	beego.Router("/v1/discovery/:appID/:serviceName/:clean", discoCtrl, "get:Get")

	memCtrl := controllers.NewMemoryCtrl(ctrlmap)
	beego.Router("/v1/memory", memCtrl, "get:Get")
	beego.Router("/v1/memory/apps", memCtrl, "get:GetApps")
}

//EnableFilter returns a ControllerMap which holds path Role requirements
func EnableFilter(s *mango.Service, host string) *control.ControllerMap {
	ctrlmap := control.CreateControlMap(s)

	emptyMap := make(secure.ActionMap)
	ctrlmap.Add("/v1/discovery", emptyMap)

	userMap := make(secure.ActionMap)
	userMap["GET"] = roletype.Admin
	ctrlmap.Add("/v1/memory", userMap)

	beego.InsertFilter("/v1/memory", beego.BeforeRouter, ctrlmap.FilterAPI, false)
	allowed := fmt.Sprintf("https://*%s", strings.TrimSuffix(host, "/"))

	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowOrigins: []string{allowed},
		AllowMethods: []string{"GET", "POST", "OPTIONS"},
	}), false)

	return ctrlmap
}

package routers

import (
	"github.com/louisevanderlith/droxolite"
	"github.com/louisevanderlith/router/controllers"

	"github.com/louisevanderlith/droxolite/roletype"
)

//Setup creates the routing paths and attaches security filters
func Setup(poxy *droxolite.Epoxy) {
	//Discovery
	discoCtrl := &controllers.DiscoveryController{}
	discoGroup := droxolite.NewRouteGroup("discovery", discoCtrl)
	discoGroup.AddRoute("/", "POST", roletype.Unknown, discoCtrl.Post)
	discoGroup.AddRoute("/{appID}/{serviceName:[a-zA-Z.]+}", "GET", roletype.Unknown, discoCtrl.GetDirty)
	discoGroup.AddRoute("/{appID}/{serviceName:[a-zA-Z.]+}/{clean:true|false}", "GET", roletype.Unknown, discoCtrl.Get)
	poxy.AddGroup(discoGroup)

	//Memory
	memCtrl := &controllers.MemoryController{}
	memGroup := droxolite.NewRouteGroup("memory", memCtrl)
	memGroup.AddRoute("/", "GET", roletype.Admin, memCtrl.Get)
	memGroup.AddRoute("/apps", "GET", roletype.Admin, memCtrl.GetApps)
	poxy.AddGroup(memGroup)
}

/*
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
*/

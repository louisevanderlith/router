package routers

import (
	"github.com/louisevanderlith/router/controllers"

	"github.com/louisevanderlith/droxolite/mix"
	"github.com/louisevanderlith/droxolite/resins"
	"github.com/louisevanderlith/droxolite/roletype"
	"github.com/louisevanderlith/droxolite/routing"
)

//Setup creates the routing paths and attaches security filters
func Setup(e resins.Epoxi) {
	//Discovery
	discoCtrl := &controllers.Discovery{}
	discoGroup := routing.NewRouteGroup("discovery", mix.JSON)
	discoGroup.AddRoute("Register API", "/", "POST", roletype.Unknown, discoCtrl.Post)
	discoGroup.AddRoute("Get URL Dirty", "/{appID}/{serviceName:[a-zA-Z.]+}", "GET", roletype.Unknown, discoCtrl.GetDirty)
	discoGroup.AddRoute("Get URL Clean", "/{appID}/{serviceName:[a-zA-Z.]+}/{clean:true|false}", "GET", roletype.Unknown, discoCtrl.Get)
	e.AddGroup(discoGroup)

	//Memory
	memCtrl := &controllers.Memory{}
	memGroup := routing.NewRouteGroup("memory", mix.JSON)
	memGroup.AddRoute("Get Registered Services", "/", "GET", roletype.Admin, memCtrl.Get)
	memGroup.AddRoute("Get Registered Services Names", "/apps", "GET", roletype.Admin, memCtrl.GetApps)
	e.AddGroup(memGroup)
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

package routers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/louisevanderlith/router/controllers"

	"github.com/louisevanderlith/droxolite/mix"
	"github.com/louisevanderlith/droxolite/resins"
	"github.com/louisevanderlith/droxolite/roletype"
)

//Setup creates the routing paths and attaches security filters
func Setup(e resins.Epoxi) {
	routr := e.Router().(*mux.Router)

	discoCtrl := &controllers.Discovery{}
	e.JoinBundle("/", roletype.Unknown, mix.JSON, discoCtrl)
	e.JoinPath(routr, "/discovery/{appID}/{serviceName:[a-zA-Z.]+}", "Get URL Dirty", http.MethodGet, roletype.Unknown, mix.JSON, discoCtrl.GetDirty)
	e.JoinPath(routr, "/discovery/{appID}/{serviceName:[a-zA-Z.]+}/{clean:true|false}", "Get URL", http.MethodGet, roletype.Unknown, mix.JSON, discoCtrl.Get)

	//Memory
	memCtrl := &controllers.Memory{}
	e.JoinBundle("/", roletype.Admin, mix.JSON, memCtrl)
	e.JoinPath(routr, "/memory/apps", "Get Registered Services Names", http.MethodGet, roletype.Admin, mix.JSON, memCtrl.GetApps)

	appl := &controllers.Applicants{}
	e.JoinPath(routr, "/applicants/{profile:[a-zA-Z]+}", "Get Registered Applications for the Proxy", http.MethodGet, roletype.Unknown, mix.JSON, appl.Get)
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

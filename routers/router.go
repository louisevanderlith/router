// @APIVersion 1.0.0
// @Title Router API
// @Description API for the Router
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/louisevanderlith/mango"
	"github.com/louisevanderlith/mango/control"
	"github.com/louisevanderlith/router/controllers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
)

func Setup(s *mango.Service) {
	ctrlmap := EnableFilter(s)

	discoCtrl := controllers.NewDiscoveryCtrl(ctrlmap)

	beego.Router("/v1/discovery", discoCtrl, "post:Post")
	beego.Router("/v1/discovery/:appID/:serviceName", discoCtrl, "get:GetDirty")
	beego.Router("/v1/discovery/:appID/:serviceName/:clean", discoCtrl, "get:Get")
	beego.Router("/v1/memory", controllers.NewMemoryCtrl(ctrlmap))
}

func EnableFilter(s *mango.Service) *control.ControllerMap {
	ctrlmap := control.CreateControlMap(s)

	emptyMap := make(control.ActionMap)

	ctrlmap.Add("/discovery", emptyMap)
	ctrlmap.Add("/memory", emptyMap)

	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:    []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Content-Type"},
		ExposeHeaders:   []string{"Content-Length", "Access-Control-Allow-Origin"},
	}))

	return ctrlmap
}

package controllers

import (
	"net/http"

	"github.com/louisevanderlith/mango/control"
	"github.com/louisevanderlith/router/logic"
)

type MemoryController struct {
	control.APIController
}

func NewMemoryCtrl(ctrlMap *control.ControllerMap) *MemoryController {
	result := &MemoryController{}
	result.SetInstanceMap(ctrlMap)

	return result
}

// @Title GetRegistered Services
// @Description Gets the serrvices registered
// @Success 200 {string} models.Service.ID
// @router / [get]
func (req *MemoryController) Get() {
	srvMap := logic.GetServiceMap()
	req.Serve(http.StatusOK, nil, srvMap)
}

// @Title GetApp Names
// @Description Gets the Names of services registered
// @Success 200 {string} models.Service.ID
// @router /apps [get]
func (req *MemoryController) GetApps() {
	srvMap := logic.GetServiceMap()

	result := make(map[string]struct{})
	for name := range srvMap {
		result[name] = struct{}{}
	}

	req.Serve(http.StatusOK, nil, result)
}

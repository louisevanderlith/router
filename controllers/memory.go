package controllers

import (
	"net/http"

	"github.com/louisevanderlith/droxolite/context"
	"github.com/louisevanderlith/router/logic"
)

type Memory struct {
}

// @Title GetRegistered Services
// @Description Gets the serrvices registered
// @Success 200 {string} models.Service.ID
// @router / [get]
func (req *Memory) Get(ctx context.Requester) (int, interface{}) {
	srvMap := logic.GetServiceMap()
	return http.StatusOK, srvMap
}

// @Title GetApp Names
// @Description Gets the Names of services registered
// @Success 200 {string} models.Service.ID
// @router /apps [get]
func (req *Memory) GetApps(ctx context.Requester) (int, interface{}) {
	srvMap := logic.GetServiceMap()

	result := make(map[string]struct{})
	for name := range srvMap {
		result[name] = struct{}{}
	}

	return http.StatusOK, result
}

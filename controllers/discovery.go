package controllers

import (
	"errors"
	"net/http"

	"strconv"

	"github.com/louisevanderlith/droxolite"
	"github.com/louisevanderlith/droxolite/xontrols"
	"github.com/louisevanderlith/router/logic"
)

type DiscoveryController struct {
	xontrols.APICtrl
}

// @Title RegisterAPI
// @Description Register an API
// @Success 200 {string} models.Service.ID
// @Failure 403 body is empty
// @router / [post]
func (req *DiscoveryController) Post() {
	service := &droxolite.Service{}
	req.Ctx.Body(service)

	appID, err := logic.AddService(service)

	if err != nil {
		req.Serve(http.StatusInternalServerError, err, nil)
	}

	req.Serve(http.StatusOK, nil, appID)
}

// @Title GetService
// @Description Gets the recommended service
// @Param	appID			path	string 	true		"the application requesting a service"
// @Param	serviceName		path 	string	true		"the name of the service you want to get"
// @Param	clean			path 	bool	false		"clean will return a user friendly URL and not the application's actual URL"
// @Success 200 {string} Service.URL
// @Failure 403 :serviceName or :appID is empty
// @router /:appID/:serviceName/:clean [get]
func (req *DiscoveryController) Get() {
	appID := req.Ctx.FindParam("appID")
	serviceName := req.Ctx.FindParam("serviceName")

	if appID == "" || serviceName == "" {
		err := errors.New("appID AND serviceName must be populated")
		req.Serve(http.StatusBadRequest, err, nil)
		return
	}

	clean, cleanErr := strconv.ParseBool(req.Ctx.FindParam("clean"))

	if cleanErr != nil {
		clean = false
	}

	url, err := logic.GetServicePath(serviceName, appID, clean)

	if err != nil {
		req.Serve(http.StatusInternalServerError, err, nil)
		return
	}

	req.Serve(http.StatusOK, nil, url)
}

// @Title GetDirtyService
// @Description Gets the recommended service
// @Param	appID			path	string 	true		"the application requesting a service"
// @Param	serviceName		path 	string	true		"the name of the service you want to get"
// @Success 200 {string} Service.URL
// @Failure 403 :serviceName or :appID is empty
// @router /:appID/:serviceName [get]
func (req *DiscoveryController) GetDirty() {
	appID := req.Ctx.FindParam("appID")
	serviceName := req.Ctx.FindParam("serviceName")

	if appID == "" || serviceName == "" {
		err := errors.New("appID AND serviceName must be populated")
		req.Serve(http.StatusBadRequest, err, nil)
		return
	}

	url, err := logic.GetServicePath(serviceName, appID, false)

	if err != nil {
		req.Serve(http.StatusInternalServerError, err, nil)
		return
	}

	req.Serve(http.StatusOK, nil, url)
}

package controllers

import (
	"errors"
	"log"
	"net/http"
	"os"

	"strconv"

	"github.com/louisevanderlith/droxolite/bodies"
	"github.com/louisevanderlith/droxolite/context"
	"github.com/louisevanderlith/router/logic"
)

type Discovery struct {
}

// @Title RegisterAPI
// @Description Register an API
// @Success 200 {string} models.Service.ID
// @Failure 403 body is empty
// @router / [post]
func (req *Discovery) Create(ctx context.Requester) (int, interface{}) {
	service := &bodies.Service{}
	ctx.Body(service)

	appID, err := logic.AddService(service)

	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, appID
}

// @Title GetService
// @Description Gets the recommended service
// @Param	appID			path	string 	true		"the application requesting a service"
// @Param	serviceName		path 	string	true		"the name of the service you want to get"
// @Param	clean			path 	bool	false		"clean will return a user friendly URL and not the application's actual URL"
// @Success 200 {string} Service.URL
// @Failure 403 :serviceName or :appID is empty
// @router /:appID/:serviceName/:clean [get]
func (req *Discovery) Get(ctx context.Requester) (int, interface{}) {
	appID := ctx.FindParam("appID")
	serviceName := ctx.FindParam("serviceName")

	if appID == "" || serviceName == "" {
		err := errors.New("appID AND serviceName must be populated")
		return http.StatusBadRequest, err
	}

	clean, cleanErr := strconv.ParseBool(ctx.FindParam("clean"))

	if cleanErr != nil {
		clean = false
	}

	url, err := logic.GetServicePath(serviceName, appID, os.Getenv("HOST"), clean)

	if err != nil {
		log.Println(err)
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, url
}

func (x *Discovery) Search(ctx context.Requester) (int, interface{}) {
	return http.StatusMethodNotAllowed, nil
}

// @Title GetDirtyService
// @Description Gets the recommended service
// @Param	appID			path	string 	true		"the application requesting a service"
// @Param	serviceName		path 	string	true		"the name of the service you want to get"
// @Success 200 {string} Service.URL
// @Failure 403 :serviceName or :appID is empty
// @router /:appID/:serviceName [get]
func (req *Discovery) GetDirty(ctx context.Requester) (int, interface{}) {
	appID := ctx.FindParam("appID")
	serviceName := ctx.FindParam("serviceName")

	if appID == "" || serviceName == "" {
		err := errors.New("appID AND serviceName must be populated")
		return http.StatusBadRequest, err
	}

	url, err := logic.GetServicePath(serviceName, appID, os.Getenv("HOST"), false)

	if err != nil {
		log.Println(err)
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, url
}

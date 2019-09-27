package controllers

import (
	"net/http"

	"github.com/louisevanderlith/droxolite/context"
	"github.com/louisevanderlith/router/logic"
)

type Applicants struct {
}

func (x *Applicants) Get(ctx context.Requester) (int, interface{}) {
	profile := ctx.FindParam("profile")

	if len(profile) == 0 {
		return http.StatusBadRequest, nil
	}

	result := logic.GetApplicants(profile)

	return http.StatusOK, result
}

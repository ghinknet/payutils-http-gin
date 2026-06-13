package httpgin

import (
	"github.com/gin-gonic/gin"

	"go.gh.ink/payutils/v3/errors"
	"go.gh.ink/payutils/v3/model"
)

type Instance struct {
	Router gin.IRoutes
}

type Driver struct{}

func (d Driver) NewInstance(instance any) (model.HttpInstance, error) {
	// Both *gin.Engine and *gin.RouterGroup implement gin.IRoutes.
	router, ok := instance.(gin.IRoutes)
	if !ok {
		return Instance{}, errors.ErrUnsupportedInstance
	}
	return Instance{Router: router}, nil
}

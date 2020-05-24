package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/p1ass/midare/lib/errors"
	"github.com/p1ass/midare/lib/logging"
)

func sendError(err error, c *gin.Context) {
	logger := logging.New()

	switch e := errors.Cause(err).(type) {
	case *errors.ServiceError:
		logger.Warn(err.Error(), logging.Error(err))
		sendServiceError(e, c)
	default:
		logger.Warn(err.Error(), logging.Error(err))
		c.String(http.StatusInternalServerError, "internal error")
	}
}

func sendServiceError(err *errors.ServiceError, c *gin.Context) {
	switch err.Code {
	case errors.NotFound:
		c.String(http.StatusNotFound, err.Error())
	case errors.BadRequest:
		c.String(http.StatusBadRequest, err.Error())
	case errors.Unauthorized:
		c.String(http.StatusUnauthorized, err.Error())
	case errors.Forbidden:
		c.String(http.StatusForbidden, err.Error())
	case errors.Unknown:
		c.String(http.StatusInternalServerError, err.Error())
	case errors.InvalidArgument:
		c.String(http.StatusBadRequest, err.Error())
	default:
		c.String(http.StatusInternalServerError, err.Error())
	}
}

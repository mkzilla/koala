package types

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

var (
	BadRequestParameter          = fmt.Errorf("bad request parameter")
	BadUserPassword              = fmt.Errorf("incorrect user password")
	FailedToGetDataFromDB        = fmt.Errorf("failed to get data from database")
	BadSystemRSAService          = fmt.Errorf("system encrypted transmission service error")
	NoSuchUser                   = fmt.Errorf("no such user")
	InBlackList                  = fmt.Errorf("user is in black list")
	NoSuchTask                   = fmt.Errorf("no such task")
	NoSuchObjective              = fmt.Errorf("no such objective")
	OAuthV2Error                 = fmt.Errorf("bad oauth 2.0 request")
	NoSuchNotification           = fmt.Errorf("no such notification")
	GroupLimit                   = fmt.Errorf("Each person can only join up to %d groups", 3)
	FailedToInsertDataToDatabase = fmt.Errorf("failed to insert data")
)

func HandleError(c *gin.Context, err error, inner error) {
	switch err {
	case BadRequestParameter:
		log.WithError(inner).Info(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
	case FailedToGetDataFromDB:
		log.WithError(inner).Warn(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
	case BadUserPassword:
		log.WithError(inner).Warn(err.Error())
		c.AbortWithStatusJSON(http.StatusUnauthorized, err.Error())
	case BadSystemRSAService:
		log.WithError(inner).Warn(err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
	case NoSuchUser:
		c.AbortWithStatusJSON(http.StatusNotFound, err.Error())
	case GroupLimit:
		c.AbortWithStatusJSON(http.StatusForbidden, err.Error())
	case FailedToInsertDataToDatabase:
		log.WithError(err).Error(err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
	case NoSuchNotification:
		c.AbortWithStatusJSON(http.StatusNotFound, err.Error())
	case NoSuchTask:
		c.AbortWithStatusJSON(http.StatusNotFound, err.Error())
	case OAuthV2Error:
		log.WithError(inner).Warn(err.Error())
		c.AbortWithStatusJSON(http.StatusForbidden, err.Error())
	case InBlackList:
		c.AbortWithStatusJSON(http.StatusForbidden, err.Error())
	default:
		c.AbortWithStatusJSON(http.StatusForbidden, err.Error())
	}
}

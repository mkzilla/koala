package status

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/mkzilla/koala/pkg/types"
	"github.com/mkzilla/koala/pkg/types/config"
)

func GetRecent(c *gin.Context) {
	var u types.User
	if c.Query("username") != "" {
		config.DBEngine.Where("username = ?", c.Query("username")).Get(&u)
	} else {
		usr, _ := c.Get(config.User)
		u = usr.(types.User)
	}
	status, err := types.GetRecentStatusByUser(u.ID)
	if err != nil {
		types.HandleError(c, types.FailedToGetDataFromDB, err)
		return
	}
	c.JSON(http.StatusOK, status)
}

func CreateNewStatus(c *gin.Context) {
	var status types.Status
	if err := c.ShouldBind(&status); err != nil {
		types.HandleError(c, types.BadRequestParameter, err)
		return
	}
	usr, _ := c.Get(config.User)
	u := usr.(types.User)
	status.UserID = u.ID
	status, err := types.InsertStatus(status)
	if err != nil {
		types.HandleError(c, types.FailedToGetDataFromDB, err)
		return
	}
	c.JSON(http.StatusOK, status)
}

package notification

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/mkzilla/koala/pkg/types"
	"github.com/mkzilla/koala/pkg/types/config"
)

func List(c *gin.Context) {
	var nts = make([]types.Notification, 0)
	usr, _ := c.Get(config.User)

	err := config.DBEngine.Desc("id").
		Where("user_id = ?", usr.(types.User).ID).Limit(c.GetInt(config.PageSize), c.GetInt(config.PageOffset)).Find(&nts)
	if err != nil {
		types.HandleError(c, types.FailedToGetDataFromDB, err)
		return
	}
	c.JSON(http.StatusOK, nts)
}

func DoRead(c *gin.Context) {
	id := c.Param("notification_id")
	usr, _ := c.Get(config.User)
	var notification = types.Notification{}

	ext, err := config.DBEngine.Table("notification").Where("id = ?", id).Get(&notification)
	if err != nil {
		types.HandleError(c, types.FailedToGetDataFromDB, err)
		return
	}
	if !ext || notification.UserID != usr.(types.User).ID {
		types.HandleError(c, types.NoSuchNotification, nil)
		return
	}
	n := notification
	n.IsRead = true
	_, err = config.DBEngine.Update(n, notification)
	if err != nil {
		types.HandleError(c, types.FailedToInsertDataToDatabase, err)
		return
	}
	c.JSON(http.StatusOK, "")
}

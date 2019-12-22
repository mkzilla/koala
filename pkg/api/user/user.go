package user

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/mkzilla/koala/pkg/types"
	"github.com/mkzilla/koala/pkg/types/config"
	"github.com/mkzilla/koala/pkg/types/data"
)

func AssignList(c *gin.Context) {
	var usrs = make([]data.User, 0)
	al := c.Query("like") + "%"
	usr, _ := c.Get(config.User)
	err := config.DBEngine.Table("user").
		Join("INNER", "group_user", "user.id = group_user.user_id and group_user.enable = true").
		Where("(username like ? or nickname like ? ) and group_user.group_id in (select group_id from group_user where user_id = ?)", al, al, usr.(types.User).ID).
		Limit(c.GetInt(config.PageSize), c.GetInt(config.PageOffset)).
		Find(&usrs)
	if err != nil {
		types.HandleError(c, types.FailedToGetDataFromDB, err)
		return
	}
	c.JSON(http.StatusOK, usrs)
}

func List(c *gin.Context) {
	var usrs = make([]types.User, 0)
	al := c.Query("like") + "%"
	err := config.DBEngine.Table("user").Desc("user.id").
		Where("username like ?", al).
		Limit(c.GetInt(config.PageSize), c.GetInt(config.PageOffset)).
		Find(&usrs)
	if err != nil {
		types.HandleError(c, types.FailedToGetDataFromDB, err)
		return
	}
	c.JSON(http.StatusOK, usrs)
}

func Create(c *gin.Context) {
	usr := types.User{}
	err := c.ShouldBind(&usr)
	if err != nil {
		types.HandleError(c, types.BadRequestParameter, err)
		return
	}
	_, err = config.DBEngine.Insert(&usr)
	if err != nil {
		types.HandleError(c, types.FailedToInsertDataToDatabase, err)
		return
	}
	c.JSON(http.StatusOK, usr)
}

func UpdateUser(c *gin.Context) {
	usr := types.User{}
	err := c.ShouldBind(&usr)
	if err != nil {
		types.HandleError(c, types.BadRequestParameter, err)
		return
	}

	var id = c.Param("user_id")
	var ou types.User
	ext, err := config.DBEngine.Where("id = ?", id).Get(&ou)
	if err != nil {
		types.HandleError(c, types.FailedToGetDataFromDB, err)
	}
	if !ext || ou.ID != usr.ID {
		types.HandleError(c, types.NoSuchUser, nil)
	}
	_, err = config.DBEngine.ID(ou.ID).AllCols().Update(usr)
	if err != nil {
		types.HandleError(c, types.FailedToInsertDataToDatabase, err)
		return
	}
	c.JSON(http.StatusOK, usr)
}

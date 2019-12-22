package group

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/mkzilla/koala/pkg/types"
	"github.com/mkzilla/koala/pkg/types/config"
)

func ListIn(c *gin.Context) {
	var grps = make([]types.Group, 0)
	al := c.Query("name")
	usr, _ := c.Get(config.User)

	enable, err := strconv.ParseBool(c.Query("enable"))
	if err != nil {
		types.HandleError(c, types.BadRequestParameter, err)
		return
	}

	err = config.DBEngine.Table("group").
		Join("INNER", "group_user", "group_user.group_id = group.id and enable = ? and user_id = ?", enable, usr.(types.User).ID).
		Desc("group.id").
		Where("name like ?", al+"%").Limit(c.GetInt(config.PageSize), c.GetInt(config.PageOffset)).Find(&grps)
	if err != nil {
		types.HandleError(c, types.FailedToGetDataFromDB, err)
		return
	}
	c.JSON(http.StatusOK, grps)
}

func GetGroup(c *gin.Context) {
	id := c.Param("group_id")
	grp, err := types.GetGroupByID(id)
	if err != nil {
		types.HandleError(c, types.FailedToGetDataFromDB, err)
		return
	}
	c.JSON(http.StatusOK, grp)

}

func GetOKRByGroup(c *gin.Context) {
	id := c.Param("group_id")
	objs, err := types.GetObjectivesByGroup(id, c.GetInt(config.PageSize), c.GetInt(config.PageOffset))
	if err != nil {
		types.HandleError(c, types.FailedToGetDataFromDB, err)
		return
	}
	c.JSON(http.StatusOK, objs)
}

func CreateGroupOKR(c *gin.Context) {
	var o types.Objective
	err := c.ShouldBind(&o)
	if err != nil {
		types.HandleError(c, types.BadRequestParameter, err)
		return
	}
	sid := c.Param("group_id")
	id, _ := strconv.ParseInt(sid, 10, 64)
	o.GroupID = id
	o, err = types.CreateObjective(o)
	if err != nil {
		types.HandleError(c, types.FailedToInsertDataToDatabase, err)
		return
	}
	c.JSON(http.StatusOK, o)
}

func EditGroupOKR(c *gin.Context) {
	var o types.Objective
	err := c.ShouldBind(&o)
	if err != nil {
		types.HandleError(c, types.BadRequestParameter, err)
		return
	}
	sid := c.Param("group_id")
	id, _ := strconv.ParseInt(sid, 10, 64)
	o.GroupID = id
	err = types.UpdateObjectiveByGroupID(sid, o)
	if err != nil {
		types.HandleError(c, types.FailedToInsertDataToDatabase, err)
		return
	}
	c.JSON(http.StatusOK, o)
}

func Create(c *gin.Context) {
	usr, _ := c.Get(config.User)
	cnt, err := config.DBEngine.Table("group_user").Where("user_id = ? and enable = true", usr.(types.User).ID).Count()
	if err != nil {
		types.HandleError(c, types.FailedToGetDataFromDB, err)
		return
	}
	if cnt >= 3 {
		types.HandleError(c, types.GroupLimit, nil)
		return
	}
	grp := types.Group{}
	err = c.ShouldBind(&grp)
	grp.AdminID = usr.(types.User).ID
	grp.Admin = usr.(types.User).Username
	if err != nil {
		types.HandleError(c, types.BadRequestParameter, err)
		return
	}
	_, err = config.DBEngine.Insert(&grp)
	if err != nil {
		types.HandleError(c, types.FailedToInsertDataToDatabase, err)
		return
	}
	var grpU = types.GroupUser{
		UserID:  grp.AdminID,
		GroupID: grp.ID,
		Enable:  true,
	}
	_, err = config.DBEngine.AllCols().Insert(&grpU)
	if err != nil {
		types.HandleError(c, types.FailedToInsertDataToDatabase, err)
		return
	}
	c.JSON(http.StatusOK, grp)
}

func ListUsers(c *gin.Context) {
	var usrs = make([]types.User, 0)
	al := c.Query("name")
	id := c.Param("group_id")
	err := config.DBEngine.Table("user").
		Join("INNER", "group_user", "user.id = user_id and group_user.enable=true").
		Where("group_id = ? and user.username like ?", id, al+"%").Limit(c.GetInt(config.PageSize), c.GetInt(config.PageOffset)).Find(&usrs)
	if err != nil {
		types.HandleError(c, types.FailedToGetDataFromDB, err)
		return
	}
	c.JSON(http.StatusOK, usrs)
}

func CheckIn(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("group_id"))
	usr, _ := c.Get(config.User)
	cnt, err := config.DBEngine.Table("group_user").Where("user_id = ? and enable = true", usr.(types.User).ID).Count()
	if err != nil {
		types.HandleError(c, types.FailedToGetDataFromDB, err)
	}
	if cnt >= 3 {
		types.HandleError(c, types.GroupLimit, nil)
	}

	var grpo = types.GroupUser{
		UserID:  usr.(types.User).ID,
		GroupID: int64(id),
	}

	var grpn = types.GroupUser{
		Enable: true,
	}

	_, err = config.DBEngine.Cols("enable").Update(grpn, grpo)
	if err != nil {
		types.HandleError(c, types.FailedToInsertDataToDatabase, err)
		return
	}
	c.JSON(http.StatusOK, usr)

}

func CheckOut(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("group_id"))
	usr, _ := c.Get(config.User)
	var grp = types.GroupUser{
		UserID:  usr.(types.User).ID,
		GroupID: int64(id),
	}
	_, err := config.DBEngine.Delete(grp)
	if err != nil {
		types.HandleError(c, types.FailedToInsertDataToDatabase, err)
		return
	}
	c.JSON(http.StatusOK, usr)
}

// 组管理员权限
func AddUser(c *gin.Context) {
	name := c.Param("username")
	id, _ := strconv.Atoi(c.Param("group_id"))
	var usr = types.User{
		Username: name,
	}
	ext, err := config.DBEngine.Get(&usr)
	if err != nil {
		types.HandleError(c, types.FailedToGetDataFromDB, err)
	}
	if !ext {
		types.HandleError(c, types.NoSuchUser, nil)
		return
	}
	var grpU = types.GroupUser{
		UserID:  usr.ID,
		GroupID: int64(id),
		Enable:  false,
	}
	_, err = config.DBEngine.Insert(&grpU)
	if err != nil {
		types.HandleError(c, types.FailedToInsertDataToDatabase, err)
		return
	}
	c.JSON(http.StatusOK, "")
}

// 组管理员权限
func RemoveUser(c *gin.Context) {
	name := c.Param("username")
	id, _ := strconv.Atoi(c.Param("group_id"))
	var usr = types.User{
		Username: name,
	}
	ext, err := config.DBEngine.Get(&usr)
	if err != nil {
		types.HandleError(c, types.FailedToGetDataFromDB, err)
	}
	if !ext {
		types.HandleError(c, types.NoSuchUser, nil)
		return
	}
	var grpU = types.GroupUser{
		UserID:  usr.ID,
		GroupID: int64(id),
	}

	_, err = config.DBEngine.Delete(&grpU)
	if err != nil {
		types.HandleError(c, types.FailedToInsertDataToDatabase, err)
		return
	}
	c.JSON(http.StatusOK, "")
}

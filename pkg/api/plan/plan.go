package plan

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
	plan, err := types.GetRecentPlanByUser(u.ID)
	if err != nil {
		types.HandleError(c, types.FailedToGetDataFromDB, err)
		return
	}
	c.JSON(http.StatusOK, plan)
}

func CreateNewPlan(c *gin.Context) {
	var plan types.Plan
	if err := c.ShouldBind(&plan); err != nil {
		types.HandleError(c, types.BadRequestParameter, err)
		return
	}
	usr, _ := c.Get(config.User)
	u := usr.(types.User)
	plan.UserID = u.ID
	plan, err := types.InsertPlan(plan)
	if err != nil {
		types.HandleError(c, types.FailedToGetDataFromDB, err)
		return
	}
	c.JSON(http.StatusOK, plan)
}

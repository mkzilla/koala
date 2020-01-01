package technique

import (
	"net/http"
	"time"

	"github.com/mkzilla/koala/pkg/types"
	"github.com/mkzilla/koala/pkg/types/config"

	"github.com/gin-gonic/gin"
)

func StartTechnique(c *gin.Context) {
	usr, _ := c.Get(config.User)
	var nw = time.Now()

	tss, err := types.GetCountTechnique(usr.(types.User).ID, nw.Add(time.Minute*-25))
	if err != nil {
		types.HandleError(c, types.FailedToGetDataFromDB, err)
		return
	}
	if len(tss) > 0 {
		c.JSON(http.StatusBadRequest, len(tss))
		return
	}
	var t = types.Technique{
		UserID:          usr.(types.User).ID,
		CreateTimeStamp: nw.Unix(),
	}
	ts, err := types.InsertTechnique(t)
	if err != nil {
		types.HandleError(c, types.FailedToInsertDataToDatabase, err)
		return
	}
	ts.CreateTime = nw
	c.JSON(http.StatusOK, ts)
}

func GetRecentTechnique(c *gin.Context) {
	usr, _ := c.Get(config.User)
	var nw = time.Now()

	tss, err := types.GetCountTechnique(usr.(types.User).ID, nw.Add(time.Minute*-25))
	if err != nil {
		types.HandleError(c, types.FailedToGetDataFromDB, err)
		return
	}
	if len(tss) > 0 {
		tss[0].CreateTime = time.Unix(tss[0].CreateTimeStamp, 0).Local()
		c.JSON(http.StatusOK, tss[0])
		return
	}
	c.JSON(http.StatusNotFound, nil)
}

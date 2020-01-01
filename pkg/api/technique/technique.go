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

	tss, err := types.GetDoingTechnique(usr.(types.User).ID)
	if err != nil {
		types.HandleError(c, types.FailedToGetDataFromDB, err)
		return
	}
	if tss != nil {
		c.JSON(http.StatusBadRequest, tss)
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
	ts, err := types.GetDoingTechnique(usr.(types.User).ID)
	if err != nil {
		types.HandleError(c, types.FailedToGetDataFromDB, err)
		return
	}
	if ts != nil {
		ts.CreateTime = time.Unix(ts.CreateTimeStamp, 0).Local()
		c.JSON(http.StatusOK, ts)
		return
	}
	c.JSON(http.StatusNotFound, nil)
}

func SetDone(c *gin.Context) {
	usr, _ := c.Get(config.User)
	err := types.SetDoneTechniqueByID(usr.(types.User).ID)
	if err != nil {
		types.HandleError(c, types.FailedToInsertDataToDatabase, err)
		return
	}
	c.JSON(http.StatusOK, nil)
}

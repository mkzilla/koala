package statistics

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/mkzilla/koala/pkg/types"
	"github.com/mkzilla/koala/pkg/types/data"
)

func Statistics(c *gin.Context) {
	var arr data.StatisticsData
	gs, err := types.GroupByState()
	if err != nil {
		types.HandleError(c, types.FailedToGetDataFromDB, err)
		return
	}
	arr.Tasks = gs
	cl, err := types.CommentTimeLine()
	if err != nil {
		types.HandleError(c, types.FailedToGetDataFromDB, err)
		return
	}
	arr.Comments = cl
	c.JSON(http.StatusOK, arr)

}

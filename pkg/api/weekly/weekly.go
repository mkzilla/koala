package weekly

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/mkzilla/koala/pkg/types"
	"github.com/mkzilla/koala/pkg/types/config"
	"github.com/mkzilla/koala/pkg/types/data"
)

type queryParam struct {
	From     *time.Time `json:"from" form:"from" time_format:"2006-01-02T15:04:05"`
	To       *time.Time `json:"to" form:"to" time_format:"2006-01-02T15:04:05"`
	Username string     `json:"username" form:"username"`
}

func ListSelf(c *gin.Context) {
	usr, _ := c.Get(config.User)

	var params queryParam

	err := c.ShouldBindQuery(&params)
	if err != nil {
		types.HandleError(c, types.BadRequestParameter, err)
		return
	}
	var u types.User
	if params.Username != "" {
		config.DBEngine.Where("username = ?", params.Username).Get(&u)
	} else {
		u = usr.(types.User)
	}
	reports, err := getReports(u.ID, params)
	if err != nil {
		types.HandleError(c, types.FailedToGetDataFromDB, err)
		return
	}
	c.JSON(http.StatusOK, reports)
}

func getReports(id int64, params queryParam) (*data.Reports, error) {
	var comms = []types.Comment{}
	var reports = data.Reports{}
	taskMapComment := make(map[int64][]types.Comment)
	if params.From == nil || params.To == nil {
		now := time.Now()
		params.To = &now
		offset := int(time.Monday - now.Weekday())
		if offset > 0 {
			offset = -6
		}
		weekStart := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, offset)
		params.From = &weekStart
	}

	err := config.DBEngine.Table("comment").
		Where("user_id = ? and type = ? and create_time > ? and create_time < ?", id, types.TaskSchedule, params.From, params.To).
		OrderBy("id desc").Find(&comms)
	if err != nil {
		return nil, err
	}
	for _, t := range comms {
		taskMapComment[t.TaskID] = append(taskMapComment[t.TaskID], t)
	}
	var usr = types.User{
		ID: id,
	}
	config.DBEngine.Get(&usr)
	reports.User = usr
	reports.From = params.From
	reports.To = params.To

	for k, v := range taskMapComment {
		tsk := types.Task{
			ID: k,
		}
		_, err := config.DBEngine.Get(&tsk)
		if err != nil {
			return nil, err
		}
		reports.Reports = append(reports.Reports, data.Report{
			Task:      tsk,
			Schedules: v,
		})
	}
	return &reports, nil

}

func ListGroup(c *gin.Context) {
	var params queryParam
	err := c.ShouldBindQuery(&params)
	if err != nil {
		types.HandleError(c, types.BadRequestParameter, err)
		return
	}
	var id = c.Param("group_id")
	var grs = []types.GroupUser{}
	err = config.DBEngine.Table("group_user").Where("group_id = ? and enable = true", id).Find(&grs)
	if err != nil {
		types.HandleError(c, types.FailedToGetDataFromDB, err)
		return
	}
	var reports = []data.Reports{}
	for _, g := range grs {
		report, err := getReports(g.UserID, params)
		if err != nil {
			types.HandleError(c, types.FailedToGetDataFromDB, err)
		}
		reports = append(reports, *report)
	}
	c.JSON(http.StatusOK, reports)
}

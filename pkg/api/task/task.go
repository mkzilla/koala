package task

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/mkzilla/koala/pkg/types"
	"github.com/mkzilla/koala/pkg/types/config"
	"github.com/mkzilla/koala/pkg/types/data"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type QueryType string

const (
	is_watch  QueryType = "is_watch"
	is_vote   QueryType = "is_vote"
	is_assign QueryType = "is_assign"

	is_doing QueryType = "is_doing"
)

type listParams struct {
	Type     QueryType       `json:"type" form:"type"`
	State    types.TaskState `json:"state" form:"state"`
	ID       int64           `json:"id" form:"id"`
	uid      int64
	Username string `json:"username" form:"username"`
}

func GetOKRByUser(c *gin.Context) {
	var params listParams
	err := c.ShouldBindQuery(&params)
	if err != nil {
		types.HandleError(c, types.BadRequestParameter, err)
		return
	}
	usr, _ := c.Get(config.User)
	if params.Username != "" {
		var u types.User
		config.DBEngine.Where("username = ?", params.Username).Get(&u)
		params.uid = u.ID
	}
	if params.uid == 0 {
		params.uid = usr.(types.User).ID
	}
	var objs []types.Objective
	if params.ID != 0 {
		var obj types.Objective
		obj, err = types.GetObjectiveByID(params.ID)
		objs = append(objs, obj)
	} else if params.Type == is_doing {
		objs, err = types.GetDoingObjectivesByUser(params.uid, c.GetInt(config.PageSize), c.GetInt(config.PageOffset))
	} else {
		objs, err = types.GetObjectivesByUser(params.uid, c.GetInt(config.PageSize), c.GetInt(config.PageOffset))
	}
	if err != nil {
		types.HandleError(c, types.FailedToGetDataFromDB, err)
		return
	}
	c.JSON(http.StatusOK, objs)
}

func CreateUserOKR(c *gin.Context) {
	var o types.Objective
	err := c.ShouldBind(&o)
	if err != nil {
		types.HandleError(c, types.BadRequestParameter, err)
		return
	}
	usr, _ := c.Get(config.User)
	o.GroupID = 0
	o.UserID = usr.(types.User).ID
	o, err = types.CreateObjective(o)
	if err != nil {
		types.HandleError(c, types.FailedToInsertDataToDatabase, err)
		return
	}
	c.JSON(http.StatusOK, o)
}

func EditUserOKR(c *gin.Context) {
	var o types.Objective
	err := c.ShouldBind(&o)
	if err != nil {
		types.HandleError(c, types.BadRequestParameter, err)
		return
	}
	usr, _ := c.Get(config.User)
	o.GroupID = 0
	o.UserID = usr.(types.User).ID
	err = types.UpdateObjectiveByUserID(o.UserID, o)
	if err != nil {
		types.HandleError(c, types.FailedToInsertDataToDatabase, err)
		return
	}
	c.JSON(http.StatusOK, o)
}

func List(c *gin.Context) {
	var ls = []data.TaskTimeLine{}
	var params listParams
	err := c.ShouldBindQuery(&params)
	if err != nil {
		types.HandleError(c, types.BadRequestParameter, err)
		return
	}
	usr, _ := c.Get(config.User)
	if params.Username != "" {
		var u types.User
		config.DBEngine.Where("username = ?", params.Username).Get(&u)
		params.uid = u.ID
	}
	if params.uid == 0 {
		params.uid = usr.(types.User).ID
	}
	switch params.Type {
	case is_watch:
		err = config.DBEngine.Table("task_watch").
			Join("LEFT", "task", "task_watch.task_id = task.id").
			Where("task_watch.user_id = ? and (task.state = ? or ? = -1)", params.uid, params.State, params.State).
			OrderBy("task_watch.id desc").Limit(c.GetInt(config.PageSize), c.GetInt(config.PageOffset)).Find(&ls)
	case is_vote:
		err = config.DBEngine.Table("task_vote").
			Join("LEFT", "task", "task_vote.task_id = task.id").
			Where("task_vote.user_id = ?  and (task.state = ? or ? = -1)", params.uid, params.State, params.State).
			OrderBy("task_vote.id desc").Limit(c.GetInt(config.PageSize), c.GetInt(config.PageOffset)).Find(&ls)
	case is_assign:
		var tasks = []types.Task{}
		err = config.DBEngine.Table("task").
			Where("assign_to = ?  and (task.state = ? or ? = -1)", params.uid, params.State, params.State).
			OrderBy("task.id desc").Limit(c.GetInt(config.PageSize), c.GetInt(config.PageOffset)).Find(&tasks)
		if err != nil {
			types.HandleError(c, types.FailedToGetDataFromDB, err)
			return
		}
		c.JSON(http.StatusOK, tasks)
		return
	default:
		FeedLog(c)
		return
	}
	if err != nil {
		types.HandleError(c, types.FailedToGetDataFromDB, err)
		return
	}
	c.JSON(http.StatusOK, ls)
}

func FeedLog(c *gin.Context) {
	usr, _ := c.Get(config.User)
	var usrs []types.User

	err := config.DBEngine.Table("user").
		Join("INNER", "group_user", "user.id = group_user.user_id and group_user.enable = true").
		Where("group_user.group_id in (select group_id from group_user where user_id = ? and group_user.enable = true)", usr.(types.User).ID).Find(&usrs)
	if err != nil {
		types.HandleError(c, types.FailedToGetDataFromDB, err)
		return
	}

	ids := fmt.Sprintf("%d", usr.(types.User).ID)
	for _, u := range usrs {
		ids = fmt.Sprintf("%s,%d", ids, u.ID)
	}
	ids = "(" + ids + ")"

	var ls = []data.TaskTimeLine{}
	err = config.DBEngine.Table("task_log").
		Join("LEFT", "task", "task_log.task_id = task.id").
		Join("LEFT", "user", "task_log.user_id = user.id").
		Where("task_log.user_id in "+ids).
		OrderBy("task_log.id desc").Limit(c.GetInt(config.PageSize), c.GetInt(config.PageOffset)).Find(&ls)
	if err != nil {
		types.HandleError(c, types.FailedToGetDataFromDB, err)
		return
	}
	c.JSON(http.StatusOK, ls)
}

func Kanban(c *gin.Context) {
	usr, _ := c.Get(config.User)
	var tbs = new(data.TaskBoard)
	err := config.DBEngine.Where("state = ? and assign_to = ?", types.WILL, usr.(types.User).ID).Find(&tbs.Wills)
	if err != nil {
		types.HandleError(c, types.FailedToGetDataFromDB, err)
		return
	}

	err = config.DBEngine.Where("state = ? and assign_to = ?", types.DOING, usr.(types.User).ID).Find(&tbs.Doing)
	if err != nil {
		types.HandleError(c, types.FailedToGetDataFromDB, err)
		return
	}

	err = config.DBEngine.Where("state = ? and assign_to = ?", types.DONE, usr.(types.User).ID).OrderBy("update_time desc").Limit(10, 0).Find(&tbs.Done)
	if err != nil {
		types.HandleError(c, types.FailedToGetDataFromDB, err)
		return
	}
	c.JSON(http.StatusOK, tbs)
}

func Recent(c *gin.Context) {
	usr, _ := c.Get(config.User)
	var ls = []data.TaskTimeLine{}
	err := config.DBEngine.Table("task").Distinct("task.*,task_log.user_id,task_log.username,task_log.task_id").
		Join("LEFT", "task_log", "task_log.task_id = task.id").
		Where("task_log.user_id = ?", usr.(types.User).ID).Limit(6, 0).
		OrderBy("task.id desc").Find(&ls)
	if err != nil {
		types.HandleError(c, types.FailedToGetDataFromDB, err)
		return
	}
	c.JSON(http.StatusOK, ls)
}

func Create(c *gin.Context) {
	var t types.Task
	err := c.ShouldBind(&t)
	if err != nil {
		types.HandleError(c, types.BadRequestParameter, err)
		return
	}
	usr, _ := c.Get(config.User)
	t.AssignTo = usr.(types.User).ID

	_, err = config.DBEngine.Insert(&t)
	if err != nil {
		types.HandleError(c, types.FailedToInsertDataToDatabase, err)
		return
	}
	_, err = config.DBEngine.Insert(types.TaskLog{
		UserID:   usr.(types.User).ID,
		Username: usr.(types.User).Username,
		TaskID:   t.ID,
		Behavior: types.CreateTaskLog,
	})
	if err != nil {
		types.HandleError(c, types.FailedToInsertDataToDatabase, err)
		return
	}
	c.JSON(http.StatusOK, t.ID)
}

func View(c *gin.Context) {
	var id = c.Param("task_id")
	var task types.Task
	_, err := config.DBEngine.Where("id = ?", id).Get(&task)
	if err != nil {
		types.HandleError(c, types.FailedToGetDataFromDB, err)
		return
	}
	var taskDetail = data.TaskDetail{
		Task: task,
	}

	usr, _ := c.Get(config.User)

	ext, _ := config.DBEngine.Table("task_vote").
		Where("user_id = ? and task_id= ?", usr.(types.User).ID, id).Count()
	if ext > 0 {
		taskDetail.IsVote = true
	}
	ext, _ = config.DBEngine.Table("task_watch").
		Where("user_id = ? and task_id= ?", usr.(types.User).ID, id).Count()
	if ext > 0 {
		taskDetail.IsWatch = true
	}

	if task.AssignTo != 0 {
		var usr = types.User{
			ID: task.AssignTo,
		}
		config.DBEngine.Get(&usr)
		taskDetail.User = usr
	}

	c.JSON(http.StatusOK, taskDetail)
}

func ListComment(c *gin.Context) {
	var id = c.Param("task_id")
	var cmt = []data.Comment{}
	config.DBEngine.Table("comment").
		Where("task_id = ?", id).Join("INNER", "user", "comment.user_id = user.id").
		OrderBy("comment.id desc").Find(&cmt)
	c.JSON(http.StatusOK, cmt)
}

func Comment(c *gin.Context) {
	var cmt types.Comment
	err := c.ShouldBind(&cmt)
	if err != nil {
		types.HandleError(c, types.BadRequestParameter, err)
		return
	}
	usr, _ := c.Get(config.User)
	cmt.UserID = usr.(types.User).ID
	cmt.Username = usr.(types.User).Username
	_, err = config.DBEngine.Insert(&cmt)
	if err != nil {
		types.HandleError(c, types.FailedToInsertDataToDatabase, err)
		return
	}

	if cmt.Type == types.TaskComment {
		_, err = config.DBEngine.Insert(types.TaskLog{
			UserID:   cmt.UserID,
			Username: cmt.Username,
			TaskID:   cmt.TaskID,
			Behavior: types.AddCommentLog,
			ToValue:  cmt.Info,
		})
	} else {
		_, err = config.DBEngine.Insert(types.TaskLog{
			UserID:   cmt.UserID,
			Username: cmt.Username,
			TaskID:   cmt.TaskID,
			Behavior: types.AddScheduleLog,
			ToValue:  cmt.Info,
		})
	}
	if err != nil {
		types.HandleError(c, types.FailedToInsertDataToDatabase, err)
		return
	}
	c.JSON(http.StatusOK, "")
}

func ListVote(c *gin.Context) {
	var id = c.Param("task_id")
	var tv = []types.TaskVote{}
	err := config.DBEngine.Where("task_id = ?", id).Find(&tv)
	if err != nil {
		types.HandleError(c, types.FailedToGetDataFromDB, err)
		return
	}
	c.JSON(http.StatusOK, tv)
}

func Vote(c *gin.Context) {
	var id = c.Param("task_id")
	n, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	usr, _ := c.Get(config.User)
	var voteLog types.TaskVote
	ext, err := config.DBEngine.Where("task_id = ? and user_id = ?", id, usr.(types.User).ID).Delete(voteLog)
	if err != nil {
		types.HandleError(c, types.FailedToGetDataFromDB, err)
		return
	}
	if ext == 0 {
		voteLog.Username = usr.(types.User).Username
		voteLog.TaskID = int64(n)
		voteLog.UserID = usr.(types.User).ID
		_, err := config.DBEngine.Insert(voteLog)
		if err != nil {
			types.HandleError(c, types.FailedToInsertDataToDatabase, err)
			return
		}
		_, err = config.DBEngine.Insert(types.TaskLog{
			UserID:   usr.(types.User).ID,
			Username: usr.(types.User).Username,
			TaskID:   int64(n),
			Behavior: types.VoteLog,
		})
		if err != nil {
			types.HandleError(c, types.FailedToInsertDataToDatabase, err)
			return
		}
		c.JSON(http.StatusOK, "")
		return

	} else {
		_, err := config.DBEngine.Delete(types.TaskLog{
			UserID:   usr.(types.User).ID,
			TaskID:   int64(n),
			Behavior: types.VoteLog,
		})
		if err != nil {
			types.HandleError(c, types.FailedToInsertDataToDatabase, err)
			return
		}
	}
	c.JSON(http.StatusOK, "")
}

func ListWatch(c *gin.Context) {
	var id = c.Param("task_id")
	var tw = []types.TaskWatch{}
	err := config.DBEngine.Where("task_id = ?", id).Find(&tw)
	if err != nil {
		types.HandleError(c, types.FailedToGetDataFromDB, err)
		return
	}
	c.JSON(http.StatusOK, tw)
}

func Watch(c *gin.Context) {
	var id = c.Param("task_id")
	n, err := strconv.Atoi(id)
	if err != nil {
		types.HandleError(c, types.BadRequestParameter, err)
		return
	}
	usr, _ := c.Get(config.User)
	var watchLog types.TaskWatch
	ext, err := config.DBEngine.Where("task_id = ? and user_id = ?", id, usr.(types.User).ID).Delete(watchLog)
	if err != nil {
		types.HandleError(c, types.FailedToGetDataFromDB, err)
		return
	}
	if ext == 0 {
		watchLog.Username = usr.(types.User).Username
		watchLog.TaskID = int64(n)
		watchLog.UserID = usr.(types.User).ID
		_, err := config.DBEngine.Insert(watchLog)
		if err != nil {
			types.HandleError(c, types.FailedToInsertDataToDatabase, err)
		}
		_, err = config.DBEngine.Insert(types.TaskLog{
			UserID:   usr.(types.User).ID,
			Username: usr.(types.User).Username,
			TaskID:   int64(n),
			Behavior: types.WatchLog,
		})
		if err != nil {
			types.HandleError(c, types.FailedToInsertDataToDatabase, err)
			return
		}
		c.JSON(http.StatusOK, "")
		return
	}
	_, err = config.DBEngine.Delete(types.TaskLog{
		UserID:   usr.(types.User).ID,
		TaskID:   int64(n),
		Behavior: types.WatchLog,
	})
	if err != nil {
		types.HandleError(c, types.FailedToInsertDataToDatabase, err)
		return
	}
	c.JSON(http.StatusOK, "")

}

func Update(c *gin.Context) {
	var t types.Task
	err := c.ShouldBind(&t)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	var now = time.Now()
	t.UpdateTime = &now
	var id = c.Param("task_id")
	var ot types.Task
	ext, err := config.DBEngine.Where("id = ?", id).Get(&ot)
	if err != nil {
		types.HandleError(c, types.FailedToGetDataFromDB, err)
	}
	if !ext || ot.ID != t.ID {
		types.HandleError(c, types.NoSuchTask, nil)
		return
	}
	_, err = config.DBEngine.Update(t, ot)
	if err != nil {
		types.HandleError(c, types.FailedToInsertDataToDatabase, err)
		return
	}
	usr, _ := c.Get(config.User)
	if ot.AssignTo != t.AssignTo {
		var u = types.User{
			ID: t.AssignTo,
		}
		config.DBEngine.Get(&u)
		_, err = config.DBEngine.Insert(types.TaskLog{
			UserID:   usr.(types.User).ID,
			Username: usr.(types.User).Username,
			TaskID:   t.ID,
			ToValue:  u.Username,
			Behavior: types.AddAssignLog,
		})
		if err != nil {
			types.HandleError(c, types.FailedToInsertDataToDatabase, nil)
			return
		}
	}
	if ot.Title != t.Title {
		_, err = config.DBEngine.Insert(types.TaskLog{
			UserID:    usr.(types.User).ID,
			Username:  usr.(types.User).Username,
			TaskID:    t.ID,
			FromValue: ot.Title,
			ToValue:   t.Title,
			Behavior:  types.EditTitleLog,
		})
	}
	if ot.Desc != t.Desc {
		_, err = config.DBEngine.Insert(types.TaskLog{
			UserID:    usr.(types.User).ID,
			Username:  usr.(types.User).Username,
			TaskID:    t.ID,
			FromValue: ot.Desc,
			ToValue:   t.Desc,
			Behavior:  types.EditDescLog,
		})
	}
	if err != nil {
		types.HandleError(c, types.FailedToInsertDataToDatabase, err)
		return
	}
	// TODO 进度管理
	c.JSON(http.StatusOK, "")
}

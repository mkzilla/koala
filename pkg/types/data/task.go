package data

import (
	"github.com/mkzilla/koala/pkg/types"
)

type TaskTimeLine struct {
	TaskLog    types.TaskLog `xorm:"extends" json:"task_log"`
	Nickname   string        `json:"nickname" xorm:"nickname"`
	Title      string        `json:"title" xorm:"'title'"`
	State      string        `json:"state" xorm:"'state'"`           // 任务状态：未开始、进行中、完成、取消、失败、暂停
	Importance int32         `json:"importance" xorm:"'importance'"` // 重要性：不重要，一般重要，重要，很重要，极其重要
	Priority   int32         `json:"priority" xorm:"'priority'"`     // 优先级：不紧急，一般紧急，紧急，很紧急，极其紧急
}

type TaskDetail struct {
	types.Task
	IsVote  bool       `json:"is_vote"`
	IsWatch bool       `json:"is_watch"`
	User    types.User `json:"user"`
}

type TaskBoard struct {
	Wills []types.Task `json:"will"`
	Doing []types.Task `json:"doing"`
	Done  []types.Task `json:"done"`
}

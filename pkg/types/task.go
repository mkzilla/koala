package types

import (
	"time"

	"encoding/json"

	"github.com/mkzilla/koala/pkg/types/config"
)

type TaskType string
type TaskState int

const (
	UNSTART = iota // 没有进展
	WILL           // 即将开始
	DOING          // 进行中
	DONE           // 完成
	GIVEUP         // 取消
	SUSPEND        //暂停
)

const (
	ISSUE TaskType = "issue"
	PLAN  TaskType = "plan"
	TASK  TaskType = "task"
)

type Task struct {
	ID         int64      `json:"id" xorm:"'id' pk autoincr"`
	CreateTime *time.Time `json:"create_time,omitempty" xorm:"'create_time' created"`
	UpdateTime *time.Time `json:"update_time,omitempty" xorm:"'update_time' created"`
	Title      string     `json:"title" xorm:"'title'"`              // 任务标题
	Desc       string     `json:"desc" xorm:"'desc' varchar(10000)"` // 任务简介
	State      TaskState  `json:"state" xorm:"'state'"`              // 任务状态：未开始、进行中、完成、取消、失败、暂停
	Importance int32      `json:"importance" xorm:"'importance'"`    // 重要性：不重要，一般重要，重要，很重要，极其重要
	Priority   int32      `json:"priority" xorm:"'priority'"`        // 优先级：不紧急，一般紧急，紧急，很紧急，极其紧急
	Type       TaskType   `json:"type" xorm:"'type'"`
	AssignTo   int64      `json:"assign_to" xorm:"assign_to"`
}

type TaskGroupCount struct {
	Count int64     `json:"count"`
	State TaskState `json:"state"`
}

func GroupByState() ([]TaskGroupCount, error) {
	var data []map[string][]byte
	var as []TaskGroupCount

	data, err := config.DBEngine.Query("select count(*) count,state from task group by state")
	for _, d := range data {
		var a TaskGroupCount
		json.Unmarshal(d["count"], &a.Count)
		json.Unmarshal(d["state"], &a.State)
		as = append(as, a)
	}
	return as, err
}

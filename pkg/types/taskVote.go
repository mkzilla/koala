package types

import (
	"time"
)

type TaskVote struct {
	ID         int64      `json:"id" xorm:"'id' pk autoincr"`
	CreateTime *time.Time `json:"create_time" xorm:"'create_time' created"`
	TaskID     int64      `json:"task_id" xorm:"'task_id' unique(tu)"`
	UserID     int64      `json:"user_id" xorm:"'user_id' unique(tu)"`
	Username   string     `json:"username" xorm:"username"`
}

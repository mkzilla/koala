package data

import (
	"time"
)

type Comment struct {
	ID         int64      `json:"id" xorm:"'id' pk autoincr"`
	CreateTime *time.Time `json:"create_time,omitempty" xorm:"'create_time' created"`
	Info       string     `json:"info" xorm:"'info' varchar(5000)"`
	UserID     int64      `json:"user_id" xorm:"'user_id'"`
	Username   string     `json:"username" xorm:"username"`
	Type       string     `json:"type" xorm:"type"`
	TaskID     int64      `json:"task_id" xorm:"task_id"`
	Nickname   string     `json:"nickname" xorm:"nickname"`
}

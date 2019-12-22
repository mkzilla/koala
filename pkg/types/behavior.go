package types

import "time"

// 关注 点赞 反对

type Behavior struct {
	ID         int64      `json:"id" xorm:"'id' pk autoincr"`
	CreateTime *time.Time `json:"createTime" xorm:"'create_time' created"`
	UpdateTime *time.Time `json:"createTime" xorm:"'update_time' created"`
	UID        int64      `json:"uid" xorm:"'uid'"`
	TaskID     int64      `json:"taskID" xorm:"'task_id'"`
	Behavior   string     `json:"behavior" xorm:"'behavior'"`
}

// issue

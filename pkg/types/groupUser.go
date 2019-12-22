package types

import (
	"time"
)

type GroupUser struct {
	ID         int64      `json:"id" xorm:"'id' pk autoincr"`
	GroupID    int64      `json:"group_id" xorm:"group_id unique(gu)"`
	UserID     int64      `json:"user_id" xorm:"user_id unique(gu)"`
	CreateTime *time.Time `json:"create_time" xorm:"'createTime' created"`
	Enable     bool       `json:"enable" xorm:"enable"` // 需要用户自己确认是否加入
}

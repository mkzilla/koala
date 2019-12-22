package types

import (
	"time"
)

type Notification struct {
	ID         int64      `json:"id" xorm:"'id' pk autoincr"`
	UserID     int64      `json:"user_id" xorm:"user_id"`
	FromID     int64      `json:"from_id" xorm:"from_id"`
	Info       string     `json:"info" xorm:"info"`
	IsRead     bool       `json:"is_read" xorm:"is_read"`
	CreateTime *time.Time `json:"create_time" xorm:"'create_time' created"`
	Route      string     `json:"route" xorm:"route"`
}

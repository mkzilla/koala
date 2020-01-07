package types

import "time"

type Label struct {
	ID         int64      `json:"id" xorm:"'id' pk autoincr"`
	CreateTime *time.Time `json:"createTime" xorm:"'createTime' created"`
	UpdateTime *time.Time `json:"updateTime" xorm:"'updateTime' updated"`
	Name       string     `json:"name" xorm:"'name'"`
}

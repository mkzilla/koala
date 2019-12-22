package types

import "time"

type Role struct {
	ID         int64      `json:"id" xorm:"'id' pk autoincr"`
	CreateTime *time.Time `json:"createTime" xorm:"'createTime' created"`
	UpdateTime *time.Time `json:"createTime" xorm:"'updateTime' created"`
	Name       string     `json:"name" xorm:"'name'"`
	Desc       string     `json:"desc" xorm:"'desc'"`
}

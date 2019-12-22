package types

import (
	"time"

	"github.com/mkzilla/koala/pkg/types/config"
)

type Group struct {
	ID         int64      `json:"id" xorm:"'id' pk autoincr"`
	CreateTime *time.Time `json:"create_time" xorm:"'createTime' created"`
	UpdateTime *time.Time `json:"update_time" xorm:"'updateTime' created"`
	Name       string     `json:"name" xorm:"'name' unique"`
	Desc       string     `json:"desc" xorm:"'desc'"`
	AdminID    int64      `json:"admin_id" xorm:"admin_id"`
	Admin      string     `json:"admin" xorm:"admin"`
}

func GetGroupByID(id interface{}) (*Group, error) {
	var grp Group
	_, err := config.DBEngine.Where("id = ?", id).Get(&grp)
	return &grp, err
}

package types

import (
	"time"

	"github.com/mkzilla/koala/pkg/types/config"
)

type Plan struct {
	ID         int64      `json:"id" xorm:"'id' pk autoincr"`
	GroupID    int64      `json:"groupID" xorm:"groupID"`
	UserID     int64      `json:"userID" xorm:"userID"`
	Desc       string     `json:"desc" xorm:"'desc' varchar(5000)"`
	CreateTime *time.Time `json:"createTime" xorm:"'createTime' created"`
	UpdateTime *time.Time `json:"updateTime" xorm:"'updateTime' updated"`
}

func InsertPlan(p Plan) (Plan, error) {
	_, err := config.DBEngine.InsertOne(&p)
	return p, err
}

func GetRecentPlanByUser(id interface{}) (Plan, error) {
	var p Plan
	_, err := config.DBEngine.Where("userID = ?", id).Desc("id").Get(&p)
	return p, err
}

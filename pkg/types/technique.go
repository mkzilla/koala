package types

import (
	"time"

	"github.com/mkzilla/koala/pkg/types/config"
)

type Technique struct {
	ID              int64     `json:"id" xorm:"'id' pk autoincr"`
	UserID          int64     `json:"userID" xorm:"userID"`
	CreateTimeStamp int64     `json:"-" xorm:"'createTime' created"`
	CreateTime      time.Time `json:"createTime" xorm:"-"`
	Done            bool      `json:"done" xorm:"done"`
}

func InsertTechnique(t Technique) (Technique, error) {
	_, err := config.DBEngine.InsertOne(&t)
	return t, err
}

func GetCountTechnique(uid int64, t time.Time) (ts []Technique, err error) {
	err = config.DBEngine.Table("technique").Where("userID = ? and createTime > ?", uid, t.Unix()).Find(&ts)
	return ts, err
}

func SetDoneByID(id interface{}) error {
	return nil
}

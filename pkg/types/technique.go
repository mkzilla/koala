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

func GetDoingTechnique(uid int64) (*Technique, error) {
	var t1 = Technique{}
	ext, err := config.DBEngine.Table("technique").Where("userID = ? and done = 0", uid).Get(&t1)
	if !ext {
		return nil, err
	}
	return &t1, err
}

func SetDoneTechniqueByID(uid int64) error {
	_, err := config.DBEngine.Exec("update technique set done = true where userID = ?", uid)
	return err
}

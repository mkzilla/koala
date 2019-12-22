package types

import (
	"github.com/mkzilla/koala/pkg/types/config"
	"time"
)

type Objective struct {
	ID      int64      `json:"id" xorm:"'id' pk autoincr"`
	GroupID int64      `json:"groupID" xorm:"groupID"`
	UserID  int64      `json:"userID" xorm:"userID"`
	Title   string     `json:"title" xorm:"'title' varchar(2000)"`
	Desc    string     `json:"desc" xorm:"'desc' varchar(5000)"`
	Start   *time.Time `json:"start" xorm:"start"`
	End     *time.Time `json:"end" xorm:"end"`
	Score   int        `json:"score" xorm:"score"`
	KRs     string     `json:"keyResults" xorm:"'keyResults' TEXT"`
}

func CreateObjective(obj Objective) (Objective, error) {
	_, err := config.DBEngine.InsertOne(&obj)
	return obj, err
}

func UpdateObjectiveByGroupID(id interface{}, obj Objective) error {
	var o1 = Objective{}
	ext, err := config.DBEngine.Where("groupID = ? and id = ?", id, obj.ID).Get(&o1)
	if err != nil {
		return err
	}
	if !ext || o1.ID != obj.ID {
		return NoSuchObjective
	}
	_, err = config.DBEngine.Update(obj, o1)
	return err
}

func UpdateObjectiveByUserID(id interface{}, obj Objective) error {
	var o1 = Objective{}
	ext, err := config.DBEngine.Where("userID = ? and id = ?", id, obj.ID).Get(&o1)
	if err != nil {
		return err
	}
	if !ext || o1.ID != obj.ID {
		return NoSuchObjective
	}
	_, err = config.DBEngine.Update(obj, o1)
	return err
}

func DeleteObjectiveByID(id int64) {
	//TODO
}

func GetObjectivesByGroup(id interface{}, ps int, po int) ([]Objective, error) {
	var objs []Objective
	err := config.DBEngine.Where("groupID = ?", id).Desc("end").Limit(ps, po).Find(&objs)
	return objs, err
}

func GetObjectivesByUser(id interface{}, ps int, po int) ([]Objective, error) {
	var objs []Objective
	err := config.DBEngine.Where("userID = ?", id).Desc("end").Limit(ps, po).Find(&objs)
	return objs, err
}

func GetDoingObjectivesByUser(id interface{}, ps int, po int) ([]Objective, error) {
	var objs []Objective
	err := config.DBEngine.Where("userID = ? and start < ? and end > ?", id, time.Now(), time.Now()).Desc("end").Limit(ps, po).Find(&objs)
	return objs, err
}

func GetObjectiveByID(id interface{}) (Objective, error) {
	var obj Objective
	_, err := config.DBEngine.Where("id = ?", id).Get(&obj)
	return obj, err
}

func IsGroupObjective(id interface{}, group interface{}) bool {
	var obj Objective
	ext, err := config.DBEngine.Where("id = ? and groupID = ?", id, group).Get(&obj)

	if err == nil && ext {
		return true
	}
	return false
}

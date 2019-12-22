package types

import (
	"encoding/json"
	"time"

	"github.com/mkzilla/koala/pkg/types/config"
)

const (
	TaskComment  = "评论"
	TaskSchedule = "进展"
)

type Comment struct {
	ID         int64      `json:"id" xorm:"'id' pk autoincr"`
	CreateTime *time.Time `json:"create_time,omitempty" xorm:"'create_time' created"`
	Info       string     `json:"info" xorm:"'info' varchar(5000)"`
	UserID     int64      `json:"user_id" xorm:"'user_id'"`
	Username   string     `json:"username" xorm:"username"`
	Type       string     `json:"type" xorm:"type"`
	TaskID     int64      `json:"task_id" xorm:"task_id"`
}

type CommentCount struct {
	Count int64  `json:"count"`
	Date  string `json:"date"`
}

func CommentTimeLine() ([]CommentCount, error) {
	var data []map[string][]byte
	var cc []CommentCount

	var now = time.Now()
	now = now.Add(time.Hour * -24 * 30)
	data, err := config.DBEngine.Query(`select 1 count,DATE_FORMAT( create_time, "%Y-%m-%d" ) date from comment where create_time > ?`, now)
	for _, d := range data {
		var a CommentCount
		json.Unmarshal(d["count"], &a.Count)
		a.Date = string(d["date"])
		if len(cc) == 0 {
			cc = append(cc, a)
		} else if cc[len(cc)-1].Date == a.Date {
			cc[len(cc)-1].Count++
		} else {
			cc = append(cc, a)
		}
	}
	return cc, err
}

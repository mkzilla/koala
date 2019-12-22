package types

import (
	"time"
)

const (
	CreateTaskLog      = "创建任务"
	UpdateTaskStateLog = "变更任务状态"
	AddCommentLog      = "添加评论"
	AddScheduleLog     = "添加进展"
	EditTitleLog       = "修改标题"
	EditDescLog        = "修改内容"
	VoteLog            = "投票"
	WatchLog           = "添加关注"
	AddAssignLog       = "指派"
)

// 聚合时间线策略设计
// 时间线：自己的任务的变动，组成员的活动
// 活动包含：创建任务，修改任务内容，编辑状态，评论、进程、投票、收藏
// 一个任务最好只出现一次，所有操作聚合在一起
// 当一个人占据大量时间线的时候，进行聚合
// 抽象行为
// 变更任务状态 xx yy
// 添加评论
// 添加进展
// 修改标题 xx yy
// 修改内容 xx yy
// （取消时删除）投票 xx
// （取消时删除）收藏 xx

type TaskLog struct {
	ID         int64      `json:"id" xorm:"'id' pk autoincr"`
	CreateTime *time.Time `json:"create_time" xorm:"'create_time' created"`
	UpdateTime *time.Time `json:"update_time" xorm:"'update_time' created"`
	TaskID     int64      `json:"task_id" xorm:"'task_id'"`
	UserID     int64      `json:"user_id" xorm:"'user_id'"`
	Username   string     `json:"username" xorm:"username"`
	Behavior   string     `json:"behavior" xorm:"'behavior'"`
	FromValue  string     `json:"from_value" xorm:"varchar(10000)"`
	ToValue    string     `json:"to_value" xorm:"varchar(10000)"`
}

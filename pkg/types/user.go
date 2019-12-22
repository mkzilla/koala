package types

import (
	"time"

	"github.com/mkzilla/koala/pkg/types/config"
)

type Common struct {
	ID         int64      `json:"id" xorm:"'id' pk autoincr"`
	CreateTime *time.Time `json:"createTime" xorm:"'createTime'"`
	UpdateTime *time.Time `json:"createTime" xorm:"'updateTime'"`
}

type User struct {
	ID         int64      `json:"id" xorm:"'id' pk autoincr"`
	CreateTime *time.Time `json:"createTime" xorm:"'createTime' created"`
	UpdateTime *time.Time `json:"createTime" xorm:"'updateTime' created"`

	Username     string `json:"username" xorm:"'username' unique"`
	Nickname     string `json:"nickname" xorm:"nickname"`
	Email        string `json:"email" xorm:"'email'"`
	GlobalUnique string `json:"globalUnique" xorm:"'globalUnique' unique"`

	Platform string `json:"platform" xorm:"platform"`
	Phone    string `json:"phone" xorm:"'phone'"`
	Enable   bool   `json:"enable" xorm:"'enable'"`
	IsAdmin  bool   `json:"is_admin" xorm:"isAdmin"`
}

func GetCountByUsername(username string) int64 {
	cnt, _ := config.DBEngine.Table("user").Where("username = ?", username).Count()
	return cnt
}

type Password struct {
	ID       int64  `json:"id" xorm:"'id' pk autoincr"`
	UserID   int64  `json:"user_id" xorm:"user_id"`
	Password string `json:"password" xorm:"'password'"`
}

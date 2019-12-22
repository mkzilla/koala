package data

type User struct {
	ID       int64  `json:"id" xorm:"'id' pk autoincr"`
	Username string `json:"username" xorm:"'username'"`
	Nickname string `json:"nickname" xorm:"nickname"`
}

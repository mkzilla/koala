package data

import (
	"github.com/mkzilla/koala/pkg/types"
	"time"
)

type Report struct {
	Task      types.Task      `json:"task"`
	Schedules []types.Comment `json:"schedules"`
}

type Reports struct {
	From    *time.Time `json:"from"`
	To      *time.Time `json:"to"`
	User    types.User `json:"user"`
	Reports []Report   `json:"reports"`
}

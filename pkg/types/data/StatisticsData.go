package data

import . "github.com/mkzilla/koala/pkg/types"

type StatisticsData struct {
	Comments []CommentCount   `json:"comments"`
	Tasks    []TaskGroupCount `json:"tasks"`
}

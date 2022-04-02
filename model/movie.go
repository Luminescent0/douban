package model

import "time"

type Movie struct {
	Id                   int
	Name                 string
	ReleaseDate          time.Time
	ReleaseArea          string
	Genre                string
	Score                float32
	EvaOfNum             string
	DirectorAndLeadActor string
	Rank                 int
	Introduction         string
	CommentNum           int `gorm:"column:comment_num"`
}

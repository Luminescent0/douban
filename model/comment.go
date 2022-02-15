package model

import "time"

type Comment struct {
	Id          int
	Username    string
	ReleaseDate time.Time
	Content     string
	MovieName   string
}

type LongComment struct {
	Id          int
	Promulgator string
	Title       string
	Content     string
	MovieName   string
	ReleaseTime time.Time
}

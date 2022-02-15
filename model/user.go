package model

type User struct {
	Id           int
	Username     string `json:"username" validate:"min=4,max=10"`
	Password     string `json:"password" validate:"min=6,max=16"`
	Introduction string `json:"introduction"`
}

type UserWant struct {
	Label     string
	MovieName string
	Comment   string
	Url       string
	MovieId   int
}

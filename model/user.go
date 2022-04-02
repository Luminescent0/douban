package model

type User struct {
	Id           int    `gorm:"primaryKey"`
	Username     string `json:"username" validate:"min=4,max=10"`
	Password     string `json:"password" validate:"min=6,max=16"`
	Introduction string `json:"introduction"`
	Avatar       string
	Address      string
}

type UserWant struct {
	Username  string
	Label     string
	MovieName string `gorm:"column:movie_name"`
	Comment   string
	Url       string
	MovieId   int `gorm:"column:movie_id"`
}

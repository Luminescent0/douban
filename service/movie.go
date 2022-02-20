package service

import (
	"douban/dao"
	"douban/model"
)

func CheckWantSee(username string) ([]model.UserWant, error) {
	return dao.CheckUserWant(username)
}

func AddWantSee(username, movieName, comment, label string, movieId int) error {
	return dao.AddWantSee(username, movieName, comment, label, movieId)
}
func DeleteWantSee(username, movieName string) error {
	return dao.DeleteWantSee(username, movieName)
}

func Seen(username string) ([]model.UserWant, error) {
	return dao.Seen(username)
}

func AddSeen(username, movieName, comment, label string, movieId int) error {
	return dao.AddSeen(username, movieName, comment, label, movieId)
}

func DeleteSeen(username, movieName string) error {
	return dao.DeleteSeen(username, movieName)
}

func GetMoviesByKeyWord(keyword string) ([]model.Movie, error) {
	return dao.SelectMovies(keyword)
}

func PostDiscussion(promulgator, title, content, movieName string) error {
	return dao.PostDiscussion(promulgator, title, content, movieName)
}

func DeleteDiscussion(promulgator, title, movieName string) error {
	return dao.DeleteDiscussion(promulgator, title, movieName)
}

func GetMovieById(movieId int) (model.Movie, error) {
	return dao.GetMovieById(movieId)
}

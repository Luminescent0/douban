package service

import (
	"douban/dao"
	"douban/model"
)

func GetComment(username string) ([]model.Comment, error) {
	return dao.GetComment(username)
}

func GetLongComment(username string) ([]model.LongComment, error) {
	return dao.GetLongComment(username)
}

func DeleteComment(username, movieName string) error {
	return dao.DeleteComment(username, movieName)
}

func PostLongComment(promulgator, title, content, movieName string) error {
	return dao.PostLongComment(promulgator, title, content, movieName)
}

func PostComment(promulgator, content, movieName string) error {
	return dao.PostComment(promulgator, content, movieName)
}

func PostDisComment(promulgator, comment, movieName, title string) error {
	return dao.PostDisComment(promulgator, comment, movieName, title)
}

func DeleteDisComment(promulgator, movieName, title string) error {
	return dao.DeleteDisComment(promulgator, movieName, title)
}

package service

import (
	"database/sql"
	"douban/dao"
	"douban/model"
)

func IsPasswordCorrect(username, password string) (bool, error) {
	user, err := dao.SelectUserByUsername(username)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}
	if user.Password != password {
		return false, nil
	}
	return true, nil
}

func IsRepeatUsername(username string) (bool, error) {
	_, err := dao.SelectUserByUsername(username)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
func Register(user model.User) error {
	err := dao.InsertUser(user)
	return err
}

func ChangePassword(username, newPassword string) error {
	err := dao.UpdatePassword(username, newPassword)
	return err
}

func CheckIntroduction(username string) (model.User, error) {
	user, err := dao.CheckIntroduction(username)
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

func ChangeIntroduction(username, introduction string) error {
	err := dao.ChangeIntroduction(username, introduction)
	return err
}

func UploadAvatar(username, fileName string) (error, string) {
	err := dao.UploadAvatar(username, fileName)
	if err != nil {
		return err, ""
	}
	return nil, fileName
}

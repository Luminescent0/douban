package service

import (
	"database/sql"
	"douban/dao"
	"douban/model"
	"errors"
	"golang.org/x/crypto/bcrypt"
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
	err := Cipher(user)
	if err != nil {
		return err
	}
	err = dao.InsertUser(user)
	if err != nil {
		return err
	}
	return nil
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

// Bcrypt 密码加盐
type Bcrypt struct {
	cost int
}

func Cipher(user model.User) error {
	hash := Bcrypt{
		cost: bcrypt.DefaultCost,
	}
	_, err := hash.Make([]byte(user.Password))
	if err != nil {
		return errors.New("加密失败")
	}
	return nil
}

func (b *Bcrypt) Make(password []byte) ([]byte, error) {
	return bcrypt.GenerateFromPassword(password, b.cost)
}

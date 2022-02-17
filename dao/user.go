package dao

import (
	"douban/model"
)

func SelectUserByUsername(username string) (model.User, error) {
	user := model.User{}

	err := dB.QueryRow("select id, password from user where username = ?", username).Scan(&user.Id, &user.Password)
	if err != nil {
		return user, err
	}
	return user, nil
}

func InsertUser(user model.User) error {
	_, err := dB.Exec("INSERT INTO user(username, password) "+"values(?, ?);", user.Username, user.Password)
	return err
}

func UpdatePassword(username, newPassword string) error {
	_, err := dB.Exec("UPDATE user SET password = ? WHERE username = ?", newPassword, username)
	return err
}

func CheckIntroduction(username string) (model.User, error) {
	var user = model.User{}
	err := dB.QueryRow("select * from user where username = ?", username).Scan(&user.Id, &user.Password, &user.Introduction)
	if err != nil {
		return user, err
	}
	return user, nil

}

func ChangeIntroduction(username, introduction string) error {
	_, err := dB.Exec("update user set introduction = ? where username = ?", introduction, username)
	return err
}

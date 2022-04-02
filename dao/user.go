package dao

import (
	"douban/model"
	"fmt"
)

func SelectUserByUsername(username string) (model.User, error) {
	user := model.User{}
	err := dB.Table("user").Where("username=?", username).Find(&user)
	//err = dB.QueryRow("select id, password from user where username = ?", username).Scan(&user.Id, &user.Password)
	fmt.Println(err.Error)
	if err.Error != nil {
		return user, err.Error
	}
	return user, nil
}

func InsertUser(user model.User) error {
	err := dB.Table("user").Select("username", "password").Create(&user)
	if err != nil {
		fmt.Println(err.Error)
		return err.Error
	}
	return nil
	//sqlStr := "INSERT INTO user(username, password) " + "values(?, ?);"
	//stmt, err := dB.Prepare(sqlStr)
	//if err != nil {
	//	fmt.Println("prepare failed,err:", err)
	//	return err
	//}
	//defer stmt.Close()
	//_, err = stmt.Exec(user.Username, user.Password)
	//if err != nil {
	//	fmt.Println("insert failed,err", err)
	//}
	//return nil
}

func UpdatePassword(username, newPassword string) error {
	user := model.User{Username: username, Password: newPassword}
	err := dB.Table("user").Model(&user).Where("username = ?", username).Update("password", newPassword)
	if err != nil {
		fmt.Println(err.Error)
		return err.Error
	}
	return nil
	//sqlStr := "UPDATE user SET password = ? WHERE username = ?"
	//stmt, err := dB.Prepare(sqlStr)
	//if err != nil {
	//	fmt.Println("prepare failed,err", err)
	//	return err
	//}
	//defer stmt.Close()
	//_, err = stmt.Exec(newPassword, username)
	//if err != nil {
	//	fmt.Println("update password failed,err", err)
	//}
	//return nil
}

func CheckIntroduction(username string) (model.User, error) {
	var user = model.User{Username: username}
	err := dB.Table("user").Where("username=?", username).Find(&user)
	if err != nil {
		fmt.Println(err.Error)
		return user, err.Error
	}
	return user, nil
	//err := dB.QueryRow("select * from user where username = ?", username).Scan(&user.Id, &user.Password, &user.Introduction)
	//if err != nil {
	//	return user, err
	//}
	//return user, nil

}

func ChangeIntroduction(username, introduction string) error {
	var user = model.User{
		Username:     username,
		Introduction: introduction,
	}
	err := dB.Table("user").Model(&user).Where("username=?", username).Update("introduction", introduction)
	if err != nil {
		fmt.Println(err.Error)
		return err.Error
	}
	return nil
	//sqlStr := "update user set introduction = ? where username = ?"
	//stmt, err := dB.Prepare(sqlStr)
	//if err != nil {
	//	fmt.Println("prepare failed,err", err)
	//	return err
	//}
	//defer stmt.Close()
	//_, err = stmt.Exec(introduction, username)
	//if err != nil {
	//	fmt.Println("change introduction failed err", err)
	//}
	//return nil
}

//func UploadAvatar(username, loadString, fileAddress string) error {
//	sqlStr := "update user set avatar = ?,address =? where username=?"
//	stmt, err := dB.Prepare(sqlStr)
//	if err != nil {
//		fmt.Println("prepare failed,err", err)
//		return err
//	}
//	defer stmt.Close()
//	_, err = stmt.Exec(loadString, fileAddress, username)
//	if err != nil {
//		fmt.Println("upload avatar failed,err", err)
//		return err
//	}
//	return nil
//}

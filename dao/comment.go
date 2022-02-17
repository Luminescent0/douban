package dao

import (
	"douban/model"
	"fmt"
)

func GetComment(username string) ([]model.Comment, error) {
	var comments []model.Comment
	sqlStr := "select movie_name,content from comment where username=?"
	stmt, err := dB.Prepare(sqlStr)
	if err != nil {
		fmt.Println("prepare failed,err:", err)
		return comments, err
	}
	defer stmt.Close()
	rows, err := stmt.Query(username)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var comment model.Comment
		err = rows.Scan(&comment.Content, &comment.MovieName)
		if err != nil {
			return nil, err

		}
		comments = append(comments, comment)
	}
	return comments, nil
}

func DeleteComment(username, movieName string) error {
	_, err := dB.Exec("delete from comment where username=? and movie_name=?", username, movieName)
	if err != nil {
		fmt.Println("delete comment failed,err", err)
		return err
	}
	var commentNum int
	err = dB.QueryRow("select evaOfNum from movies where name=?", movieName).Scan(&commentNum)
	if err != nil {
		return err
	}
	commentNum -= 1
	_, err = dB.Exec("update movies set comment_num=?where name=?", commentNum, movieName)
	if err != nil {
		return err
	}

	return nil
}

func PostLongComment(promulgator, title, content, movieName string) error {
	sqlStr := "insert into long_comment(promulgator, title, content, movie_name)value(?,?,?,?)"
	stmt, err := dB.Prepare(sqlStr)
	if err != nil {
		fmt.Println("prepare failed,err:", err)
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(promulgator, title, content, movieName)
	if err != nil {
		fmt.Println("post long comment failed,err:", err)
		return err
	}
	return nil
}

func GetLongComment(username string) ([]model.LongComment, error) {
	var comments []model.LongComment
	sqlStr := "select title,movie_name,content from long_comment where promulgator=?"
	stmt, err := dB.Prepare(sqlStr)
	if err != nil {
		fmt.Println("prepare failed,err:", err)
		return comments, err
	}
	defer stmt.Close()
	rows, err := stmt.Query(username)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var comment model.LongComment
		err = rows.Scan(&comment.Title, &comment.Content, &comment.MovieName)
		if err != nil {
			return nil, err

		}
		comments = append(comments, comment)
	}
	return comments, nil
}

func PostComment(promulgator, content, movieName string) error {
	sqlStr := "insert into comment(username, content, movie_name)value(?,?,?)"
	stmt, err := dB.Prepare(sqlStr)
	if err != nil {
		fmt.Println("prepare failed,err", err)
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(promulgator, content, movieName)
	if err != nil {
		fmt.Println("post comment failed,err", err)
		return err
	}
	return nil
}

func PostDisComment(promulgator, comment, movieName, title string) error {
	sqlstr := "insert into dis_comment(promulgator,content,movie_name,title)value (?,?,?,?)"
	stmt, err := dB.Prepare(sqlstr)
	if err != nil {
		fmt.Println("prepare failed,err", err)
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(promulgator, comment, movieName, title)
	if err != nil {
		fmt.Println("post disComment failed,err:", err)
	}
	return nil
}

func DeleteDisComment(promulgator, movieName, title string) error {
	sqlStr := "delete from dis_comment where promulgator=? and movie_name=?and title=?"
	stmt, err := dB.Prepare(sqlStr)
	if err != nil {
		fmt.Println("prepare failed,err", err)
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(promulgator, movieName, title)
	if err != nil {
		fmt.Println("delete disComment,err:", err)
		return err
	}
	return nil

}

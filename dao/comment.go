package dao

import (
	"douban/model"
	"fmt"
)

func GetComment(username string) ([]model.Comment, error) {
	var comments []model.Comment
	rows, err := dB.Query("select movie_name,content from comment where username=?", username)
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
	_, err := dB.Exec("insert into long_comment(promulgator, title, content, movie_name)value(?,?,?,?)", promulgator, title, content, movieName)
	return err
}

func GetLongComment(username string) ([]model.LongComment, error) {
	var comments []model.LongComment
	rows, err := dB.Query("select title,movie_name,content from long_comment where promulgator=?", username)
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
	_, err := dB.Exec("insert into comment(username, content, movie_name)value(?,?,?)", promulgator, content, movieName)
	return err
}

func PostDisComment(promulgator, comment, movieName, title string) error {
	_, err := dB.Exec("insert into dis_comment(promulgator,content,movie_name,title)value (?,?,?,?)", promulgator, comment, movieName, title)
	return err
}

func DeleteDisComment(promulgator, movieName, title string) error {
	_, err := dB.Exec("delete from dis_comment where promulgator=? and movie_name=?and title=?", promulgator, movieName, title)
	return err

}

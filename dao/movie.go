package dao

import (
	"douban/model"
	"fmt"
	"strconv"
)

func CheckUserWant(username string) ([]model.UserWant, error) {
	var wants []model.UserWant
	sqlStr := "select label,comment,movie_id,movie_name from wantsee where username=?"
	stmt, err := dB.Prepare(sqlStr)
	if err != nil {
		fmt.Println("prepare failed,err", err)
	}
	defer stmt.Close()
	rows, err := stmt.Query(username)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var want model.UserWant
		err = rows.Scan(&want.Label, &want.Comment, &want.MovieId, &want.MovieName)
		want.Url = "http://121.4.229.95:8080/movieGet/" + strconv.Itoa(want.MovieId)
		if err != nil {
			return nil, err

		}
		wants = append(wants, want)
	}
	return wants, nil
}

func AddWantSee(username, movieName, comment, label string, movieId int) error {
	sqlStr := "insert into wantsee(movie_name,username, label, comment,movie_id) value (?,?,?,?,?)"
	stmt, err := dB.Prepare(sqlStr)
	if err != nil {
		fmt.Println("prepare failed,err:", err)
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(movieName, username, label, comment, movieId)
	if err != nil {
		fmt.Println("insert failed err", err)
		return err
	}

	return nil
}
func DeleteWantSee(username, movieName string) error {
	sqlStr := "delete from wantsee where username=? and movie_name=?"
	stmt, err := dB.Prepare(sqlStr)
	if err != nil {
		fmt.Println("prepare failed,err:", err)
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(username, movieName)
	if err != nil {
		fmt.Println("delete  wantSee failed,err", err)
		return err
	}
	return nil
}

func Seen(username string) ([]model.UserWant, error) {
	var seens []model.UserWant
	sqlStr := "select  label,comment,movie_name,movie_id from wantsee where username=?"
	stmt, err := dB.Prepare(sqlStr)
	if err != nil {
		fmt.Println("prepare failed,err:", err)
		return seens, err
	}
	defer stmt.Close()
	rows, err := stmt.Query(username)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var seen model.UserWant
		err = rows.Scan(&seen.Label, &seen.Comment, &seen.MovieName, &seen.MovieId)
		if err != nil {
			return nil, err
		}
		seen.Url = "http://121.4.229.95/movieGet/" + strconv.Itoa(seen.MovieId)
		seens = append(seens, seen)
	}
	return seens, nil
}

func AddSeen(username, movieName, comment, label string, movieId int) error {
	sqlStr := "insert into seen(movie_name,username, label, comment,movie_id) value (?,?,?,?,?)"
	stmt, err := dB.Prepare(sqlStr)
	if err != nil {
		fmt.Println("prepare failed,err:", err)
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(movieName, username, label, comment, movieId)
	if err != nil {
		fmt.Println("insert failed err", err)
		return err
	}
	return nil
}

func DeleteSeen(username, movieName string) error {
	sqlStr := "delete from seen where username=? and movie_name=?"
	stmt, err := dB.Prepare(sqlStr)
	if err != nil {
		fmt.Println("prepare failed,err:", err)
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(username, movieName)
	if err != nil {
		fmt.Println("delete seen failed,err", err)
		return err
	}
	return nil
}

func SelectMovies(keyword string) ([]model.Movie, error) {
	var movies []model.Movie
	rows, err := dB.Query("select * from movies where movies.name like ? or directorAndLead_actor like ?;",
		"%"+keyword+"%", "%"+keyword+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var movie model.Movie
		err = rows.Scan(&movie.Id, &movie.DirectorAndLeadActor, &movie.Genre, &movie.EvaOfNum, &movie.Score, &movie.Introduction, &movie.Name, &movie.Rank, &movie.ReleaseDate, &movie.ReleaseArea)
		if err != nil {
			return nil, err
		}
		movies = append(movies, movie)

	}
	return movies, nil
}

func PostDiscussion(promulgator, title, content, movieName string) error {
	sqlStr := "insert into discussion(promulgator, title, content, movie_name)value(?,?,?,?)"
	stmt, err := dB.Prepare(sqlStr)
	if err != nil {
		fmt.Println("prepare failed,err:", err)
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(promulgator, title, content, movieName)
	if err != nil {
		fmt.Println("post discussion failed,err", err)
		return err
	}
	return nil
}

func DeleteDiscussion(promulgator, title, movieName string) error {
	sqlStr := "delete from discussion where promulgator=? and movie_name=? and title=?"
	stmt, err := dB.Prepare(sqlStr)
	if err != nil {
		fmt.Println("prepare failed,err:", err)
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(promulgator, movieName, title)
	if err != nil {
		fmt.Println("delete discussion failed,err", err)
		return err
	}
	return nil
}

func GetMovie(movieName string) (model.Movie, error) {
	var movie = model.Movie{}
	err := dB.QueryRow("select * from movies where name = ?", movieName).Scan(&movie.Id, &movie.Score, &movie.Rank, &movie.Genre, &movie.Introduction,
		&movie.ReleaseArea, &movie.ReleaseDate, &movie.EvaOfNum, &movie.DirectorAndLeadActor)
	if err != nil {
		return movie, err
	}
	return movie, nil
}

package dao

import (
	"douban/model"
	"fmt"
	"strconv"
)

func CheckUserWant(username string) ([]model.UserWant, error) {
	var wants []model.UserWant
	rows, err := dB.Query("select label,comment,movie_id,movie_name from wantsee where username=?", username)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var want model.UserWant
		err = rows.Scan(&want.Label, &want.Comment, &want.MovieId, &want.MovieName)
		want.Url = "http://127.0.0.1:8080/movieGet/" + strconv.Itoa(want.MovieId)
		if err != nil {
			return nil, err

		}
		wants = append(wants, want)
	}
	return wants, nil
}

func AddWantSee(username, movieName, comment, label string, movieId int) error {
	ret, err := dB.Exec("insert into wantsee(movie_name,username, label, comment,movie_id) value (?,?,?,?,?)",
		movieName, username, label, comment, movieId)
	if err != nil {
		fmt.Println("insert failed err", err)
		return err
	}
	_, err = ret.LastInsertId()
	if err != nil {
		fmt.Println("get lastinsert ID failed,err", err)
		return err
	}
	return nil
}
func DeleteWantSee(username, movieName string) error {
	_, err := dB.Exec("delete from wantsee where username=? and movie_name=?", username, movieName)
	if err != nil {
		fmt.Println("delete  wantSee failed,err", err)
		return err
	}
	return nil
}

func Seen(username string) ([]model.UserWant, error) {
	var seens []model.UserWant
	rows, err := dB.Query("select  label,comment,movie_name,movie_id from wantsee where username=?", username)
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
		seen.Url = "http://127.0.0.1/movieGet/" + strconv.Itoa(seen.MovieId)
		seens = append(seens, seen)
	}
	return seens, nil
}

func AddSeen(username, movieName, comment, label string, movieId int) error {
	_, err := dB.Exec("insert into seen(movie_name,username, label, comment,movie_id) value (?,?,?,?,?)", movieName, username, label, comment, movieId)
	if err != nil {
		fmt.Println("insert failed err", err)
		return err
	}
	//_,err = ret.LastInsertId()
	//if err != nil {
	//	fmt.Println("get lastinsert ID failed,err",err)
	//	return err
	//}
	return nil
}

func DeleteSeen(username, movieName string) error {
	_, err := dB.Exec("delete from seen where username=? and movie_name=?", username, movieName)
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
	_, err := dB.Exec("insert into discussion(promulgator, title, content, movie_name)value(?,?,?,?)", promulgator, title, content, movieName)
	return err
}

func DeleteDiscussion(promulgator, title, movieName string) error {
	_, err := dB.Exec("delete from discussion where promulgator=? and movie_name=? and title=?", promulgator, movieName, title)
	if err != nil {
		fmt.Println("delete discussion failed,err", err)
		return err
	}
	return nil
}

func GetMovie(movieName string) (model.Movie, error) {
	var movie = model.Movie{}
	id := dB.QueryRow("select * from movies where name = ?", movieName)
	if id.Err() != nil {
		return movie, id.Err()
	}
	err := id.Scan(&movie.Id, &movie.Score, &movie.Rank, &movie.Genre, &movie.Introduction,
		&movie.ReleaseArea, &movie.ReleaseDate, &movie.EvaOfNum, &movie.DirectorAndLeadActor)
	if err != nil {
		return movie, err
	}
	return movie, nil
}

func GetMovieById(movieId int) (model.Movie, error) {
	var movie = model.Movie{}
	id := dB.QueryRow("select * from movies where id = ?", movieId)
	if id.Err() != nil {
		return movie, id.Err()
	}
	err := id.Scan(&movie.Id, &movie.Score, &movie.Rank, &movie.Genre, &movie.Introduction,
		&movie.ReleaseArea, &movie.ReleaseDate, &movie.EvaOfNum, &movie.DirectorAndLeadActor)
	if err != nil {
		return movie, err
	}
	return movie, nil
}

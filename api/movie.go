package api

import (
	"douban/service"
	"douban/tool"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

func wantSee(ctx *gin.Context) {
	username := ctx.Param("username")
	wants, err := service.CheckWantSee(username)
	if err != nil {
		fmt.Println("get user's wantSee failed err:", err)
		tool.RespInternalError(ctx)
	}
	tool.RespSuccessfulWithDate(ctx, wants)
}

func addWantSee(ctx *gin.Context) {
	iUsername, _ := ctx.Get("username")
	username := iUsername.(string)
	label := ctx.PostForm("label")
	comment := ctx.PostForm("comment")
	movieID := ctx.Param("movieId")
	movieId, _ := strconv.Atoi(movieID)
	movie, err := service.GetMovieById(movieId)
	if err != nil {
		tool.RespErrorWithDate(ctx, "数据库中查询不到该电影")
	}
	movieName := movie.Name
	err = service.AddWantSee(username, movieName, comment, label, movieId)
	if err != nil {
		fmt.Println("add wantSee failed err", err)
		tool.RespInternalError(ctx)
	}
	tool.RespSuccessful(ctx)
}

func deleteWantSee(ctx *gin.Context) {
	iUsername, _ := ctx.Get("username")
	username := iUsername.(string)
	movieID := ctx.Param("movieId")
	movieId, _ := strconv.Atoi(movieID)
	movie, err := service.GetMovieById(movieId)
	if err != nil {
		tool.RespErrorWithDate(ctx, "数据库中查询不到该电影")
	}
	movieName := movie.Name
	err = service.DeleteWantSee(username, movieName)
	if err != nil {
		fmt.Println("delete Wantsee failed err", err)
		tool.RespInternalError(ctx)
	}
	tool.RespSuccessful(ctx)
}

func seen(ctx *gin.Context) {
	username := ctx.Param("username")
	seen, err := service.Seen(username)
	if err != nil {
		fmt.Println("get user's seen failed err:", err)
		tool.RespInternalError(ctx)
	}
	tool.RespSuccessfulWithDate(ctx, seen)
}

func addSeen(ctx *gin.Context) {
	iUsername, _ := ctx.Get("username")
	username := iUsername.(string)
	label := ctx.PostForm("label")
	comment := ctx.PostForm("comment")
	movieID := ctx.Param("movieId")
	movieId, _ := strconv.Atoi(movieID)
	movie, err := service.GetMovieById(movieId)
	if err != nil {
		tool.RespErrorWithDate(ctx, "数据库中查询不到该电影")
	}
	movieName := movie.Name
	err = service.AddSeen(username, movieName, comment, label, movieId)
	if err != nil {
		fmt.Println("add seen failed err", err)
		tool.RespInternalError(ctx)
	}
	tool.RespSuccessful(ctx)
}

func deleteSeen(ctx *gin.Context) {
	iUsername, _ := ctx.Get("username")
	username := iUsername.(string)
	movieID := ctx.Param("movieId")
	movieId, _ := strconv.Atoi(movieID)
	movie, err := service.GetMovieById(movieId)
	if err != nil {
		tool.RespErrorWithDate(ctx, "数据库中查询不到该电影")
	}
	movieName := movie.Name
	err = service.DeleteSeen(username, movieName)
	if err != nil {
		fmt.Println("delete seen failed err", err)
		tool.RespInternalError(ctx)
	}
	tool.RespSuccessful(ctx)
}

func search(ctx *gin.Context) {
	keyword := ctx.PostForm("")
	movies, err := service.GetMoviesByKeyWord(keyword)
	if err != nil {
		fmt.Println("get movies err :", err)
		tool.RespInternalError(ctx)
		return
	}
	if movies == nil {
		fmt.Println("not found movie")
		tool.RespSuccessfulWithDate(ctx, gin.H{"msg": "没有找到相关电影，换个搜索词试试吧。"})
	}
	tool.RespSuccessfulWithDate(ctx, movies)
}

func postDiscussion(ctx *gin.Context) {
	iUsername, _ := ctx.Get("username")
	promulgator := iUsername.(string)
	title := ctx.PostForm("title")
	content := ctx.PostForm("content")
	movieID := ctx.Param("movieId")
	movieId, _ := strconv.Atoi(movieID)
	movie, err := service.GetMovieById(movieId)
	if err != nil {
		tool.RespErrorWithDate(ctx, "数据库中查询不到该电影")
	}
	movieName := movie.Name
	err = service.PostDiscussion(promulgator, title, content, movieName)
	if err != nil {
		tool.RespSuccessfulWithDate(ctx, "上传失败")
	}
	tool.RespSuccessfulWithDate(ctx, "上传成功")

}
func deleteDiscussion(ctx *gin.Context) {
	iUsername, _ := ctx.Get("username")
	promulgator := iUsername.(string)
	title := ctx.PostForm("title")
	movieID := ctx.Param("movieId")
	movieId, _ := strconv.Atoi(movieID)
	movie, err := service.GetMovieById(movieId)
	if err != nil {
		tool.RespErrorWithDate(ctx, "数据库中查询不到该电影")
	}
	movieName := movie.Name
	err = service.DeleteDiscussion(promulgator, title, movieName)
	if err != nil {
		tool.RespSuccessfulWithDate(ctx, "删除失败")
	}
	tool.RespSuccessfulWithDate(ctx, "删除成功")
}

func getMovieInfo(ctx *gin.Context) {
	movieId := ctx.Param("movieId")
	movieID, err := strconv.Atoi(movieId)
	if err != nil {
		tool.RespInternalError(ctx)
		return
	}
	movie, err1 := service.GetMovieById(movieID)
	if err1 != nil {
		tool.RespSuccessfulWithDate(ctx, "无法获取电影相关信息")
		return
	}
	tool.RespSuccessfulWithDate(ctx, movie)

}

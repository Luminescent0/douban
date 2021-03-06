package api

import (
	"douban/service"
	"douban/tool"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

func getComment(ctx *gin.Context) {
	username := ctx.Param("username")
	comments, err := service.GetComment(username)
	if err != nil {
		tool.RespInternalError(ctx)
	}
	if comments == nil {
		tool.RespSuccessfulWithDate(ctx, gin.H{"msg": "您还没有发布过短评"})
	}
	tool.RespSuccessfulWithDate(ctx, comments)

}

func getLongComment(ctx *gin.Context) {
	username := ctx.Param("username")
	comments, err := service.GetLongComment(username)
	if err != nil {
		tool.RespInternalError(ctx)
	}
	if comments == nil {
		tool.RespSuccessfulWithDate(ctx, gin.H{"msg": "您还没有发布过影评"})
	}
	tool.RespSuccessfulWithDate(ctx, comments)

}

func deleteComment(ctx *gin.Context) {
	iUsername, _ := ctx.Get("username")
	username := iUsername.(string)
	movieID := ctx.Param("movieId")
	movieId, _ := strconv.Atoi(movieID)
	movie, err := service.GetMovieById(movieId)
	if err != nil {
		tool.RespErrorWithDate(ctx, "数据库中查询不到该电影")
	}
	movieName := movie.Name
	err = service.DeleteComment(username, movieName)
	if err != nil {
		fmt.Println("delete comment failed err", err)
		tool.RespSuccessfulWithDate(ctx, "删除失败")
	}
	tool.RespSuccessful(ctx)
}

func postLongComment(ctx *gin.Context) {
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
	err = service.PostLongComment(promulgator, title, content, movieName)
	if err != nil {
		tool.RespSuccessfulWithDate(ctx, "上传失败")
	}
	tool.RespSuccessfulWithDate(ctx, "上传成功")
}

func postComment(ctx *gin.Context) {
	iUsername, _ := ctx.Get("username")
	promulgator := iUsername.(string)
	content := ctx.PostForm("content")
	movieID := ctx.Param("movieId")
	movieId, _ := strconv.Atoi(movieID)
	movie, err := service.GetMovieById(movieId)
	if err != nil {
		tool.RespErrorWithDate(ctx, "数据库中查询不到该电影")
	}
	movieName := movie.Name
	err = service.PostComment(promulgator, content, movieName)
	if err != nil {
		tool.RespSuccessfulWithDate(ctx, "上传失败")
	}
	tool.RespSuccessfulWithDate(ctx, "上传成功")
}

func deleteLongComment(ctx *gin.Context) {
	iUsername, _ := ctx.Get("username")
	username := iUsername.(string)
	movieID := ctx.Param("movieId")
	movieId, _ := strconv.Atoi(movieID)
	movie, err := service.GetMovieById(movieId)
	if err != nil {
		tool.RespErrorWithDate(ctx, "数据库中查询不到该电影")
	}
	movieName := movie.Name
	err = service.DeleteComment(username, movieName)
	if err != nil {
		fmt.Println("delete comment failed err", err)
		tool.RespSuccessfulWithDate(ctx, "删除失败")
	}
	tool.RespSuccessfulWithDate(ctx, "删除成功")
}

func postDisComment(ctx *gin.Context) {
	title := ctx.PostForm("title")     //讨论的标题
	comment := ctx.PostForm("comment") //评论内容
	iUsername, _ := ctx.Get("username")
	movieID := ctx.Param("movieId")
	movieId, _ := strconv.Atoi(movieID)
	movie, err := service.GetMovieById(movieId)
	if err != nil {
		tool.RespErrorWithDate(ctx, "数据库中查询不到该电影")
	}
	movieName := movie.Name
	promulgator := iUsername.(string)
	err = service.PostDisComment(promulgator, comment, movieName, title)
	if err != nil {
		fmt.Println(err)
		tool.RespSuccessfulWithDate(ctx, "评论失败")
	}
	tool.RespSuccessfulWithDate(ctx, "评论成功")
}

func deleteDisComment(ctx *gin.Context) {
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
	err = service.DeleteDisComment(promulgator, movieName, title)
	if err != nil {
		fmt.Println(err)
		tool.RespSuccessfulWithDate(ctx, "删除失败")
	}
	tool.RespSuccessfulWithDate(ctx, "删除成功")
}

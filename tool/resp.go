package tool

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RespErrorWithDate(ctx *gin.Context, date interface{}) {
	ctx.JSON(http.StatusOK, gin.H{
		"info": date,
	})
}

func RespInternalError(ctx *gin.Context) {
	ctx.JSON(http.StatusInternalServerError, gin.H{"info": "服务器错误"})
}

func RespSuccessful(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"info": "成功"})
}

func RespSuccessfulWithDate(ctx *gin.Context, date interface{}) {
	ctx.JSON(http.StatusOK, gin.H{
		"info": "成功",
		"date": date,
	})
}

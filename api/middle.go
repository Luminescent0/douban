package api

import (
	"douban/tool"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"strings"
)

type MyClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

var MySecret = []byte("xian1")

func CreateToken(username string) (string, error) {
	//创建自己的声明

	CreateClaims := MyClaims{
		username,
		jwt.StandardClaims{
			NotBefore: time.Now().Unix(),
			ExpiresAt: time.Now().Unix() + 60*60*2,
			Issuer:    "douban",
			Subject:   "xian",
		},
	}
	//使用指定的签名方法创建对象
	reqClaim := jwt.NewWithClaims(jwt.SigningMethodHS256, CreateClaims)
	token, err := reqClaim.SignedString(MySecret)
	if err != nil {
		return "", err
	}
	return token, nil
}

func JwtAuthMiddleware(ctx *gin.Context) {
	//假设token放在Header的Authorization中，并使用Bearer开头
	authHeader := ctx.Request.Header.Get("Authorization")
	if authHeader == "" {
		tool.RespSuccessfulWithDate(ctx, gin.H{"msg": "请求头中auth为空"})
		ctx.Abort()
		return
	}
	//按空格分割
	parts := strings.SplitN(authHeader, "", 2)
	if len(parts) != 2 && parts[0] != "Bearer" {
		tool.RespSuccessfulWithDate(ctx, gin.H{"msg": "请求头中auth格式有误"})
		ctx.Abort()
		return
	}
	token, err := jwt.ParseWithClaims(authHeader, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return MySecret, nil
	})
	username := token.Claims.(*MyClaims).Username
	//Time := token.Claims.(*MyClaims).ExpiresAt
	//if Time < time.Now().Unix() {
	//	tool.RespSuccessfulWithDate(ctx, gin.H{"msg": "token过期"})
	//	ctx.Abort()
	//	return
	//}
	if err != nil {
		fmt.Println("parse token failed err", err)
		tool.RespInternalError(ctx)
		return
	}
	if token.Valid == false {
		tool.RespSuccessfulWithDate(ctx, gin.H{"msg": "token无效"})
		ctx.Abort()
		return
	}

	//将当前请求的username信息保存到请求的上下文中
	ctx.Set("username", username)
	ctx.Next() //后续的处理函数可以通过ctx.Get()来获取当前请求的用户信息

}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		if method != "" {
			c.Header("Access-Control-Allow-Origin", c.GetHeader("origin"))
			//c.Header("Access-Control-Allow-Origin","*")
			c.Header("Access-Control-Allow-Methods", "POST,GET,DELETE,OPTIONS,PUT")
			c.Header("Access-Control-Allow-Headers", "Origin,X-Requested-With,Content-Type,Accept,Authorization")
			c.Header("Access-Control-Expose-Headers", "Content-Length,Access-Control-Allow-Origin,Access-Control-Allow-Headers,Cache-Control,Content-Language")
			c.Header("Access-Control-Allow-Credentials", "true")
			//c.Set("content-type","application/json")
		}
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}

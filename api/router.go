package api

import (
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func InitEngine() {
	engine := gin.Default()

	engine.Use(static.Serve("/", static.LocalFile("/usr/share/nginx/html", false)))

	engine.Use(Cors())
	engine.POST("/login", login)       //登录
	engine.POST("/register", register) //注册

	userGroup := engine.Group("/user")
	{
		userGroup.Use(JwtAuthMiddleware)
		userGroup.POST("/password", changePassword)                   //修改密码
		userGroup.POST("/avatar", uploadAvatar)                       //上传头像
		userGroup.POST("/:username/introduction", changeIntroduction) //自我介绍

	}
	userInfo := engine.Group("/user/:username")
	{
		userInfo.GET("/avatar", Avatar)              //头像
		userInfo.GET("/introduction", introduction)  //自我介绍
		userInfo.GET("/wantSee", wantSee)            //用户想看
		userInfo.GET("/seen", seen)                  //用户看过
		userInfo.GET("/comment", getComment)         //用户短评
		userInfo.GET("/longComment", getLongComment) //用户影评

	}
	movie := engine.Group("/movie/:movieId")
	{
		movie.Use(JwtAuthMiddleware)
		movie.POST("/wantSee", addWantSee) //想看
		movie.DELETE("/wantSee", deleteWantSee)

		movie.POST("/seen", addSeen) //看过
		movie.DELETE("/seen", deleteSeen)

		movie.DELETE("/comment", deleteComment) //短评
		movie.POST("/comment", postComment)

		movie.POST("/longComment", postLongComment) //影评
		movie.DELETE("/longComment", deleteLongComment)

		movieDis := movie.Group("/discussion/:movieId")
		{
			movieDis.Use(JwtAuthMiddleware)
			movieDis.POST("/", postDiscussion) //讨论区
			movieDis.DELETE("/", deleteDiscussion)
			movieDis.POST("/dis_comment", postDisComment) //讨论区评论
			movieDis.DELETE("/dis_comment", deleteDisComment)
		}

	}

	movieGetGroup := engine.Group("/movieGet")
	{
		movieGetGroup.GET("/search", search)         //搜索
		movieGetGroup.GET("/:movieId", getMovieInfo) //获取电影详情

	}

	err := engine.Run(":8090")
	if err != nil {
		return
	}
}

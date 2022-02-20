package api

import (
	"github.com/gin-gonic/gin"
)

func InitEngine() {
	engine := gin.Default()

	engine.Use(Cors())
	engine.POST("/login", login)       //登录
	engine.POST("/register", register) //注册

	userGroup := engine.Group("/user")
	{
		userGroup.Use(JwtAuthMiddleware)
		userGroup.POST("/password", changePassword) //修改密码
		userGroup.POST("/avatar", uploadAvatar)     //上传头像

		userGroup.GET("/:username/menu/introduction", introduction) //个人介绍
		userGroup.POST("/:username/menu/introduction", changeIntroduction)

		userGroup.GET("/:username/wantSee", wantSee)            //用户想看
		userGroup.GET("/:username/seen", seen)                  //用户看过
		userGroup.GET("/:username/comment", getComment)         //用户短评
		userGroup.GET("/:username/longComment", getLongComment) //用户影评

	}
	movie := engine.Group("/movie")
	{
		movie.Use(JwtAuthMiddleware)
		movie.POST("/:username/:movieId/wantSee", addWantSee) //想看
		movie.DELETE("/:username/:movieId/wantSee", deleteWantSee)

		movie.POST("/:username/:movieId/seen", addSeen) //看过
		movie.DELETE("/:username/:movieId/seen", deleteSeen)

		movie.DELETE("/:username/:movieId/comment", deleteComment) //短评
		movie.POST("/:username/:movieId/comment", postComment)

		movie.POST("/:username/:movieId/longComment", postLongComment) //影评
		movie.DELETE("/:username/:movieId/longComment", deleteLongComment)

		movieDis := movie.Group("/discussion")
		{
			movieDis.POST("/:username/discussion", postDiscussion) //讨论区
			movieDis.DELETE("/:username/discussion", deleteDiscussion)
			movieDis.POST("/:username/dis_comment", postDisComment) //讨论区评论
			movieDis.DELETE("/:username/dis_comment", deleteDisComment)
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

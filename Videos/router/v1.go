package router

import (
	"github.com/gin-gonic/gin"
	"github.com/videos/API"
	"github.com/videos/middleware"
)

// Setup 启动路由
func Setup() {
	// router
	r := gin.Default()

	// v1版 http请求响应
	v1 := r.Group("/api/v1")
	{
		//用户模块
		v1.POST("/login", API.LoginUser)                        //登录
		v1.POST("/register", API.RegisterUser)                  //注册
		v1.Use(middleware.JWT1()).POST("/setUser", API.SetInfo) //修改个人资料
		v1.Use(middleware.JWT1()).GET("/showUser", API.Show)    //展示个人资料
		//视频模块
		v1.Use(middleware.JWT1()).POST("/upload", API.Upload)        //上传视频
		v1.Use(middleware.JWT1()).POST("/operate", API.VideoOperate) //视频操作
		//搜索
		v1.Use(middleware.JWT1()).GET("/searchVideo", API.SearchVideo) //搜索视频
		v1.Use(middleware.JWT1()).GET("/searchUser", API.SearchUser)   //搜索用户
		v1.Use(middleware.JWT1()).GET("/selectVideo", API.SelectVideo) //筛选视频
		v1.Use(middleware.JWT1()).GET("/showRecord", API.ShowRecord)   //历史搜索记录
	}
	v2 := r.Group("/api/v1/manager")
	{
		//管理员模块
		v2.POST("/login", API.LoginManager)                                 //管理员登录
		v2.POST("/register", API.RegisterManager)                           //管理员注册
		v2.Use(middleware.JWT2()).POST("/passVideo", API.PassVideo)         //审核视频
		v2.Use(middleware.JWT2()).POST("/operateUser", API.ManageUser)      //管理视频
		v2.Use(middleware.JWT2()).POST("/deleteComment", API.DeleteComment) //删除评论

	}

	r.Run(":8080")
}

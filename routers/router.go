package routers

import (
	"net/http"

	"github.com/YogeLiu/CloudDisk/middle"
	"github.com/YogeLiu/CloudDisk/routers/api"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	engine := gin.Default()
	v1 := engine.Group("/pan/v1")
	site := v1.Group("site")
	{
		site.GET("ping", func(ctx *gin.Context) { ctx.JSON(http.StatusOK, gin.H{"msg": "pong"}) })
	}
	file := v1.Group("file")
	{
		upload := file.Group("upload").Use(middle.CheckPerm(gin.Mode()))
		{
			upload.POST("session")
			upload.DELETE("/:seesion_id")
		}
		// 更新文件
		file.PUT("/update/:id")
		// 删除文件
		file.DELETE("/:id")
		// 获取文件
		file.GET("")
	}
	directory := v1.Group("direction").Use(middle.CheckPerm(gin.Mode()))
	{
		// 创建目录
		directory.POST("")
	}
	user := v1.Group("user")
	{
		user.POST("register", api.Register)
		user.POST("login", api.Login)
	}
	// 回调
	callback := v1.Group("callback").Use(middle.CheckPerm(gin.Mode()))
	{
		// 上传回调
		callback.POST("/:session_id")
	}
	return engine
}

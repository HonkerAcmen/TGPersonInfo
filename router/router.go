package router

import (
	"TGPersonInfo/api"
	"TGPersonInfo/middleware"

	"github.com/gin-gonic/gin"
)

func CollectRouter(r *gin.Engine) *gin.Engine {
	r.Use(middleware.CORSMiddleware())
	v1 := r.Group("/api")
	{
		// v1.POST("/userinfo", middleware.AuthmiddleWare(), controller.DTOUserInfo())
		v1.GET("/getuser", api.GetUserInfo())
		v1.POST("/setuser", api.SetUserInfo())
	}
	return r
}

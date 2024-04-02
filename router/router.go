package router

import (
	"TGPersonInfo/controller"
	"TGPersonInfo/middleware"

	"github.com/gin-gonic/gin"
)

func CollectRouter(r *gin.Engine) *gin.Engine {
	r.Use(middleware.CORSMiddleware())
	r.Use(middleware.IsCreateUserTable())
	v1 := r.Group("/api")
	{
		// v1.POST("/userinfo", middleware.AuthmiddleWare(), controller.DTOUserInfo())
		// v1.GET("/getuser", api.GetUserInfo())
		// v1.POST("/setuser", api.SetUserInfo())
		v1.POST("/register", controller.UserRegister)
		v1.GET("/login", controller.UserLogin)
	}
	return r
}

package routers

import (
	"github.com/gin-gonic/gin"
	"self/app/http/controller"
)

func UserRegister(router *gin.RouterGroup)  {
	//router.GET("/user/login", controller.GetAdminUserList)
	router.POST("/user/login", controller.Login)
	router.GET("/user/info", controller.Info)
}
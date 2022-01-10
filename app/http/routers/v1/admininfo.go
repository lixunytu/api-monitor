package routers

import (
	"github.com/gin-gonic/gin"
	"self/app/http/controller"
)

func AdminRegister(router *gin.RouterGroup)  {
	router.POST("/admininfo", controller.CreateAdminInfo)
	router.GET("/admininfo", controller.JWTAuth(),controller.GetAdminInfoList)
	router.PUT("/admininfo/:id", controller.UpdateAAdminInfo)
	router.DELETE("/admininfo/:id", controller.DeleteAAdminInfo)
}
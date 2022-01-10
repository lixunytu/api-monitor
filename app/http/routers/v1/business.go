package routers

import (
	"github.com/gin-gonic/gin"
	"self/app/http/controller"
)

func BusinessRegister(router *gin.RouterGroup)  {
	router.POST("/getdata",controller.GetBodyData)
	router.POST("/datatest",controller.GetDataTest)
	router.POST("/httptest",controller.GetHttpTest)
	router.GET("/prasecron",controller.ParseCron)
}

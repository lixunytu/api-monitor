package routers

import (
	"github.com/gin-gonic/gin"
	"self/app/http/controller"
)

func DepartMentInfoRegister(router *gin.RouterGroup)  {
	router.POST("/departmentinfo", controller.CreateDepartMentInfo)
	router.GET("/departmentinfo", controller.GetDepartMentInfoList)
	router.PUT("/departmentinfo/:id", controller.UpdateADepartMentInfo)
	router.DELETE("/departmentinfo/:id", controller.DeleteADepartMentInfo)
}

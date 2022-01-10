package routers

import (
	"github.com/gin-gonic/gin"
	"self/app/http/controller"
)

func EmployeeRegister(router *gin.RouterGroup)  {
	router.POST("/employee", controller.CreateEmployee)
	router.DELETE("/employee/:id", controller.DelEmployee)
	router.PUT("/employee/:id", controller.UpdateEmployee)
	router.GET("/employee", controller.EmployeeList)
}

package routers

import (
	"github.com/gin-gonic/gin"
	"self/app/http/controller"
)

func MonitorInfoRegister(router *gin.RouterGroup)  {
	router.POST("/monitorinfo", controller.CreateMonitorInfo)
	router.PUT("/monitorinfo", controller.CreateAlarmInfo)
	router.GET("/monitorinfo", controller.MonitorList)
	router.GET("/monitorinfo/get", controller.MonitorInfoGet)
	router.POST("/monitorinfo/start", controller.MonitorInfoStart)
	router.POST("/monitorinfo/stop", controller.MonitorInfoStop)
	router.POST("/monitorinfo/del", controller.MonitorInfoDel)
	router.PUT("/monitorinfo/:monitorId", controller.UpdateMonitorInfo)
}

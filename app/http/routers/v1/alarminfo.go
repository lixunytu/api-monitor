package routers

import (
	"github.com/gin-gonic/gin"
	"self/app/http/controller"
)

func AlarmInfoRegister(router *gin.RouterGroup)  {
	router.GET("/alarmAInfo", controller.GetAlarmInfoList)
	router.GET("/alarmSampInfo", controller.GetAlarmSampInfo)
	router.POST("/alarmAInfo", controller.CreateAlarmInfo)
	router.GET("/alarmAInfo/:alarmId", controller.GetAlarmInfoById)
	router.DELETE("/alarmAInfo/:alarmId", controller.DeleteAAlarmInfo)
	router.PUT("/alarmAInfo/:alarmId", controller.UpdateAAlarmInfo)
}
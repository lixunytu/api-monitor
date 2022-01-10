package routers

import (
	"self/app/config"
	"self/app/http/routers/v1"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	if config.Conf.Release {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.Default()

	v1Group := r.Group("v1")
	routers.UserRegister(v1Group)
	routers.AdminRegister(v1Group)
	routers.BusinessRegister(v1Group)
	routers.DepartMentInfoRegister(v1Group)
	routers.AlarmInfoRegister(v1Group)
	routers.MonitorInfoRegister(v1Group)
	routers.EmployeeRegister(v1Group)
	return r
}

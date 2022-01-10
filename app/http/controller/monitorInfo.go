package controller

import (
	"net/http"
	"strconv"

	"self/app/models"

	"github.com/gin-gonic/gin"
)

func CreateMonitorInfo(c *gin.Context) {
	var m models.MonitorInfo
	err :=c.BindJSON(&m)
	if err!=nil {
		c.JSON(http.StatusOK, Fail(err.Error()))
		return
	}
	m.Status=1
	err=models.MonitorInfo{}.CreateA(&m)
	if err != nil{
		c.JSON(http.StatusOK, Fail(err.Error()))
	}else{
		c.JSON(http.StatusOK, Success(m))
	}
}

func UpdateMonitorInfo(c *gin.Context) {

	monitorId, ok := c.Params.Get("monitorId")
	if !ok {
		c.JSON(http.StatusOK, Fail("无效的id"))
		return
	}

	m,err := models.MonitorInfo{}.GetA([]string{monitorId})

	if err!=nil {
		c.JSON(http.StatusOK, Fail(err.Error()))
		return
	}

	err =c.BindJSON(m)
	if err!=nil {
		c.JSON(http.StatusOK, Fail(err.Error()))
		return
	}
	err=models.MonitorInfo{}.UpdateA(monitorId,m)
	if err != nil{
		c.JSON(http.StatusOK, Fail(err.Error()))
	}else{
		c.JSON(http.StatusOK, Success("更新成功"))
	}
}

func MonitorList(c *gin.Context) {
	pageNum,err := strconv.Atoi(c.DefaultQuery("pageNum","0"))
	if err!=nil {
		c.JSON(http.StatusOK,Fail(err.Error()))
		return
	}
	pageSize,err := strconv.Atoi(c.DefaultQuery("pageSize","0"))
	if err!=nil {
		c.JSON(http.StatusOK,Fail(err.Error()))
		return
	}

	all, err := models.MonitorInfo{}.GetAllForParams(pageNum,pageSize)
	if err!=nil {
		c.JSON(http.StatusOK,Fail(err.Error()))
		return
	}

	count, err := models.MonitorInfo{}.GetAllCount()
	if err!=nil {
		c.JSON(http.StatusOK,Fail(err.Error()))
		return
	}

	type resp struct {
		Count int `json:"count"`
		Data []*models.MonitorInfo `json:"data"`
	}

	var res = resp{
		Count: count,
		Data:  all,
	}

	c.JSON(http.StatusOK,Success(res))
}

func MonitorInfoGet(c *gin.Context) {
	monitorId := c.Query("monitorId")
	a, err := models.MonitorInfo{}.GetA([]string{monitorId})
	if err!=nil {
		c.JSON(http.StatusOK,Fail(err.Error()))
		return
	}
	c.JSON(http.StatusOK,Success(a))
}

func MonitorInfoStart(c *gin.Context) {

	m := make(map[string]interface{})
	c.BindJSON(&m)
	err := models.MonitorInfo{}.StartStatus(m["monitorId"])
	if err!=nil {
		c.JSON(http.StatusOK,Fail(err.Error()))
		return
	}
	c.JSON(http.StatusOK,Success("启用成功"))
}

func MonitorInfoStop(c *gin.Context) {
	m := make(map[string]interface{})
	c.BindJSON(&m)
	err := models.MonitorInfo{}.StopStatus(m["monitorId"])
	if err!=nil {
		c.JSON(http.StatusOK,Fail(err.Error()))
		return
	}
	c.JSON(http.StatusOK,Success("停止成功"))
}

func MonitorInfoDel(c *gin.Context) {
	m := make(map[string]interface{})
	c.BindJSON(&m)
	err := models.MonitorInfo{}.DeleteA(m["monitorId"])
	if err!=nil {
		c.JSON(http.StatusOK,Fail(err.Error()))
		return
	}
	c.JSON(http.StatusOK,Success("删除成功"))
}


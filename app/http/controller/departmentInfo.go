package controller

import (
	"net/http"

	"self/app/models"

	"github.com/gin-gonic/gin"
)

func CreateDepartMentInfo(c *gin.Context) {
	// 前端页面填写待办事项 点击提交 会发请求到这里
	// 1. 从请求中把数据拿出来
	var d models.DepartMentInfo
	c.BindJSON(&d)
	// 2. 存入数据库
	err:=models.DepartMentInfo{}.CreateA(&d)
	if err != nil{
		c.JSON(http.StatusOK, Fail(err.Error()))
	}else{
		c.JSON(http.StatusOK, Success(d))
		//c.JSON(http.StatusOK, gin.H{
		//	"code": 2000,
		//	"msg": "success",
		//	"data": todo,
		//})
	}
}

func GetDepartMentInfoList(c *gin.Context) {
	// 查询todo这个表里的所有数据
	DepartMentInfoList, err := models.DepartMentInfo{}.GetAll()
	if err!= nil {
		c.JSON(http.StatusOK, Fail(err.Error()))
	}else {
		c.JSON(http.StatusOK, Success(DepartMentInfoList))
	}
}

func UpdateADepartMentInfo(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, Fail("无效的id"))
		return
	}
	d, err := models.DepartMentInfo{}.GetA(id)
	if err != nil {
		c.JSON(http.StatusOK, Fail(err.Error()))
		return
	}
	c.BindJSON(&d)

	err = models.DepartMentInfo{}.UpdateA(d);
	if  err!= nil{
		c.JSON(http.StatusOK, Fail(err.Error()))
	}else{
		c.JSON(http.StatusOK, Success(d))
	}
}

func DeleteADepartMentInfo(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, Fail("无效的id"))
		return
	}

	err := models.DepartMentInfo{}.DeleteA(id);
	if err!=nil{
		c.JSON(http.StatusOK, Fail(err.Error()))
	}else{
		c.JSON(http.StatusOK, Success(id+":deleted"))
	}
}

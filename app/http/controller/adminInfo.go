package controller

import (
	"net/http"

	"self/app/models"

	"github.com/gin-gonic/gin"

)

//AdminInfo
func CreateAdminInfo(c *gin.Context) {
	// 前端页面填写待办事项 点击提交 会发请求到这里
	// 1. 从请求中把数据拿出来
	var admin models.AdminInfo
	c.BindJSON(&admin)
	// 2. 存入数据库
	err := models.AdminInfo{}.CreateA(&admin)
	if err != nil {
		c.JSON(http.StatusOK, Fail(err.Error()))
	} else {
		c.JSON(http.StatusOK, Success(admin))
		//c.JSON(http.StatusOK, gin.H{
		//	"code": 2000,
		//	"msg": "success",
		//	"data": todo,
		//})
	}
}

func GetAdminInfoList(c *gin.Context) {
	// 查询todo这个表里的所有数据
	AdminInfoList, err := models.AdminInfo{}.GetAll()
	if err != nil {
		c.JSON(http.StatusOK, Fail(err.Error()))
	} else {
		c.JSON(http.StatusOK, Success(AdminInfoList))
	}
}

func UpdateAAdminInfo(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, Fail("无效的id"))
		return
	}
	admin, err := models.AdminInfo{}.GetA(id)
	if err != nil {
		c.JSON(http.StatusOK, Fail(err.Error()))
		return
	}
	c.BindJSON(&admin)

	err = models.AdminInfo{}.UpdateA(admin)

	if err != nil {
		c.JSON(http.StatusOK, Fail(err.Error()))
	} else {
		c.JSON(http.StatusOK, Success(admin))
	}
}

func DeleteAAdminInfo(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, Fail("无效的id"))
		return
	}
	err := models.AdminInfo{}.DeleteA(id)
	if err != nil {
		c.JSON(http.StatusOK, Fail(err.Error()))
	} else {
		c.JSON(http.StatusOK, Success(id+"deleted"))
	}
}

package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"self/app/models"
	"strconv"
)

func CreateEmployee(c *gin.Context) {
	var e models.EmployeeInfo
	c.BindJSON(&e)
	err := models.EmployeeInfo{}.CreateA(&e)
	if err != nil {
		c.JSON(http.StatusOK, Fail(err.Error()))
	} else {
		c.JSON(http.StatusOK, Success(e))
	}
}

func EmployeeList(c *gin.Context) {
	keyWords := c.DefaultQuery("employeeKeywords","")
	if keyWords!="" {
		words, err := models.EmployeeInfo{}.GetSumByKeyWords(keyWords)
		if err!=nil {
			c.JSON(http.StatusOK,Fail(err.Error()))
			return
		}
		c.JSON(http.StatusOK,Success(words))
		return
	}

	currentPage, err := strconv.Atoi(c.DefaultQuery("currentPage", "0"))
	if err != nil {
		c.JSON(http.StatusOK, Fail(err.Error()))
		return
	}
	pageSize, err := strconv.Atoi(c.DefaultQuery("pageSize", "0"))
	if err != nil {
		c.JSON(http.StatusOK, Fail(err.Error()))
		return
	}
	all, err := models.EmployeeInfo{}.GetAll(currentPage, pageSize)
	if err != nil {
		c.JSON(http.StatusOK, Fail(err.Error()))
		return
	}

	type resStuc struct {
		Count int                    `json:"count"`
		Data  []*models.EmployeeInfo `json:"data"`
	}

	count, err := models.EmployeeInfo{}.GetCount()
	if err != nil {
		c.JSON(http.StatusOK, Fail(err.Error()))
		return
	}

	var data = resStuc{
		Count: count,
		Data:  all,
	}

	c.JSON(http.StatusOK, Success(data))

}

func UpdateEmployee(c *gin.Context) {
	eId, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, Fail("无效的id"))
		return
	}
	e, err := models.EmployeeInfo{}.GetA(eId)
	if err != nil {
		c.JSON(http.StatusOK, Fail(err.Error()))
		return
	}
	err = c.BindJSON(&e)
	if err != nil {
		c.JSON(http.StatusOK, Fail(err.Error()))
		return
	}
	err = models.EmployeeInfo{}.UpdateA(e)
	if err != nil {
		c.JSON(http.StatusOK, Fail(err.Error()))
		return
	}
	c.JSON(http.StatusOK, Success("修改成功"))
}

func DelEmployee(c *gin.Context) {
	eId, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, Fail("无效的id"))
		return
	}
	err := models.EmployeeInfo{}.DeleteA(eId)
	if err != nil {
		c.JSON(http.StatusOK, Fail(err.Error()))
	} else {
		c.JSON(http.StatusOK, Success("删除成功"))
	}
}

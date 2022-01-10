package controller

import (
	"net/http"
	"self/app/models"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

//AlarmInfo
func CreateAlarmInfo(c *gin.Context) {


	var alarm models.AlarmInfo
	err := c.BindJSON(&alarm)
	if err != nil {
		c.JSON(http.StatusOK, Fail("参数错误："+err.Error()))
		return
	}
	err = models.AlarmInfo{}.CreateA(&alarm)
	if err != nil {
		c.JSON(http.StatusOK, Fail(err.Error()))
	} else {
		c.JSON(http.StatusOK, Success(alarm))
	}
}

func GetAlarmInfoList(c *gin.Context) {

	pageNum, err := strconv.Atoi(c.DefaultQuery("pageNum", "0"))
	if err != nil {
		c.JSON(http.StatusOK, Fail(err.Error()))
		return
	}
	pageSize, err := strconv.Atoi(c.DefaultQuery("pageSize", "0"))
	if err != nil {
		c.JSON(http.StatusOK, Fail(err.Error()))
		return
	}

	all, err := models.AlarmInfo{}.GetAll(pageNum,pageSize)
	if err != nil {
		c.JSON(http.StatusOK, Fail(err.Error()))
	} else {
		count, err := models.AlarmInfo{}.GetAllCount()
		if err != nil {
			c.JSON(http.StatusOK, Fail(err.Error()))
			return
		}
		type alarmInfoforUser struct {
			AlarmType []int `json:"alarm_type"`
			models.AlarmInfo
			ReceptionUsers map[string][]*models.EmployeeInfo `json:"receptionUsers"`
		}
		type resStr struct {
			Count int                 `json:"count"`
			Data  []*alarmInfoforUser `json:"data"`
		}

		var alarmInfo = make([]*alarmInfoforUser, 0)

		for _, v := range all {
			e, err := splitString(v.AlarmType, map[string]string{"1": v.RecipientMail, "2": v.DingGroup, "3": v.RecipientMessage, "4": v.RecipientPhone})
			if err != nil {
				c.JSON(http.StatusOK, Fail(err.Error()))
				return
			}

			at := make([]int, 0)

			if v.AlarmType != "" {
				for _, v := range strings.Split(v.AlarmType[1:len(v.AlarmType)-1], ",") {
					v, e := strconv.Atoi(v)
					if e != nil {
						continue
					}
					at = append(at, v)
				}
			}

			a := &alarmInfoforUser{
				AlarmType:      at,
				AlarmInfo:      *v,
				ReceptionUsers: e,
			}
			alarmInfo = append(alarmInfo, a)
		}

		res := resStr{
			Count: count,
			Data:  alarmInfo,
		}

		c.JSON(http.StatusOK, Success(res))
	}
}

func GetAlarmInfoById(c *gin.Context)  {
	alarmId,ok := c.Params.Get("alarmId")
	if !ok {
		c.JSON(http.StatusOK,Fail("alarmId is not exist"))
		return
	}
	v, err := models.AlarmInfo{}.GetA(alarmId)
	if err!=nil {
		c.JSON(http.StatusOK,Fail(err.Error()))
		return
	}

	type alarmInfoforUser struct {
		AlarmType []int `json:"alarm_type"`
		models.AlarmInfo
		ReceptionUsers map[string][]*models.EmployeeInfo `json:"receptionUsers"`
	}

	e,err := splitString(v.AlarmType, map[string]string{"1": v.RecipientMail, "2": v.DingGroup, "3": v.RecipientMessage, "4": v.RecipientPhone})

	if err != nil {
		c.JSON(http.StatusOK, Fail(err.Error()))
		return
	}
	at := make([]int, 0)

	if v.AlarmType != "" {
		for _, v := range strings.Split(v.AlarmType[1:len(v.AlarmType)-1], ",") {
			v, e := strconv.Atoi(v)
			if e != nil {
				continue
			}
			at = append(at, v)
		}
	}

	a := &alarmInfoforUser{
		AlarmType:      at,
		AlarmInfo:      *v,
		ReceptionUsers: e,
	}


	c.JSON(http.StatusOK,Success(a))
}

func GetAlarmSampInfo(c *gin.Context) {
	all, err := models.AlarmInfo{}.GetAllWithIdName()
	if err != nil {
		c.JSON(http.StatusOK, Fail(err.Error()))
	} else {
		c.JSON(http.StatusOK, Success(all))
	}
}

func UpdateAAlarmInfo(c *gin.Context) {
	id, ok := c.Params.Get("alarmId")
	if !ok {
		c.JSON(http.StatusOK, Fail("无效的id"))
		return
	}
	alarm, err := models.AlarmInfo{}.GetA(id)
	if err != nil {
		c.JSON(http.StatusOK, Fail(err.Error()))
		return
	}

	err = c.BindJSON(&alarm)
	if err!=nil {
		c.JSON(http.StatusOK,Fail(err.Error()))
		return
	}
	err = models.AlarmInfo{}.UpdateA(id, alarm)
	if err != nil {
		c.JSON(http.StatusOK, Fail(err.Error()))
	} else {
		c.JSON(http.StatusOK, Success(alarm))
	}
}

func DeleteAAlarmInfo(c *gin.Context) {
	id, ok := c.Params.Get("alarmId")
	if !ok {
		c.JSON(http.StatusOK, Fail("无效的id"))
		return
	}

	err := models.AlarmInfo{}.DeleteA(id)
	if err != nil {
		c.JSON(http.StatusOK, Fail(err.Error()))
	} else {
		c.JSON(http.StatusOK, Success(id+":deleted"))
	}
}

func splitString(at string, s map[string]string) (r map[string][]*models.EmployeeInfo, err error) {
	r = make(map[string][]*models.EmployeeInfo, 4)
	if at != "" {
		at = at[1 : len(at)-1]
	}

	for _, v := range strings.Split(at, ",") {
		if s[v] != "" {
			s[v] = s[v][1 : len(s[v])-1]
			sum, err := models.EmployeeInfo{}.GetSum(strings.Split(s[v], ","))
			if err != nil {
				return nil, err
			}
			r[v] = sum
		}
	}
	return
}

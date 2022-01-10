package controller

import (
	"io/ioutil"
	"net/http"
	"self/app/business/execute"
	"strconv"

	"self/app/common/httpcommon"
	"self/app/common/util"

	"github.com/gin-gonic/gin"
)

const (
	URL           = "url"
	REQUESTHEADER = "request_header"
	//SERVERURL     = "server_url"
	POSTDATA     = "post_data"
	GETDATA      = "get_data"
	RT           = "RT"
	BODY         = "body"
	REMOTEADDR   = "remote_addr"
	RULES        = "rules"
	ISTRUE       = "isTrue"
	ERRORSTRING  = "errorString"
	ERRORRULES   = "errorRules"
	EXPECTEDCODE = "expectedCode"
)

func GetBodyData(c *gin.Context) {
	buf, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusOK, Fail(err.Error()))
		return
	}

	bdm, err := util.JsonToMap(string(buf))
	if err != nil {
		c.JSON(http.StatusOK, Fail(err.Error()))
		return
	}

	_, rt, _, rs, ra, err := httpcommon.HttpCLientCommen(bdm[URL], 2, bdm[REQUESTHEADER], bdm[POSTDATA], bdm[GETDATA], httpcommon.NEEDBODY)
	if err != nil {
		c.JSON(http.StatusOK, Fail(err.Error()))
		return
	}

	rm := make(map[string]string)
	rm[RT] = strconv.Itoa(int(rt))
	rm[BODY] = rs
	rm[REMOTEADDR] = ra

	//json, err := util.MapToJson(rm)
	//if err != nil {
	//	c.JSON(http.StatusOK, Fail(err.Error()))
	//	return
	//}

	c.JSON(http.StatusOK, Success(rm))
}

func GetDataTest(c *gin.Context) {
	buf, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusOK, Fail(err.Error()))
		return
	}

	bdm, err := util.JsonToMap(string(buf))
	if err != nil {
		c.JSON(http.StatusOK, Fail(err.Error()))
		return
	}
	_, rt, _, rs, ra, err := httpcommon.HttpCLientCommen(bdm[URL], 2, bdm[REQUESTHEADER], bdm[POSTDATA], bdm[GETDATA], httpcommon.NEEDBODY)
	if err != nil {

		c.JSON(http.StatusOK, Fail(err.Error()))
		return
	}

	isHealth, errs, errru, err := execute.DataControl(bdm[RULES], rs)
	if err != nil {

		c.JSON(http.StatusOK, Fail(err.Error()))
		return
	}

	rm := make(map[string]string)
	rm[RT] = strconv.Itoa(int(rt))
	//rm[BODY] = rs
	rm[REMOTEADDR] = ra
	rm[ISTRUE] = strconv.FormatBool(isHealth)
	rm[ERRORSTRING] = errs
	rm[ERRORRULES] = errru

	c.JSON(http.StatusOK, Success(rm))
}

func GetHttpTest(c *gin.Context) {
	buf, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusOK, Success(err.Error()))
		return
	}

	bdm, err := util.JsonToMap(string(buf))
	if err != nil {
		c.JSON(http.StatusOK, Success(err.Error()))
		return
	}

	r, _, _, _, _, err := httpcommon.HttpCLientCommen(bdm[URL], 2, bdm[REQUESTHEADER], bdm[POSTDATA], bdm[GETDATA], httpcommon.NOTNEEDBODY)
	defer r.Body.Close()
	if err != nil {
		c.JSON(http.StatusOK, Success(err.Error()))
		return
	}
	ecm := execute.ExposeCodeTomap(bdm[EXPECTEDCODE])

	_,ok := ecm[strconv.Itoa(r.StatusCode)]

	if ok {
		c.JSON(http.StatusOK, Success("校验成功"))
	}else {
		c.JSON(http.StatusOK,Success("实际返回码："+strconv.Itoa(r.StatusCode)))
	}


}

func ParseCron(c *gin.Context)  {
	e := c.Query("e")
	if e==""{
		c.JSON(http.StatusOK,Fail("请输入cron表达式"))
		return
	}

	cron ,err:= util.ParseCron(e)
	if err!=nil {
		c.JSON(http.StatusOK,Fail("解析失败"))
		return
	}
	c.JSON(http.StatusOK,Success(cron))
}

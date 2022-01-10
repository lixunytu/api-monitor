package execute

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"self/app/business/control"
	"self/app/common/httpcommon"
	"self/app/common/util"
)

func httpexecute(c control.Control,isHttp bool) (err error) {
	var (
		outmap   = make(map[string]string, 3)
		response = &http.Response{}
		rt, st   int64
		retry    = c.Monitor.MRetry
		wt       = c.Monitor.MWaitTime
		isHealth = false
		ra       string
		ecm      = ExposeCodeTomap(c.Monitor.ExpectedCode)
		code     int
	)

	if retry == 0 && !isHttp {
		retry = 1
	}

	for count := 0; count < retry; count++ {
		response, rt, st, _, ra, err = httpcommon.HttpCLientCommen(c.Monitor.ServerUrl, c.Monitor.MTimeout, c.Monitor.RequestHeaders,
			c.Monitor.PostData, c.Monitor.GetData, false)
		if err == nil {
			defer response.Body.Close()
			code = response.StatusCode
			_, isHealth = ecm[strconv.Itoa(response.StatusCode)]
		} else {
			fmt.Println(err)
			//TODO：记录日志
		}
		if isHealth {
			break
		}
		time.Sleep(time.Duration(wt) * time.Second)
	}

	if err != nil {
		outmap["Reson"] = err.Error()
		outmap["Remote_Addr"] = ra
	}

	outputstring, err := util.MapToJson(outmap)
	if err != nil {
		return err
	}

	e := ExecuteInfo{
		Id:           c.Monitor.Id,
		AId:          c.Monitor.AlarmId,
		Type:         c.Monitor.MType,
		IsHealthy:    isHealth,
		Status_Code:  code,
		Expect_Code:  c.Monitor.ExpectedCode,
		ResponseTime: rt,
		RunTime:      st,
		Output:       outputstring,
	}

	Producer(e)

	return nil
}

func ExposeCodeTomap(s string) (mp map[string]bool) {
	mp = make(map[string]bool, 5)
	s = s[1 : len(s)-1]
	for _, i := range strings.Split(s, ",") {
		mp[i] = true
	}
	return
}

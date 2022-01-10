package execute

import (
	"encoding/json"
	"fmt"
	"net/http"
	"self/app/common/util"
	"strconv"
	"time"

	"self/app/business/control"
	"self/app/common/httpcommon"

	"github.com/tidwall/gjson"
)

const (
	COUNT    = "count"
	EQUAL    = "equal"
	KEYEXIST = "key-exist"
	NOTNULL  = "not-null"
	NOTEQUAL = "not-equal"
	INCLUDE  = "include"
	GT       = "gt"
	LT       = "lt"
)

func dataexecute(c control.Control, isHttp bool) (err error) {

	var (
		outmap              = make(map[string]string, 3)
		response            = &http.Response{}
		rt, st              int64
		retry               = c.Monitor.MRetry
		wt                  = c.Monitor.MWaitTime
		isHealth            = false
		code                int
		ra, bd, errs, errru string
	)

	if retry == 0 && !isHttp {
		retry = 1
	}

	for count := 0; count < retry; count++ {
		response, rt, st, bd, ra, err = httpcommon.HttpCLientCommen(c.Monitor.ServerUrl, c.Monitor.MTimeout, c.Monitor.RequestHeaders,
			c.Monitor.PostData, c.Monitor.GetData, httpcommon.NEEDBODY)

		if err == nil {
			//判断
			isHealth, errs, errru, err = DataControl(c.Monitor.MRules, bd)
			if isHealth {
				code = response.StatusCode
				break
			}

		} else {
			//TODO：判断
		}
		time.Sleep(time.Duration(wt) * time.Second)
	}

	if err != nil {
		outmap["Reson"] = err.Error()
		outmap["RemoteAddr"] = ra
		outString, _ := util.MapToJson(outmap)

		e := ExecuteInfo{
			Id:           c.Monitor.Id,
			AId:          c.Monitor.AlarmId,
			Type:         DATA_MONITOR,
			IsHealthy:    isHealth,
			Expect_Code:  "",
			Status_Code:  code,
			ResponseTime: rt,
			RunTime:      st,
			Output:       outString,
		}

		ALARM <- e

		return err
	}

	response_map := make(map[string]string, len(response.Header))
	for k, v := range response.Header {
		response_map[k] = SliceToString(v)
	}
	respString, _ := util.MapToJson(response_map)

	outmap["ErrRules"] = errru
	outmap["ErrString"] = errs
	outmap["Body"] = bd
	outmap["ErrString"] = errs

	outmap["ResponseHeader"] = respString
	outmap["Reason"] = "true"
	outString, _ := util.MapToJson(outmap)

	e := ExecuteInfo{
		Id:           c.Monitor.Id,
		AId:          c.Monitor.AlarmId,
		Type:         DATA_MONITOR,
		IsHealthy:    isHealth,
		Expect_Code:  "",
		Status_Code:  code,
		ResponseTime: rt,
		RunTime:      st,
		Output:       outString,
	}

	ALARM <- e

	defer response.Body.Close()

	return nil
}

func DataControl(r, d string) (isHealth bool, errs, errru string, err error) {
	isHealth = true
	m := rulesToMap(r)

	for i := 0; i < len(m); i++ {
		switch m[i]["rule"] {
		case COUNT:
			isHealth := KeyExist(m[i]["arg1"], d)
			if !isHealth {
				return isHealth, m[i]["arg1"] + " key not exist ", m[i]["rule"] + "," + m[i]["arg1"] + "," + m[i]["arg2"], nil
			}
			isHealth, err, errs = DataCount(m[i]["arg1"], m[i]["arg2"], d)
			if err != nil {
				fmt.Println(err)
			}
			if !isHealth {
				return isHealth, errs, m[i]["rule"] + "," + m[i]["arg1"] + "," + m[i]["arg2"], err
			}
		case EQUAL:
			isHealth := KeyExist(m[i]["arg1"], d)
			if !isHealth {
				return isHealth, m[i]["arg1"] + " key not exist ", m[i]["rule"] + "," + m[i]["arg1"] + "," + m[i]["arg2"], nil
			}
			isHealth, err, errstring := Equal(m[i]["arg1"], m[i]["arg2"], d)
			if err != nil {
				fmt.Println(err)
			}
			if !isHealth {
				return isHealth, errstring, m[i]["rule"] + "," + m[i]["arg1"] + "," + m[i]["arg2"], err
			}
		case KEYEXIST:
			isHealth := KeyExist(m[i]["arg1"], d)
			if !isHealth {
				return isHealth, m[i]["arg1"] + " key not exist ", m[i]["rule"] + "," + m[i]["arg1"] + "," + m[i]["arg2"], nil
			}
		case NOTNULL:
			isHealth = NotNull(m[i]["arg1"], m[i]["arg2"], d)
			if !isHealth {
				return isHealth, m[i]["arg1"] + " key is null or key is not exist", m[i]["rule"] + "," + m[i]["arg1"] + "," + m[i]["arg2"], nil
			}
		case NOTEQUAL:
			isHealth := KeyExist(m[i]["arg1"], d)
			if !isHealth {
				return isHealth, m[i]["arg1"] + " key not exist ", m[i]["rule"] + "," + m[i]["arg1"] + "," + m[i]["arg2"], nil
			}
			isHealth, err, errstring := NotEqual(m[i]["arg1"], m[i]["arg2"], d)
			if err != nil {
				fmt.Println(err)
			}
			if !isHealth {
				return isHealth, errstring, m[i]["rule"] + "," + m[i]["arg1"] + "," + m[i]["arg2"], err
			}
		case INCLUDE:
			isHealth := KeyExist(m[i]["arg1"], d)
			if !isHealth {
				return isHealth, m[i]["arg1"] + " key not exist ", m[i]["rule"] + "," + m[i]["arg1"] + "," + m[i]["arg2"], nil
			}
			isHealth, err, errstring := Include(m[i]["arg1"], m[i]["arg2"], d)
			if err != nil {
				fmt.Println(err)
			}
			if !isHealth {
				return isHealth, errstring, m[i]["rule"] + "," + m[i]["arg1"] + "," + m[i]["arg2"], err
			}
		case GT:
			isHealth := KeyExist(m[i]["arg1"], d)
			if !isHealth {
				return isHealth, m[i]["arg1"] + " key not exist ", m[i]["rule"] + "," + m[i]["arg1"] + "," + m[i]["arg2"], nil
			}
			isHealth, err, errstring := GTF(m[i]["arg1"], m[i]["arg2"], d)
			if err != nil {
				fmt.Println(err)
			}
			if !isHealth {
				return isHealth, errstring, m[i]["rule"] + "," + m[i]["arg1"] + "," + m[i]["arg2"], err
			}
		case LT:
			isHealth := KeyExist(m[i]["arg1"], d)
			if !isHealth {
				return isHealth, m[i]["arg1"] + " key not exist ", m[i]["rule"] + "," + m[i]["arg1"] + "," + m[i]["arg2"], nil
			}
			isHealth, err, errstring := LTF(m[i]["arg1"], m[i]["arg2"], d)
			if err != nil {
				fmt.Println(err)
			}
			if !isHealth {
				return isHealth, errstring, m[i]["rule"] + "," + m[i]["arg1"] + "," + m[i]["arg2"], err
			}
		default:
		}
	}

	return isHealth, "", "null",err
}

func rulesToMap(rules string) (s []map[string]string) {
	length := gjson.Get(rules, "#")
	for i := 0; i < int(length.Int()); i++ {
		m := make(map[string]string)
		data := gjson.Get(rules, strconv.Itoa(i))
		err := json.Unmarshal([]byte(data.String()), &m)
		if err != nil {
			fmt.Println(err)
		}
		s = append(s, m)
	}
	return
}

func SliceToString(a []string) string {
	str := ""
	for _, v := range a {
		str = str + v + " "
	}
	return str
}

package alarm

import (
	"fmt"
	"github.com/pkg/errors"
	"self/app/business/execute"
	"self/app/models"
	"strconv"
	"strings"
)

const (
	DINGDING = "1"
	MAIL     = "2"
	MESSAGE  = "3"
	PHONE    = "4"
)

var (
	alarmList = make(map[int]int)
)

func Start() {
	for e := range execute.ALARM {

		fmt.Println(e)
		err := alarmControl(e)
		if err!=nil {
			//TODO:记录日志
			fmt.Println(err)
		}
	}
}

func alarmControl(e execute.ExecuteInfo) (err error) {
	a, err := models.AlarmInfo{}.GetA(strconv.Itoa(e.AId))
	if err != nil {
		//TODO：记录日志
		return errors.Wrap(err, "alarmControl: select alarmInfo error")
	}

	atpm,member, err := delalarm(*a)

	if err != nil {
		//TODO：记录日志
		return errors.Wrap(err, "alarmControl: delalarm error")
	}

	if !e.IsHealthy {
		alarmList[e.Id]++
		for k, v := range atpm {
			switch k {
			//TODO：这里会重复报警，后续需要处理
			case DINGDING:
				user, err := selectUser(member[k], DINGDING)
				if err!=nil {
					return errors.Wrap(err,"alarmControl: select ding User error")
				}
				if alarmList[e.Id] >= v {
					err := SendDingMsg("test报警", user)
					if err!= nil {
						//TODO:日志
					}
				}
			case MAIL:
				user, err := selectUser(member[k], MAIL)
				if err!=nil {
					return errors.Wrap(err,"alarmControl: select ding User error")
				}
				if alarmList[e.Id] >= v {
					SendDingMsg("", user)
				}
			case MESSAGE:
				user, err := selectUser(member[k], MESSAGE)
				if err!=nil {
					return errors.Wrap(err,"alarmControl: select ding User error")
				}
				if alarmList[e.Id] >= v {
					SendDingMsg("", user)
				}
			case PHONE:
				user, err := selectUser(member[k], PHONE)
				if err!=nil {
					return errors.Wrap(err,"alarmControl: select ding User error")
				}
				if alarmList[e.Id] >= v {
					SendDingMsg("", user)
				}

			}
		}
	} else {
		alarmList[e.Id] = 0
	}

	return
}

func delalarm(a models.AlarmInfo) (mp map[string]int,mem map[string]string ,err error) {
	mp = make(map[string]int,4)
	mem = make(map[string]string,4)
	if len(a.AlarmType) <= 2 {
		return nil, nil ,fmt.Errorf("alarm type is error")
	}

	a.AlarmType = a.AlarmType[1 : len(a.AlarmType)-1]
	for _, i := range strings.Split(a.AlarmType, ",") {
		switch i {
		case DINGDING:
			mp[i] = a.DingGroupFailNum
			mem[i] = a.DingGroup[1:len(a.DingGroup)-1]
		case MAIL:
			mp[i] = a.DingGroupFailNum
			mem[i] = a.RecipientMail[1:len(a.RecipientMail)-1]
		case MESSAGE:
			mp[i] = a.MessageFailNum
			mem[i] = a.RecipientMessage[1:len(a.RecipientMessage)-1]
		case PHONE:
			mp[i] = a.PhoneFailNum
			mem[i] = a.RecipientPhone[1:len(a.RecipientPhone)-1]
		}
	}

	return

}

func selectUser(s string,atype string) (r []string,err error) {
	split := strings.Split(s, ",")
	sum, err := models.EmployeeInfo{}.GetSum(split)
	if err != nil {
		return nil,errors.Wrap(err,"selectUser: error")
	}
	if atype == DINGDING {
		for _,v := range sum {
			r = append(r,v.DingUrl)
		}
	} else if atype == MAIL {
		for _,v := range sum {
			r = append(r,v.Email)
		}
	}else if atype == MESSAGE || atype == PHONE {
		for _,v := range sum {
			r = append(r,v.PhoneNumber)
		}
	}
	return
}

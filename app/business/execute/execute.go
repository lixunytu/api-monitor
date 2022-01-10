package execute

import (
	"fmt"

	"self/app/business/control"
)

const (
	HTTP_MONITOR = 1
	DATA_MONITOR = 2
	PORT_MONITOR = 3
	PID_MONITOR = 4
	KAFKA_MONITOR = 5
	HTTPREQ = true
	CONTRALREQ = false

)

func Start()  {

	for c := range control.Mq {
		fmt.Println(c)
		err := Distribute(c)
		if err!=nil {
			//TODO:记录日志
			fmt.Println(err)
		}
	}

}

func Distribute(c control.Control) (err error)  {

	switch c.Monitor.MType {
	case HTTP_MONITOR:
		go httpexecute(c,CONTRALREQ)
	case DATA_MONITOR:
		go dataexecute(c,CONTRALREQ)
	}

	return err
}

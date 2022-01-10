package execute

var ALARM chan ExecuteInfo = make(chan ExecuteInfo,200)

func MqStart()  {
	//TODO
}

func Producer(i ExecuteInfo) {
	ALARM <- i
}

func Consumer() ExecuteInfo {
	return <-ALARM
}
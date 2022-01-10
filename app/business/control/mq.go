package control

var Mq chan Control = make(chan Control,200)

func MqStart()  {
	//TODO
}

func Producer(i Control) {
	Mq <- i
}

func Consumer() Control {
	return <-Mq
}

package main

import (
	"fmt"
	"os"
	"self/app/business/alarm"
	"self/app/business/control"
	"self/app/business/execute"
	"self/app/config"
	"self/app/databases"
	"self/app/databases/mysql"
	"self/app/http/routers"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage： conf/config.ini")
		return
	}
	// 加载配置文件
	if err := config.Init(os.Args[1]); err != nil {
		fmt.Printf("load config from file failed, err:%v\n", err)
		return
	}

	if err := databases.MysqlStart(); err != nil {
		fmt.Printf("init mysql, err:%v\n", err)
		return
	}
	defer mysql.Close()

	//control business
	if err := control.InitControl(); err != nil {
		fmt.Printf("init control, err:%v\n", err)
		return
	}
	//exector business
	go execute.Start()
	go alarm.Start()

	// 注册路由
	r := routers.SetupRouter()
	if err := r.Run(fmt.Sprintf(":%d", config.Conf.Port)); err != nil {
		fmt.Printf("server startup failed, err:%v\n", err)
	}
}

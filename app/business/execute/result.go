package execute

type ExecuteInfo struct {
	// 监控 id
	Id int
	//报警id
	AId int
	//监控类型
	Type int
	//是否成功
	IsHealthy bool

	Expect_Code string
	//状态
	Status_Code int
	//返回时间
	ResponseTime int64
	//运行时间
	RunTime int64
	//输出的内容
	Output string
}

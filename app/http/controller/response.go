package controller

type HttpResponse struct {
	Status int         `json:"status"`
	Mes    string      `json:"mes,omitempty"`
	Data   interface{} `json:"data,omitempty"`
}

func Success(data ...interface{}) *HttpResponse {
	resp := &HttpResponse{Status: 0, Mes: "success"}
	if len(data) > 0 {
		if data[0] == nil {
			resp.Data = make([]string, 0)
		} else {
			resp.Data = data[0]
		}
	} else {
		resp.Data = make([]string, 0)
	}
	return resp
}

func Fail(msg string) *HttpResponse {
	resp := &HttpResponse{Status: 1, Mes: msg}
	return resp
}

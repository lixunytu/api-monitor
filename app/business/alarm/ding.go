package alarm

import (
	"net/http"
	"strings"
)

func SendDingMsg(msg string, dg []string) error {

	for _,dingId := range dg {
		webHook := "https://oapi.dingtalk.com/robot/send?access_token=" + dingId
		content := `{"msgtype": "text",
		"text": {"content": "`+ msg + `"}
	}`
		req, err := http.NewRequest("POST", webHook, strings.NewReader(content))
		if err != nil {
			return err
		}

		client := &http.Client{
		}
		req.Header.Set("Content-Type", "application/json; charset=utf-8")
		resp, err := client.Do(req)

		if err != nil {
			return err
		}
		defer resp.Body.Close()
	}

	return nil
}

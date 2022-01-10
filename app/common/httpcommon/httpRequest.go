package httpcommon

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"strconv"
	"time"

	"github.com/pkg/errors"
	"self/app/common/util"

)

const (
	NEEDBODY = true
	NOTNEEDBODY = false
)

func HttpCLientCommen(urls string, t int, hs string, pd string,
	gd string, iresp bool) (r *http.Response, rt, endt int64, rs string, ra string, err error) {

	if gd != "" {
		gdm, err := util.JsonTointerfaceMap(gd)
		if err != nil {
			return nil, 0, time.Now().UnixNano() / 1e6, "", "", errors.Wrap(err, "HttpCLientCommen: util.JsonTointerfaceMap error")
		}
		urls, err = util.Urlanalysis(urls, gdm)
		if err != nil {
			return nil, 0, time.Now().UnixNano() / 1e6, "", "", errors.Wrap(err, "HttpCLientCommen: util.Urlanalysis error")
		}
	}

	request := &http.Request{}

	if len(pd) == 0 {
		request, err = http.NewRequest("GET", urls, nil)
		if err != nil {
			return nil, 0, time.Now().UnixNano() / 1e6, "", "", errors.Wrap(err, "HttpCLientCommen: GET request error")
		}
	} else {
		pd, err = util.PostAnalysis(pd)
		if err != nil {
			return nil, 0, time.Now().UnixNano() / 1e6, "", "", errors.Wrap(err, "HttpCLientCommen: util.PostAnalysis error")
		}

		request, err = http.NewRequest("POST", urls, bytes.NewReader([]byte(pd)))
		request.Header.Add("Content-Length", strconv.Itoa(int(request.ContentLength)))
		if err != nil {
			return nil, 0, time.Now().UnixNano() / 1e6, "", "", errors.Wrap(err, "HttpCLientCommen: POST request error")
		}
	}

	client := &http.Client{
		CheckRedirect: RedirectFunc,
		Timeout:       time.Duration(time.Duration(t) * time.Second),
		Transport: &http.Transport{
			DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
				conn, err := net.Dial(network, addr)
				if err != nil {
					request.RemoteAddr = ""
				} else {
					request.RemoteAddr = conn.RemoteAddr().String()
				}
				return conn, err
			},
		},
	}
	//get 请求
	//设置 header
	request.Header.Add("Cache-Control", "no-cache")
	//request.Header.Add("Host", "play.google.com")
	request.Header.Add("Accept", "*/*")
	//request.Header.Add("Content-Type", "gzip, deflate, br")
	request.Header.Add("Accept-Encoding", "application/x-www-form-urlencoded")
	request.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/535.1 (KHTML, like Gecko) Chrome/14.0.835.163 Safari/535.1")
	//设置自定义请求
	var hm = make(map[string]string)
	if len(hs) != 0 {
		hm, err = util.JsonToMap(hs)
		if err != nil {
			return nil, 0, time.Now().UnixNano() / 1e6, "", "", errors.Wrap(err, "HttpCLientCommen: Json to map error")
		}

		for k, v := range hm {
			request.Header.Del(k)
			request.Header.Add(k, v)
		}
	}

	var st, et int64
	//开始时间,执行时间,结束时间
	r = new(http.Response)
	st = time.Now().UnixNano() / 1e6
	r, err = client.Do(request)
	if err != nil {
		return nil, et - st, st, "", request.RemoteAddr, errors.Wrap(err, "HttpCLientCommen: client.do error")
	}
	et = time.Now().UnixNano() / 1e6

	//是否需要 body
	if iresp {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			return nil, et - st, st, "", request.RemoteAddr, errors.Wrap(err, "HttpCLientCommen: get body error")
		}
		return r, et, st, string(body), request.RemoteAddr, err
	}

	return r, et - st, st, "", request.RemoteAddr, err

}

func RedirectFunc(req *http.Request, via []*http.Request) error {
	fmt.Println(req.RequestURI)
	// 如果返回 非nil 则禁止向下重定向 返回nil 则 一直向下请求 10 次 重定向
	return http.ErrUseLastResponse
}

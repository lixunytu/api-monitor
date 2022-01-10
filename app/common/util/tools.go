package util

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/robfig/cron"

	//"github.com/lixunytu/cron"
	"io"
	"math/rand"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/pkg/errors"
)

func JsonTointerfaceMap(sj string) (mp map[string]interface{}, err error) {
	mp = make(map[string]interface{}, 20)
	if err = json.Unmarshal([]byte(sj), &mp); err != nil {
		return nil, errors.Wrap(err, "func JsonTointerfaceMap error")
	}

	return
}

func Urlanalysis(urls string, gdm map[string]interface{}) (delurl string, err error) {
	//解析 传过来的 url
	u, err := url.Parse(urls)
	if err != nil {
		return "", errors.Wrap(err, "Urlanalysis error: url.Parse err")
	}
	//获取 get 参数
	urlparam := u.RawQuery
	//get 参数转换为 map
	m, err := url.ParseQuery(urlparam)
	if err != nil {
		return "", errors.Wrap(err, "Urlanalysis error: url.ParseQuery err")
	}

	//处理 get 数据
	gdm2 := InterfaceToString(gdm)
	//定义当前时间戳,后续使用
	timeUnix := time.Now().Unix()
	timestamp := strconv.Itoa(int(timeUnix))

	count := 0
	for k, v := range *gdm2 {
		k = strings.TrimSpace(k) //去掉空格
		if v == "${timestamp}" {
			m.Set(k, timestamp)
		} else if len(v) > 0 && string(v[0]) == "$" {
			if v == "${md5NoxMobile}" {
				m.Set(k, md5Encode(gdm["timestamp"].(string)+gdm["url"].(string), timestamp))
			} else if v == "${getRandomStr}" {
				m.Set(k, GetRandomStr(count))
			} else {
				m.Set(k, v)
			}
		} else {
			m.Set(k, v)
		}
		count++
	}
	u.RawQuery = m.Encode()
	return u.String(), nil
}

func InterfaceToString(in map[string]interface{}) *map[string]string {
	m := make(map[string]string, len(in))
	for k, v := range in {
		switch v.(type) {
		case string:
			m[k] = v.(string)
		case int:
			m[k] = strconv.Itoa(v.(int))
		case bool:
			m[k] = strconv.FormatBool(v.(bool))
		}
	}
	return &m

}

func PostAnalysis(pd string) (pds string, err error) {

	pdm, err := JsonTointerfaceMap(pd)
	if err != nil {
		return pd, errors.Wrap(err, "PostAnalysis error")
	}

	//定义当前时间戳,后续使用
	timeUnix := time.Now().Unix()
	timestamp := strconv.Itoa(int(timeUnix))

	count := 0
	for k, v := range pdm {
		k = strings.TrimSpace(k)
		var v1 string
		ok := false
		switch v.(type) {
		case string:
			ok = true
			v1 = v.(string)
		}

		if ok {
			if v1 == "${timestamp}" {
				pdm[k] = timestamp
			} else if len(v1) > 0 && string(v1[0]) == "$" {
				if v1 == "${md5NoxMobile}" {
					pdm[k] = md5Encode(timestamp+pdm["url"].(string), "nox_video")
				} else if v1 == "${getRandomStr}" {
					pdm[k] = GetRandomStr(count)
				}
			}
		}
		count++

	}
	return MapTointerfaceJson(pdm)
}

func md5Encode(tmp string, key string) string {
	tmp = key + tmp + key
	w := md5.New()
	io.WriteString(w, tmp)                   //将str写入到w中
	md5str2 := fmt.Sprintf("%x", w.Sum(nil)) //w.Sum(nil)将w的hash转成[]byte格式
	return md5str2
}

func GetRandomStr(key int) string {

	str := "ABCDEFGHIGKLMNOPQRSTUVWXYZabcdefghigklmnopqrstuvwxyz0123456789"
	result_str := ""
	rand.Seed(time.Now().Unix() + int64(key))

	for i := 0; i < 16; i++ {
		index := rand.Intn(62)
		result_str = string(str[index]) + result_str
	}
	return result_str
}

func MapTointerfaceJson(stringmap map[string]interface{}) (jsonstring string, err error) {
	jsonbyte, err := json.Marshal(stringmap)
	if err != nil {
		return "", errors.Wrap(err, "MapTointerface error")
	}
	jsonstring = string(jsonbyte)
	return
}

func JsonToMap(stringjson string) (mp map[string]string, err error) {
	mp = make(map[string]string, 20)
	err = json.Unmarshal([]byte(stringjson), &mp)
	return
}

func MapToJson(stringmap map[string]string) (jsonstring string, err error) {
	if len(stringmap) == 0 {
		return "",nil
	}
	jsonbyte, err := json.Marshal(stringmap)
	if err != nil {
		return "", errors.Wrap(err,"MapToJson: error")
	}
	jsonstring = string(jsonbyte)
	return
}

func Md5V(str string) string  {
	str=str+"admin_imooc_node"
	h := md5.New()
	h.Write([]byte(str))
	return strings.ToUpper(hex.EncodeToString(h.Sum(nil)))
}

func ParseCron(e string) (cronString [10]string,err error) {

	var p cron.Schedule
	if len(strings.Split(e, " ")) == 5{
		p, err = cron.ParseStandard(e)
	} else if len(strings.Split(e, " ")) == 6 {
		p, err = cron.Parse(e)
	} else {
		return cronString,errors.New("表达式有问题，请重新输入")
	}
	if err != nil {
		return cronString,err
	}
	t := time.Now()
	k:=0
	for i := 10; i > 0; i-- {
		t = p.Next(t)
		cronString[k]=t.Format("2006-01-02 15:04:05")
		k++
	}
	return cronString,nil
}
package control

import (
	"math/rand"
	"strconv"
	"strings"
	"time"

	"self/app/models"

	"github.com/lixunytu/cron"
	"github.com/pkg/errors"
)

type Control struct {
	Monitor models.MonitorInfo
	CronTab *cron.Cron
	Spec    string
}

const (
	second = 60
	hour   = 3600
)

var (
	AllControlList []Control
	AllIdList      []int
)

func InitControl() (err error) {
	m, err := models.MonitorInfo{}.GetAll()
	if err != nil {
		//TODO
		return errors.Wrap(err, "InitControl error")
	}

	for k, info := range m {
		spec := ""
		var c Control
		if info.MIsdefault == 0 {
			rand.Seed(time.Now().Unix() + int64(k))
			s := rand.Intn(second)
			m := rand.Intn(second)
			spec = GetSpecString(info.MFrequency, m, s)
		} else {
			spec = strings.TrimSpace(info.MCronString)
		}

		c = Control{
			Monitor: *info,
			CronTab: cron.New(),
			Spec:    spec,
		}

		AllControlList = append(AllControlList, c)
		AllIdList = append(AllIdList, info.Id)
	}

	for _, c := range AllControlList {
		if c.Monitor.Status == 0 {
			StartCron(c)
		}
	}


	return
}

func GetSpecString(f, m, s int) string {
	spec := ""
	if f >= hour {
		spec = strconv.Itoa(int(s)) + " " + strconv.Itoa(m) + " */" + strconv.Itoa(int(f)/second/second) + " * * *"
	} else if f > second {
		spec = strconv.Itoa(int(s)) + " */" + strconv.Itoa(int(f)/second) + " * * * *"
	} else {
		spec = "*/" + strconv.Itoa(int(f)) + " * * * * *"
	}
	return spec
}

func StartCron(c Control) {
	c.CronTab.AddFunc(c.Spec, func() {
		Producer(c)
	})
	c.CronTab.Start()
}

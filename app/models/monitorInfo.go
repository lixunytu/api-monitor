package models

import (
	"self/app/databases/mysql"
	"time"
)

type MonitorInfo struct {
	Id               int       `json:"id,omitempty"`
	ProductId        int       `json:"product_id,omitempty"`
	AlarmId          int       `json:"alarm_id,omitempty"`
	Name             string    `json:"name,omitempty"`
	MRank            int       `json:"m_rank,omitempty"`
	MType            int       `json:"m_type,omitempty"`
	Status           int       `json:"status"`
	MachineRegion    string    `json:"machine_region,omitempty"`
	ServerUrl        string    `json:"server_url,omitempty"`
	MWaitTime        int       `json:"m_wait_time,omitempty"`
	MRetry           int       `json:"m_retry,omitempty"`
	MIsdefault       int       `json:"m_isdefault"`
	MCronString      string    `json:"m_cron_string,omitempty"`
	RequestHeaders   string    `json:"request_headers,omitempty"`
	PostData         string    `json:"post_data,omitempty"`
	GetData          string    `json:"get_data,omitempty"`
	ExpectedCode     string    `json:"expected_code,omitempty"`
	MRules           string    `json:"m_rules,omitempty"`
	MRort            int       `json:"m_port,omitempty"`
	ProcessName      string    `json:"process_name,omitempty"`
	IpAddress        string    `json:"ip_address,omitempty"`
	Remarks          string    `json:"remarks"`
	AlarmCallbackApi string    `json:"alarm_callback_api,omitempty"`
	MTimeout         int       `json:"m_timeout,omitempty"`
	MFrequency       int       `json:"m_frequency,omitempty"`
	LastRunTime      int       `json:"last_run_time,omitempty"`
	Zookeeper        string    `json:"zookeeper,omitempty"`
	TopicName        string    `json:"topic_name,omitempty"`
	GroupName        string    `json:"group_name,omitempty"`
	KafkaLag         int       `json:"kafka_lag,omitempty"`
	UpdatedAt        time.Time `json:"updated_at,omitempty"`
	CreatedAt        time.Time `json:"created_at,omitempty"`
	CreatedUser      string    `json:"created_user,omitempty"`
	UpdatedUser      string    `json:"updated_user,omitempty"`
}

func (d MonitorInfo) CreateA(m *MonitorInfo) (err error) {
	err = mysql.DB.Create(&m).Error
	return
}

func (d MonitorInfo) GetAll() (m []*MonitorInfo, err error) {
	if err = mysql.DB.Where("status != -1").Find(&m).Error; err != nil {
		return nil, err
	}
	return
}

func (d MonitorInfo) GetAllForParams(num, size int) (m []*MonitorInfo, err error) {
	offset := (num - 1) * size
	if err = mysql.DB.Where("status != -1").Order("id desc").Limit(size).Offset(offset).Find(&m).Error; err != nil {
		return nil, err
	}
	return
}

func (d MonitorInfo) GetAllCount() (num int, err error) {
	if err = mysql.DB.Model(&MonitorInfo{}).Where("status != -1").Count(&num).Error; err != nil {
		return 0, err
	}
	return
}

func (d MonitorInfo) GetA(id []string) (m *MonitorInfo, err error) {
	m = new(MonitorInfo)
	if err = mysql.DB.Debug().Where("id IN (?)", id).First(m).Error; err != nil {
		return nil, err
	}
	return
}

func (d MonitorInfo) UpdateA(id string, m *MonitorInfo) (err error) {
	err = mysql.DB.Model(MonitorInfo{}).Where("id = ?", id).Update(m).Error
	return
}

func (d MonitorInfo) DeleteA(id interface{}) (err error) {
	err = mysql.DB.Model(&MonitorInfo{}).Where("id IN (?)", id).Update("status", -1).Error
	return
}

func (d MonitorInfo) StartStatus(id interface{}) (err error) {
	err = mysql.DB.Model(&MonitorInfo{}).Where("id IN (?)", id).Update("status", 0).Error
	return
}

func (d MonitorInfo) StopStatus(id interface{}) (err error) {
	err = mysql.DB.Model(&MonitorInfo{}).Where("id IN (?)", id).Update("status", 1).Error
	return
}

package models

import (
	"self/app/databases/mysql"
)

type AlarmInfo struct {
	Id                   int    `json:"id,omitempty"`
	Name                 string `json:"name,omitempty"`
	RealTime             int    `json:"real_time,omitempty"`
	RecipientMail        string `json:"recipient_mail,omitempty"`
	AlarmType            string `json:"alarm_type,omitempty"`
	DingGroup            string `json:"ding_group,omitempty"`
	RecipientMessage     string `json:"recipient_message,omitempty"`
	RecipientPhone       string `json:"recipient_phone,omitempty"`
	DingGroupFailNum     int    `json:"ding_group_fail_num,omitempty"`
	DingGroupAlarmNum    int    `json:"ding_group_alarm_num,omitempty"`
	DingGroupSummaryTime int    `json:"ding_group_summary_time,omitempty"`
	MessageFailNum       int    `json:"message_fail_num,omitempty"`
	MessageAlarmNum      int    `json:"message_alarm_num,omitempty"`
	MessageSummaryTime   int    `json:"message_summary_time,omitempty"`
	MessageOffDayAlarm   int    `json:"message_off_day_alarm,omitempty"`
	PhoneFailNum         int    `json:"phone_fail_num,omitempty"`
	PhoneAlarmNum        int    `json:"phone_alarm_num,omitempty"`
	PhoneSummaryTime     int    `json:"phone_summary_time,omitempty"`
	PhoneOffDayAlarm     int    `json:"phone_off_day_alarm,omitempty"`
	CreatedAt            int    `json:"created_at,omitempty"`
	UpdatedAt            int    `json:"updated_at,omitempty"`
	CreateUser           int    `json:"create_user,omitempty"`
	UpdateUser           int    `json:"update_user,omitempty"`
	Remarks              string `json:"remarks,omitempty"`
	MailTitle            string `json:"mail_title,omitempty"`
	DingGroupContent     string `json:"ding_group_content,omitempty"`
	MessageContent       string `json:"message_content,omitempty"`
	Status               int    `json:"status,omitempty"`
}

func (a AlarmInfo) CreateA(alarm *AlarmInfo) (err error) {
	err = mysql.DB.Create(&alarm).Error
	return
}

func (a AlarmInfo) GetAll(pageNum,pageSize int) (alarm []*AlarmInfo, err error) {
	offset := (pageNum-1)*pageSize
	if err = mysql.DB.Offset(offset).Limit(pageSize).Find(&alarm).Error; err != nil {
		return nil, err
	}
	return
}

func (a AlarmInfo) GetAllCount() (n int, err error) {
	err = mysql.DB.Model(AlarmInfo{}).Count(&n).Error
	return
}

func (a AlarmInfo) GetA(id string) (alarm *AlarmInfo, err error) {
	alarm = new(AlarmInfo)
	if err = mysql.DB.Debug().Where("id=?", id).First(alarm).Error; err != nil {
		return nil, err
	}
	return
}

func (a AlarmInfo) UpdateA(id string,alarm *AlarmInfo) (err error) {
	err = mysql.DB.Where("id=?",id).Save(alarm).Error
	return
}

func (a AlarmInfo) DeleteA(id string) (err error) {
	err = mysql.DB.Where("id=?", id).Delete(&AlarmInfo{}).Error
	return
}

func (a AlarmInfo) GetAllWithIdName() (alarm []*AlarmInfo, err error) {
	if err = mysql.DB.Select([]string{"id", "name"}).Where("status=?", 0).Find(&alarm).Error; err != nil {
		return nil, err
	}
	return
}


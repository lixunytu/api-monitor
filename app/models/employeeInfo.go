package models

import (
	"database/sql"
	"github.com/pkg/errors"
	"math/big"
	"self/app/databases/mysql"
)

type EmployeeInfo struct {
	Id           int     `json:"id"`
	Type         int     `json:"type""`
	DepartMentId int     `json:"depart_ment_id"`
	Name         string  `json:"name"`
	Email        string  `json:"email"`
	PhoneNumber  string  `json:"phone_number"`
	DingUrl      string  `json:"ding_url"`
	CreatedAt    big.Int `json:"created_at"`
	UpdatedAt    big.Int `json:"updated_at"`
}


func (a EmployeeInfo) CreateA (employ *EmployeeInfo) (err error) {
	err = mysql.DB.Create(&employ).Error
	return
}

func (a EmployeeInfo) GetAll(num,size int) (employ []*EmployeeInfo, err error) {
	offset := (num-1)*size
	if err = mysql.DB.Offset(offset).Limit(size).Find(&employ).Error; err != nil {
		if errors.Is(err,sql.ErrNoRows) {
			return nil,nil
		}
		return nil, err
	}
	return
}

func (a EmployeeInfo) GetCount() (n int, err error) {
	err = mysql.DB.Model(EmployeeInfo{}).Count(&n).Error
	if err != nil {
		return 0, err
	}
	return
}

func (a EmployeeInfo) GetSum(id []string) (employ []*EmployeeInfo, err error) {

	if err = mysql.DB.Where("id IN (?)",id).Find(&employ).Error; err != nil {
		return nil, err
	}

	return
}

func (a EmployeeInfo) GetA(id string) (employ *EmployeeInfo, err error) {
	employ = new(EmployeeInfo)
	if err = mysql.DB.Debug().Where("id=?", id).First(employ).Error; err != nil {
		return nil, err
	}
	return
}

func (a EmployeeInfo) UpdateA (employ *EmployeeInfo) (err error) {
	err = mysql.DB.Model(EmployeeInfo{}).Where("id=?",employ.Id).Update(employ).Error
	return
}

func (a EmployeeInfo) DeleteA (id string) (err error) {
	err = mysql.DB.Where("id=?", id).Delete(&EmployeeInfo{}).Error
	return
}

func (a EmployeeInfo) GetSumByKeyWords(key string) (employ []*EmployeeInfo, err error) {
	if err = mysql.DB.Debug().Where("email like ?","%"+key+"%").
		Or(" name like ?","%"+key+"%").
		Or(" phone_number like ?","%"+key+"%").Find(&employ).
		Error; err != nil {
		return nil, err
	}
	return
}
package models

import "self/app/databases/mysql"

type DepartMentInfo struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}



func (d DepartMentInfo)CreateA(department *DepartMentInfo) (err error) {
	err = mysql.DB.Create(&department).Error
	return
}

func (d DepartMentInfo)GetAll() (department []*DepartMentInfo, err error) {
	if err = mysql.DB.Find(&department).Error; err != nil {
		return nil, err
	}
	return
}

func (d DepartMentInfo)GetA(id string) (department *DepartMentInfo, err error) {
	department = new(DepartMentInfo)
	if err = mysql.DB.Debug().Where("id=?", id).First(department).Error; err != nil {
		return nil, err
	}
	return
}

func (d DepartMentInfo)UpdateA(department *DepartMentInfo) (err error) {
	err = mysql.DB.Save(department).Error
	return
}

func (d DepartMentInfo)DeleteA(id string) (err error) {
	err = mysql.DB.Where("id=?", id).Delete(&DepartMentInfo{}).Error
	return
}
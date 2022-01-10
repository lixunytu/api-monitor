package models

import (
	"self/app/databases/mysql"
)

type Executor struct {
	Id           int     `json:"id"`
	Name string `json:"name"`
	Ipaddr string `json:"ipaddr"`
}


func (a Executor) CreateA (exec *Executor) (err error) {
	err = mysql.DB.Create(&exec).Error
	return
}

func (a Executor) GetAll() (exec []*Executor, err error) {
	if err = mysql.DB.Find(&exec).Error; err != nil {
		return nil, err
	}
	return
}

func (a Executor) GetA(id string) (exec *Executor, err error) {
	exec = new(Executor)
	if err = mysql.DB.Debug().Where("id=?", id).First(exec).Error; err != nil {
		return nil, err
	}
	return
}

func (a Executor) UpdateA (exec *Executor) (err error) {
	err = mysql.DB.Save(exec).Error
	return
}

func (a Executor) DeleteA (id string) (err error) {
	err = mysql.DB.Where("id=?", id).Delete(&Executor{}).Error
	return
}
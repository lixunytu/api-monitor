package models

import "self/app/databases/mysql"

type ProductInfo struct {
	Id           int    `json:"id"`
	DepartMentId int    `json:"depart_ment_id"`
	Name         string `json:"name"`
	Pid          int    `json:"pid"`
	Level        int    `json:"level"`
	ProductId    int    `json:"product_id"`
}

func (d ProductInfo) CreateA(m *ProductInfo) (err error) {
	err = mysql.DB.Create(&m).Error
	return
}

func (d ProductInfo) GetAll() (m []*ProductInfo, err error) {
	if err = mysql.DB.Find(&m).Error; err != nil {
		return nil, err
	}
	return
}

func (d ProductInfo) GetA(id string) (m *ProductInfo, err error) {
	m = new(ProductInfo)
	if err = mysql.DB.Debug().Where("id=?", id).First(m).Error; err != nil {
		return nil, err
	}
	return
}

func (d ProductInfo) UpdateA(m *ProductInfo) (err error) {
	err = mysql.DB.Save(m).Error
	return
}

func (d ProductInfo) DeleteA(id string) (err error) {
	err = mysql.DB.Where("id=?", id).Delete(&ProductInfo{}).Error
	return
}

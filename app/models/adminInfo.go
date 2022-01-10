package models

import (
	"math/big"
	"self/app/databases/mysql"
)

// Todo Model
type AdminInfo struct {
	ID           int     `json:"id"`
	Role         int     `json:"role"`
	DepartmentId int     `json:"department_id"`
	ProductId    int     `json:"product_id"`
	EmployeeId   int     `json:"employee_id"`
	CreatedAt    big.Int `json:"created_at"`
	UpdatedAt    big.Int `json:"updated_at"`
}

func (a AdminInfo) CreateA (admin *AdminInfo) (err error) {
	err = mysql.DB.Create(&admin).Error
	return
}

func (a AdminInfo) GetAll() (admin []*AdminInfo, err error) {
	if err = mysql.DB.Find(&admin).Error; err != nil {
		return nil, err
	}
	return
}

func (a AdminInfo) GetA(id string) (admin *AdminInfo, err error) {
	admin = new(AdminInfo)
	if err = mysql.DB.Debug().Where("id=?", id).First(admin).Error; err != nil {
		return nil, err
	}
	return
}

func (a AdminInfo) UpdateA (admin *AdminInfo) (err error) {
	err = mysql.DB.Save(admin).Error
	return
}

func (a AdminInfo) DeleteA (id string) (err error) {
	err = mysql.DB.Where("id=?", id).Delete(&AdminInfo{}).Error
	return
}

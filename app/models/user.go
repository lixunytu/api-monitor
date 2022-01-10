package models

import (
	"database/sql"
	"self/app/common/util"
	"self/app/databases/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type LoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AdminUser struct {
	Id       int      `json:"id,omitempty"`
	Username string   `json:"username,omitempty"`
	Password string   `json:"password,omitempty"`
	Role     string   `json:"role,omitempty"`
	Ncikname string   `json:"ncikname,omitempty"`
	Avatar   string   `json:"avatar,omitempty"`
	Roles    []string `json:"roles,omitempty" gorm:"-"`
}

func (a AdminUser) CreateA(m *AdminUser) (err error) {
	err = mysql.DB.Create(&m).Error
	return
}

func (a AdminUser) GetAll() (m []*AdminUser, err error) {
	if err = mysql.DB.Find(&m).Error; err != nil {
		return nil, err
	}
	return
}

func (d AdminUser) GetA(id string) (m *AdminUser, err error) {
	m = new(AdminUser)
	if err = mysql.DB.Debug().Where("id=?", id).First(m).Error; err != nil {
		return nil, err
	}
	return
}

func (d AdminUser) GetAByName(username string) (m *AdminUser, err error) {
	m = new(AdminUser)
	if err = mysql.DB.Select([]string{"id","username","ncikname","role","avatar"}).Where("username=?", username).First(m).Error; err != nil {
		return nil, err
	}
	m.Roles = []string{m.Role}
	return
}

func (d AdminUser) GetAByNameAndPass(username, password string) (m *AdminUser, err error) {
	m = new(AdminUser)
	if err = mysql.DB.Debug().Where("username=? and password=?", username, password).First(m).Error; err != nil {
		return nil, err
	}

	if m==nil {
		return nil,sql.ErrNoRows
	}

	m.Roles = []string{m.Role}
	return
}

func (d AdminUser) UpdateA(m *AdminUser) (err error) {
	err = mysql.DB.Save(m).Error
	return
}

func (d AdminUser) DeleteA(id string) (err error) {
	err = mysql.DB.Where("id=?", id).Delete(&AdminUser{}).Error
	return
}

func LoginCheck(req LoginReq) (isPass bool, user *AdminUser, err error) {
	username := req.Username
	password := util.Md5V(req.Password)

	user, err = AdminUser{}.GetAByNameAndPass(username, password)

	return true, user, err
}

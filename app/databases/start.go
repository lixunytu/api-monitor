package databases

import (
	"fmt"

	"self/app/config"
	"self/app/databases/mysql"
	"self/app/models"

	"github.com/pkg/errors"

)



func MysqlStart() error {
	// 创建数据库
	// sql: CREATE DATABASE bubble;
	// 连接数据库
	err := mysql.InitMySQL(config.Conf.MySQLConfig)
	if err != nil {
		fmt.Printf("init mysql failed, err:%v\n", err)
		return errors.Wrap(err,"mysql start error")
	}
	//defer mysql.Close() // 程序退出关闭数据库连接
	// 模型绑定

	mysql.DB.AutoMigrate(&models.AdminInfo{})
	mysql.DB.AutoMigrate(&models.AlarmInfo{})
	mysql.DB.AutoMigrate(&models.DepartMentInfo{})
	mysql.DB.AutoMigrate(&models.Executor{})
	mysql.DB.AutoMigrate(&models.MonitorInfo{})
	mysql.DB.AutoMigrate(&models.ProductInfo{})
	mysql.DB.AutoMigrate(&models.AdminUser{})
	mysql.DB.AutoMigrate(&models.EmployeeInfo{})


	return nil

}

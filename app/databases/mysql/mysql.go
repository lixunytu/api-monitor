package mysql

import (
	"self/app/config"

	"github.com/jinzhu/gorm"
	//_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var (
	DB *gorm.DB
)

func InitMySQL(cfg *config.MySQLConfig) (err error) {
	//dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
	//	cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DB)
	//DB, err = gorm.Open("mysql", dsn)
	DB, err = gorm.Open("sqlite3", "monitor.db")
	if err != nil {
		return
	}
	return DB.DB().Ping()
}

func Close() {
	DB.Close()
}

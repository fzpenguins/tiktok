package db

import (
	"log"
	"strings"
	"tiktok/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	var err error
	dsn := strings.Join([]string{config.SqlUserName, ":", config.SqlPassword,
		"@tcp(" + config.MysqlIP + ")/", config.DataBase,
		"?charset=utf8mb4&parseTime=True&loc=Local"}, "")
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	err = Migrate()
	if err != nil {
		panic(err)
	}
}

func Migrate() (err error) {
	if err = DB.AutoMigrate(&Comment{}); err != nil {
		log.Println("auto migrate failed", err)
		return
	}
	if err = DB.AutoMigrate(&Like{}); err != nil {
		log.Println("auto migrate failed", err)
		return
	}
	return
}

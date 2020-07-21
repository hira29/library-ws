package config

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func ConfigSql() *gorm.DB {
	user := "root"
	password := "12345678"
	host := "127.0.0.1:3306"
	dbName := "perpustakaan"

	db, err := gorm.Open("mysql", user+":"+password+"@tcp("+host+")/"+dbName+"?parseTime=true")
	fmt.Println("connect")

	if err != nil {
		fmt.Println(err)
	}
	return db
}

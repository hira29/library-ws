package config

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func ConfigSql() *gorm.DB {
	//user := "root"
	//password := "12345678"
	//host := "127.0.0.1:3306"
	//dbName := "perpustakaan"

	user := "l6icdro9m3gvbqxo"
	password := "mbo228x5phcf09zs"
	host := "un0jueuv2mam78uv.cbetxkdyhwsb.us-east-1.rds.amazonaws.com:3306"
	dbName := "vf6ax87hwl544kc6"

	db, err := gorm.Open("mysql", user+":"+password+"@tcp("+host+")/"+dbName+"?parseTime=true")
	fmt.Println("connect")

	if err != nil {
		fmt.Println(err)
	}
	return db
}

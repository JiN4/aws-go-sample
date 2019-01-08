package model

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var db = NewDBConn()

func NewDBConn() *gorm.DB {
	db, err := gorm.Open(GetDBConfig())
	db.LogMode(true)

	if err != nil {
		panic(err)
	}

	return db
}

func GetDBConn() *gorm.DB {
	return db
}

func GetDBConfig() (string, string) {
	return "mysql", "root:kobedenshi@tcp(gs-group-weet-mysql.cnvtlozgscam.ap-northeast-1.rds.amazonaws.com:3306)/gs_group_weet?charset=utf8&parseTime=True&loc=Local"
}

package main

import (
	_ "github.com/jinzhu/gorm"
	"github.com/weet/model"
)

func main() {

	db := model.GetDBConn()

	// migrate
	db.CreateTable(&model.Sample{})

}

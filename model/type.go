package model

import (
	_ "time"

	"github.com/jinzhu/gorm"
)

// //model内で共通のfieldをここに抜き出す
//type //model struct {
//    ID      uint `gorm:"primary_key" json:"id"`
//    CreatedAt time.Time
//    UpdatedAt time.Time
//}

// 自分のユーザー情報
type Sample struct {
	//Model
	gorm.Model
	Name    string `gorm:"not null"`
	Comment string `gorm:"not null"`
}

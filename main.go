import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	gorm.Model
	Name     string `gorm:"size:255"`
	Password string `gorm:"size:255"`
	Email    string `gorm:"size:255"`
}

db, err := gorm.Open("mysql", "root:kobedenshi@tcp(gs-group-weet-mysql.cnvtlozgscam.ap-northeast-1.rds.amazonaws.com:3306)/gs_group_weet?charset=utf8&parseTime=True&loc=Local")
if err != nil {
    panic("failed to connect database")
}

// 実行完了後DB接続を閉じる
defer db.Close() 

// ログ出力を有効にする
db.LogMode(true) 

sample.go
// DBエンジンを「InnoDB」に設定
db.Set("gorm:table_options", "ENGINE=InnoDB")

// マイグレーション
db.AutoMigrate(&User{})
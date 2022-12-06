package model

import (
	"fmt"
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
)

var db *gorm.DB

func init() {
	// [ユーザ名]:[パスワード]@tcp([ホスト名]:[ポート番号])/[データベース名]?charset=[文字コード]
	dbconf := "user:password@tcp(db_xb:3306)/xb-map?charset=utf8mb4"
	var err error
	db, err = gorm.Open(mysql.Open(dbconf), &gorm.Config{})
	if err != nil {
		fmt.Println("データベース接続失敗")
	} else {
		fmt.Println("データベース接続成功")
	}
	
	// dbをmigrateする
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Shop{})
}
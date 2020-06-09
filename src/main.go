package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

type Product struct {
	gorm.Model
	Code string
	// 64bitの符号なし整数
	Price uint
}

func main() {
	db, err := gorm.Open("mysql", "root:rootpass@tcp(mysql:3306)/gorm_db")
	if err != nil {
		panic("cannot connect db.")
	}

	log.Print("success in connecting db.")
	log.Print("%T", db)

	defer db.Close()

	db.AutoMigrate(&Product{}) // マイグレーション

	db.Create(&Product{ // データの挿入
		Code: "L1212",
		Price: 1000,
	})
}
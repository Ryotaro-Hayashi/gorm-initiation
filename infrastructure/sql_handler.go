package infrastructure

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gorm-initiation/domain"
	"gorm-initiation/interface/database"
)

type SqlHandler struct {
	Conn *gorm.DB
}

// database.SqlHandlerはinterfaceであり、ポインタ型
func NewSqlHandler() database.SqlHandler {
	db, err := gorm.Open("mysql", "root:rootpass@tcp(mysql:3306)/gorm_db")
	if err != nil {
		panic(err.Error())
	}

	db.AutoMigrate(domain.User{}) // マイグレーション

	db.Create(domain.User{ // データの挿入
		Name: "testman",
	})

	sqlHandler := new(SqlHandler)
	sqlHandler.Conn = db

	return sqlHandler
}

// interface/databaseとusecaseのinterfaceの実装
func (sqlHandler *SqlHandler) Find() (users domain.Users, err error){
	// データを取得して、users と err に格納
	err = sqlHandler.Conn.Find(&users).Error

	return
}
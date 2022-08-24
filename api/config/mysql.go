package config

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewMysqlConfig() *gorm.DB {
	dsn := "user:password@tcp(db:3306)/db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	fmt.Println(err)
	fmt.Println(db)

	return db
}

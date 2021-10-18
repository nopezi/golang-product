package configs

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func Database() (*gorm.DB, error) {
	dbname := "golang-products"
	host := "127.0.0.1"
	user := "root"
	pass := "my-secret-pw"
	port := "3306"

	konek := user + `:` + pass + `@tcp(`+ host + `:` + port +`)/` + dbname +`?charset=utf8mb4&parseTime=True&loc=Local`
	db, err := gorm.Open(mysql.Open(konek), &gorm.Config{})

	if err != nil {
		fmt.Println("db error", err.Error())
	}

	return db, err
}
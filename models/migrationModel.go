package models

import (
	"fmt"
	"golang-produk/configs"
	"golang-produk/structs"
	"io/ioutil"
)

// cek jika table images sudah ada
func CekTableImage() structs.CekTableImage {
	var Table structs.CekTableImage
	db, _ := configs.Database()
	sql := `SELECT table_schema
			FROM information_schema.tables
			WHERE table_schema = 'golang-products' 
				AND table_name = 'images'
			LIMIT 1;`
	db.Raw(sql).Scan(&Table)
	return Table
}

// cek jika table products sudah ada
func CekTableProducts() structs.CekTableImage {
	var Table structs.CekTableImage
	db, _ := configs.Database()
	sql := `SELECT table_schema
			FROM information_schema.tables
			WHERE table_schema = 'golang-products' 
				AND table_name = 'products'
			LIMIT 1;`
	db.Raw(sql).Scan(&Table)
	return Table
}

//  cek database
func CekDatabase() structs.CekTableImage {
	var Table structs.CekTableImage
	db, _ := configs.Database()
	sql := `SELECT table_schema
			FROM information_schema.tables
			WHERE table_schema = 'golang-products'
			LIMIT 1;`
	db.Raw(sql).Scan(&Table)
	return Table
}

func MigrationDatabase()  {
	db, _ := configs.DatabaseMigration()
	db.Exec("CREATE DATABASE IF NOT EXISTS `golang-products`;")
}

func TableImage()  {
	db, _ := configs.Database()
	cek, err := ioutil.ReadFile("migration/create_images_table.sql")
	if err != nil {
        fmt.Print(err)
    }
	// cek table images
	ada := CekTableImage()
	// jika table nya kosong, maka bikin baru
	if ada.TableSchema == "" {
		db.Exec(string(cek))
	}
}

func TableProducts()  {
	db, _ := configs.Database()
	cek, err := ioutil.ReadFile("migration/create_product_table_up.sql")
	if err != nil {
        fmt.Print(err)
    }
	// cek table images
	ada := CekTableProducts()
	// jika table nya kosong, maka bikin baru
	if ada.TableSchema == "" {
		db.Exec(string(cek))
	}
}
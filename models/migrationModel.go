package models

import (
	"golang-produk/configs"
	"golang-produk/structs"
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
	db, _ := configs.Database()
	db.Migrator().CurrentDatabase()
}
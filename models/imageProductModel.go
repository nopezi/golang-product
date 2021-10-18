package models

import (
	"golang-produk/configs"
	"golang-produk/structs"
)

func GetImageProduct(id int64) structs.ImageProduct {
	var hasil structs.ImageProduct
	// ambil fungsi db
	db, _ := configs.Database()
	//  ambil data yang sudah di update
	db.Table("images").Where("id = ?", id).Find(&hasil)
	// kirim data hasil
	return hasil
}

func AddImageProduct(input structs.Images)  {
	db, _ := configs.Database()
	db.Table("images").Create(&input)
}

func UpdateImageProduct(input structs.UpdateImages) {
	// ambil fungsi db
	db, _ := configs.Database()
	//  update data nya
	db.Table("images").Updates(&input)
}

func DeleteImageProduct(id int64)  {
	var imageProduct structs.Images
	db, _ := configs.Database()

	db.Where("id = ?", id).Delete(&imageProduct)
}
package models

import (
	// "fmt"
	"golang-produk/configs"
	"golang-produk/structs"
)

func GetProduct() []structs.Product {
	var product []structs.Product
	db, _ := configs.Database()
	sql := `SELECT 
				products.*,
				concat('assets/', images.name) as image 
			FROM 
				products
			LEFT JOIN
				images on images.id_product = products.id`
	db.Raw(sql).Scan(&product)
	
	return product
}

func AddProduct(input structs.InputProduct)  {
	db, _ := configs.Database()
	db.Table("products").Create(&input)
}

func UpdateProduct(input structs.UpdateProduct) structs.UpdateProduct {
	var hasil structs.UpdateProduct
	// ambil fungsi db
	db, _ := configs.Database()
	//  update data nya
	db.Table("products").Updates(&input)
	//  ambil data yang sudah di update
	db.Table("products").Where("id = ?", input.Id).Find(&hasil)
	// kirim data hasil
	return hasil
}

func DeleteProduct(id int)  {
	db, _ := configs.Database()
	var Product structs.Product

	db.Where("id = ?", id).Delete(Product)
}
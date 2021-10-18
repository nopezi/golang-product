package main

import (
	"fmt"
	"golang-produk/models"
	"golang-produk/routers"
)

func main()  {

	// migrasi database dan table
	models.MigrationDatabase()
	models.TableImage()
	models.TableProducts()

	fmt.Println("berjalan di port 9090")
	routers.HandleRequests()
}
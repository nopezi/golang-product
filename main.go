package main

import (
	"fmt"
	"golang-produk/models"
	// "golang-produk/routers"
)

func main()  {
	fmt.Println("berjalan di port 9090")
	// routers.HandleRequests()
	models.MigrationDatabase()
}
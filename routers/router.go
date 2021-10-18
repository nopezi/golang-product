package routers

import (
	"log"
	"net/http"
	"golang-produk/controllers"
)

func HandleRequests()  {

	http.HandleFunc("/getToken", controllers.TokenPage)
	http.Handle("/coba", isAuthorized(controllers.CobaPage))

	http.Handle("/products", isAuthorized(controllers.GetProduct))
	http.Handle("/add_product", isAuthorized(controllers.AddProduct))
	http.Handle("/update_product", isAuthorized(controllers.UpdateProduct))
	http.Handle("/delete_product", isAuthorized(controllers.DeleteProduct))

	http.Handle("/upload_image_product", isAuthorized(controllers.UploadImageProduct))
	http.Handle("/update_image_product", isAuthorized(controllers.EditImageProduct))
	http.Handle("/delete_image_product", isAuthorized(controllers.DeleteImageProduct))

	log.Fatal(http.ListenAndServe(":9090", nil))
}
package controllers

import (
	"encoding/json"
	"golang-produk/configs"
	"golang-produk/models"
	"golang-produk/structs"
	"io/ioutil"
	"net/http"
	"time"
)

func GetProduct(w http.ResponseWriter, r *http.Request)  {
	var res structs.Result
	data := models.GetProduct()

	if data == nil {
		res.Status = "false"
		res.Message = "Data produk kosong"
	} else {
		res.Status = "true"
		res.Message = "Berhasil mendapatkan data produk"
	}
	
	res.Data = data

	configs.ResultOk(w, res)
}

func AddProduct(w http.ResponseWriter, r *http.Request)  {
	input, _ := ioutil.ReadAll(r.Body)
	var res structs.Result
	var produkInput structs.InputProduct

	json.Unmarshal(input, &produkInput)
	currentTime := time.Now()
	waktu := currentTime.Format("2006-01-02 15:04:05")
	produkInput.CreatedAt = waktu

	if produkInput.NameProduct == "" {
		res.Status = "false"
		res.Message = "nama produk tidak boleh kosong"
	} else if produkInput.Price == 0 {
		res.Status = "false"
		res.Message = "harga tidak boleh kosong"
	} else {
		res.Status = "true"
		res.Message = "berhasil tambah produk"
		res.Data = produkInput
		models.AddProduct(produkInput)
	}

	configs.ResultOk(w, res)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request)  {
	input, _ := ioutil.ReadAll(r.Body)
	var res structs.Result
	var produkUpdate structs.UpdateProduct

	json.Unmarshal(input, &produkUpdate)
	hasil := models.UpdateProduct(produkUpdate)

	res.Status = "true"
	res.Message = "Berhasil update produk"
	res.Data = hasil

	configs.ResultOk(w, res)
}

func DeleteProduct(w http.ResponseWriter, r *http.Request)  {
	input, _ := ioutil.ReadAll(r.Body)
	var res structs.ResultDua
	var produk structs.UpdateProduct

	json.Unmarshal(input, &produk)
	var kosong interface{}

	if produk.Id == 0 {
		res.Status = false
		res.Message = "Gagal hapus data, Id produk tidak boleh kosong"
		res.Data = kosong
	} else {
		models.DeleteProduct(produk.Id)

		res.Status = true
		res.Message = "Berhasil delete produk"
		res.Data = produk.Id
	}

	configs.Result(w, res)
}
package controllers

import (
	"encoding/json"
	"fmt"
	"golang-produk/configs"
	"golang-produk/models"
	"golang-produk/structs"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

func UploadImageProduct(w http.ResponseWriter, r *http.Request)  {
	var res structs.Result
	var dataUpload structs.Images
	// upload file 10 mb
	r.ParseMultipartForm(10 << 20)
	file, handler, err := r.FormFile("image")

	if err != nil {
        fmt.Println("Error Retrieving the File")
        fmt.Println(err)
        return
    }
    defer file.Close()
    fmt.Printf("Uploaded File: %+v\n", handler.Filename)
    fmt.Printf("File Size: %+v\n", handler.Size)
    fmt.Printf("MIME Header: %+v\n", handler.Header)

	a, err := strconv.ParseInt(r.FormValue("id_produk")[0:], 10, 64);
	if err != nil {
		fmt.Fprintf(w, "error", err.Error())
	}

	dataUpload.Name = handler.Filename
	dataUpload.IdProduct = a

	// tempFile, err := ioutil.TempFile("assets", "upload-*.png")
	// tempFile, err := ioutil.TempFile("assets", handler.Filename)
    // if err != nil {
    //     fmt.Println(err)
    // }
    // defer tempFile.Close()
	fileLocation := filepath.Join("assets", handler.Filename)
	dst, err := os.Create(fileLocation)
	defer dst.Close()

	// fileBytes, err := ioutil.ReadAll(file)
    // if err != nil {
    //     fmt.Println(err)
    // }
	// tempFile.Write(fileBytes)
	models.AddImageProduct(dataUpload)
	res.Status = "true"
	res.Message = "Berhasil upload gambar"
	res.Data = dataUpload
	configs.ResultOk(w, res)
}

func EditImageProduct(w http.ResponseWriter, r *http.Request)  {
	var res structs.ResultDua
	var updateImageProduct structs.UpdateImages

	id, err := strconv.ParseInt(r.FormValue("id")[0:], 10, 64)
	if err != nil {
		fmt.Fprintf(w, "error", err.Error())
		res.Status = false
		res.Message = err.Error()
		configs.Result(w, res)
		return
	}

	cekData := models.GetImageProduct(id)
	if cekData.Id == 0 {
		res.Status = false
		res.Message = "data gambar tidak ditemukan"
		configs.Result(w, res)
		return
	}

	// hapus gambar sebelum nya
	gambarSebelum := "assets/" + cekData.Name
	os.Remove(gambarSebelum)

	// upload file 10 mb
	r.ParseMultipartForm(10 << 20)
	file, handler, err := r.FormFile("image")

	if err != nil {
        fmt.Println("Error Retrieving the File")
        fmt.Println(err)
        return
    }
    defer file.Close()
	// amgil nama file, ukuran dan mime dari gambar
    fmt.Printf("Uploaded File: %+v\n", handler.Filename)
    fmt.Printf("File Size: %+v\n", handler.Size)
    fmt.Printf("MIME Header: %+v\n", handler.Header)

	// upload ke folder assets
	fileLocation := filepath.Join("assets", handler.Filename)
	// dst, err := os.OpenFile(fileLocation, os.O_WRONLY|os.O_CREATE, 0666)
	dst, err := os.Create(fileLocation)
	defer dst.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	updateImageProduct.Id = id
	updateImageProduct.Name = handler.Filename
	models.UpdateImageProduct(updateImageProduct)

	cekHasil := models.GetImageProduct(id)
	res.Status = true
	res.Message = "berhasil update gambar produk"
	res.Data = cekHasil
	configs.Result(w, res)
}

func DeleteImageProduct(w http.ResponseWriter, r *http.Request)  {
	var res structs.ResultDua
	var delete structs.DeleteImageProduct

	// id, err := strconv.ParseInt(r.FormValue("id")[0:], 10, 64)
	input, _ := ioutil.ReadAll(r.Body) 
	json.Unmarshal(input, &delete)
	
	if delete.Id == 0 {
		res.Status = false
		res.Message = "id tidak boleh kosong "
		configs.Result(w, res)
		return
	}

	cek := models.GetImageProduct(delete.Id)

	if cek.Id == 0 {
		res.Status = false
		res.Message = "data tidak ditemukan, silahkan cek id anda"
		configs.Result(w, res)
		return
	}
	// delete gambarnya
	os.Remove("assets/" + cek.Name)
	// hapus datanya
	models.DeleteImageProduct(delete.Id)

	res.Status = true
	res.Message = "berhasil hapus gambar"
	configs.Result(w, res)
}
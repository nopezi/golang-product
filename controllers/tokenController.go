package controllers

import (
	"fmt"
	"golang-produk/configs"
	"golang-produk/structs"
	"net/http"
)

func TokenPage(w http.ResponseWriter, r *http.Request)  {
	validToken, err := configs.GenerateJWT()
	var res structs.Result
	var tokenResult structs.ResultToken
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	// // // fmt.Fprintf(w, validToken)
	tokenResult.Token = validToken
	res.Status = "true"
	res.Message = "berhasil mendapatkan token"
	res.Data = tokenResult
	configs.ResultOk(w, res)
}
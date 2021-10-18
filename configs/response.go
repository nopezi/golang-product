package configs

import (
	"encoding/json"
	"golang-produk/structs"
	"net/http"
)

func ResultOk(w http.ResponseWriter, res structs.Result)  {
	hasil, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(hasil)
}

func Result(w http.ResponseWriter, res structs.ResultDua)  {
	hasil, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(hasil)
}
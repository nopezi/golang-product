package routers

import (
	"encoding/json"
	"net/http"
)

func ResultOk(w http.ResponseWriter, res Result)  {
	hasil, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(hasil)
}
package controllers

import (
	"fmt"
	"net/http"
	"time"
)

func CobaPage(w http.ResponseWriter, r *http.Request)  {
	// layout := "YYYY-MM-dd"
	// input := "2017-08-31"
	// layout := "2006-01-02"
	// waktu, _ := time.Parse(layout)
	currentTime := time.Now()
	waktu := currentTime.Format("2006-01-02 15:04:05")
	fmt.Fprintf(w, "Super Secret Information ", waktu)
}
package controller

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"library-ws/config"
	"library-ws/dao"
	"library-ws/fungsi"
	"library-ws/model"
	"net/http"
	"time"
)

func MainAPI(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	Data := make(map[string]interface{})
	Data["message"] = "helloWorld"
	Data["status"] = true
	Data["data"] = nil
	json.NewEncoder(w).Encode(Data)
}

func AddCategory(w http.ResponseWriter, r *http.Request) {
	db := config.ConfigSql()
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var kategori model.Kategori
	kategori.Id = fungsi.ToCRC32(time.Now().String())
	kategori.Kategori = params["categoryName"]
	json.NewEncoder(w).Encode(dao.AddCategory(kategori, db))
	defer db.Close()
}

func DeleteCategory(w http.ResponseWriter, r *http.Request) {
	db := config.ConfigSql()
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	json.NewEncoder(w).Encode(dao.DeleteCategory(params["categoryName"], db))
	defer db.Close()
}

func ListCategory(w http.ResponseWriter, r *http.Request) {
	db := config.ConfigSql()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(dao.ListCategory(db))
	defer db.Close()
}

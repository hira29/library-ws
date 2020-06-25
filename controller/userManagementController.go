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

func UMCreate(w http.ResponseWriter, r *http.Request) {
	db := config.ConfigSql()
	w.Header().Set("Content-Type", "application/json")
	var dataUM model.Data_UserManagement
	now := time.Now()
	now.Format(time.RFC3339)
	dataUM.Tanggal_ditambahkan = now.Format("2006-01-02 15:04:05")
	json.NewDecoder(r.Body).Decode(&dataUM)
	dataUM.Password = fungsi.ToMd5(dataUM.Password)
	json.NewEncoder(w).Encode(dao.UMCreate(dataUM, db))
	defer db.Close()
}

func UMList(w http.ResponseWriter, r *http.Request) {
	db := config.ConfigSql()
	var page model.Paging
	w.Header().Set("Content-Type", "application/json")
	json.NewDecoder(r.Body).Decode(&page)
	json.NewEncoder(w).Encode(dao.UMList(page, db))
	defer db.Close()
}

func UMUpdate(w http.ResponseWriter, r *http.Request) {
	db := config.ConfigSql()
	var dataUM model.Data_UserManagement
	w.Header().Set("Content-Type", "application/json")
	json.NewDecoder(r.Body).Decode(&dataUM)
	json.NewEncoder(w).Encode(dao.UMUpdate(dataUM, db))
	defer db.Close()
}

func UMDelete(w http.ResponseWriter, r *http.Request) {
	db := config.ConfigSql()
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	json.NewEncoder(w).Encode(dao.UMDelete(params["id"], db))
	defer db.Close()
}

func UMLogin(w http.ResponseWriter, r *http.Request) {
	db := config.ConfigSql()
	var Login model.UMLogin
	w.Header().Set("Content-Type", "application/json")
	json.NewDecoder(r.Body).Decode(&Login)
	Login.Password = fungsi.ToMd5(Login.Password)
	json.NewEncoder(w).Encode(dao.UMLogin(Login.Username, Login.Password, db))
	defer db.Close()
}


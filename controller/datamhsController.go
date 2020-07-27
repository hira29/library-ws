package controller

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"library-ws/config"
	"library-ws/dao"
	"library-ws/fungsi"
	"library-ws/model"
	"net/http"
	"path"
	"time"
)

func CreateMhs(w http.ResponseWriter, r *http.Request) {
	db := config.ConfigSql()
	w.Header().Set("Content-Type", "application/json")
	var dataMhs model.Data_mahasiswa
	now := time.Now()
	now.Format(time.RFC3339)
	//dataMhs.Tanggal_ditambahkan = now.Format("2006-01-02 15:04:05")
	json.NewDecoder(r.Body).Decode(&dataMhs)
	dataMhs.Password = fungsi.ToMd5(dataMhs.Password)
	dataMhs.Mhs_id = fungsi.ToCRC32(time.Now().String())
	json.NewEncoder(w).Encode(dao.CreateMhs(dataMhs, db))
	defer db.Close()
}

func ViewMhsById(w http.ResponseWriter, r *http.Request) {
	db := config.ConfigSql()
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	json.NewEncoder(w).Encode(dao.ViewMhsById(params["mhsId"], db))
	defer db.Close()
}

func ListMhs(w http.ResponseWriter, r *http.Request) {
	db := config.ConfigSql()
	var page model.Paging
	w.Header().Set("Content-Type", "application/json")
	json.NewDecoder(r.Body).Decode(&page)
	json.NewEncoder(w).Encode(dao.ListMhs(page, db))
	defer db.Close()
}

func UpdateMhsById(w http.ResponseWriter, r *http.Request) {
	db := config.ConfigSql()
	var dataMhs model.Data_mahasiswa
	w.Header().Set("Content-Type", "application/json")
	json.NewDecoder(r.Body).Decode(&dataMhs)
	json.NewEncoder(w).Encode(dao.UpdateMhsById(dataMhs, db))
	defer db.Close()
}

func DeleteMhsById(w http.ResponseWriter, r *http.Request) {
	db := config.ConfigSql()
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	json.NewEncoder(w).Encode(dao.DeleteMhsById(params["mhsId"], db))
	defer db.Close()
}

func Login(w http.ResponseWriter, r *http.Request) {
	db := config.ConfigSql()
	var Login model.Login
	w.Header().Set("Content-Type", "application/json")
	json.NewDecoder(r.Body).Decode(&Login)
	Login.Password = fungsi.ToMd5(Login.Password)
	json.NewEncoder(w).Encode(dao.Login(Login.Mhs_id, Login.Password, db))
	defer db.Close()
}

func RegisterMhs(w http.ResponseWriter, r *http.Request) {
	db := config.ConfigSql()
	var Register model.Data_mahasiswa
	w.Header().Set("Content-Type", "application/json")
	json.NewDecoder(r.Body).Decode(&Register)
	json.NewEncoder(w).Encode(dao.RegisterMhs(Register, db))
	defer db.Close()
}

func CreateListMhs(w http.ResponseWriter, r *http.Request) {
	db := config.ConfigSql()
	var page model.Paging
	w.Header().Set("Content-Type", "application/json")
	json.NewDecoder(r.Body).Decode(&page)
	json.NewEncoder(w).Encode(dao.CreateListMhs(page, db))
	defer db.Close()
}

func DownloadListMhs(w http.ResponseWriter, r *http.Request) {
	fp := path.Join("application", "../DataMhs.xlsx")
	http.ServeFile(w, r, fp)
}

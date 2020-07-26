package controller

import (
	"encoding/json"
	"library-ws/config"
	"library-ws/dao"
	"library-ws/fungsi"
	"library-ws/model"
	"net/http"
	"time"
)

func SetPengembalian(w http.ResponseWriter, r *http.Request) {
	db := config.ConfigSql()
	w.Header().Set("Content-Type", "application/json")
	var dataKembali model.SetPengembalian
	json.NewDecoder(r.Body).Decode(&dataKembali)
	json.NewEncoder(w).Encode(dao.SetPengembalian(dataKembali, db))
	defer db.Close()
}

func GetPengembalian(w http.ResponseWriter, r *http.Request) {
	db := config.ConfigSql()
	w.Header().Set("Content-Type", "application/json")
	var dataKembali model.Data_peminjaman
	json.NewDecoder(r.Body).Decode(&dataKembali)
	json.NewEncoder(w).Encode(dao.GetPengembalian(dataKembali, db))
	defer db.Close()
}

func SetAdminPengembalian(w http.ResponseWriter, r *http.Request) {
	db := config.ConfigSql()
	w.Header().Set("Content-Type", "application/json")
	var dataKembali model.Riwayat_peminjaman
	dataKembali.Id_peminjaman = fungsi.ToCRC32(time.Now().String())
	json.NewDecoder(r.Body).Decode(&dataKembali)
	json.NewEncoder(w).Encode(dao.SetAdminPengembalian(dataKembali, db))
	defer db.Close()
}

func ListPengembalian(w http.ResponseWriter, r *http.Request) {
	db := config.ConfigSql()
	var page model.PeminjamanPaging
	w.Header().Set("Content-Type", "application/json")
	json.NewDecoder(r.Body).Decode(&page)
	json.NewEncoder(w).Encode(dao.ListPengembalian(page, db))
	defer db.Close()
}

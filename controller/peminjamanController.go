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

func Pinjam(w http.ResponseWriter, r *http.Request) {
	db := config.ConfigSql()
	w.Header().Set("Content-Type", "application/json")
	var dataPinjam model.Data_peminjaman
	dataPinjam.Id_peminjaman = fungsi.ToCRC32(time.Now().String())
	json.NewDecoder(r.Body).Decode(&dataPinjam)
	json.NewEncoder(w).Encode(dao.Pinjam(dataPinjam, db))
	defer db.Close()
}

func Kembali(w http.ResponseWriter, r *http.Request) {
	db := config.ConfigSql()
	w.Header().Set("Content-Type", "application/json")
	var dataKembali model.Riwayat_peminjaman
	dataKembali.Id_peminjaman = fungsi.ToCRC32(time.Now().String())
	json.NewDecoder(r.Body).Decode(&dataKembali)
	json.NewEncoder(w).Encode(dao.Kembali(dataKembali, db))
	defer db.Close()
}

func Riwayat(w http.ResponseWriter, r *http.Request) {
	db := config.ConfigSql()
	var page model.PeminjamanPaging
	w.Header().Set("Content-Type", "application/json")
	json.NewDecoder(r.Body).Decode(&page)
	json.NewEncoder(w).Encode(dao.Riwayat(page, db))
	defer db.Close()
}

func Berlangsung(w http.ResponseWriter, r *http.Request) {
	db := config.ConfigSql()
	var page model.PeminjamanPaging
	w.Header().Set("Content-Type", "application/json")
	json.NewDecoder(r.Body).Decode(&page)
	json.NewEncoder(w).Encode(dao.Berlangsung(page, db))
	defer db.Close()
}

func ViewByIdPeminjaman(w http.ResponseWriter, r *http.Request) {
	db := config.ConfigSql()
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	json.NewEncoder(w).Encode(dao.ViewByIdPeminjaman(params["idPeminjaman"], db))
	defer db.Close()
}
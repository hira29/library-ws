package controller

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"library-ws/config"
	"library-ws/dao"
	"library-ws/model"
	"net/http"
	"path"
)

func CreateBuku(w http.ResponseWriter, r *http.Request) {
	db := config.ConfigSql()
	w.Header().Set("Content-Type", "application/json")
	var dataBuku model.Data_buku
	json.NewDecoder(r.Body).Decode(&dataBuku)
	json.NewEncoder(w).Encode(dao.CreateBuku(dataBuku, db))
	defer db.Close()
}

func ViewBukuById(w http.ResponseWriter, r *http.Request) {
	db := config.ConfigSql()
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	json.NewEncoder(w).Encode(dao.ViewBukuById(params["bukuId"], db))
	defer db.Close()
}

func ListBuku(w http.ResponseWriter, r *http.Request) {
	db := config.ConfigSql()
	var page model.Paging
	w.Header().Set("Content-Type", "application/json")
	json.NewDecoder(r.Body).Decode(&page)
	json.NewEncoder(w).Encode(dao.ListBuku(page, db))
	defer db.Close()
}

func UpdateBukuById(w http.ResponseWriter, r *http.Request) {
	db := config.ConfigSql()
	var dataBuku model.Data_buku
	w.Header().Set("Content-Type", "application/json")
	json.NewDecoder(r.Body).Decode(&dataBuku)
	json.NewEncoder(w).Encode(dao.UpdateBukuById(dataBuku, db))
	defer db.Close()
}

func UpdateStokById(w http.ResponseWriter, r *http.Request) {
	db := config.ConfigSql()
	var stok model.UpdateStok
	w.Header().Set("Content-Type", "application/json")
	json.NewDecoder(r.Body).Decode(&stok)
	json.NewEncoder(w).Encode(dao.UpdateStokById(stok, db))
	defer db.Close()
}

func DeleteBukuById(w http.ResponseWriter, r *http.Request) {
	db := config.ConfigSql()
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	json.NewEncoder(w).Encode(dao.DeleteBukuById(params["bukuId"], db))
	defer db.Close()
}

func CreateListBuku(w http.ResponseWriter, r *http.Request) {
	db := config.ConfigSql()
	var page model.Paging
	w.Header().Set("Content-Type", "application/json")
	json.NewDecoder(r.Body).Decode(&page)
	json.NewEncoder(w).Encode(dao.CreateListBuku(page, db))
	defer db.Close()
}

func DownloadListBuku(w http.ResponseWriter, r *http.Request) {
	fp := path.Join("application", "../DataBooks.xlsx")
	http.ServeFile(w, r, fp)
}

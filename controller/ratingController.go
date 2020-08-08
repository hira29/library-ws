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

func CreateRating(w http.ResponseWriter, r *http.Request) {
	db := config.ConfigSql()
	w.Header().Set("Content-Type", "application/json")
	var dataRating model.Data_Rating
	json.NewDecoder(r.Body).Decode(&dataRating)
	json.NewEncoder(w).Encode(dao.CreateRating(dataRating, db))
	defer db.Close()
}

func ListRating(w http.ResponseWriter, r *http.Request) {
	db := config.ConfigSql()
	var page model.Rating_Paging
	w.Header().Set("Content-Type", "application/json")
	json.NewDecoder(r.Body).Decode(&page)
	json.NewEncoder(w).Encode(dao.ListRating(page, db))
	defer db.Close()
}

func DeleteRating(w http.ResponseWriter, r *http.Request) {
	db := config.ConfigSql()
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	json.NewEncoder(w).Encode(dao.DeleteRating(params["ratingId"], db))
	defer db.Close()
}

func ViewByRatingId(w http.ResponseWriter, r *http.Request) {
	db := config.ConfigSql()
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	json.NewEncoder(w).Encode(dao.ViewByRatingId(params["ratingId"], db))
	defer db.Close()
}

func CreateListRating(w http.ResponseWriter, r *http.Request) {
	db := config.ConfigSql()
	var page model.Rating_Paging
	w.Header().Set("Content-Type", "application/json")
	json.NewDecoder(r.Body).Decode(&page)
	json.NewEncoder(w).Encode(dao.CreateListRating(page, db))
	defer db.Close()
}

func DownloadListRating(w http.ResponseWriter, r *http.Request) {
	fp := path.Join("application", "../DataBooksRating.xlsx")
	http.ServeFile(w, r, fp)
}

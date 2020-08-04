package controller

import (
	"encoding/json"
	"library-ws/config"
	"library-ws/dao"
	"library-ws/model"
	"net/http"
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

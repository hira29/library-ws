package controller

import (
	"encoding/json"
	"library-ws/config"
	"library-ws/dao"
	"net/http"
)

func BookSummary(w http.ResponseWriter, r *http.Request) {
	db := config.ConfigSql()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(dao.SummaryBook(db))
	defer db.Close()
}

func LoanSummary(w http.ResponseWriter, r *http.Request) {
	db := config.ConfigSql()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(dao.LoanSummary(db))
	defer db.Close()
}

func ReturnSummary(w http.ResponseWriter, r *http.Request) {
	db := config.ConfigSql()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(dao.ReturnSummary(db))
	defer db.Close()
}
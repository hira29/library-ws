package model

import "time"

type Data_Rating struct {
	Id_Rating     string    `json:"id_rating"`
	Id_Peminjaman string    `json:"id_peminjaman"`
	Id_Buku       string    `json:"id_buku"`
	Id_Mhs        string    `json:"id_mhs"`
	Nama          string    `json:"nama"`
	Rating        float64   `json:"rating"`
	Komentar      string    `json:"komentar"`
	Tanggal       time.Time `json:"tanggal"`
}

type Rating_Paging struct {
	Id_buku string `json:"id_buku"`
	Page    int    `json:"page"`
	Size    int    `json:"size"`
}

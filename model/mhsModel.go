package model

import "time"

type Data_mahasiswa struct {
	Mhs_id        string    `json:"mhs_id"`
	Nama          string    `json:"nama"`
	Tempat_lahir  string    `json:"tempat_lahir"`
	Tanggal_lahir time.Time `json:"tanggal_lahir"`
	Email         string    `json:"email"`
	Password      string    `json:"password"`
	Active        int64     `json:"active"`
}

type Login struct {
	Mhs_id   string `json:"mhs_id"`
	Password string `json:"password"`
}

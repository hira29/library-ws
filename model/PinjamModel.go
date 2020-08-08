package model

import "time"

type Riwayat_peminjaman struct {
	Id_peminjaman        string    `json:"id_peminjaman"`
	Id_buku              string    `json:"id_buku"`
	Judul_buku           string    `json:"judul_buku"`
	Id_mhs               string    `json:"id_mhs"`
	Id_rating            string    `json:"id_rating"`
	Tanggal_peminjaman   time.Time `json:"tanggal_peminjaman"`
	Tanggal_pengembalian time.Time `json:"tanggal_pengembalian"`
	Tanggal_kembali      time.Time `json:"tanggal_kembali"`
}

type Data_peminjaman struct {
	Id_peminjaman        string    `json:"id_peminjaman"`
	Id_buku              string    `json:"id_buku"`
	Judul_buku           string    `json:"judul_buku"`
	Id_mhs               string    `json:"id_mhs"`
	Tanggal_peminjaman   time.Time `json:"tanggal_peminjaman"`
	Tanggal_pengembalian time.Time `json:"tanggal_pengembalian"`
	Tanggal_kembali      time.Time `json:"tanggal_kembali"`
	Pengembalian         int       `json:"pengembalian"`
}

type PeminjamanPaging struct {
	Id     string `json:"id"`
	Search string `json:"search"`
	Page   int    `json:"page"`
	Size   int    `json:"size"`
}

type SetPengembalian struct {
	Id_peminjaman string `json:"id_peminjaman"`
}

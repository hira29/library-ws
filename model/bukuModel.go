package model

type Data_buku struct {
	Buku_id 			string 	`json:"buku_id"`
	Judul				string	`json:"judul"`
	Penulis 			string	`json:"penulis"`
	Penerbit			string	`json:"penerbit"`
	Jumlah_Eksemplar	string	`json:"jumlah_eksemplar"`
	Kategori			string	`json:"kategori"`
	Letak_Buku			string	`json:"letak_buku"`
	Gambar				string	`json:"gambar"`
	Tanggal_ditambahkan	string	`json:"tanggal_ditambahkan"`
	Stok 				int64	`json:"stok"`
}

type UpdateStok struct {
	Buku_id 	string `json:"buku_id"`
	Stok 		int64  `json:"stok"`
}

type Paging struct {
	Search  string	`json:"search"`
	Category string `json:"category"`
	Page 	int		`json:"page"`
	Size 	int		`json:"size"`
}
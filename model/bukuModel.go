package model

type Data_buku struct {
	Buku_id             string `json:"buku_id"`
	Judul               string `json:"judul"`        //
	Edisi               string `json:"edisi"`        //
	Pengarang           string `json:"pengarang"`    //
	Kota_terbit         string `json:"kota_terbit"`  //
	Penerbit            string `json:"penerbit"`     //
	Tahun_Terbit        string `json:"tahun_terbit"` //
	Isbn                string `json:"isbn"`
	Klasifikasi         string `json:"klasifikasi"`
	Kategori            string `json:"kategori"`
	Umum_res            string `json:"umum_res"`
	Bahasa              string `json:"bahasa"`
	Deskripsi           string `json:"deskripsi"`
	Lokasi              string `json:"lokasi"`
	Gambar              string `json:"gambar"`
	Tanggal_ditambahkan string `json:"tanggal_ditambahkan"`
	Stok                int64  `json:"stok"`
}

type UpdateStok struct {
	Buku_id string `json:"buku_id"`
	Stok    int64  `json:"stok"`
}

type Paging struct {
	Search   string `json:"search"`
	Category string `json:"category"`
	Page     int    `json:"page"`
	Size     int    `json:"size"`
}

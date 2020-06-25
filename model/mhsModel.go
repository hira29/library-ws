package model

type Data_mahasiswa struct{
	Mhs_id 			string `json:"mhs_id"`
	Password 		string `json:"password"`
	Nama			string	`json:"nama"`
	Jurusan			string 	`json:"jurusan"`
	Tahun_masuk 	string	`json:"tahun_masuk"`
	Tanggal_ditambahkan	string	`json:"tanggal_ditambahkan"`
}

type Login struct{
	Mhs_id			string `json:"mhs_id"`
	Password		string `json:"password"`
}

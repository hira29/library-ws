package model

type Data_UserManagement struct {
	Id      		string `json:"id"`
	Username 		string `json:"username"`
	Password 		string `json:"password"`
	Role			string `json:"role"`
	Nama			string	`json:"nama"`
	Email			string 	`json:"email"`
	Phone		 	string	`json:"Phone"`
	Tanggal_ditambahkan	string	`json:"tanggal_ditambahkan"`
}

type UMLogin struct{
	Username		string `json:"username"`
	Password		string `json:"password"`
}
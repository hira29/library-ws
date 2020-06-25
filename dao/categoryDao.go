package dao

import (
	"github.com/jinzhu/gorm"
	"library-ws/model"
)

func AddCategory(kategori model.Kategori,db *gorm.DB) model.Return {
	var message string
	var status bool
	db.Create(&kategori)
	message = "Data Berhasil Ditambahkan"
	status 	= true
	return model.Return{Status: status, Data: kategori, Message: message}
}

func DeleteCategory(id string, db *gorm.DB) model.Return{
	var message string
	var status bool

	data := db.Where("kategori = ?", id).Delete(&model.Kategori{})
	if data.Error != nil {
		message = "Data Gagal DiHapus"
		status 	= false
		data.Value = nil
	} else {
		message = "Data Berhasil DiHapus"
		status = true
		data.Value = data.RowsAffected
	}
	return model.Return{Status: status, Data: data.Value, Message: message}
}

func ListCategory(db *gorm.DB) model.Return{
	var message string
	var status bool
	var Kategori []model.Kategori

	data := db.Find(&Kategori)
	message = "Berhasil"
	status = true

	return model.Return{Status: status, Data: data.Value, Message: message}
}

package dao

import (
	"github.com/jinzhu/gorm"
	"library-ws/model"
	"time"
)

func SummaryBook(db *gorm.DB) model.Return {
	var message string
	var status bool
	var result int64

	db.Table("data_bukus").Count(&result)

	message = "Data Berhasil Ditemukan"
	status 	= true
	return model.Return{Status: status, Data: result, Message: message}
}

func LoanSummary(db *gorm.DB) model.Return {
	var message string
	var status bool
	var result int64


	db.Model(&model.Data_peminjaman{}).Count(&result)

	message = "Data Berhasil Ditemukan"
	status 	= true
	return model.Return{Status: status, Data: result, Message: message}
}

func ReturnSummary(db *gorm.DB) model.Return {
	var message string
	var status bool
	var result int64
	var last time.Time

	now := time.Now().Local()
	last = now.AddDate(0, 0, -7)


	db.Model(&model.Riwayat_peminjaman{}).Where("tanggal_kembali BETWEEN ? AND ?", last, now).Count(&result)

	message = "Data Berhasil Ditemukan"
	status 	= true
	return model.Return{Status: status, Data: result, Message: message}
}
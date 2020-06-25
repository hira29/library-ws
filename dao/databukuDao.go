package dao

import (
	"github.com/biezhi/gorm-paginator/pagination"
	"github.com/jinzhu/gorm"
	"library-ws/fungsi"
	"library-ws/model"
	"time"
)

func CreateBuku(buku model.Data_buku, db *gorm.DB) model.Return {
	buku.Buku_id = fungsi.ToCRC32(time.Now().String())
	now := time.Now()
	now.Format(time.RFC3339)
	buku.Tanggal_ditambahkan = now.Format("2006-01-02 15:04:05")
	var message string
	var status bool
	check := db.NewRecord(buku)
	if check == true {
		db.Create(&buku)
		message = "Data Berhasil Ditambahkan"
		status 	= true
	} else {
		message = "Data Gagal Ditambahkan"
		status  = false
	}

	return model.Return{Status: status, Data: buku, Message: message}
}

func ViewBukuById(id string, db *gorm.DB) model.Return {
	var message string
	var status bool

	data := db.Where("buku_id = ?", id).First(&model.Data_buku{})
	if data.Error != nil {
		message = "Data Gagal Ditemukan"
		status 	= false
		data.Value = nil
	} else {
		message = "Data Berhasil Ditemukan"
		status = true
	}

	return model.Return{Status: status, Data: data.Value, Message: message}
}

func ListBuku(page model.Paging,db *gorm.DB) model.Return {
	var message string
	var status bool
	var data_bukus []model.Data_buku

	var DataBase *gorm.DB
	DataBase = db

	data := db.Find(&model.Data_buku{})
	if data.Error != nil {
		message = "Data Gagal Ditemukan"
		status 	= false
		data.Value = nil
	} else {
		message = "Data Berhasil Ditemukan"
		status = true
	}


	if page.Search == "" {
		if page.Category == "All"{
			DataBase = db
		} else if page.Category == "" {
			DataBase = db
		} else {
			DataBase = db.Where("kategori LIKE ?", "%"+page.Category+"%")
		}
	} else {
		if page.Category == "All"{
			DataBase = db.Where("judul LIKE ?", "%"+page.Search+"%")
		} else if page.Category == "" {
			DataBase = db.Where("judul LIKE ?", "%"+page.Search+"%")
		} else {
			DataBase = db.Where("judul LIKE ? AND kategori LIKE ?", "%"+page.Search+"%", "%"+page.Category+"%")
		}
	}

	paginator := pagination.Paging(&pagination.Param{
		DB:      DataBase,
		Page:    page.Page,
		Limit:   page.Size,
		ShowSQL: true,
	}, &data_bukus)

	return model.Return{Status: status, Data: paginator, Message: message}
}

func UpdateBukuById(data_buku model.Data_buku, db *gorm.DB) model.Return {
	var message string
	var status bool

	data := db.Model(&model.Data_buku{}).Where("buku_id = ?", data_buku.Buku_id).Update(data_buku)
	if data.Error != nil {
		message = "Data Gagal DiUpdate"
		status 	= false
		data.Value = nil
	} else {
		message = "Data Berhasil DiUpdate"
		status = true
		data.Value = data.RowsAffected
	}
	return model.Return{Status: status, Data: data.Value, Message: message}
}

func UpdateStokById(stok model.UpdateStok, db *gorm.DB) model.Return{
	var message string
	var status bool

	data := db.Model(&model.Data_buku{}).Where("buku_id = ?", stok.Buku_id).Update(map[string]interface{}{"stok":stok.Stok})
	if data.Error != nil {
		message = "Data Gagal DiUpdate"
		status 	= false
		data.Value = nil
	} else {
		message = "Data Berhasil DiUpdate"
		status = true
		data.Value = data.RowsAffected
	}
	return model.Return{Status: status, Data: data.Value, Message: message}
}

func DeleteBukuById(id string, db *gorm.DB) model.Return{
	var message string
	var status bool

	data := db.Where("buku_id = ?", id).Delete(&model.Data_buku{})
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
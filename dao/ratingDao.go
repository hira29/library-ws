package dao

import (
	"encoding/json"
	"github.com/biezhi/gorm-paginator/pagination"
	"github.com/jinzhu/gorm"
	"library-ws/fungsi"
	"library-ws/model"
	"math"
	"time"
)

func CreateRating(Data_Rating model.Data_Rating, db *gorm.DB) model.Return {
	var message string
	var status bool
	var InterfaceBuku model.Data_buku
	var InterfaceMhs model.Data_mahasiswa
	var result float64
	var curRating, nextRating, getRating float64

	now := time.Now()
	now.Format(time.RFC3339)
	Data_Rating.Tanggal = now
	Mhs := db.Where("mhs_id = ?", Data_Rating.Id_Mhs).First(&model.Data_mahasiswa{})
	dataMhs, _ := json.Marshal(Mhs.Value)
	_ = json.Unmarshal(dataMhs, &InterfaceMhs)
	Data_Rating.Nama = InterfaceMhs.Nama
	Data_Rating.Id_Rating = fungsi.ToCRC32(time.Now().String())

	Buku := db.Where("buku_id = ?", Data_Rating.Id_Buku).First(&model.Data_buku{})
	if Buku.Error != nil {
		message = "Data Gagal Ditemukan"
		status = false
		Buku.Value = nil
	} else {
		dataBuku, _ := json.Marshal(Buku.Value)
		_ = json.Unmarshal(dataBuku, &InterfaceBuku)
		db.Model(&model.Data_Rating{}).Where("id_buku = ?", Data_Rating.Id_Buku).Count(&result)

		curRating = InterfaceBuku.Rating * result
		getRating = Data_Rating.Rating
		nextRating = (getRating + curRating) / (result + 1)
		nextRating = math.Ceil(nextRating*10) / 10

		check := db.NewRecord(Data_Rating)
		if check == true {
			_ = db.Model(&model.Data_buku{}).Where("buku_id = ?", Data_Rating.Id_Buku).Update(map[string]interface{}{"rating": nextRating})
			db.Create(&Data_Rating)
			message = "Data Berhasil Ditambahkan"
			status = true
		} else {
			message = "Data Gagal Ditambahkan"
			status = false
		}

	}

	return model.Return{Status: status, Data: Data_Rating, Message: message}
}

func ListRating(page model.Rating_Paging, db *gorm.DB) model.Return {
	var message string
	var status bool
	var data_rating []model.Data_Rating

	var DataBase *gorm.DB
	DataBase = db

	if page.Id_buku == "" {
		DataBase = db
	} else {
		DataBase = db.Where("id_buku = ?", page.Id_buku)
	}

	paginator := pagination.Paging(&pagination.Param{
		DB:      DataBase,
		Page:    page.Page,
		Limit:   page.Size,
		ShowSQL: true,
	}, &data_rating)

	if paginator.TotalRecord == 0 {
		message = "Tidak Ada Rating"
		status = false
	} else {
		message = "Success"
		status = true
	}

	return model.Return{Status: status, Data: paginator, Message: message}
}

func DeleteRating(id string, db *gorm.DB) model.Return {
	var message string
	var status bool
	var InterfaceRating model.Data_Rating
	var InterfaceBuku model.Data_buku
	var result float64
	var curRating, nextRating, getRating float64

	data := db.Where("id_rating = ?", id).First(&model.Data_Rating{})
	dataRating, _ := json.Marshal(data.Value)
	_ = json.Unmarshal(dataRating, &InterfaceRating)

	Buku := db.Where("buku_id = ?", InterfaceRating.Id_Buku).First(&model.Data_buku{})
	if Buku.Error != nil {
		message = "Data Gagal Ditemukan"
		status = false
		Buku.Value = nil
	} else {
		dataBuku, _ := json.Marshal(Buku.Value)
		_ = json.Unmarshal(dataBuku, &InterfaceBuku)
		db.Model(&model.Data_Rating{}).Where("id_buku = ?", InterfaceRating.Id_Buku).Count(&result)

		curRating = InterfaceBuku.Rating * result
		getRating = InterfaceRating.Rating
		nextRating = (curRating - getRating) / (result - 1)
		nextRating = math.Ceil(nextRating*10) / 10
	}

	Delete := db.Where("id_rating = ?", id).Delete(&model.Data_Rating{})
	if Delete.Error != nil {
		message = "Data Gagal DiHapus"
		status = false
		data.Value = nil
	} else {
		message = "Data Berhasil DiHapus"
		status = true
		Delete.Value = data.RowsAffected
		_ = db.Model(&model.Data_buku{}).Where("buku_id = ?", InterfaceRating.Id_Buku).Update(map[string]interface{}{"rating": nextRating})
	}
	return model.Return{Status: status, Data: Delete.Value, Message: message}
}

func CreateListRating(page model.Rating_Paging, db *gorm.DB) model.Return {
	var message string
	var status bool
	var data_rating []model.Data_Rating

	var DataBase *gorm.DB
	DataBase = db

	if page.Id_buku == "" {
		DataBase = db
	} else {
		DataBase = db.Where("id_buku = ?", page.Id_buku)
	}

	paginator := pagination.Paging(&pagination.Param{
		DB:      DataBase,
		Page:    page.Page,
		Limit:   page.Size,
		ShowSQL: true,
	}, &data_rating)

	if paginator.TotalRecord == 0 {
		message = "Tidak Ada Rating"
		status = false
	} else {
		message = "Success"
		status = true

		var Records []model.Data_Rating
		getRecords, _ := json.Marshal(paginator.Records)
		_ = json.Unmarshal(getRecords, &Records)

		data := db.Where("buku_id = ?", page.Id_buku).First(&model.Data_buku{})
		var RecordBook model.Data_buku
		getBooksRecords, _ := json.Marshal(data.Value)
		_ = json.Unmarshal(getBooksRecords, &RecordBook)

		fungsi.Excelsize_Rating(RecordBook, Records)
	}

	return model.Return{Status: status, Data: paginator, Message: message}
}

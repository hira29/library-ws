package dao

import (
	"encoding/json"
	"github.com/biezhi/gorm-paginator/pagination"
	"github.com/jinzhu/gorm"
	"library-ws/model"
)

func CreateMhs(mhs model.Data_mahasiswa, db *gorm.DB) model.Return {
	var message string
	var status bool
	check := db.NewRecord(mhs)
	if check == true {
		data := db.Create(&mhs)
		if data.Error != nil {
			message = "Duplikasi NRP, Data Gagal Ditambahkan"
			status  = false
		} else {
			message = "Data Berhasil Ditambahkan"
			status 	= true
		}
		mhs.Password = "AuthGuard Protected!"
	} else {
		message = "Data Gagal Ditambahkan"
		status  = false
	}

	return model.Return{Status: status, Data: mhs, Message: message}
}

func ViewMhsById(id string, db *gorm.DB) model.Return {
	var message string
	var status bool
	var mhsInterface model.Data_mahasiswa

	data := db.Where("mhs_id = ?", id).First(&model.Data_mahasiswa{})
	if data.Error != nil {
		message = "Data Gagal Ditemukan"
		status 	= false
		data.Value = nil
	} else {
		message = "Data Berhasil Ditemukan"
		status = true
		setCord, _ := json.Marshal(data.Value)
		_ = json.Unmarshal(setCord, &mhsInterface)
		mhsInterface.Password = "AuthGuard Protected!"
	}

	return model.Return{Status: status, Data: data.Value, Message: message}
}

func ListMhs(page model.Paging,db *gorm.DB) model.Return {
	var message string
	var status bool
	var data_mahasiswas []model.Data_mahasiswa
	var Records []model.Data_mahasiswa
	var DataBase *gorm.DB

	data := db.Find(&model.Data_mahasiswa{})
	if data.Error != nil {
		message = "Data Gagal Ditemukan"
		status 	= false
		data.Value = nil
	} else {
		message = "Data Berhasil Ditemukan"
		status = true
	}

	if page.Search == "" {
		DataBase = db
	} else {
		DataBase = db.Where("nama LIKE ? OR mhs_id LIKE ? OR jurusan LIKE ?", "%"+page.Search+"%", "%"+page.Search+"%", "%"+page.Search+"%")
	}

	paginator := pagination.Paging(&pagination.Param{
		DB:      DataBase,
		Page:    page.Page,
		Limit:   page.Size,
		ShowSQL: true,
	}, &data_mahasiswas)

	getRecords, _ := json.Marshal(paginator.Records)
	_ = json.Unmarshal(getRecords, &Records)

	var UpdatedRecords []model.Data_mahasiswa
	for _, records := range Records {
		records.Password = "AuthGuard Protected!"
		UpdatedRecords = append(UpdatedRecords, records)
	}

	return model.Return{Status: status, Data: paginator, Message: message}
}

func UpdateMhsById(data_mhs model.Data_mahasiswa, db *gorm.DB) model.Return {
	var message string
	var status bool

	updateData := make(map[string]interface{})
	if data_mhs.Nama != "" {
		updateData["nama"] = data_mhs.Nama
	}
	if data_mhs.Jurusan != "" {
		updateData["jurusan"] = data_mhs.Jurusan
	}
	if data_mhs.Tahun_masuk != "" {
		updateData["tahun_masuk"] = data_mhs.Tahun_masuk
	}


	data := db.Model(&model.Data_mahasiswa{}).Where("mhs_id = ?", data_mhs.Mhs_id).Update(updateData)
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

func DeleteMhsById(id string, db *gorm.DB) model.Return{
	var message string
	var status bool

	data := db.Where("mhs_id = ?", id).Delete(&model.Data_mahasiswa{})
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

func Login(id string, password string, db *gorm.DB) model.Return {
	var message string
	var status bool
	var mhsInterface model.Data_mahasiswa

	data := db.Where("mhs_id = ? AND password = ?", id, password).Find(&model.Data_mahasiswa{})
	if data.Error != nil {
		message = "Username atau password salah"
		status 	= false
		data.Value = nil
	} else {
		message = "Anda berhasil Login"
		status = true
		setCord, _ := json.Marshal(data.Value)
		_ = json.Unmarshal(setCord, &mhsInterface)
		mhsInterface.Password = "AuthGuard Protected!"
	}
	return model.Return{Status: status, Data: mhsInterface, Message: message}
}

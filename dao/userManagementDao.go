package dao

import (
"encoding/json"
"github.com/biezhi/gorm-paginator/pagination"
"github.com/jinzhu/gorm"
	"library-ws/fungsi"
	"library-ws/model"
	"time"
)

func UMCreate(data_UM model.Data_UserManagement, db *gorm.DB) model.Return {
	var message string
	var status bool
	var result int64
	data_UM.Id = fungsi.ToCRC32(time.Now().String())

	db.Model(&model.Data_UserManagement{}).Where("username = ?", data_UM.Username).Count(&result)
	if result == 0 {
		data := db.Create(&data_UM)
		if data.Error != nil {
			message = "Duplikasi Id, Data Gagal Ditambahkan"
			status  = false
		} else {
			message = "Data Berhasil Ditambahkan"
			status 	= true
		}
		data_UM.Password = "AuthGuard Protected!"
	} else {
		message = "Duplikasi username, Data Gagal Ditambahkan"
		status  = false
	}


	return model.Return{Status: status, Data: data_UM, Message: message}
}

func UMList(page model.Paging,db *gorm.DB) model.Return {
	var message string
	var status bool
	var data_UM []model.Data_UserManagement
	var Records []model.Data_UserManagement
	var DataBase *gorm.DB

	data := db.Find(&model.Data_UserManagement{})
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
		DataBase = db.Where("nama LIKE ? OR username LIKE ? OR role LIKE ?", "%"+page.Search+"%", "%"+page.Search+"%", "%"+page.Search+"%")
	}

	paginator := pagination.Paging(&pagination.Param{
		DB:      DataBase,
		Page:    page.Page,
		Limit:   page.Size,
		ShowSQL: true,
	}, &data_UM)

	getRecords, _ := json.Marshal(paginator.Records)
	_ = json.Unmarshal(getRecords, &Records)

	var UpdatedRecords []model.Data_UserManagement
	for _, records := range Records {
		records.Password = "AuthGuard Protected!"
		UpdatedRecords = append(UpdatedRecords, records)
	}

	return model.Return{Status: status, Data: paginator, Message: message}
}

func UMUpdate(data_UM model.Data_UserManagement, db *gorm.DB) model.Return {
	var message string
	var status bool

	updateData := make(map[string]interface{})
	if data_UM.Nama != "" {
		updateData["nama"] = data_UM.Nama
	}
	if data_UM.Role != "" {
		updateData["role"] = data_UM.Role
	}
	if data_UM.Phone != "" {
		updateData["phone"] = data_UM.Phone
	}
	if data_UM.Username != "" {
		updateData["username"] = data_UM.Username
	}
	if data_UM.Email != "" {
		updateData["email"] = data_UM.Email
	}

	data := db.Model(&model.Data_UserManagement{}).Where("id = ?", data_UM.Id).Update(updateData)
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

func UMDelete(id string, db *gorm.DB) model.Return{
	var message string
	var status bool

	data := db.Where("id = ?", id).Delete(&model.Data_UserManagement{})
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

func UMLogin(username string, password string, db *gorm.DB) model.Return {
	var message string
	var status bool
	var UmInterface model.Data_UserManagement

	data := db.Where("username = ? AND password = ?", username, password).Find(&model.Data_UserManagement{})
	if data.Error != nil {
		message = "Username atau password salah"
		status 	= false
		data.Value = nil
	} else {
		message = "Anda berhasil Login"
		status = true
		setCord, _ := json.Marshal(data.Value)
		_ = json.Unmarshal(setCord, &UmInterface)
		UmInterface.Password = "AuthGuard Protected!"
	}
	return model.Return{Status: status, Data: UmInterface, Message: message}
}


package dao

import (
	"encoding/json"
	"fmt"
	"github.com/biezhi/gorm-paginator/pagination"
	"github.com/jinzhu/gorm"
	"library-ws/fungsi"
	"library-ws/model"
	"time"
)

func Pinjam(peminjaman model.Data_peminjaman, db *gorm.DB) model.Return {
	var message string
	var status bool
	var outputBuku model.Data_buku
	datas := make(map[string]interface{})

	mahasiswa := db.Where("mhs_id = ?", peminjaman.Id_mhs).First(&model.Data_mahasiswa{})
	if mahasiswa.Error != nil {
		message = "Data Mahasiswa Gagal Ditemukan"
		status = false
		mahasiswa.Value = nil
		datas = nil
	} else {
		buku := db.Where("buku_id = ?", peminjaman.Id_buku).First(&model.Data_buku{})
		if buku.Error != nil {
			message = "Data Buku Gagal Ditemukan"
			status = false
			buku.Value = nil
			datas = nil
			peminjaman.Judul_buku = ""
		} else {
			dataBuku, _ := json.Marshal(buku.Value)
			_ = json.Unmarshal(dataBuku, &outputBuku)
			peminjaman.Judul_buku = outputBuku.Judul

			if outputBuku.Stok == 0 {
				message = "Stok Buku Habis"
				status = false
				buku.Value = nil
				datas = nil
			} else {
				stok := model.UpdateStok{Stok: outputBuku.Stok - 1, Buku_id: outputBuku.Buku_id}
				data := db.Model(&model.Data_buku{}).Where("buku_id = ?", stok.Buku_id).Update(map[string]interface{}{"stok": stok.Stok})
				if data.Error != nil {
					message = "Error: Gagal Update data Stok"
					status = false
					data.Value = nil
				} else {
					status = true
					outputBuku.Stok = outputBuku.Stok - 1
					datas["buku"] = outputBuku
					datas["mahasiswa"] = mahasiswa.Value
				}
			}
		}
	}

	if status == true {
		check := db.NewRecord(peminjaman)
		if check == true {
			data := db.Create(&peminjaman)
			if data.Error != nil {
				message = "Data Gagal Ditambahkan"
				datas = nil
				status = false
			} else {
				message = "Data Berhasil Ditambahkan"
				status = true
				datas["peminjaman"] = peminjaman
			}
		} else {
			message = "Data Gagal Ditambahkan"
			datas = nil
			status = false
		}
	}

	return model.Return{Status: status, Data: datas, Message: message}
}

func Kembali(Kembali model.Riwayat_peminjaman, db *gorm.DB) model.Return {
	var message string
	var status bool
	datas := make(map[string]interface{})
	var outputBuku model.Data_buku

	dataPinjam := db.Where("id_peminjaman = ?", Kembali.Id_peminjaman).First(&model.Data_peminjaman{})
	if dataPinjam.Error != nil {
		message = "Data Pinjam Gagal Ditemukan"
		status = false
		datas = nil
	} else {
		peminjaman, _ := json.Marshal(dataPinjam.Value)
		_ = json.Unmarshal(peminjaman, &Kembali)

		mahasiswa := db.Where("mhs_id = ?", Kembali.Id_mhs).First(&model.Data_mahasiswa{})
		if mahasiswa.Error != nil {
			message = "Data Mahasiswa Gagal Ditemukan"
			status = false
			mahasiswa.Value = nil
			datas = nil
		} else {
			buku := db.Where("buku_id = ?", Kembali.Id_buku).First(&model.Data_buku{})
			if buku.Error != nil {
				message = "Data Buku Gagal Ditemukan"
				status = false
				buku.Value = nil
				datas = nil
				Kembali.Judul_buku = ""
			} else {
				dataBuku, _ := json.Marshal(buku.Value)
				_ = json.Unmarshal(dataBuku, &outputBuku)
				Kembali.Judul_buku = outputBuku.Judul

				stok := model.UpdateStok{Stok: outputBuku.Stok + 1, Buku_id: outputBuku.Buku_id}
				data := db.Model(&model.Data_buku{}).Where("buku_id = ?", stok.Buku_id).Update(map[string]interface{}{"stok": stok.Stok})
				if data.Error != nil {
					message = "Error: Gagal Update data Stok"
					status = false
					data.Value = nil
				} else {
					status = true
					outputBuku.Stok = outputBuku.Stok + 1
					datas["buku"] = outputBuku
					datas["mahasiswa"] = mahasiswa.Value
				}
			}
		}
	}

	if status == true {
		Kembali.Tanggal_kembali = time.Now().Local()
		check := db.NewRecord(Kembali)
		if check == true {
			data := db.Create(&Kembali)
			if data.Error != nil {
				message = "Buku Gagal Dikembalikan"
				datas = nil
				status = false
			} else {
				message = "Buku Berhasil Dikembalikan"
				status = true
				datas["pengembalian"] = Kembali

				_ = db.Where("id_peminjaman = ?", Kembali.Id_peminjaman).Delete(&model.Data_peminjaman{})
			}
		} else {
			message = "Buku Gagal Dikembalikan"
			datas = nil
			status = false
		}
	}

	return model.Return{Status: status, Data: datas, Message: message}
}

func Riwayat(page model.PeminjamanPaging, db *gorm.DB) model.Return {
	var message string
	var RiwayatPeminjaman []model.Riwayat_peminjaman
	var Records []model.Riwayat_peminjaman

	var DataBase *gorm.DB
	DataBase = db

	if page.Id == "" {
		if page.Search == "" {
			DataBase = db
		} else {
			DataBase = db.Where("id_peminjaman LIKE ? OR id_buku LIKE ? OR judul_buku LIKE ?", "%"+page.Search+"%", "%"+page.Search+"%", "%"+page.Search+"%")
		}
	} else {
		if page.Search == "" {
			DataBase = db.Where("id_mhs LIKE ?", "%"+page.Id+"%")
		} else {
			DataBase = db.Where("(id_mhs LIKE ? AND id_peminjaman LIKE ?) OR (id_mhs LIKE ? AND id_buku LIKE ?) OR (id_mhs LIKE ? AND judul_buku LIKE ?)", "%"+page.Id+"%", "%"+page.Search+"%", "%"+page.Id+"%", "%"+page.Search+"%", "%"+page.Id+"%", "%"+page.Search+"%")
		}
	}

	paginator := pagination.Paging(&pagination.Param{
		DB:      DataBase,
		Page:    page.Page,
		Limit:   page.Size,
		OrderBy: []string{"tanggal_kembali desc"},
		ShowSQL: true,
	}, &RiwayatPeminjaman)

	getRecords, _ := json.Marshal(paginator.Records)
	_ = json.Unmarshal(getRecords, &Records)

	var UpdatedRecords []map[string]interface{}
	for _, records := range Records {
		set := make(map[string]interface{})
		buku := db.Where("buku_id = ?", records.Id_buku).First(&model.Data_buku{})
		if buku.Error != nil {
			buku.Value = nil
			set["data_peminjaman"] = records
			set["detail_buku"] = buku.Value
		} else {
			set["data_peminjaman"] = records
			set["detail_buku"] = buku.Value
			mhs := db.Where("mhs_id = ?", records.Id_mhs).First(&model.Data_mahasiswa{})
			if mhs.Error != nil {
				mhs.Value = nil
				set["detail_mhs"] = mhs.Value
			} else {
				setCord, _ := json.Marshal(mhs.Value)
				var mhsInterface model.Data_mahasiswa
				_ = json.Unmarshal(setCord, &mhsInterface)
				mhsInterface.Password = "AuthGuard Protected!"
				fmt.Print(mhsInterface)
				set["detail_mhs"] = mhsInterface
			}
		}
		UpdatedRecords = append(UpdatedRecords, set)
	}

	paginator.Records = UpdatedRecords
	return model.Return{Status: true, Data: paginator, Message: message}
}

func Berlangsung(page model.PeminjamanPaging, db *gorm.DB) model.Return {
	var message string
	var DataPeminjaman []model.Data_peminjaman
	var Records []model.Data_peminjaman

	var DataBase *gorm.DB
	DataBase = db

	if page.Id == "" {
		if page.Search == "" {
			DataBase = db
		} else {
			DataBase = db.Where("id_peminjaman LIKE ? OR id_buku LIKE ? OR judul_buku LIKE ?", "%"+page.Search+"%", "%"+page.Search+"%", "%"+page.Search+"%")
		}
	} else {
		if page.Search == "" {
			DataBase = db.Where("id_mhs LIKE ?", "%"+page.Id+"%")
		} else {
			DataBase = db.Where("(id_mhs LIKE ? AND id_peminjaman LIKE ?) OR (id_mhs LIKE ? AND id_buku LIKE ?) OR (id_mhs LIKE ? AND judul_buku LIKE ?)", "%"+page.Id+"%", "%"+page.Search+"%", "%"+page.Id+"%", "%"+page.Search+"%", "%"+page.Id+"%", "%"+page.Search+"%")
		}
	}

	paginator := pagination.Paging(&pagination.Param{
		DB:      DataBase,
		Page:    page.Page,
		Limit:   page.Size,
		OrderBy: []string{"tanggal_peminjaman desc"},
		ShowSQL: true,
	}, &DataPeminjaman)

	getRecords, _ := json.Marshal(paginator.Records)
	_ = json.Unmarshal(getRecords, &Records)

	var UpdatedRecords []map[string]interface{}
	for _, records := range Records {
		set := make(map[string]interface{})
		buku := db.Where("buku_id = ?", records.Id_buku).First(&model.Data_buku{})
		if buku.Error != nil {
			buku.Value = nil
			set["data_peminjaman"] = records
			set["detail_buku"] = buku.Value
		} else {
			set["data_peminjaman"] = records
			set["detail_buku"] = buku.Value
			mhs := db.Where("mhs_id = ?", records.Id_mhs).First(&model.Data_mahasiswa{})
			if mhs.Error != nil {
				mhs.Value = nil
				set["detail_mhs"] = mhs.Value
			} else {
				setCord, _ := json.Marshal(mhs.Value)
				var mhsInterface model.Data_mahasiswa
				_ = json.Unmarshal(setCord, &mhsInterface)
				mhsInterface.Password = "AuthGuard Protected!"
				fmt.Print(mhsInterface)
				set["detail_mhs"] = mhsInterface
			}
		}
		UpdatedRecords = append(UpdatedRecords, set)
	}

	paginator.Records = UpdatedRecords
	return model.Return{Status: true, Data: paginator, Message: message}
}

func ViewByIdPeminjaman(id string, db *gorm.DB) model.Return {
	var message string
	var status bool
	var Records model.Data_peminjaman

	data := db.Where("id_peminjaman = ?", id).First(&model.Data_peminjaman{})
	if data.Error != nil {
		message = "Data Gagal Ditemukan"
		status = false
		data.Value = nil
	} else {
		message = "Data Berhasil Ditemukan"
		status = true

	}
	set := make(map[string]interface{})
	if status == true {
		getRecords, _ := json.Marshal(data.Value)
		_ = json.Unmarshal(getRecords, &Records)

		buku := db.Where("buku_id = ?", Records.Id_buku).First(&model.Data_buku{})
		if buku.Error != nil {
			buku.Value = nil
			set["data_peminjaman"] = Records
			set["detail_buku"] = buku.Value
		} else {
			set["data_peminjaman"] = Records
			set["detail_buku"] = buku.Value
			mhs := db.Where("mhs_id = ?", Records.Id_mhs).First(&model.Data_mahasiswa{})
			if mhs.Error != nil {
				mhs.Value = nil
				set["detail_mhs"] = mhs.Value
			} else {
				set["detail_mhs"] = mhs.Value
			}
		}
	}

	return model.Return{Status: status, Data: set, Message: message}
}

func CreateListPeminjamanBerlangsung(page model.PeminjamanPaging, db *gorm.DB) model.Return {
	var message string
	var DataPeminjaman []model.Data_peminjaman
	var Records []model.Data_peminjaman

	var DataBase *gorm.DB
	DataBase = db

	if page.Id == "" {
		if page.Search == "" {
			DataBase = db
		} else {
			DataBase = db.Where("id_peminjaman LIKE ? OR id_buku LIKE ? OR judul_buku LIKE ?", "%"+page.Search+"%", "%"+page.Search+"%", "%"+page.Search+"%")
		}
	} else {
		if page.Search == "" {
			DataBase = db.Where("id_mhs LIKE ?", "%"+page.Id+"%")
		} else {
			DataBase = db.Where("(id_mhs LIKE ? AND id_peminjaman LIKE ?) OR (id_mhs LIKE ? AND id_buku LIKE ?) OR (id_mhs LIKE ? AND judul_buku LIKE ?)", "%"+page.Id+"%", "%"+page.Search+"%", "%"+page.Id+"%", "%"+page.Search+"%", "%"+page.Id+"%", "%"+page.Search+"%")
		}
	}

	paginator := pagination.Paging(&pagination.Param{
		DB:      DataBase,
		Page:    page.Page,
		Limit:   page.Size,
		OrderBy: []string{"tanggal_peminjaman desc"},
		ShowSQL: true,
	}, &DataPeminjaman)

	getRecords, _ := json.Marshal(paginator.Records)
	_ = json.Unmarshal(getRecords, &Records)

	var mhsInterface model.Data_mahasiswa
	var UpdatedRecords []map[string]interface{}
	for _, records := range Records {
		set := make(map[string]interface{})
		buku := db.Where("buku_id = ?", records.Id_buku).First(&model.Data_buku{})
		if buku.Error != nil {
			buku.Value = nil
			set["data_peminjaman"] = records
			set["detail_buku"] = buku.Value
		} else {
			set["data_peminjaman"] = records
			set["detail_buku"] = buku.Value
			mhs := db.Where("mhs_id = ?", records.Id_mhs).First(&model.Data_mahasiswa{})
			if mhs.Error != nil {
				mhs.Value = nil
				set["detail_mhs"] = mhs.Value
			} else {
				setCord, _ := json.Marshal(mhs.Value)
				_ = json.Unmarshal(setCord, &mhsInterface)
				mhsInterface.Password = "AuthGuard Protected!"
				fmt.Print(mhsInterface)
				set["detail_mhs"] = mhsInterface
				setCord, _ = json.Marshal(buku.Value)
				_ = json.Unmarshal(setCord, &mhsInterface)
			}
		}
		UpdatedRecords = append(UpdatedRecords, set)
	}

	paginator.Records = UpdatedRecords

	fungsi.Excelsize_PeminjamanBerlangsung(mhsInterface, Records)

	return model.Return{Status: true, Data: paginator, Message: message}
}

func CreateListPeminjamanRiwayat(page model.PeminjamanPaging, db *gorm.DB) model.Return {
	var message string
	var RiwayatPeminjaman []model.Riwayat_peminjaman
	var Records []model.Riwayat_peminjaman

	var DataBase *gorm.DB
	DataBase = db

	if page.Id == "" {
		if page.Search == "" {
			DataBase = db
		} else {
			DataBase = db.Where("id_peminjaman LIKE ? OR id_buku LIKE ? OR judul_buku LIKE ?", "%"+page.Search+"%", "%"+page.Search+"%", "%"+page.Search+"%")
		}
	} else {
		if page.Search == "" {
			DataBase = db.Where("id_mhs LIKE ?", "%"+page.Id+"%")
		} else {
			DataBase = db.Where("(id_mhs LIKE ? AND id_peminjaman LIKE ?) OR (id_mhs LIKE ? AND id_buku LIKE ?) OR (id_mhs LIKE ? AND judul_buku LIKE ?)", "%"+page.Id+"%", "%"+page.Search+"%", "%"+page.Id+"%", "%"+page.Search+"%", "%"+page.Id+"%", "%"+page.Search+"%")
		}
	}

	paginator := pagination.Paging(&pagination.Param{
		DB:      DataBase,
		Page:    page.Page,
		Limit:   page.Size,
		OrderBy: []string{"tanggal_kembali desc"},
		ShowSQL: true,
	}, &RiwayatPeminjaman)

	getRecords, _ := json.Marshal(paginator.Records)
	_ = json.Unmarshal(getRecords, &Records)

	var mhsInterface model.Data_mahasiswa
	var UpdatedRecords []map[string]interface{}
	for _, records := range Records {
		set := make(map[string]interface{})
		buku := db.Where("buku_id = ?", records.Id_buku).First(&model.Data_buku{})
		if buku.Error != nil {
			buku.Value = nil
			set["data_peminjaman"] = records
			set["detail_buku"] = buku.Value
		} else {
			set["data_peminjaman"] = records
			set["detail_buku"] = buku.Value
			mhs := db.Where("mhs_id = ?", records.Id_mhs).First(&model.Data_mahasiswa{})
			if mhs.Error != nil {
				mhs.Value = nil
				set["detail_mhs"] = mhs.Value
			} else {
				setCord, _ := json.Marshal(mhs.Value)
				_ = json.Unmarshal(setCord, &mhsInterface)
				mhsInterface.Password = "AuthGuard Protected!"
				fmt.Print(mhsInterface)
				set["detail_mhs"] = mhsInterface
			}
		}
		UpdatedRecords = append(UpdatedRecords, set)
	}

	paginator.Records = UpdatedRecords

	fungsi.Excelsize_PeminjamanRiwayat(mhsInterface, Records)

	return model.Return{Status: true, Data: paginator, Message: message}
}

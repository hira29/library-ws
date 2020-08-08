package dao

import (
	"encoding/json"
	"fmt"
	"github.com/biezhi/gorm-paginator/pagination"
	"github.com/jinzhu/gorm"
	"library-ws/model"
	"time"
)

func SetAdminPengembalian(Kembali model.Riwayat_peminjaman, db *gorm.DB) model.Return {
	var message string
	var status bool
	datas := make(map[string]interface{})
	var outputBuku model.Data_buku
	Kembali.Id_rating = "0"

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

func ListPengembalian(page model.PeminjamanPaging, db *gorm.DB) model.Return {
	var message string
	var DataPeminjaman []model.Data_peminjaman
	var Records []model.Data_peminjaman

	var DataBase *gorm.DB
	DataBase = db

	if page.Id == "" {
		if page.Search == "" {
			DataBase = db.Where("pengembalian = ?", 1)
		} else {
			DataBase = db.Where("pengembalian = ? AND id_peminjaman LIKE ? OR id_buku LIKE ? OR judul_buku LIKE ?", 1, "%"+page.Search+"%", "%"+page.Search+"%", "%"+page.Search+"%")
		}
	} else {
		if page.Search == "" {
			DataBase = db.Where("pengembalian = ? AND id_mhs LIKE ?", 1, "%"+page.Id+"%")
		} else {
			DataBase = db.Where("(pengembalian = ? AND id_mhs LIKE ? AND id_peminjaman LIKE ?) OR (pengembalian = ? AND id_mhs LIKE ? AND id_buku LIKE ?) OR (pengembalian = ? AND id_mhs LIKE ? AND judul_buku LIKE ?)", 1, "%"+page.Id+"%", "%"+page.Search+"%", 1, "%"+page.Id+"%", "%"+page.Search+"%", 1, "%"+page.Id+"%", "%"+page.Search+"%")
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

func SetPengembalian(data model.SetPengembalian, db *gorm.DB) model.Return {
	Time_now := time.Now().Local()
	var message string
	var status bool
	get := db.Model(&model.Data_peminjaman{}).Where("id_peminjaman = ?", data.Id_peminjaman).Update(map[string]interface{}{"pengembalian": 1, "tanggal_kembali": Time_now})
	get = db.Where("id_peminjaman = ?", data.Id_peminjaman).Find(&model.Data_peminjaman{})
	if get.Error != nil {
		message = "Error: Gagal Update data"
		status = false
		get.Value = nil
	} else {
		setCord, _ := json.Marshal(get.Value)
		var peminjamanInterface model.Data_peminjaman
		_ = json.Unmarshal(setCord, &peminjamanInterface)

		if peminjamanInterface.Pengembalian == 0 {
			message = "Error: Gagal Update data"
			status = false
			get.Value = nil
		} else {
			message = "Buku Berhasil Dikembalikan, Menunggu persetujuan admin Perpustakaan"
			status = true
		}
	}
	return model.Return{Status: status, Data: get.Value, Message: message}
}

func GetPengembalian(data model.Data_peminjaman, db *gorm.DB) model.Return {
	var message string
	var status bool

	mahasiswa := db.Where("mhs_id = ?", data.Id_mhs).First(&model.Data_mahasiswa{})
	if mahasiswa.Error != nil {
		mahasiswa.Value = nil
	}
	buku := db.Where("buku_id = ?", data.Id_buku).First(&model.Data_buku{})
	if buku.Error != nil {
		buku.Value = nil
	}

	get := db.Where("id_buku = ? AND id_mhs = ?", data.Id_buku, data.Id_mhs).Find(&model.Data_peminjaman{})
	if get.Error != nil {
		message = "Error: Gagal Menemukan data"
		status = false
		get.Value = nil
	} else {
		setCord, _ := json.Marshal(get.Value)
		var peminjamanInterface model.Data_peminjaman
		_ = json.Unmarshal(setCord, &peminjamanInterface)

		if peminjamanInterface.Pengembalian == 0 {
			message = "Data Berhasil Ditemukan!"
			status = true
		} else {
			message = "Buku Telah dikembalikan, Mengunggu persetujuan admin Perpustakaan"
			status = false
		}
	}
	return model.Return{Status: status, Data: map[string]interface{}{"data_peminjaman": get.Value, "detail_mhs": mahasiswa.Value, "detail_buku": buku.Value}, Message: message}
}

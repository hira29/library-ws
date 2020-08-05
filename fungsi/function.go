package fungsi

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"hash/crc32"
	"library-ws/model"
	"strconv"
	"strings"

	"github.com/360EntSecGroup-Skylar/excelize"
)

func ToMd5(string string) string {
	data := []byte(string)
	b := md5.Sum(data)
	pass := hex.EncodeToString(b[:])
	return pass
}

func ToCRC32(string string) string {
	data := []byte(string)
	crc32InUint32 := crc32.ChecksumIEEE([]byte(data))
	crc32InString := strconv.FormatUint(uint64(crc32InUint32), 16)
	return crc32InString
}

func CategoryParseEngine(Klasifikasi string) string {
	classifiers := strings.Split(Klasifikasi, ",")
	var Category string
	for x, class := range classifiers {
		c, _ := strconv.ParseFloat(class, 32)
		if c == 4 {
			if x == 0 {
				Category = Category + "Data processing, computer science"
			} else {
				Category = Category + ", Data processing, computer science"
			}
		} else if c >= 200 && c < 300 {
			if x == 0 {
				Category = Category + "Religion"
			} else {
				Category = Category + ", Religion"
			}
		} else if c >= 300 && c < 400 {
			if x == 0 {
				Category = Category + "Social science"
			} else {
				Category = Category + ", Social science"
			}
		} else if c >= 400 && c < 500 {
			if x == 0 {
				Category = Category + "Language"
			} else {
				Category = Category + ", Language"
			}
		} else if c >= 500 && c < 600 {
			if x == 0 {
				Category = Category + "Natural science"
			} else {
				Category = Category + ", Natural science"
			}
		} else if c >= 620 && c < 621 {
			if x == 0 {
				Category = Category + "Engineering & allied operation"
			} else {
				Category = Category + ", Engineering & allied operation"
			}
		} else if c >= 621.3 && c < 621.4 {
			if x == 0 {
				Category = Category + "Electric, electronic, magnetic, communications"
			} else {
				Category = Category + ", Electric, electronic, magnetic, communications"
			}
		} else if c >= 623 && c < 624 {
			if x == 0 {
				Category = Category + "Ship"
			} else {
				Category = Category + ", Ship"
			}
		} else if c >= 604 && c < 605 {
			if x == 0 {
				Category = Category + "Technical drawing"
			} else {
				Category = Category + ", Technical drawing"
			}
		} else if c >= 621.4 && c < 621.9 {
			if x == 0 {
				Category = Category + "Heat engineering, Pump, Pneumatic, Machine engineering"
			} else {
				Category = Category + ", Heat engineering, Pump, Pneumatic, Machine engineering"
			}
		} else if c >= 658 && c < 659 {
			if x == 0 {
				Category = Category + "General Management"
			} else {
				Category = Category + ", General Management"
			}
		} else if c >= 629 && c < 629.9 {
			if x == 0 {
				Category = Category + "Other branches of engineering; automatic control engineering"
			} else {
				Category = Category + ", Other branches of engineering; automatic control engineering"
			}
		} else if c >= 669 && c < 672 {
			if x == 0 {
				Category = Category + "Metallurgy, manufacturing, metalworking processes"
			} else {
				Category = Category + ", Metallurgy, manufacturing, metalworking processes"
			}
		} else if c >= 700 && c < 701 {
			if x == 0 {
				Category = Category + "The Arts"
			} else {
				Category = Category + ", The Arts"
			}
		} else if c >= 800 && c < 1000 {
			if x == 0 {
				Category = Category + "Literature, history and geography"
			} else {
				Category = Category + ", Literature, history and geography"
			}
		} else {
			if x == 0 {
				Category = Category + "Unknown Category"
			} else {
				Category = Category + ", Unknown Category"
			}
		}
	}
	return Category
}

func Excelsize_Books(Records []model.Data_buku) {
	f := excelize.NewFile()
	// Create a new sheet.
	index := f.NewSheet("Sheet1")
	// Set value of a cell.
	// Set active sheet of the workbook.
	f.SetActiveSheet(index)
	// Save xlsx file by the given path.
	f.SetCellValue("Sheet1", "A1", "No.")
	f.SetCellValue("Sheet1", "B1", "ID Buku")
	f.SetCellValue("Sheet1", "C1", "Judul")
	f.SetCellValue("Sheet1", "D1", "Edisi")
	f.SetCellValue("Sheet1", "E1", "Pengarang")
	f.SetCellValue("Sheet1", "F1", "Kota Terbit")
	f.SetCellValue("Sheet1", "G1", "Penerbit")
	f.SetCellValue("Sheet1", "H1", "Tahun Terbit")
	f.SetCellValue("Sheet1", "I1", "ISBN")
	f.SetCellValue("Sheet1", "J1", "Klasifikasi")
	f.SetCellValue("Sheet1", "K1", "Kategori")
	f.SetCellValue("Sheet1", "L1", "Umum/Res")
	f.SetCellValue("Sheet1", "M1", "Bahasa")
	f.SetCellValue("Sheet1", "N1", "Deskripsi")
	f.SetCellValue("Sheet1", "O1", "Lokasi")
	f.SetCellValue("Sheet1", "P1", "Tanggal Ditambahkan")
	f.SetCellValue("Sheet1", "Q1", "Stok")

	for i, records := range Records {
		num := strconv.Itoa(i + 2)
		f.SetCellValue("Sheet1", "A"+num, strconv.Itoa(i+1))
		f.SetCellValue("Sheet1", "B"+num, records.Buku_id)
		f.SetCellValue("Sheet1", "C"+num, records.Judul)
		f.SetCellValue("Sheet1", "D"+num, records.Edisi)
		f.SetCellValue("Sheet1", "E"+num, records.Pengarang)
		f.SetCellValue("Sheet1", "F"+num, records.Kota_terbit)
		f.SetCellValue("Sheet1", "G"+num, records.Penerbit)
		f.SetCellValue("Sheet1", "H"+num, records.Tahun_Terbit)
		f.SetCellValue("Sheet1", "I"+num, records.Isbn)
		f.SetCellValue("Sheet1", "J"+num, records.Klasifikasi)
		f.SetCellValue("Sheet1", "K"+num, records.Kategori)
		f.SetCellValue("Sheet1", "L"+num, records.Umum_res)
		f.SetCellValue("Sheet1", "M"+num, records.Bahasa)
		f.SetCellValue("Sheet1", "N"+num, records.Deskripsi)
		f.SetCellValue("Sheet1", "O"+num, records.Lokasi)
		f.SetCellValue("Sheet1", "P"+num, records.Tanggal_ditambahkan)
		f.SetCellValue("Sheet1", "Q"+num, records.Stok)
	}

	if err := f.SaveAs("DataBooks.xlsx"); err != nil {
		fmt.Println(err)
	}

}

func Excelsize_Mhs(Records []model.Data_mahasiswa) {
	f := excelize.NewFile()
	// Create a new sheet.
	index := f.NewSheet("Sheet1")
	// Set value of a cell.
	// Set active sheet of the workbook.
	f.SetActiveSheet(index)
	// Save xlsx file by the given path.
	f.SetCellValue("Sheet1", "A1", "No.")
	f.SetCellValue("Sheet1", "B1", "ID Mahasiswa")
	f.SetCellValue("Sheet1", "C1", "Nama")
	f.SetCellValue("Sheet1", "D1", "Email")
	f.SetCellValue("Sheet1", "E1", "Password")
	f.SetCellValue("Sheet1", "F1", "Tempat Lahir")
	f.SetCellValue("Sheet1", "G1", "Tanggal Lahir")
	f.SetCellValue("Sheet1", "H1", "Aktif")

	for i, records := range Records {
		num := strconv.Itoa(i + 2)
		f.SetCellValue("Sheet1", "A"+num, strconv.Itoa(i+1))
		f.SetCellValue("Sheet1", "B"+num, records.Mhs_id)
		f.SetCellValue("Sheet1", "C"+num, records.Nama)
		f.SetCellValue("Sheet1", "D"+num, records.Email)
		f.SetCellValue("Sheet1", "E"+num, records.Password)
		f.SetCellValue("Sheet1", "F"+num, records.Tempat_lahir)
		f.SetCellValue("Sheet1", "G"+num, records.Tanggal_lahir)
		f.SetCellValue("Sheet1", "H"+num, records.Active)
	}

	if err := f.SaveAs("DataMhs.xlsx"); err != nil {
		fmt.Println(err)
	}

}

func Excelsize_PeminjamanBerlangsung(RecordsMhs model.Data_mahasiswa, Records []model.Data_peminjaman) {
	f := excelize.NewFile()
	// Create a new sheet.
	index := f.NewSheet("Sheet1")
	// Set value of a cell.
	// Set active sheet of the workbook.
	f.SetActiveSheet(index)
	// Save xlsx file by the given path.
	f.SetCellValue("Sheet1", "A1", "ID")
	f.SetCellValue("Sheet1", "A2", "Nama")
	f.SetCellValue("Sheet1", "A3", "Email")
	f.SetCellValue("Sheet1", "B1", RecordsMhs.Mhs_id)
	f.SetCellValue("Sheet1", "B2", RecordsMhs.Nama)
	f.SetCellValue("Sheet1", "B3", RecordsMhs.Email)

	f.SetCellValue("Sheet1", "A5", "No.")
	f.SetCellValue("Sheet1", "B5", "ID Peminjaman")
	f.SetCellValue("Sheet1", "C5", "Id Buku")
	f.SetCellValue("Sheet1", "D5", "Judul Buku")
	f.SetCellValue("Sheet1", "E5", "Tanggal Peminjaman")
	f.SetCellValue("Sheet1", "F5", "Tanggal Pengembalian")
	f.SetCellValue("Sheet1", "G5", "Tanggal Kembali")

	for i, records := range Records {
		num := strconv.Itoa(i + 6)
		f.SetCellValue("Sheet1", "A"+num, strconv.Itoa(i+1))
		f.SetCellValue("Sheet1", "B"+num, records.Id_peminjaman)
		f.SetCellValue("Sheet1", "C"+num, records.Id_buku)
		f.SetCellValue("Sheet1", "D"+num, records.Judul_buku)
		f.SetCellValue("Sheet1", "F"+num, records.Tanggal_peminjaman)
		f.SetCellValue("Sheet1", "G"+num, records.Tanggal_pengembalian)
		f.SetCellValue("Sheet1", "H"+num, records.Tanggal_kembali)
	}

	if err := f.SaveAs("DataPeminjamanBerlangsung.xlsx"); err != nil {
		fmt.Println(err)
	}
}

func Excelsize_PeminjamanRiwayat(RecordsMhs model.Data_mahasiswa, Records []model.Riwayat_peminjaman) {
	f := excelize.NewFile()
	// Create a new sheet.
	index := f.NewSheet("Sheet1")
	// Set value of a cell.
	// Set active sheet of the workbook.
	f.SetActiveSheet(index)
	// Save xlsx file by the given path.
	f.SetCellValue("Sheet1", "A1", "ID")
	f.SetCellValue("Sheet1", "A2", "Nama")
	f.SetCellValue("Sheet1", "A3", "Email")
	f.SetCellValue("Sheet1", "B1", RecordsMhs.Mhs_id)
	f.SetCellValue("Sheet1", "B2", RecordsMhs.Nama)
	f.SetCellValue("Sheet1", "B3", RecordsMhs.Email)

	f.SetCellValue("Sheet1", "A5", "No.")
	f.SetCellValue("Sheet1", "B5", "ID Peminjaman")
	f.SetCellValue("Sheet1", "C5", "Id Buku")
	f.SetCellValue("Sheet1", "D5", "Judul Buku")
	f.SetCellValue("Sheet1", "E5", "Tanggal Peminjaman")
	f.SetCellValue("Sheet1", "F5", "Tanggal Pengembalian")
	f.SetCellValue("Sheet1", "G5", "Tanggal Kembali")

	for i, records := range Records {
		num := strconv.Itoa(i + 6)
		f.SetCellValue("Sheet1", "A"+num, strconv.Itoa(i+1))
		f.SetCellValue("Sheet1", "B"+num, records.Id_peminjaman)
		f.SetCellValue("Sheet1", "C"+num, records.Id_buku)
		f.SetCellValue("Sheet1", "D"+num, records.Judul_buku)
		f.SetCellValue("Sheet1", "E"+num, records.Tanggal_peminjaman)
		f.SetCellValue("Sheet1", "F"+num, records.Tanggal_pengembalian)
		f.SetCellValue("Sheet1", "G"+num, records.Tanggal_kembali)
	}

	if err := f.SaveAs("DataPeminjamanRiwayat.xlsx"); err != nil {
		fmt.Println(err)
	}

}

func Excelsize_Rating(RecordsBooks model.Data_buku, Records []model.Data_Rating) {
	f := excelize.NewFile()
	// Create a new sheet.
	index := f.NewSheet("Sheet1")
	// Set value of a cell.
	// Set active sheet of the workbook.
	f.SetActiveSheet(index)
	// Save xlsx file by the given path.
	f.SetCellValue("Sheet1", "A1", "ID")
	f.SetCellValue("Sheet1", "A2", "Judul")
	f.SetCellValue("Sheet1", "A3", "Pengarang")
	f.SetCellValue("Sheet1", "A4", "Penerbit / Tahun Terbit")
	f.SetCellValue("Sheet1", "A5", "Rating")
	f.SetCellValue("Sheet1", "B1", RecordsBooks.Buku_id)
	f.SetCellValue("Sheet1", "B2", RecordsBooks.Judul)
	f.SetCellValue("Sheet1", "B3", RecordsBooks.Pengarang)
	f.SetCellValue("Sheet1", "B4", RecordsBooks.Penerbit+"/"+RecordsBooks.Tahun_Terbit)
	f.SetCellValue("Sheet1", "B5", RecordsBooks.Rating)

	f.SetCellValue("Sheet1", "A7", "No.")
	f.SetCellValue("Sheet1", "B7", "ID Rating")
	f.SetCellValue("Sheet1", "C7", "Nama")
	f.SetCellValue("Sheet1", "D7", "ID Mhs")
	f.SetCellValue("Sheet1", "E7", "Rating")
	f.SetCellValue("Sheet1", "F7", "Komentar")
	f.SetCellValue("Sheet1", "G7", "Tanggal")

	for i, records := range Records {
		num := strconv.Itoa(i + 8)
		f.SetCellValue("Sheet1", "A"+num, strconv.Itoa(i+1))
		f.SetCellValue("Sheet1", "B"+num, records.Id_Rating)
		f.SetCellValue("Sheet1", "C"+num, records.Nama)
		f.SetCellValue("Sheet1", "D"+num, records.Id_Mhs)
		f.SetCellValue("Sheet1", "E"+num, records.Rating)
		f.SetCellValue("Sheet1", "F"+num, records.Komentar)
		f.SetCellValue("Sheet1", "G"+num, records.Tanggal)
	}

	if err := f.SaveAs("DataBooksRating.xlsx"); err != nil {
		fmt.Println(err)
	}

}

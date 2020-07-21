package main

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"library-ws/controller"

	//httpSwagger "github.com/swaggo/http-swagger"
	"log"
	"net/http"
)

func main() {
	port := "6996"
	r := mux.NewRouter()

	headers := handlers.AllowedHeaders([]string{
		"X-Requested-With", "Accept", "Authorization", "Content-Type", "X-CSRF-Token",
	})
	origins := handlers.AllowedOrigins([]string{"*"})
	methods := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS", "DELETE"})

	api := r.PathPrefix("/perpustakaan/api/v1").Subrouter()
	buku := api.PathPrefix("/data_buku").Subrouter()
	buku.HandleFunc("/create", controller.CreateBuku).Methods("POST")
	buku.HandleFunc("/list", controller.ListBuku).Methods("POST")
	buku.HandleFunc("/view/{bukuId}", controller.ViewBukuById).Methods("GET")
	buku.HandleFunc("/update", controller.UpdateBukuById).Methods("PUT")
	buku.HandleFunc("/update/stok", controller.UpdateStokById).Methods("POST")
	buku.HandleFunc("/delete/{bukuId}", controller.DeleteBukuById).Methods("DELETE")

	mhs := api.PathPrefix("/data_mhs").Subrouter()
	mhs.HandleFunc("/create", controller.CreateMhs).Methods("POST")
	mhs.HandleFunc("/list", controller.ListMhs).Methods("POST")
	mhs.HandleFunc("/view/{mhsId}", controller.ViewMhsById).Methods("GET")
	mhs.HandleFunc("/update", controller.UpdateMhsById).Methods("PUT")
	mhs.HandleFunc("/delete/{bukuId}", controller.DeleteMhsById).Methods("DELETE")
	mhs.HandleFunc("/login", controller.Login).Methods("POST")
	mhs.HandleFunc("/register", controller.RegisterMhs).Methods("POST")

	peminjaman := api.PathPrefix("/peminjaman").Subrouter()
	peminjaman.HandleFunc("/pinjam", controller.Pinjam).Methods("POST")
	peminjaman.HandleFunc("/riwayat", controller.Riwayat).Methods("POST")
	peminjaman.HandleFunc("/berlangsung", controller.Berlangsung).Methods("POST")
	peminjaman.HandleFunc("/kembali", controller.Kembali).Methods("POST")
	peminjaman.HandleFunc("/view/{idPeminjaman}", controller.ViewByIdPeminjaman).Methods("GET")

	kategori := api.PathPrefix("/kategori").Subrouter()
	kategori.HandleFunc("/tambah/{categoryName}", controller.AddCategory).Methods("GET")
	kategori.HandleFunc("/hapus/{categoryName}", controller.DeleteCategory).Methods("DELETE")
	kategori.HandleFunc("/list", controller.ListCategory).Methods("GET")

	summary := api.PathPrefix("/summary").Subrouter()
	summary.HandleFunc("/buku", controller.BookSummary).Methods("GET")
	summary.HandleFunc("/peminjaman", controller.LoanSummary).Methods("GET")
	summary.HandleFunc("/pengembalian", controller.ReturnSummary).Methods("GET")

	userManagement := api.PathPrefix("/usermanagement").Subrouter()
	userManagement.HandleFunc("/login", controller.UMLogin).Methods("POST")
	userManagement.HandleFunc("/create", controller.UMCreate).Methods("POST")
	userManagement.HandleFunc("/list", controller.UMList).Methods("POST")
	userManagement.HandleFunc("/update", controller.UMUpdate).Methods("PUT")
	userManagement.HandleFunc("/delete/{id}", controller.UMDelete).Methods("DELETE")

	log.Println("API STARTED!")
	_ = http.ListenAndServe(":"+port, handlers.CORS(headers, origins, methods)(r))
}

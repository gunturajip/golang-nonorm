package models

type MahasiswaHobi struct {
	IDMahasiswa int `json:"id_mahasiswa" form:"id_mahasiswa"`
	IDHobi      int `json:"id_hobi" form:"id_hobi"`
}

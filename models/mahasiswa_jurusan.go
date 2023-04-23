package models

type MahasiswaJurusan struct {
	IDMahasiswa int `json:"id_mahasiswa" form:"id_mahasiswa"`
	IDJurusan   int `json:"id_jurusan" form:"id_jurusan"`
}

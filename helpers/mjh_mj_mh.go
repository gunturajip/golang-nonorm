package helpers

import "golang-nonorm/models"

// MJHMJMH struct stands for "Mahasiswa, Jurusan, Hobi, MahasiswaJurusan, MahasiswaHobi" which combine
// mahasiswa, jurusan, hobi, mahasiswa_jurusan, and mahasiswa_hobi data
type MJHMJMH struct {
	Mahasiswa        models.Mahasiswa        `json:"mahasiswa"`
	Jurusan          models.Jurusan          `json:"jurusan"`
	Hobi             models.Hobi             `json:"hobi"`
	MahasiswaJurusan models.MahasiswaJurusan `json:"mahasiswa_jurusan"`
	MahasiswaHobi    models.MahasiswaHobi    `json:"mahasiswa_hobi"`
}

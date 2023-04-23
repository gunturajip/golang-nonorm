package helpers

import "golang-nonorm/models"

// MJH struct stands for "Mahasiswa, Jurusan, Hobi" which combine
// mahasiswa data with nama_jurusan and nama_hobi of each mahasiswa
type MJH struct {
	Mahasiswa   models.Mahasiswa `json:"mahasiswa"`
	NamaJurusan string           `json:"nama_jurusan"`
	NamaHobi    string           `json:"nama_hobi"`
}

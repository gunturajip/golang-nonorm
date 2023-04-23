package models

type Jurusan struct {
	ID          int    `json:"id"`
	NamaJurusan string `json:"nama_jurusan" form:"nama_jurusan"`
}

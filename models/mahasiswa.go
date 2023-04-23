package models

import "time"

type Mahasiswa struct {
	ID                int       `json:"id"`
	Nama              string    `json:"nama" form:"nama"`
	Usia              int       `json:"usia" form:"usia"`
	Gender            bool      `json:"gender" form:"gender"`
	TanggalRegistrasi time.Time `json:"tanggal_registrasi" form:"tanggal_registrasi"`
}

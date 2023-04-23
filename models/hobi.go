package models

type Hobi struct {
	ID       int    `json:"id"`
	NamaHobi string `json:"nama_hobi" form:"nama_hobi"`
}

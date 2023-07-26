package models

type Users struct {
	Id      int    `gorm:"primaryKey" json:"id"`
	Nim     int    `gorm:"unique" json:"nim" validate:"required"`
	Nama    string `gorm:"type: varChar(20)" json:"nama" validate:"required"`
	Jurusan string `gorm:"type: varChar(20)" json:"jurusan" validate:"required"`
}

package models

import "time"

type Absensi struct {
	ID         int64     `gorm:"primaryKey" json:"id"`
	KaryawanID uint      `json:"karyawan_id"`
	Status     string    `gorm:"type:varchar(50)" json:"status"`
	Keterangan string    `gorm:"type:varchar(100)" json:"keterangan"`
	Date       time.Time `gorm:"autoCreateTime" json:"date"`
	Karyawan   Karyawan  `gorm:"foreignKey:KaryawanID" json:"karyawan"`
}

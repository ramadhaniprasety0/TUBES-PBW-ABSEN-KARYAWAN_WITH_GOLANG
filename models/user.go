package models

type User struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	Username string `gorm:"unique"`
	Password string
	Nama     string `gorm:"type:varchar(300)" json:"nama_admin"`
}

package models

type User struct {
	iduser uint `gorm:"primaryKey"`
	email string
	password string
}
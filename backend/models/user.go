package models

type User struct {
	Iduser uint `gorm:"primaryKey"`
	Email string
	Password string
}
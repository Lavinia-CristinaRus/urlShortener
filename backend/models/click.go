package models
import "time"

type Click struct {
	Idclick uint `gorm:"primaryKey"`
	Timestamp time.Time
	Ip_address string
	Country string
	City string
	Idurl uint
}
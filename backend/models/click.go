package models
import "time"

type Url struct {
	idclick uint `gorm:"primaryKey"`
	timestamp time.Time
	ip_address string
	country string
	city string
	idurl uint
}
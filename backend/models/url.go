package models
import "time"

type Url struct {
	idurl uint `gorm:"primaryKey"`
	long_url string
	short_url string `gorm:"uniqueIndex"`
	created_at time.Time
	expires_at *time.Time `gorm:"type:date"`
	iduser uint
}
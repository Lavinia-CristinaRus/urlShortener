package models
import "time"

type Url struct {
	Idurl uint `gorm:"primaryKey"`
	Long_url string
	Short_url string `gorm:"uniqueIndex"`
	Created_at time.Time
	Expires_at time.Time `gorm:"type:date"`
	Iduser uint
}
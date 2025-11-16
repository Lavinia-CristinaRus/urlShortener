package database
import (
	"log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := "root:admin@tcp(127.0.0.1:3306)/url-shortener?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Could not connect to MySQL:", err)
	}
	DB = db
	log.Println("Connected to MySQL")
}
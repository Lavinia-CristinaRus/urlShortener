package url

import (
	"math/rand"
	"time"
	"gorm.io/gorm"
	"log"

	"github.com/gin-gonic/gin"
	"url-shortener/database"
	"url-shortener/models"
)

func generateShortUrl() string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	rand.Seed(time.Now().UnixNano())

	code := make([]byte, 7)
	for i := range code {
		code[i] = letters[rand.Intn(len(letters))]
	}

	return string(code)
}

func GenerateUrl(c *gin.Context) {
	var body struct {
		Url        string `json:"Url"`
		Iduser	   uint `json:"Iduser"`
	}

	if err := c.BindJSON(&body); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	const maxRetries = 2000
	i := -1
	Short_url := ""
	for i = 0; i < maxRetries; i++ {
		Short_url = generateShortUrl()

		var found models.Url
		err := database.DB.Where("short_url = ?", Short_url).First(&found).Error
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				break;
			}
			log.Println("DB error while checking short URL:", err)
			continue
		}
	}
	log.Println(i)
	if i == 2000 {
		c.JSON(500, gin.H{"error": "Could not create URL, all 2000 generated options are already in use"})
		return
	}

	url := models.Url{
		Long_url:    body.Url,
		Short_url:   Short_url,
		Created_at:  time.Now(),
		Expires_at:  time.Now().AddDate(0, 6, 0),
		Iduser:      body.Iduser,
	}

	if err := database.DB.Create(&url).Error; err != nil {
		c.JSON(500, gin.H{"error": "Could not create URL, a database error occurred"})
		return
	}

	c.JSON(200, gin.H{"short_url": Short_url})
}

func CustomizeUrl(c *gin.Context) {
	var body struct {
		Idurl        uint `json:"Idurl"`
		Customurl	 string `json:"Customurl"`
	}

	if err := c.BindJSON(&body); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}
	// log.Println("body:", body)

	var found models.Url
	err := database.DB.Where("short_url = ?", body.Customurl).First(&found).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			if err := database.DB.Model(&models.Url{}).Where("idurl = ?", body.Idurl).Update("short_url", body.Customurl).Error; err != nil {
				c.JSON(400, gin.H{"error": "Failed to update the url"})
				log.Println("Update failed:", err)
				return
			}
			c.JSON(200, gin.H{"message": "Url successfully customized"})
			return
		}
		c.JSON(400, gin.H{"error": "Failed to update the url"})
		log.Println("DB error while checking short URL:", err)
		return
	}

	c.JSON(400, gin.H{"error": "Short url provided is already in use"})
	log.Println("Url already in use")
	return
}
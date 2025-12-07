package url

import (
	"math/rand"
	"time"

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
		Url        string `json:"url"`
	}

	if err := c.BindJSON(&body); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	const maxRetries = 2000 {

	}

	for i := 0; i < maxRetries; i++ {
		Short_url = generateShortUrl()

		var found models.Url
		err := database.DB.Where("short_url = ?", shortCode).First(&found).Error
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				break;
			}
			log.Println("DB error while checking short URL:", err)
			continue
		}
	}
	if i == 2000 {
		c.JSON(500, gin.H{"error": "Could not create URL, all 2000 generated options are already in use"})
		return
	}

	url := models.Url{
		Long_url:   body.Url,
		Short_url: Short_url,
	}

	if err := database.DB.Create(&url).Error; err != nil {
		c.JSON(500, gin.H{"error": "Could not create URL, a database error occurred"})
		return
	}

	c.JSON(200, gin.H{"short_url": Short_url})
}
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
		Short_url string `json:"custom_code"`
	}

	if err := c.BindJSON(&body); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	Short_url := body.Short_url
	if Short_url == "" {
		Short_url = generateShortUrl()
	}

	url := models.Url{
		Long_url:   body.Url,
		Short_url: Short_url,
	}

	if err := database.DB.Create(&url).Error; err != nil {
		c.JSON(500, gin.H{"error": "Could not create URL"})
		return
	}

	c.JSON(200, gin.H{"short_url": Short_url})
}
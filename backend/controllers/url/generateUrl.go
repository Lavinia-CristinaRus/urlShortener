package url

import (
	"math/rand"
	"time"
	"gorm.io/gorm"
	"log"

	"github.com/gin-gonic/gin"
	"url-shortener/database"
	"url-shortener/models"
	"url-shortener/controllers/user"
)

func randRange(min, max int) int {
    return rand.Intn(max-min) + min
}

func generateShortUrl() string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	rand.Seed(time.Now().UnixNano())
	length := randRange(6, 8)

	code := make([]byte, length)
	for i := range code {
		code[i] = letters[rand.Intn(len(letters))]
	}

	return string(code)
}

func GenerateUrl(c *gin.Context) {
	var body struct {
		Url        string `json:"url"`
	}
	email := c.MustGet("email").(string)
	iduser := user.GetUserIdByEmail(email)
	if iduser == -1 {
		c.JSON(400, gin.H{"error": "Invalid token"})
		return
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
		Iduser:      uint(iduser),
	}

	if err := database.DB.Create(&url).Error; err != nil {
		c.JSON(500, gin.H{"error": "Could not create URL, a database error occurred"})
		return
	}

	c.JSON(200, gin.H{"short_url": Short_url})
}
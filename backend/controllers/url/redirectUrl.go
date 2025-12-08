package url

import (
	"gorm.io/gorm"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"url-shortener/database"
	"url-shortener/models"
)

func RedirectUrl(c *gin.Context) {
    shortCode := c.Param("short")

    var url models.Url
    err := database.DB.Where("short_url = ?", shortCode).First(&url).Error
    if err != nil {
        if err == gorm.ErrRecordNotFound {
            c.JSON(404, gin.H{"error": "Short URL not found"})
            return
        }
		log.Println("Database error: ", err)
        c.JSON(500, gin.H{"error": "Database error"})
        return
    }

    if time.Now().After(url.Expires_at) {
        c.JSON(400, gin.H{"error": "URL has expired"})
        return
    }

    c.Redirect(302, url.Long_url)
}
package url

import (
	"time"
	"gorm.io/gorm"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"url-shortener/database"
	"url-shortener/models"
	"url-shortener/controllers/user"
)

func CustomizeUrl(c *gin.Context) {
	var body struct {
		Idurl        uint `json:"Idurl"`
		Customurl	 string `json:"Customurl"`
	}

	if err := c.BindJSON(&body); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

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

func SetUrlExpirationDate(c *gin.Context) {
	var body struct {
		Idurl        uint `json:"Idurl"`
		Expires_at	 time.Time `json:"Expires_at"`
	}

	if err := c.BindJSON(&body); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}
	if time.Now().After(body.Expires_at) {
		c.JSON(400, gin.H{"error": "Invalid expiration date, it can't be a past date"})
		return
	}

	if err := database.DB.Model(&models.Url{}).Where("idurl = ?", body.Idurl).Update("Expires_at", body.Expires_at).Error; err != nil {
		c.JSON(400, gin.H{"error": "Failed to update the url expiration date"})
		log.Println("Update failed:", err)
		return
	}
	c.JSON(200, gin.H{"message": "Url expiration date is set"})
	return

}

func GetMyURLs(c *gin.Context) {
	email := c.MustGet("email").(string)
	iduser := user.GetUserIdByEmail(email)
	if iduser == -1 {
		c.JSON(400, gin.H{"error": "Invalid token"})
		return
	}
	var urls []models.Url
	if err := database.DB.Where("iduser = ?", iduser).Find(&urls).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
            c.JSON(200, gin.H{"message": "You don't have any urls at the moment"})
            return
        }
		c.JSON(400, gin.H{"error": "Failed to get the urls"})
		log.Println("Get urls failed:", err)
		return
	}

	c.JSON(http.StatusOK, urls)
}
package url

import (
	"log"

	"github.com/gin-gonic/gin"
	qrcode "github.com/skip2/go-qrcode"
)

func GenerateQrCode(c *gin.Context) {
	var body struct {
		Shorturl        string `json:"ShortUrl"`
	}

	if err := c.BindJSON(&body); err != nil {
		c.JSON(400, gin.H{"error": "Error while retrieving the short url"})
		return
	}

	err := qrcode.WriteFile("http://127.0.0.1:4000/" + body.Shorturl, qrcode.Medium, 256, body.Shorturl + ".png")
    if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{"error": "Could not generate the QR code"})
        return
    }
	c.JSON(200, gin.H{"message": "QR code generated successfully"})
    return
}
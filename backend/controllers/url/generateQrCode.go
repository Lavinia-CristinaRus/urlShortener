package url

import (
	"log"

	"github.com/gin-gonic/gin"
	qrcode "github.com/skip2/go-qrcode"
)

func GenerateQrCode(c *gin.Context) {
	shortUrl := c.Param("shortUrl")

	png, _ := qrcode.Encode("http://127.0.0.1:4000/" + shortUrl, qrcode.Medium, 256)
    if png == nil {
		log.Println("err")
		c.JSON(500, gin.H{"error": "Could not generate the QR code"})
        return
    }
	log.Println("png")
	c.Data(200, "image/png", png)
    return
}
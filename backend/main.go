package main
import (
	"fmt"
	"log"
	"url-shortener/database"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"url-shortener/controllers/user"
	"url-shortener/controllers/url"
	"time"
)

func main() {
	fmt.Println("Main")
	database.Connect()

	r := gin.Default()

	r.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"http://localhost:5173"},
        AllowMethods:     []string{"POST", "GET", "OPTIONS"},
        AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
        MaxAge: 12 * time.Hour,
    }))

	r.POST("/api/signup", user.Signup)
	r.POST("/api/signin", user.Signin)
	r.POST("/api/generateUrl", url.GenerateUrl)
	r.POST("/api/customizeUrl", url.CustomizeUrl)
	r.POST("/api/setUrlExpirationDate", url.SetUrlExpirationDate)
	r.POST("/api/generateQrCode", url.GenerateQrCode)
	r.GET("/:short", url.RedirectUrl)
	log.Fatal(r.Run(":4000"))
}
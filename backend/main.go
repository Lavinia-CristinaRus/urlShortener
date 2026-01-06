package main
import (
	"fmt"
	"log"
	"url-shortener/database"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"url-shortener/controllers/user"
	"url-shortener/controllers/url"
	"url-shortener/utils"
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
	r.GET("/:short", url.RedirectUrl)

	auth := r.Group("/api")
    auth.Use(utils.AuthMiddleware())
    {
        auth.GET("/urls", url.GetMyURLs)
		auth.POST("/api/generateUrl", url.GenerateUrl)
		auth.POST("/api/customizeUrl", url.CustomizeUrl)
		auth.POST("/api/setUrlExpirationDate", url.SetUrlExpirationDate)
		auth.POST("/api/generateQrCode", url.GenerateQrCode)
    }

	log.Fatal(r.Run(":4000"))
}
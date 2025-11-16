package main
import (
	"fmt"
	"log"
	"url-shortener/database"
	"github.com/gin-gonic/gin"
	"url-shortener/controllers/user"
)

func main() {
	fmt.Println("Main")
	database.Connect()

	r := gin.Default()
	r.POST("/api/signup", user.Signup)
	r.POST("/api/signin", user.Signin)
	log.Fatal(r.Run(":4000"))
}
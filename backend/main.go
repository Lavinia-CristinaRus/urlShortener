package main
import (
	"fmt"
	"log"
	"url-shortener/database"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Main")
	database.Connect()

	r := gin.Default()
	log.Fatal(r.Run(":4000"))
}
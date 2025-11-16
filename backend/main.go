package main
import (
	"fmt"
	"log"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Main")

	r := gin.Default()

	log.Fatal(r.Run(":4000"))
}
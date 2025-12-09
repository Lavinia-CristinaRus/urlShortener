package url

import (
	"gorm.io/gorm"
	"log"
	"time"
	"net/netip"
	"net/http"
	"io"

	"github.com/oschwald/maxminddb-golang/v2"
	"github.com/gin-gonic/gin"
	"url-shortener/database"
	"url-shortener/models"
)

func getPublicIP() string {
    resp, err := http.Get("https://api.ipify.org")
    if err != nil {
        log.Fatal(err)
    }
    defer resp.Body.Close()

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        log.Fatal(err)
    }
    return string(body)
}

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

	ipAddr := c.ClientIP()
	if ipAddr == "::1" {
		ipAddr = getPublicIP()
	}
	log.Println("IP: ", ipAddr)

	var count int64

    err = database.DB.Model(&models.Click{}).
        Where("ip_address = ?", ipAddr).
        Where("timestamp > ?", time.Now().Add(-time.Minute)).
        Count(&count).Error

	if (err != nil) && (err != gorm.ErrRecordNotFound) {
		log.Println(err)
		c.JSON(500, gin.H{"error": "Could not read entries, a database error occurred"})
		return
	}

	if count >= 10 {
		c.JSON(400, gin.H{"error": "Maximum number of clicks reached in the last minute, please try again later"})
		return
	}

	db, err := maxminddb.Open("./geoLite/GeoLite2-City.mmdb")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	ip, err := netip.ParseAddr(ipAddr)
	if err != nil {
		log.Fatal(err)
	}

	var record struct {
		Country struct {
			ISOCode string            `maxminddb:"iso_code"`
			Names   map[string]string `maxminddb:"names"`
		} `maxminddb:"country"`
		City struct {
			Names map[string]string `maxminddb:"names"`
		} `maxminddb:"city"`
	}

	err = db.Lookup(ip).Decode(&record)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Country: ", record.Country.Names["en"], record.Country.ISOCode)
	log.Println("City: ", record.City.Names["en"])

	click := models.Click{
		Timestamp:  time.Now(),
		Ip_address: ipAddr,
		Country:	record.Country.Names["en"],
		City:		record.City.Names["en"],
		Idurl:		url.Idurl,
	}

	if err := database.DB.Create(&click).Error; err != nil {
		c.JSON(500, gin.H{"error": "Could not create object, a database error occurred"})
		return
	}
    c.Redirect(302, url.Long_url)
}
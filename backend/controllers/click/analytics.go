package click

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"url-shortener/database"
)

func GetStatisticsByUrlId(c *gin.Context) {
	idurl := c.Param("id")

	var totalClicks int64
	var uniqueVisitors int64

	database.DB.
		Table("clicks").
		Where("idurl = ?", idurl).
		Count(&totalClicks)


	database.DB.
		Table("clicks").
		Where("idurl = ?", idurl).
		Distinct("ip_address").
		Count(&uniqueVisitors)

	type GeoStat struct {
		Country string
		City    string
		Clicks  int64
	}

	geoStats := make([]GeoStat, 0)

	database.DB.Raw(`
		SELECT country, city, COUNT(*) as clicks
		FROM clicks
		WHERE idurl = ?
		GROUP BY country, city
		ORDER BY clicks DESC
	`, idurl).Scan(&geoStats)

	type VisitorClicks struct {
		IpAddress string `json:"ip"`
		Clicks    int64  `json:"clicks"`
	}
	type CityVisitors struct {
		City     string `json:"city"`
		Visitors int64  `json:"visitors"`
	}

	clicksPerVisitor := make([]VisitorClicks, 0)
	visitorsPerCity := make([]CityVisitors, 0)

	database.DB.Raw(`
		SELECT ip_address, COUNT(*) as clicks
		FROM clicks
		WHERE idurl = ?
		GROUP BY ip_address
	`, idurl).Scan(&clicksPerVisitor)

	database.DB.Raw(`
		SELECT city, COUNT(DISTINCT ip_address) as visitors
		FROM clicks
		WHERE idurl = ?
		GROUP BY city
	`, idurl).Scan(&visitorsPerCity)

	c.JSON(http.StatusOK, gin.H{
		"total_clicks":       totalClicks,
		"unique_visitors":    uniqueVisitors,
		"geo":                geoStats,
		"clicks_per_visitor": clicksPerVisitor,
		"visitors_per_city":  visitorsPerCity,
	})
}
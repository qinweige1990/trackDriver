package logic

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"trackDriver/DB"
)

func StartJourney(c *gin.Context) {
	start, _ := c.GetQuery("start")
	end, _ := c.GetQuery("end")
	driver, _ := c.GetQuery("driver")
	DB.CreateJourney(start, end, driver)
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func EndJourney(c *gin.Context) {
	id, _ := c.GetQuery("id")
	idInt, _ := strconv.Atoi(id)
	DB.UpdateJourney(idInt)
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func GetEndJourney(c *gin.Context) {
	journeys := DB.GetEndJourney()
	c.JSON(200, gin.H{
		"journeys": journeys,
	})
}
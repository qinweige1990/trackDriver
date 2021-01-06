
package main

import (
	"github.com/gin-gonic/gin"
	"trackDriver/logic"
)

func main() {
	router := gin.Default()

	router.GET("/start", logic.StartJourney)
	router.GET("/end", logic.EndJourney)
	router.GET("/journeys", logic.GetEndJourney)

	router.Run()
}

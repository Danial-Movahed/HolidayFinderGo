package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type holiday struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func startApi() {
	router := gin.Default()
	router.GET("/holiday", getHoliday)

	router.Run("localhost:8888")
}

func main() {
	startApi()
}

func getHoliday(c *gin.Context) {
	day, err := c.GetQuery("day")
	if !err {
		fmt.Println(err)
	}
	month, err := c.GetQuery("month")
	if !err {
		fmt.Println(err)
	}
	year, err := c.GetQuery("year")
	if !err {
		fmt.Println(err)
	}
	hol := grpcClient(day, month, year)

	c.IndentedJSON(http.StatusOK, hol)

	// c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

package main

import (
	"YOU-ARE-BOMBED/utils"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

func home(c *gin.Context) {
	c.String(http.StatusOK, "You're BOMBED!!!")
}

func main() {
	Isgenerate := os.Args[1]
	if Isgenerate == "generate" {
		if os.Args[2] == "nested" {
			levels, err := strconv.Atoi(os.Args[3])
			utils.Check(err)
			utils.GenerateNest(levels)
		} else if os.Args[2] == "flat" {
			count, err := strconv.Atoi(os.Args[3])
			utils.Check(err)
			utils.GenerateFlat(count)
		}
	} else if Isgenerate == "host" {
		router := gin.Default()
		router.GET("/", home)
		router.Static("/static", "./bomb")
		router.Run("localhost:8080")
	}
}

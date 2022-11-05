package main

import (
	"ARE-YOU-BOMBED/utils"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func home(c *gin.Context) {
	c.String(http.StatusOK, "You're BOMBED!!!")
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	if len(os.Args) < 2 {
		usage := "- Generate flat zip bomb\nCommand: go run main.go generate flat <number of files>\n\n- Generate nested zip bomb\nCommand: go run main.go generate nested <depth>\n\n- Run the server\nCommand: go run main.go host"
		fmt.Println(usage)
		return
	}
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
		router.Run("localhost:" + os.Getenv("PORT"))
	}
}

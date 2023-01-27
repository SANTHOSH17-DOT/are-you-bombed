package main

import (
	"ARE-YOU-BOMBED/utils"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func home(c *fiber.Ctx) error {
	return c.SendString("You're BOMBED!!!")
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
		app := fiber.New()
		app.Get("/", home)
		app.Static("/static", "./bomb")
		log.Fatal(app.Listen(":" + os.Getenv("PORT")))
	}
}

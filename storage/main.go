package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

func main() {
	viper.GetViper().SetConfigFile(".env")
	viper.ReadInConfig()

	app := fiber.New()

	app.Get("/api/v1/storage/health", func(c *fiber.Ctx) error { return c.SendString("alive storage") })

	log.Fatal(app.Listen(fmt.Sprintf(":%d", viper.GetInt("HTTP_PORT"))))
}

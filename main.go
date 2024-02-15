package main

import (
	"fmt"
	"log"

	"github.com/JMjirapat/qrthrough-api/api"
	"github.com/JMjirapat/qrthrough-api/config"
	"github.com/JMjirapat/qrthrough-api/infrastructure"
	"github.com/gofiber/fiber/v2"
)

var cfg config.LocalConfig

func init() {
	// config
	config.InitConfig()
	cfg = config.Config

	// lineBot
	infrastructure.InitLineBot()

	// postgresql
	infrastructure.InitDB()
}

func main() {
	app := fiber.New()
	api.InitAPI(app)

	addr := getAddress()
	log.Printf("%v started at %v", cfg.Name, cfg.Port)
	if err := app.Listen(addr); err != nil {
		log.Fatal(err)
	}
}

func getAddress() string {
	addr := ":8000"
	if cfg.Port != "" {
		addr = fmt.Sprintf(":%v", cfg.Port)
	}
	return addr
}

package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/minmaxmar/bankapp/config"
	"github.com/minmaxmar/bankapp/database"
	"github.com/minmaxmar/bankapp/logger"
	"github.com/rs/zerolog/log"
)

func main() {

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to load configuration")
	}
	logger.InitLogger(cfg.LogLevel)
	log.Info().Msgf("Log level set to: %s", cfg.LogLevel)

	database.ConnectDb(cfg.DatabaseURL)
	app := fiber.New()
	setupRoutes(app)
	app.Listen(":3000")
}

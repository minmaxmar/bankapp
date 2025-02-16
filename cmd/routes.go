package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/minmaxmar/bankapp/handlers"
)

func setupRoutes(app *fiber.App) {

	app.Get("/facts", handlers.ListFacts)
	app.Post("/fact", handlers.CreateFact)

	app.Get("/banks", handlers.ListBanks)
	app.Post("/bank", handlers.CreateBank)

}

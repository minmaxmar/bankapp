package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/minmaxmar/bankapp/database"
	"github.com/minmaxmar/bankapp/models"
	"github.com/rs/zerolog/log"
)

func ListBanks(c *fiber.Ctx) error {
	log.Debug().Msg("ListBanks")
	banks := []models.Bank{}
	database.DB.Db.Find(&banks)

	return c.Status(200).JSON(banks)
}

func CreateBank(c *fiber.Ctx) error {

	bank := new(models.Bank)

	log.Debug().Str("Body", string(c.Body())).Msg("Request received")
	if err := c.BodyParser(bank); err != nil {
		log.Error().Err(err).Msg("Error parsing body")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	log.Debug().Msgf("Parsed bank: %+v\n", bank)
	database.DB.Db.Create(&bank)

	return c.Status(200).JSON(bank)
}

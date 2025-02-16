package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/minmaxmar/bankapp/database"
	"github.com/minmaxmar/bankapp/models"
	"github.com/rs/zerolog/log"
)

func ListFacts(c *fiber.Ctx) error {
	// return c.SendString("hi bankapp")
	facts := []models.Fact{}
	database.DB.Db.Find(&facts)

	return c.Status(200).JSON(facts)
}

func CreateFact(c *fiber.Ctx) error {

	fact := new(models.Fact)

	// fmt.Println("Raw request body:", string(c.Body()))
	log.Debug().Str("Body", string(c.Body())).Msg("Request received")

	if err := c.BodyParser(fact); err != nil {
		log.Error().Err(err).Msg("Error parsing body")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	// logger.GetLogger().Debug().Str("Result", fact).Msg("Parsed")
	log.Debug().Msgf("Parsed fact: %+v\n", fact)
	// fmt.Printf("Parsed fact: %+v\n", fact)

	database.DB.Db.Create(&fact)

	return c.Status(200).JSON(fact)
}

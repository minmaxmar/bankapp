package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/minmaxmar/bankapp/database"
	"github.com/minmaxmar/bankapp/models"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

func ListClients(c *fiber.Ctx) error {
	log.Debug().Msg("ListClients")
	clients := []models.Client{}
	database.DB.Db.Find(&clients)

	return c.Status(200).JSON(clients)
}

func CreateClient(c *fiber.Ctx) error {

	client := new(models.Client)

	log.Debug().Str("Body", string(c.Body())).Msg("Request received")
	if err := c.BodyParser(client); err != nil {
		log.Error().Err(err).Msg("Error parsing body")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	log.Debug().Msgf("Parsed client: %+v\n", client)

	if client.FirstName == "" || client.LastName == "" || client.Email == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "First name, last name, and email are required.",
		})
	}

	result := database.DB.Db.Create(&client)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": fmt.Sprintf("Error creating client: %s", result.Error.Error()),
		})
	}

	return c.Status(200).JSON(client)
}

func CreateBankClient(c *fiber.Ctx) error {

	req := new(models.CreateBankClientRequest)

	log.Debug().Str("Body", string(c.Body())).Msg("Request received")
	if err := c.BodyParser(req); err != nil {
		log.Error().Err(err).Msg("Error parsing body")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	log.Debug().Msgf("Parsed request: %+v\n", req)
	// validate
	if req.ClientEmail == "" || req.BankName == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Client email and Bank name are required.",
		})
	}
	// find client
	var client models.Client
	result := database.DB.Db.Where("email = ?", req.ClientEmail).First(&client)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": fmt.Sprintf("Client with email '%s' not found.", req.ClientEmail),
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": fmt.Sprintf("Error finding client: %s", result.Error.Error()),
		})
	}
	// find bank     TODO : move to models
	var bank models.Bank
	result = database.DB.Db.Where("name = ?", req.BankName).First(&bank)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": fmt.Sprintf("Bank with name '%s' not found.", req.BankName),
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": fmt.Sprintf("Error finding bank: %s", result.Error.Error()),
		})
	}
	// if exists
	var existingClientBank models.Bank
	exists := database.DB.Db.Model(&client).Association("Banks").Find(&existingClientBank, bank.ID)
	if exists == nil {
		if existingClientBank.ID == bank.ID {
			return c.Status(fiber.StatusConflict).JSON(fiber.Map{
				"message": fmt.Sprintf("Client with email '%s' is already a client of bank '%s'.", req.ClientEmail, req.BankName),
			})
		}
	}
	// update
	err := database.DB.Db.Model(&client).Association("Banks").Append(&bank)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": fmt.Sprintf("Error creating bank-client relationship: %s", err.Error()),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": fmt.Sprintf("Client '%s' successfully added to bank '%s'.", req.ClientEmail, req.BankName),
	})

}

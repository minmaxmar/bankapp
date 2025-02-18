package handlers

import (
	"fmt"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/minmaxmar/bankapp/database"
	"github.com/minmaxmar/bankapp/models"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

func CreateCard(c *fiber.Ctx) error {

	req := new(models.CreateCardRequest)

	log.Debug().Str("Body", string(c.Body())).Msg("Request received")
	if err := c.BodyParser(req); err != nil {
		log.Error().Err(err).Msg("Error parsing body")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	log.Debug().Msgf("Parsed request: %+v\n", req)
	// validate
	if req.CardNumber == "" || req.ExpiryDate == "" || req.ClientEmail == "" || req.BankName == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "card_number, expiry_date, client_email, bank_name are required.",
		})
	}
	// validate card - TODO 16-digit card number
	cardNumber, err := strconv.ParseInt(req.CardNumber, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid card number. Must be a number.",
		})
	}

	// validate expiry_date : MM/YY
	var parsedExpiryDate time.Time
	if len(req.ExpiryDate) == 5 && (req.ExpiryDate[2] == '/') {
		parsedExpiryDate, err = time.Parse("01/06", req.ExpiryDate)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Invalid expiry date format. Expected MM/YY.",
			})
		}
		req.ExpiryDate = fmt.Sprintf("%02d/%02d", parsedExpiryDate.Month(), parsedExpiryDate.Year()%100)
	} else {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid expiry date format. Expected MM/YY.",
		})
	}

	// find bank     TODO : move to models
	var bank models.Bank
	if err := database.DB.Db.Where("name = ?", req.BankName).First(&bank).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": fmt.Sprintf("Bank with name '%s' not found.", req.BankName),
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": fmt.Sprintf("Error finding bank: %s", err.Error()),
		})
	}

	// find client   TODO : move to models
	var client models.Client
	if err := database.DB.Db.Where("email = ?", req.ClientEmail).First(&client).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": fmt.Sprintf("Client with email '%s' not found.", req.ClientEmail),
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": fmt.Sprintf("Error finding client: %s", err.Error()),
		})
	}

	// Create Card
	card := models.Card{
		CardNumber: cardNumber,
		ExpiryDate: req.ExpiryDate,
		Total:      0,
		BankID:     bank.ID,
		ClientID:   client.ID,
	}

	result := database.DB.Db.Create(&card)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": fmt.Sprintf("Error creating card: %s", result.Error.Error()),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(card)
}

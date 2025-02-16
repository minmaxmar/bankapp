package models

import (
	"math/big"

	"gorm.io/gorm"
)

// dummy
type Fact struct {
	gorm.Model
	Question string `json:"question" gorm:"type:text;not null;default:null"`
	Answer   string `json:"answer" gorm:"text;not null;default:null"`
}

type Card struct {
	gorm.Model         // Adds ID, CreatedAt, UpdatedAt, DeletedAt fields
	CardNumber big.Int `json:"card_number" gorm:"type:varchar(20);unique;not null"` // Example: "1234-5678-9012-3456"
	ExpiryDate string  `json:"expiry_date" gorm:"type:varchar(10);not null"`        // Example: "12/25"

	BankID uint `json:"bank_id"`
	Bank   Bank `json:"bank"`

	ClientID uint   `json:"client_id"`
	Client   Client `json:"client"`
}

type Client struct {
	gorm.Model
	FirstName string `json:"first_name" gorm:"type:varchar(255);not null"`
	LastName  string `json:"last_name" gorm:"type:varchar(255);not null"`
	Email     string `json:"email" gorm:"type:varchar(255);unique;not null"`
	Banks     []Bank `gorm:"many2many:bank_clients;"`
	Cards     []Card `gorm:"foreignKey:ClientID"`
}

type Bank struct {
	gorm.Model

	Name    string   `json:"name" gorm:"type:varchar(255);not null"`
	Clients []Client `gorm:"many2many:bank_clients;"`
	Cards   []Card   `gorm:"foreignKey:BankID"`
}

package associates

import (
	"gorm.io/gorm"
)

type Associate struct {
	gorm.Model
	First  string `form:"name" validate:"required,min=2"`
	Middle string `form:"middle" validate:"required,min=2"`
	Last string `form:"last" validate:"required,min=2"`
	Email string `form:"email" gorm:"unique" validate:"required,email"`
	Cell string `form:"phone" gorm:"unique" validate:"required,phone"`
	Password string
	// Rg string `gorm:"unique"`
	// Cpf string `gorm:"unique"`
	// Street string
	// District string
	// City string
	// Cep string `form:"CEP" validate:"required,cepx"`
	// State string

	// Bank string
	// BankNumber string
	// Account string
	// Pix string `form:"PIX" gorm:"unique" validate:"required,pix"`
	// EmergencyName string `form:"name" validate:"required,min=2"`
	// EmergencyEmail string `form:"email" validate:"required,email"`
	// EmergencyCell string
}
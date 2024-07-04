package types

import (
	"time"
)

type TaxShelter = string

const (
	TAXABLE          TaxShelter = "TAXABLE"
	TRADITIONAL      TaxShelter = "TRADITIONAL"
	ROTH             TaxShelter = "ROTH"
	HSA              TaxShelter = "HSA"
	FIVE_TWENTY_NINE TaxShelter = "529"
)

type User struct {
	User_id      int       `json:"user_id,omitempty"`
	Email        string    `json:"email"`
	Enc_password string    `json:"-"`
	Created_at   time.Time `json:"created_at"`
	Updated_at   time.Time `json:"updated_at"`
}

type Account struct {
	Account_id  int        `json:"account_id,omitempty"`
	Name        string     `json:"name"`
	Description string     `json:"description,omitempty"`
	Tax_shelter TaxShelter `json:"shelter_type"`
	Institution string     `json:"institution"`
	Is_closed   bool       `json:"is_closed"`
	User_id     int        `json:"user_id"`
	Created_at  time.Time  `json:"created_at"`
	Updated_at  time.Time  `json:"updated_at"`
}

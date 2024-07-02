package types

import "time"

type TaxShelter = string

const (
	TAXABLE          TaxShelter = "TAXABLE"
	TRADITIONAL      TaxShelter = "TRADITIONAL"
	ROTH             TaxShelter = "ROTH"
	HSA              TaxShelter = "HSA"
	FIVE_TWENTY_NINE TaxShelter = "529"
)

type User struct {
	User_id      string    `json:"user_id"`
	Email        string    `json:"email"`
	Enc_password string    `json:"-"`
	Created_at   time.Time `json:"created_at"`
	Updated_at   time.Time `json:"updated_at"`
}

type Account struct {
	Account_id   string     `json:"account_id"`
	Name         string     `json:"name"`
	Description  string     `json:"description"`
	Shelter_type TaxShelter `json:"shelter_type"`
	Institution  string     `json:"institution"`
	User_id      string     `json:"user_id"`
	Created_at   time.Time  `json:"created_at"`
	Updated_at   time.Time  `json:"updated_at"`
}

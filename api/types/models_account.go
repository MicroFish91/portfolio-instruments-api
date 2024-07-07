package types

import (
	"time"
)

type TaxShelter = string

const (
	TAXABLE          TaxShelter = "taxable"
	TRADITIONAL      TaxShelter = "traditional"
	ROTH             TaxShelter = "roth"
	HSA              TaxShelter = "hsa"
	FIVE_TWENTY_NINE TaxShelter = "529"
)

type Account struct {
	Account_id    int        `json:"account_id,omitempty"`
	Name          string     `json:"name"`
	Description   string     `json:"description,omitempty"`
	Tax_shelter   TaxShelter `json:"shelter_type"`
	Institution   string     `json:"institution"`
	Is_deprecated bool       `json:"is_deprecated"`
	User_id       int        `json:"user_id"`
	Created_at    time.Time  `json:"created_at"`
	Updated_at    time.Time  `json:"updated_at"`
}

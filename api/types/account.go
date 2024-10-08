package types

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v3"
)

type TaxShelter = string

const (
	TAXABLE          TaxShelter = "TAXABLE"
	TRADITIONAL      TaxShelter = "TRADITIONAL"
	ROTH             TaxShelter = "ROTH"
	HSA              TaxShelter = "HSA"
	FIVE_TWENTY_NINE TaxShelter = "529"
)

var ValidTaxShelters = []interface{}{
	TAXABLE, TRADITIONAL, ROTH, HSA, FIVE_TWENTY_NINE,
}

type Account struct {
	Account_id    int        `json:"account_id,omitempty"`
	Name          string     `json:"name"`
	Description   string     `json:"description,omitempty"`
	Tax_shelter   TaxShelter `json:"tax_shelter"`
	Institution   string     `json:"institution"`
	Is_deprecated bool       `json:"is_deprecated"`
	User_id       int        `json:"user_id"`
	Created_at    time.Time  `json:"created_at"`
	Updated_at    time.Time  `json:"updated_at"`
}

type AccountHandler interface {
	CreateAccount(fiber.Ctx) error
	GetAccounts(fiber.Ctx) error
	GetAccountById(fiber.Ctx) error
	UpdateAccount(fiber.Ctx) error
	DeleteAccount(fiber.Ctx) error
}

type AccountStore interface {
	CreateAccount(context.Context, *Account) (Account, error)
	GetAccounts(ctx context.Context, userId int, options *GetAccountsStoreOptions) ([]Account, PaginationMetadata, error)
	GetAccountById(ctx context.Context, userId, accountId int) (Account, error)
	GetAccountByName(ctx context.Context, name string, userId int) (Account, error)
	UpdateAccount(context.Context, *Account) (Account, error)
	DeleteAccount(ctx context.Context, userId, accountId int) (Account, error)
}

type GetAccountsStoreOptions struct {
	AccountIds    []int
	TaxShelter    TaxShelter
	Institution   string
	Is_deprecated string
	Current_page  int
	Page_size     int
}

// ---- Account Response Types ----

type CreateAccountResponse struct {
	Data struct {
		Account Account `json:"account"`
	} `json:"data"`
	Error string `json:"error"`
}

type GetAccountsResponse struct {
	Data struct {
		Accounts   []Account          `json:"accounts"`
		Pagination PaginationMetadata `json:"pagination"`
	} `json:"data"`
	Error string `json:"error"`
}

type GetAccountResponse struct {
	Data struct {
		Account Account `json:"account"`
	} `json:"data"`
	Error string `json:"error"`
}

type UpdateAccountResponse struct {
	Data struct {
		Account Account `json:"account"`
	} `json:"data"`
	Error string `json:"error"`
}

type DeleteAccountResponse struct {
	Data struct {
		Message string  `json:"message"`
		Account Account `json:"account"`
	} `json:"data"`
	Error string `json:"error"`
}

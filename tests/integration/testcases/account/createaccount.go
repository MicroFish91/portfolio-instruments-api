package account

import (
	"testing"

	"github.com/MicroFish91/portfolio-instruments-api/api/services/account"
	"github.com/MicroFish91/portfolio-instruments-api/tests/integration/shared"
	"github.com/MicroFish91/portfolio-instruments-api/tests/utils"
	"github.com/gofiber/fiber/v3"
)

func GetCreateAccountTests(t *testing.T, userId int, email string) []shared.TestCase {
	tok401, _, err := utils.Generate40xTokens(userId, email)
	if err != nil {
		t.Fatal(err)
	}

	return []shared.TestCase{
		// ---- 201 ----
		{
			Title: "201 Std",
			Payload: account.CreateAccountPayload{
				Name:        "VAN001",
				Description: "Vanguard Taxable Brokerage",
				Tax_shelter: "TAXABLE",
				Institution: "Vanguard",
			},
			ExpectedStatusCode: fiber.StatusCreated,
		},
		{
			Title: "201 No Description",
			Payload: account.CreateAccountPayload{
				Name:          "VAN002",
				Tax_shelter:   "TAXABLE",
				Institution:   "Vanguard",
				Is_deprecated: false,
			},
			ExpectedStatusCode: fiber.StatusCreated,
		},
		{
			Title: "201 Tax Shelter",
			Payload: account.CreateAccountPayload{
				Name:        "VAN003",
				Description: "Vanguard Taxable Brokerage",
				Tax_shelter: "ROTH",
				Institution: "Fidelity",
			},
			ExpectedStatusCode: fiber.StatusCreated,
		},
		{
			Title: "201 Institution",
			Payload: account.CreateAccountPayload{
				Name:        "VAN004",
				Description: "Vanguard Taxable Brokerage",
				Tax_shelter: "TAXABLE",
				Institution: "Fidelity",
			},
			ExpectedStatusCode: fiber.StatusCreated,
		},
		{
			Title: "201 Is Deprecated",
			Payload: account.CreateAccountPayload{
				Name:          "VAN005",
				Description:   "Vanguard Taxable Brokerage",
				Tax_shelter:   "TAXABLE",
				Institution:   "Vanguard",
				Is_deprecated: true,
			},
			ExpectedStatusCode: fiber.StatusCreated,
		},

		{
			Title: "401",
			Payload: account.CreateAccountPayload{
				Name:        "VAN006",
				Tax_shelter: "TAXABLE",
				Institution: "Vanguard",
			},
			ReplacementToken:   tok401,
			ExpectedStatusCode: fiber.StatusUnauthorized,
		},
		{
			Title: "409",
			Payload: account.CreateAccountPayload{
				Name:        "vAn001",
				Tax_shelter: "TAXABLE",
				Institution: "Vanguard",
			},
			ExpectedStatusCode: fiber.StatusConflict,
		},

		// ---- 400 ----
		{
			Title: "400 No Name",
			Payload: account.CreateAccountPayload{
				Tax_shelter: "TAXABLE",
				Institution: "Vanguard",
			},
			ExpectedStatusCode: fiber.StatusBadRequest,
		},
		{
			Title: "400 Bad Name",
			Payload: map[string]any{
				"Name":        1,
				"Tax_shelter": "TAXABLE",
				"Institution": "Vanguard",
			},
			ExpectedStatusCode: fiber.StatusBadRequest,
		},
		{
			Title: "400 No Tax Shelter",
			Payload: account.CreateAccountPayload{
				Name:        "VAN009",
				Tax_shelter: "",
				Institution: "Vanguard",
			},
			ExpectedStatusCode: fiber.StatusBadRequest,
		},
		{
			Title: "400 Wrong Tax Shelter 1",
			Payload: account.CreateAccountPayload{
				Name:        "VAN010",
				Tax_shelter: "NO_TAX",
				Institution: "Vanguard",
			},
			ExpectedStatusCode: fiber.StatusBadRequest,
		},
		{
			Title: "400 Wrong Tax Shelter 2",
			Payload: map[string]any{
				"Name":        "VAN011",
				"Tax_shelter": 5,
				"Institution": "Vanguard",
			},
			ExpectedStatusCode: fiber.StatusBadRequest,
		},
		{
			Title: "400 No Institution 1",
			Payload: account.CreateAccountPayload{
				Name:        "VAN012",
				Tax_shelter: "TAXABLE",
			},
			ExpectedStatusCode: fiber.StatusBadRequest,
		},
		{
			Title: "400 No Institution 2",
			Payload: map[string]any{
				"Name":        "VAN013",
				"Tax_shelter": "TAXABLE",
				"Institution": true,
			},
			ExpectedStatusCode: fiber.StatusBadRequest,
		},
	}
}

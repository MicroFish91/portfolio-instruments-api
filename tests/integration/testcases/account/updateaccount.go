package account

import (
	"testing"

	"github.com/MicroFish91/portfolio-instruments-api/api/services/account"
	"github.com/MicroFish91/portfolio-instruments-api/tests/integration/shared"
	"github.com/MicroFish91/portfolio-instruments-api/tests/utils"
	"github.com/gofiber/fiber/v3"
)

func GetUpdateAccountTests(t *testing.T, accountId int, userId int, email string) []shared.PutTestCase {
	tok401, _, err := utils.Generate40xTokens(userId, email)
	if err != nil {
		t.Fatal(err)
	}

	return []shared.PutTestCase{
		// ---- 200 ----
		{
			Title:       "200 Std",
			ParameterId: accountId,
			Payload: account.UpdateAccountPayload{
				Name:        "VAN031",
				Description: "Vanguard Taxable Brokerage",
				Tax_shelter: "TAXABLE",
				Institution: "Vanguard",
			},
			ExpectedStatusCode: fiber.StatusOK,
		},
		{
			Title:       "200 No Description",
			ParameterId: accountId,
			Payload: account.UpdateAccountPayload{
				Name:          "VAN032",
				Tax_shelter:   "TAXABLE",
				Institution:   "Vanguard",
				Is_deprecated: false,
			},
			ExpectedStatusCode: fiber.StatusOK,
		},
		{
			Title:       "200 Tax Shelter",
			ParameterId: accountId,
			Payload: account.UpdateAccountPayload{
				Name:        "VAN033",
				Description: "Vanguard Taxable Brokerage",
				Tax_shelter: "ROTH",
				Institution: "Fidelity",
			},
			ExpectedStatusCode: fiber.StatusOK,
		},
		{
			Title:       "200 Institution",
			ParameterId: accountId,
			Payload: account.UpdateAccountPayload{
				Name:        "VAN034",
				Description: "Vanguard Taxable Brokerage",
				Tax_shelter: "TAXABLE",
				Institution: "Fidelity",
			},
			ExpectedStatusCode: fiber.StatusOK,
		},
		{
			Title:       "200 Is Deprecated",
			ParameterId: accountId,
			Payload: account.UpdateAccountPayload{
				Name:          "VAN035",
				Description:   "Vanguard Taxable Brokerage",
				Tax_shelter:   "TAXABLE",
				Institution:   "Vanguard",
				Is_deprecated: true,
			},
			ExpectedStatusCode: fiber.StatusOK,
		},

		{
			Title:       "401",
			ParameterId: accountId,
			Payload: account.UpdateAccountPayload{
				Name:        "VAN036",
				Tax_shelter: "TAXABLE",
				Institution: "Vanguard",
			},
			ReplacementToken:   tok401,
			ExpectedStatusCode: fiber.StatusUnauthorized,
		},
		{
			Title:       "404",
			ParameterId: 9999,
			Payload: account.UpdateAccountPayload{
				Name:        "VAN037",
				Tax_shelter: "TAXABLE",
				Institution: "Vanguard",
			},
			ExpectedStatusCode: fiber.StatusNotFound,
		},

		// ---- 400 ----
		{
			Title:       "400 No Name",
			ParameterId: accountId,
			Payload: account.UpdateAccountPayload{
				Tax_shelter: "TAXABLE",
				Institution: "Vanguard",
			},
			ExpectedStatusCode: fiber.StatusBadRequest,
		},
		{
			Title:       "400 Bad Name",
			ParameterId: accountId,
			Payload: map[string]any{
				"Name":        1,
				"Tax_shelter": "TAXABLE",
				"Institution": "Vanguard",
			},
			ExpectedStatusCode: fiber.StatusBadRequest,
		},
		{
			Title:       "400 No Tax Shelter",
			ParameterId: accountId,
			Payload: account.CreateAccountPayload{
				Name:        "VAN009",
				Tax_shelter: "",
				Institution: "Vanguard",
			},
			ExpectedStatusCode: fiber.StatusBadRequest,
		},
		{
			Title:       "400 Wrong Tax Shelter 1",
			ParameterId: accountId,
			Payload: account.CreateAccountPayload{
				Name:        "VAN010",
				Tax_shelter: "NO_TAX",
				Institution: "Vanguard",
			},
			ExpectedStatusCode: fiber.StatusBadRequest,
		},
		{
			Title:       "400 Wrong Tax Shelter 2",
			ParameterId: accountId,
			Payload: map[string]any{
				"Name":        "VAN011",
				"Tax_shelter": 5,
				"Institution": "Vanguard",
			},
			ExpectedStatusCode: fiber.StatusBadRequest,
		},
		{
			Title:       "400 No Institution 1",
			ParameterId: accountId,
			Payload: account.CreateAccountPayload{
				Name:        "VAN012",
				Tax_shelter: "TAXABLE",
			},
			ExpectedStatusCode: fiber.StatusBadRequest,
		},
		{
			Title:       "400 No Institution 2",
			ParameterId: accountId,
			Payload: map[string]any{
				"Name":        "VAN013",
				"Tax_shelter": "TAXABLE",
				"Institution": true,
			},
			ExpectedStatusCode: fiber.StatusBadRequest,
		},
	}
}

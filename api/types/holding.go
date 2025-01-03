package types

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v3"
)

type AssetCategory = string

const (
	CASH        AssetCategory = "CASH"
	BILLS       AssetCategory = "BILLS"
	STB         AssetCategory = "STB"
	ITB         AssetCategory = "ITB"
	LTB         AssetCategory = "LTB"
	COMMODITIES AssetCategory = "COMMODITIES"
	GOLD        AssetCategory = "GOLD"
	REITS       AssetCategory = "REITS"
	TSM         AssetCategory = "TSM"
	DLCB        AssetCategory = "DLCB"
	DLCG        AssetCategory = "DLCG"
	DLCV        AssetCategory = "DLCV"
	DMCB        AssetCategory = "DMCB"
	DMCG        AssetCategory = "DMCG"
	DMCV        AssetCategory = "DMCV"
	DSCG        AssetCategory = "DSCG"
	DSCB        AssetCategory = "DSCB"
	DSCV        AssetCategory = "DSCV"
	ILCB        AssetCategory = "ILCB"
	ILCG        AssetCategory = "ILCG"
	ILCV        AssetCategory = "ILCV"
	IMCB        AssetCategory = "IMCB"
	IMCG        AssetCategory = "IMCG"
	IMCV        AssetCategory = "IMCV"
	ISCB        AssetCategory = "ISCB"
	ISCG        AssetCategory = "ISCG"
	ISCV        AssetCategory = "ISCV"
	CRYPTO      AssetCategory = "CRYPTO"
	OTHER       AssetCategory = "OTHER"
)

var ValidAssetCategories = []interface{}{
	CASH, BILLS, STB, ITB, LTB, COMMODITIES, GOLD, REITS, TSM,
	DLCB, DLCG, DLCV, DMCB, DMCG, DMCV, DSCG, DSCB, DSCV,
	ILCB, ILCG, ILCV, IMCB, IMCG, IMCV, ISCB, ISCG, ISCV, CRYPTO, OTHER,
}

type Holding struct {
	Holding_id        int           `json:"holding_id,omitempty"`
	Name              string        `json:"name"`
	Ticker            string        `json:"ticker,omitempty"`
	Asset_category    AssetCategory `json:"asset_category"`
	Expense_ratio_pct float32       `json:"expense_ratio_pct,omitempty"`
	Maturation_date   string        `json:"maturation_date,omitempty"`
	Interest_rate_pct float32       `json:"interest_rate_pct,omitempty"`
	Is_deprecated     bool          `json:"is_deprecated"`
	User_id           int           `json:"user_id"`
	Created_at        time.Time     `json:"created_at"`
	Updated_at        time.Time     `json:"updated_at"`
}

type HoldingHandler interface {
	CreateHolding(fiber.Ctx) error
	GetHoldings(fiber.Ctx) error
	GetHoldingById(fiber.Ctx) error
	UpdateHolding(fiber.Ctx) error
	DeleteHolding(fiber.Ctx) error
}

type HoldingStore interface {
	CreateHolding(context.Context, *Holding) (Holding, error)
	GetHoldings(ctx context.Context, userId int, options *GetHoldingsStoreOptions) ([]Holding, PaginationMetadata, error)
	GetHoldingById(ctx context.Context, userId, holdingId int) (Holding, error)
	GetHoldingByName(ctx context.Context, name string, userId int) (Holding, error)
	GetHoldingByTicker(ctx context.Context, ticker string, userId int) (Holding, error)
	UpdateHolding(context.Context, *Holding) (Holding, error)
	DeleteHolding(ctx context.Context, userId, holdingId int) (Holding, error)
}

type GetHoldingsStoreOptions struct {
	Holding_ids              []int
	Ticker                   string
	Asset_category           AssetCategory
	Has_maturation_remaining string // String representation of a bool
	Is_deprecated            string // String representation of a bool
	Current_page             int
	Page_size                int
}

// ---- Holding Response Types ----

type CreateHoldingResponse struct {
	Data struct {
		Holding Holding `json:"holding"`
	} `json:"data"`
	Error string `json:"error"`
}

type GetHoldingsResponse struct {
	Data struct {
		Holdings   []Holding          `json:"holdings"`
		Pagination PaginationMetadata `json:"pagination"`
	} `json:"data"`
	Error string `json:"error"`
}

type GetHoldingResponse struct {
	Data struct {
		Holding Holding `json:"holding"`
	} `json:"data"`
	Error string `json:"error"`
}

type UpdateHoldingResponse struct {
	Data struct {
		Holding Holding `json:"holding"`
	} `json:"data"`
	Error string `json:"error"`
}

type DeleteHoldingResponse struct {
	Data struct {
		Message string  `json:"message"`
		Holding Holding `json:"holding"`
	} `json:"data"`
	Error string `json:"error"`
}

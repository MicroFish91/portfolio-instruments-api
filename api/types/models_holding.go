package types

import "time"

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
)

type Holding struct {
	Holding_id     int           `json:"holding_id,omitempty"`
	Name           string        `json:"name"`
	Ticker         string        `json:"ticker"`
	Asset_category AssetCategory `json:"asset_category"`
	Expense_ratio  float32       `json:"expense_ratio"`
	Is_deprecated  bool          `json:"is_deprecated"`
	User_id        int           `json:"user_id"`
	Created_at     time.Time     `json:"created_at"`
	Updated_at     time.Time     `json:"updated_at"`
}

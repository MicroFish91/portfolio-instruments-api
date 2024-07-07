package types

import "time"

type AssetCategory = string

const (
	CASH        AssetCategory = "cash"
	BILLS       AssetCategory = "bills"
	STB         AssetCategory = "stb"
	ITB         AssetCategory = "itb"
	LTB         AssetCategory = "ltb"
	COMMODITIES AssetCategory = "commodities"
	GOLD        AssetCategory = "gold"
	REITS       AssetCategory = "reits"
	TSM         AssetCategory = "tsm"
	DLCB        AssetCategory = "dlcb"
	DLCG        AssetCategory = "dlcg"
	DLCV        AssetCategory = "dlcv"
	DMCB        AssetCategory = "dmcb"
	DMCG        AssetCategory = "dmcg"
	DMCV        AssetCategory = "dmcv"
	DSCG        AssetCategory = "dscg"
	DSCB        AssetCategory = "dscb"
	DSCV        AssetCategory = "dscv"
	ILCB        AssetCategory = "ilcb"
	ILCG        AssetCategory = "ilcg"
	ILCV        AssetCategory = "ilcv"
	IMCB        AssetCategory = "imcb"
	IMCG        AssetCategory = "imcg"
	IMCV        AssetCategory = "imcv"
	ISCB        AssetCategory = "iscb"
	ISCG        AssetCategory = "iscg"
	ISCV        AssetCategory = "iscv"
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

package snapshot

import (
	"errors"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type TallyCategory string

const (
	BY_ACCOUNT_NAME           TallyCategory = "ACCOUNT_NAME"
	BY_ACCOUNT_INSTITUTION    TallyCategory = "ACCOUNT_INSTITUTION"
	BY_TAX_SHELTER            TallyCategory = "TAX_SHELTER"
	BY_ASSET_CATEGORY         TallyCategory = "ASSET_CATEGORY"
	BY_WEIGHTED_EXPENSE_RATIO TallyCategory = "EXPENSE_RATIO"
)

type GetSnapshotByIdQuery struct {
	Tally_by string `json:"tally_by"` // Easier to analyze this as a string so we don't have to worry about setting up a reflection case for this as an enum
}

type GetSnapshotByIdParams struct {
	Id int `json:"id"`
}

func (p GetSnapshotByIdParams) Validate() error {
	return validation.ValidateStruct(&p,
		validation.Field(&p.Id, validation.Min(1)),
	)
}

func (q GetSnapshotByIdQuery) Validate() error {
	switch {
	case q.Tally_by == "":
		break
	case TallyCategory(q.Tally_by) == BY_ACCOUNT_NAME:
		break
	case TallyCategory(q.Tally_by) == BY_ACCOUNT_INSTITUTION:
		break
	case TallyCategory(q.Tally_by) == BY_TAX_SHELTER:
		break
	case TallyCategory(q.Tally_by) == BY_ASSET_CATEGORY:
		break
	case TallyCategory(q.Tally_by) == BY_WEIGHTED_EXPENSE_RATIO:
		break
	default:
		return errors.New("provide a valid tally_by category in all caps")
	}

	return nil
}

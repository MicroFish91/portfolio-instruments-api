package snapshot

import (
	"errors"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type GroupByCategory string

const (
	BY_ACCOUNT_NAME        GroupByCategory = "ACCOUNT_NAME"
	BY_ACCOUNT_INSTITUTION GroupByCategory = "ACCOUNT_INSTITUTION"
	BY_TAX_SHELTER         GroupByCategory = "TAX_SHELTER"
	BY_ASSET_CATEGORY      GroupByCategory = "ASSET_CATEGORY"
)

type GetSnapshotByIdQuery struct {
	Group_by string `json:"group_by"` // Easier to analyze this as a string so we don't have to worry about setting up a reflection case for this as an enum
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
	case q.Group_by == "":
		break
	case GroupByCategory(q.Group_by) == BY_ACCOUNT_NAME:
		break
	case GroupByCategory(q.Group_by) == BY_ACCOUNT_INSTITUTION:
		break
	case GroupByCategory(q.Group_by) == BY_TAX_SHELTER:
		break
	case GroupByCategory(q.Group_by) == BY_ASSET_CATEGORY:
		break
	default:
		return errors.New("provide a valid group_by category in all caps")
	}

	return nil
}

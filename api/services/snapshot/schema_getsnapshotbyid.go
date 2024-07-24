package snapshot

import validation "github.com/go-ozzo/ozzo-validation/v4"

type TallyCategory string

const (
	BY_ACCOUNT        TallyCategory = "BY_ACCOUNT"
	BY_TAX_SHELTER    TallyCategory = "BY_TAX_SHELTER"
	BY_ASSET_CATEGORY TallyCategory = "BY_ASSET_CATEGORY"
)

var validTallyCategories = []any{BY_ACCOUNT, BY_TAX_SHELTER, BY_ASSET_CATEGORY}

type GetSnapshotByIdQuery struct {
	Tally_by TallyCategory `json:"tally_by"`
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
	return validation.ValidateStruct(&q,
		validation.Field(&q.Tally_by, validation.In(validTallyCategories...)),
	)
}

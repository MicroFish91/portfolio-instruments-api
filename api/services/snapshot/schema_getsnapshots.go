package snapshot

import (
	"errors"
	"regexp"

	"github.com/MicroFish91/portfolio-instruments-api/api/types"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type GetSnapshotsQuery struct {
	Snap_ids        []int  `json:"snap_ids"`
	Snap_date_lower string `json:"snap_date_lower"`
	Snap_date_upper string `json:"snap_date_upper"`
	Order_date_by   string `json:"order_date"`

	types.PaginationQuery
}

func (q GetSnapshotsQuery) Validate() error {
	if err := q.PaginationQuery.Validate(); err != nil {
		return err
	}

	date := regexp.MustCompile(`^\d{2}/\d{2}/\d{4}$`)

	if q.Snap_date_lower != "" && !date.Match([]byte(q.Snap_date_lower)) {
		return errors.New("snap_date_lower must follow string format mm/dd/yyyy")
	}

	if q.Snap_date_upper != "" && !date.Match([]byte(q.Snap_date_upper)) {
		return errors.New("snap_date_upper must follow string format mm/dd/yyyy")
	}

	return validation.ValidateStruct(&q,
		validation.Field(&q.Snap_ids),
		validation.Field(&q.Snap_date_lower),
		validation.Field(&q.Snap_date_upper),
		validation.Field(&q.Order_date_by, validation.In("ASC", "DESC").Error(`value may be either "ASC" or "DESC"`)),
	)
}

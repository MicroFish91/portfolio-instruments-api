package snapshotvalue

import (
	"errors"
	"fmt"
	"strings"

	validation "github.com/go-ozzo/ozzo-validation"
)

type UpdateSnapshotValuePayload struct {
	Account_id     int     `json:"account_id"`
	Holding_id     int     `json:"holding_id"`
	Total          float64 `json:"total"`
	Skip_rebalance bool    `json:"skip_rebalance"`
}

func (p UpdateSnapshotValuePayload) Validate() error {
	if err := validation.ValidateStruct(&p,
		validation.Field(&p.Account_id, validation.Required, validation.Min(1)),
		validation.Field(&p.Holding_id, validation.Required, validation.Min(1)),
		validation.Field(&p.Total, validation.Required),
		validation.Field(&p.Skip_rebalance),
	); err != nil {
		return err
	}

	totalString := fmt.Sprintf("%v", p.Total)
	totalComp := strings.Split(totalString, ".")

	if len(totalComp) != 2 || len(totalComp[1]) != 2 {
		return errors.New("the total should represent a dollar value format to 2 decimals")
	}
	return nil
}

type UpdateSnapshotValueParams struct {
	Snap_id     int `json:"snap_id"`
	Snap_val_id int `json:"snap_val_id"`
}

func (p UpdateSnapshotValueParams) Validate() error {
	return validation.ValidateStruct(&p,
		validation.Field(&p.Snap_id, validation.Required, validation.Min(1)),
		validation.Field(&p.Snap_val_id, validation.Required, validation.Min(1)),
	)
}

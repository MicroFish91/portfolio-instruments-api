package utils

import (
	"time"
)

// Calculate the target date by specifying years and months from today's date
// Returns in the form "mm/dd/yyyy"
func Calc_target_date(years int, months int) string {
	now := time.Now()
	target := now.AddDate(years, months, 0)
	return target.Format("01/02/2006")
}

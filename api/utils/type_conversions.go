package utils

import "database/sql"

func ConvertIntSliceToAny(intSlice []int) []any {
	anySlice := make([]any, len(intSlice))
	for i, v := range intSlice {
		anySlice[i] = v
	}
	return anySlice
}

func ConvertNullIntToInt(nullInt sql.NullInt64) int {
	if nullInt.Valid {
		return int(nullInt.Int64)
	} else {
		return 0
	}
}

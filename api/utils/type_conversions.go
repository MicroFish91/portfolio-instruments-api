package utils

func IntSliceToAnySlice(intSlice []int) []any {
	anySlice := make([]any, len(intSlice))
	for i, v := range intSlice {
		anySlice[i] = v
	}
	return anySlice
}

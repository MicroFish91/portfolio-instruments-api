package utils

import "fmt"

var e_idx int = 0

func GetRotatingEmail() string {
	e_idx += 1
	return fmt.Sprintf("test_user%d@gmail.com", e_idx)
}

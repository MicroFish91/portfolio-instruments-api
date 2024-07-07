package querybuilder

import "strings"

func FillWithEmptyPositionals(length int) string {
	p := make([]string, length)
	for i := range p {
		p[i] = "$x"
	}
	return strings.Join(p, ", ")
}

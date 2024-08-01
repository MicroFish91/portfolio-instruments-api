package querybuilder

import (
	"fmt"
	"strconv"
)

type PgxQueryBuilder struct {
	Query            string
	QueryParams      []any
	positionalParams int
}

func NewPgxQueryBuilder() *PgxQueryBuilder {
	return &PgxQueryBuilder{}
}

// Appends a new SQL query raw without any positional parameters
func (q *PgxQueryBuilder) AddQuery(s string) {
	q.Query = fmt.Sprintf("%s\n%s", q.Query, s)
}

// Todo: Add better error handling

// Appends a new SQL query with positional parameters
// to the existing query in PgxQueryBuilder. It first replaces placeholders
// in the format "$x" with incrementing positional parameters (e.g., "$1", "$2", etc.)
func (q *PgxQueryBuilder) AddQueryWithPositionals(query string, values []any) error {
	query, n := q.replaceWithIncrementingPositionals(query)
	if n != len(values) {
		return fmt.Errorf("internal: found %d positional params, but only %d matching values", n, len(values))
	}

	q.QueryParams = append(q.QueryParams, values...)
	q.Query = fmt.Sprintf("%s\n%s", q.Query, query)
	return nil
}

// Before: "Hello $x $x"
// After: "Hello $1 $2"
func (q *PgxQueryBuilder) replaceWithIncrementingPositionals(query string) (newQuery string, incremented int) {
	var (
		before             int = q.positionalParams
		current, oneBefore rune
	)

	queryRunes := []rune(query)
	for c := 0; c < len(queryRunes); c += 1 {
		char := queryRunes[c]
		oneBefore = current
		current = char

		if oneBefore == '$' && current == 'x' {
			q.positionalParams += 1

			f := string(queryRunes[:c])
			m := fmt.Sprintf("%d", q.positionalParams)

			// Avoid an out of bounds error if c is at the last indexed position
			var e string
			if len(queryRunes)-1 == c {
				e = ""
			} else {
				e = string(queryRunes[c+1:])
			}

			query = fmt.Sprintf("%s%s%s", f, m, e)
			queryRunes = []rune(query)

			// We have to increment the c an extra amount based on the number of new characters we just added
			d := numDigits(q.positionalParams)
			if d > 1 {
				c += d - 1
			}
		}
	}

	return query, q.positionalParams - before
}

func numDigits(n int) int {
	numstring := strconv.Itoa(n)
	return len(numstring)
}

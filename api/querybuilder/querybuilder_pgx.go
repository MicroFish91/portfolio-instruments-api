package querybuilder

import (
	"fmt"
	"unicode/utf8"
)

type PgxQueryBuilder struct {
	Query            string
	QueryParams      []any
	positionalParams int
}

func NewPgxQueryBuilder() *PgxQueryBuilder {
	return &PgxQueryBuilder{}
}

func (q *PgxQueryBuilder) AddQuery(s string) {
	q.Query = fmt.Sprintf("%s\n%s", q.Query, s)
}

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

	for c, char := range query {
		oneBefore = current
		current = char

		if oneBefore == '$' && current == 'x' {
			q.positionalParams += 1

			f := query[:c]
			m := fmt.Sprintf("%d", q.positionalParams)

			// Avoid an out of bounds error if c is at the last indexed position
			var e string
			if utf8.RuneCountInString(query)-1 == c {
				e = ""
			}
			e = query[c+1:]

			query = fmt.Sprintf("%s%s%s", f, m, e)
		}
	}

	return query, q.positionalParams - before
}

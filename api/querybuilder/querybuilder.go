package querybuilder

type QueryBuilder interface {
	AddQuery(string)
	AddQueryWithPositionals(string, []any) error
}

package querybuilder

type QueryBuilder interface {
	AddRaw(string)
	AddWhere(string, []any) error
}

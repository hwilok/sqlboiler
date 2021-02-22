{{- $alias := .Aliases.Table .Table.Name}}

func (q {{$alias.DownSingular}}Query) Expr() string {
	s, _ := queries.BuildQuery(q.Query)
	return s[:len(s)-1]
}

func (q {{$alias.DownSingular}}Query) ExprArgs() (string, []interface{}) {
    queries.DisableIndexPlaceholders(q.Query)
    s, args := queries.BuildQuery(q.Query)
    return s[:len(s)-1], args
}

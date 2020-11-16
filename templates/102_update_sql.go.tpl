{{- $alias := .Aliases.Table .Table.Name -}}
{{- $schemaTable := .Table.Name | .SchemaTable}}

// UpdateSQL updates rows with raw SET
func (q {{$alias.DownSingular}}Query) UpdateSQL({{if .NoContext}}exec boil.Executor{{else}}ctx context.Context, exec boil.ContextExecutor{{end}}, set string, args ...interface{}) {{if .NoRowsAffected}}error{{else}}(int64, error){{end -}} {
    where, whereArgs := queries.WhereClause(q.Query, 1)

	sql := fmt.Sprintf("UPDATE {{$schemaTable}} SET %s %s", set, where)

	if len(whereArgs) != 0 {
        args = append(args, whereArgs...)
    }

	{{if .NoRowsAffected -}}
		{{if .NoContext -}}
	_, err := exec.Exec(sql, args...)
		{{else -}}
	_, err := exec.ExecContext(ctx, sql, args...)
		{{end -}}
	{{else -}}
		{{if .NoContext -}}
	result, err := exec.Exec(sql, args...)
		{{else -}}
	result, err := exec.ExecContext(ctx, sql, args...)
		{{end -}}
	{{end -}}
	if err != nil {
		return {{if not .NoRowsAffected}}0, {{end -}} errors.Wrap(err, "{{.PkgName}}: unable to set update for {{.Table.Name}}")
	}

	{{if not .NoRowsAffected -}}
	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "{{.PkgName}}: unable to retrieve rows affected for {{.Table.Name}}")
	}

	{{end -}}

	return {{if not .NoRowsAffected}}rowsAff, {{end -}} nil
}
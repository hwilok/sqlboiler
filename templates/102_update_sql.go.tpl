{{- $alias := .Aliases.Table .Table.Name -}}
{{- $schemaTable := .Table.Name | .SchemaTable}}

// UpdateSQL updates rows with raw SET
func (q {{$alias.DownSingular}}Query) UpdateSQL(set string, {{if .NoContext}}exec boil.Executor{{else}}ctx context.Context, exec boil.ContextExecutor{{end}}) {{if .NoRowsAffected}}error{{else}}(int64, error){{end -}} {
    where, whereArgs := queries.WhereClause(q.Query, 1)

	var sql string
	if where == "" {
		sql = fmt.Sprintf("UPDATE {{$schemaTable}} SET %s", set)
	}else {
		sql = fmt.Sprintf("UPDATE {{$schemaTable}} SET %s WHERE %s", set, where)
	}

	{{if .NoRowsAffected -}}
		{{if .NoContext -}}
	_, err := exec.Exec(sql, whereArgs...)
		{{else -}}
	_, err := exec.ExecContext(ctx, sql, whereArgs...)
		{{end -}}
	{{else -}}
		{{if .NoContext -}}
	result, err := exec.Exec(sql, whereArgs...)
		{{else -}}
	result, err := exec.ExecContext(ctx, sql, whereArgs...)
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
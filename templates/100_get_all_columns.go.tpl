{{- $alias := .Aliases.Table .Table.Name -}}
{{- $schemaTable := .Table.Name | .SchemaTable}}
func {{$alias.UpSingular}}GetAllColumns() []string {
    return {{$alias.DownSingular}}AllColumns
}

{{- $cols := .Table.Columns | columnNames }}
{{- $colsPrefix := prefixStringSlice $alias.DownSingular (prefixStringSlice "." $cols) }}
{{- $colsSchema := .Table.Columns | columnNames | prefixStringSlice (print $schemaTable ".") }}

func {{$alias.UpSingular}}SelectAllColumnsAs() qm.QueryMod {
    return qm.Select( {{ joinSlices " as " $colsSchema $colsPrefix | stringMap .StringFuncs.quoteWrap | join ", " }} )
}

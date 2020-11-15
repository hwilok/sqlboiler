{{- $alias := .Aliases.Table .Table.Name -}}
{{- $schemaTable := .Table.Name | .SchemaTable}}
func {{$alias.UpSingular}}GetAllColumns() []string {
    return {{$alias.DownSingular}}AllColumns
}

{{/*
func {{$alias.UpSingular}}SelectAllColumnsAs(name string ) []string {
    return {{.Table.Columns | columnNames | stringMap .StringFuncs.quoteWrap | join ", "}}
}

{{- $alias := .Aliases.Table .Table.Name -}}
{{- $schemaTable := .Table.Name | .SchemaTable}}
func {{$alias.UpSingular}}GetAllColumns() []string {
    newList := make([]string, len({{$alias.DownSingular}}AllColumns))
    copy(newList, {{$alias.DownSingular}}AllColumns)
    return newList
}

*/}}
{{- $alias := .Aliases.Table .Table.Name -}}

// {{$alias.UpSingular}}Rels is where relationship names are stored.
var {{$alias.UpSingular}}RelsWhere = struct {
	{{range .Table.FKeys -}}
	{{- $relAlias := $alias.Relationship .Name -}}
	{{$relAlias.Foreign}} func(mods ... qm.QueryMod) qm.QueryModFunc
	{{end -}}

	{{range .Table.ToOneRelationships -}}
	{{- $ftable := $.Aliases.Table .ForeignTable -}}
	{{- $relAlias := $ftable.Relationship .Name -}}
	{{$relAlias.Local}} func(mods ...qm.QueryMod) qm.QueryModFunc
	{{end -}}

	{{range .Table.ToManyRelationships -}}
	{{- $relAlias := $.Aliases.ManyRelationship .ForeignTable .Name .JoinTable .JoinLocalFKeyName -}}
	{{$relAlias.Local}} func(mods ...qm.QueryMod) qm.QueryModFunc
	{{end -}}{{/* range tomany */}}
}{
	{{range $fkey := .Table.FKeys -}}
	{{- $relAlias := $alias.Relationship .Name -}}
	    {{- $ltable := $.Aliases.Table $fkey.Table -}}
        {{- $ftable := $.Aliases.Table $fkey.ForeignTable -}}
        {{- $rel := $ltable.Relationship $fkey.Name -}}
        {{- $canSoftDelete := (getTable $.Tables $fkey.ForeignTable).CanSoftDelete }}
	{{$relAlias.Foreign}}: func(mods ... qm.QueryMod) qm.QueryModFunc {
        return func(q *queries.Query) {
            queryMods := []qm.QueryMod{
        		qm.Where("{{$fkey.ForeignTable | $.Quotes}}.{{$fkey.ForeignColumn | $.Quotes}} = {{$.Table.Name | $.SchemaTable}}. {{$fkey.Column}}"),
        		{{if and $.AddSoftDeletes $canSoftDelete -}}
        		qmhelper.WhereIsNull("deleted_at"),
        		{{- end}}
        	}

        	queryMods = append(queryMods, mods...)
			qSub := {{$ftable.UpPlural}}(queryMods...)

			res, args := queries.BuildQuery(qSub.Query)

			qm.Apply(q, qm.Where(fmt.Sprintf("EXISTS(%s)", res[:len(res) - 1]), args ...))
        }
	},
	{{end -}}

	{{range $rel := .Table.ToOneRelationships -}}
	{{- $ltable := $.Aliases.Table $rel.Table -}}
    		{{- $ftable := $.Aliases.Table $rel.ForeignTable -}}
    		{{- $relAlias := $ftable.Relationship $rel.Name -}}
    		{{- $canSoftDelete := (getTable $.Tables $rel.ForeignTable).CanSoftDelete }}
    		{{$relAlias.Local}}: func(mods ... qm.QueryMod) qm.QueryModFunc {
    		return func(q *queries.Query) {
                queryMods := []qm.QueryMod{
                    qm.Where("{{$rel.ForeignTable | $.Quotes}}.{{$rel.ForeignColumn | $.Quotes}} = {{$.Table.Name | $.SchemaTable}}.{{$rel.Column | $.Quotes}}"),
                    {{if and $.AddSoftDeletes $canSoftDelete -}}
                    qmhelper.WhereIsNull("deleted_at"),
                    {{- end}}
                }
                queryMods = append(queryMods, mods...)
                qSub := {{$ftable.UpPlural}}(queryMods...)

                res, args := queries.BuildQuery(qSub.Query)

                qm.Apply(q, qm.Where(fmt.Sprintf("EXISTS(%s)", res[:len(res) - 1]), args ...))
            }
    		},
	{{end -}}

	{{range  $rel := .Table.ToManyRelationships -}}
	    {{- $ltable := $.Aliases.Table $rel.Table -}}
        {{- $ftable := $.Aliases.Table $rel.ForeignTable -}}
        {{- $relAlias := $.Aliases.ManyRelationship $rel.ForeignTable $rel.Name $rel.JoinTable $rel.JoinLocalFKeyName -}}
        {{- $schemaForeignTable := .ForeignTable | $.SchemaTable -}}
        {{- $canSoftDelete := (getTable $.Tables .ForeignTable).CanSoftDelete }}
	{{$relAlias.Local}}: func(mods ... qm.QueryMod) qm.QueryModFunc {
    		return func(q *queries.Query) {
                var queryMods []qm.QueryMod
                if len(mods) != 0 {
                    queryMods = append(queryMods, mods...)
                }
                queryMods = append(queryMods,
                    qm.Where("{{$schemaForeignTable}}.{{$rel.ForeignColumn | $.Quotes}}={{$.Table.Name | $.SchemaTable}}.{{$rel.Column | $.Quotes}}"),
                    {{if and $.AddSoftDeletes $canSoftDelete -}}
                    qmhelper.WhereIsNull("{{$schemaForeignTable}}.{{"deleted_at" | $.Quotes}}"),
                    {{- end}}
                )

                qSub := {{$ftable.UpPlural}}(queryMods...)

                res, args := queries.BuildQuery(qSub.Query)


                qm.Apply(q, qm.Where(fmt.Sprintf("EXISTS(%s)", res[:len(res) - 1]), args ...))
            }
            },
	{{end -}}{{/* range tomany */}}
}


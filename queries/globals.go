package queries

func WhereClause(q *Query, startAt int) (string, []interface{}) {
	return whereClause(q, startAt)
}

// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.12.0
// source: query.sql

package querytest

import (
	"context"
	"database/sql"
)

const starExpansion = `-- name: StarExpansion :many
SELECT a, b, a, b, foo.a, foo.b FROM foo
`

type StarExpansionRow struct {
	A   sql.NullString
	B   sql.NullString
	A_2 sql.NullString
	B_2 sql.NullString
	A_3 sql.NullString
	B_3 sql.NullString
}

func (q *Queries) StarExpansion(ctx context.Context) ([]StarExpansionRow, error) {
	rows, err := q.db.Query(ctx, starExpansion)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []StarExpansionRow
	for rows.Next() {
		var i StarExpansionRow
		if err := rows.Scan(
			&i.A,
			&i.B,
			&i.A_2,
			&i.B_2,
			&i.A_3,
			&i.B_3,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

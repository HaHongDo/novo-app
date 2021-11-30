// Code generated by sqlc. DO NOT EDIT.
// source: queries.sql

package db

import (
	"context"
)

const listBookGroups = `-- name: ListBookGroups :many
SELECT id, title, description, date_created, owner_id, primary_cover_art_id FROM book_groups
FETCH FIRST $1 ROWS ONLY
`

func (q *Queries) ListBookGroups(ctx context.Context, limit int32) ([]BookGroup, error) {
	rows, err := q.db.Query(ctx, listBookGroups, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []BookGroup
	for rows.Next() {
		var i BookGroup
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Description,
			&i.DateCreated,
			&i.OwnerID,
			&i.PrimaryCoverArtID,
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

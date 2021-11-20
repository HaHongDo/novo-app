// Code generated by sqlc. DO NOT EDIT.
// source: user_role.sql

package db

import (
	"context"
	"database/sql"
)

const deleteRole = `-- name: DeleteRole :exec
DELETE FROM roles
WHERE name = $1
`

func (q *Queries) DeleteRole(ctx context.Context, name string) error {
	_, err := q.db.Exec(ctx, deleteRole, name)
	return err
}

const insertNewRole = `-- name: InsertNewRole :one
INSERT INTO roles (name, description)
VALUES ($1, $2)
RETURNING id, name, description, can_modify_role, can_modify_book_author, can_modify_book_genre, can_modify_book_group, can_modify_book_chapter, can_create_comment, can_update_comment, can_delete_comment
`

type InsertNewRoleParams struct {
	Name        string         `json:"name"`
	Description sql.NullString `json:"description"`
}

func (q *Queries) InsertNewRole(ctx context.Context, arg InsertNewRoleParams) (Role, error) {
	row := q.db.QueryRow(ctx, insertNewRole, arg.Name, arg.Description)
	var i Role
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.CanModifyRole,
		&i.CanModifyBookAuthor,
		&i.CanModifyBookGenre,
		&i.CanModifyBookGroup,
		&i.CanModifyBookChapter,
		&i.CanCreateComment,
		&i.CanUpdateComment,
		&i.CanDeleteComment,
	)
	return i, err
}

const roleIdByName = `-- name: RoleIdByName :one
SELECT id FROM roles
WHERE name = $1
FETCH FIRST ROWS ONLY
`

func (q *Queries) RoleIdByName(ctx context.Context, name string) (int32, error) {
	row := q.db.QueryRow(ctx, roleIdByName, name)
	var id int32
	err := row.Scan(&id)
	return id, err
}

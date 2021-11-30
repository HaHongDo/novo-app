// Code generated by sqlc. DO NOT EDIT.
// source: book_chapter_images.sql

package db

import (
	"context"
)

const insertBookChapterImage = `-- name: InsertBookChapterImage :exec
INSERT INTO book_chapter_images(book_chapter_id, image_id, rank) VALUES($1, $2, $3)
`

type InsertBookChapterImageParams struct {
	BookChapterID int32 `json:"bookChapterID"`
	ImageID       int32 `json:"imageID"`
	Rank          int32 `json:"rank"`
}

func (q *Queries) InsertBookChapterImage(ctx context.Context, arg InsertBookChapterImageParams) error {
	_, err := q.db.Exec(ctx, insertBookChapterImage, arg.BookChapterID, arg.ImageID, arg.Rank)
	return err
}
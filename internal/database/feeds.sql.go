// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.29.0
// source: feeds.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createFeed = `-- name: CreateFeed :one
INSERT INTO feeds
    (id, created_at, updated_at, name,url,user_id)
VALUES
    ($1, $2, $3, $4, $5, $6)
RETURNING id, created_at, updated_at, name, url, user_id
`

type CreateFeedParams struct {
	ID        pgtype.UUID
	CreatedAt pgtype.Timestamp
	UpdatedAt pgtype.Timestamp
	Name      string
	Url       string
	UserID    pgtype.UUID
}

func (q *Queries) CreateFeed(ctx context.Context, arg CreateFeedParams) (Feed, error) {
	row := q.db.QueryRow(ctx, createFeed,
		arg.ID,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.Name,
		arg.Url,
		arg.UserID,
	)
	var i Feed
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
		&i.Url,
		&i.UserID,
	)
	return i, err
}

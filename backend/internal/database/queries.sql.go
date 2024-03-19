// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: queries.sql

package database

import (
	"context"
	"database/sql"
)

const createLink = `-- name: CreateLink :execresult
INSERT INTO links (redirect, link, random) 
VALUES (?, ?, ?)
`

type CreateLinkParams struct {
	Redirect string       `json:"redirect"`
	Link     string       `json:"link"`
	Random   sql.NullBool `json:"random"`
}

func (q *Queries) CreateLink(ctx context.Context, arg CreateLinkParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, createLink, arg.Redirect, arg.Link, arg.Random)
}

const deleteLink = `-- name: DeleteLink :exec
DELETE FROM links
WHERE id = ?
`

func (q *Queries) DeleteLink(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteLink, id)
	return err
}

const getLink = `-- name: GetLink :one
SELECT id, redirect, link, clicked, random, created_at 
FROM links
WHERE id = ? 
LIMIT 1
`

func (q *Queries) GetLink(ctx context.Context, id int32) (Link, error) {
	row := q.db.QueryRowContext(ctx, getLink, id)
	var i Link
	err := row.Scan(
		&i.ID,
		&i.Redirect,
		&i.Link,
		&i.Clicked,
		&i.Random,
		&i.CreatedAt,
	)
	return i, err
}

const listAllLinks = `-- name: ListAllLinks :many
SELECT id, redirect, link, clicked, random, created_at 
FROM links
ORDER BY created_at
`

func (q *Queries) ListAllLinks(ctx context.Context) ([]Link, error) {
	rows, err := q.db.QueryContext(ctx, listAllLinks)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Link{}
	for rows.Next() {
		var i Link
		if err := rows.Scan(
			&i.ID,
			&i.Redirect,
			&i.Link,
			&i.Clicked,
			&i.Random,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

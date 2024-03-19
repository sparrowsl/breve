// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package database

import (
	"database/sql"
	"time"
)

type Link struct {
	ID        int32         `json:"id"`
	Redirect  string        `json:"redirect"`
	Url       string        `json:"url"`
	Clicked   sql.NullInt32 `json:"clicked"`
	Random    sql.NullBool  `json:"random"`
	CreatedAt time.Time     `json:"created_at"`
}

// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package repo

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Message struct {
	ID        string           `json:"id"`
	Sender    string           `json:"sender"`
	Content   string           `json:"content"`
	CreatedAt pgtype.Timestamp `json:"created_at"`
	ThreadID  *string          `json:"thread_id"`
}

type Thread struct {
	ID        string           `json:"id"`
	Topic     *string          `json:"topic"`
	CreatedAt pgtype.Timestamp `json:"created_at"`
}

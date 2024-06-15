// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: messages.sql

package sqlc

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createMessage = `-- name: CreateMessage :exec
INSERT INTO messages (
    id, 
    created_at,
    raw,
    raw_jsonb
) VALUES (
    DEFAULT, NOW(), $1, $2
)
`

type CreateMessageParams struct {
	Raw      pgtype.Text
	RawJsonb []byte
}

func (q *Queries) CreateMessage(ctx context.Context, arg CreateMessageParams) error {
	_, err := q.db.Exec(ctx, createMessage, arg.Raw, arg.RawJsonb)
	return err
}

const getMessage = `-- name: GetMessage :one
SELECT id, raw, raw_jsonb, created_at FROM messages WHERE id = $1 LIMIT 1
`

func (q *Queries) GetMessage(ctx context.Context, id int64) (Message, error) {
	row := q.db.QueryRow(ctx, getMessage, id)
	var i Message
	err := row.Scan(
		&i.ID,
		&i.Raw,
		&i.RawJsonb,
		&i.CreatedAt,
	)
	return i, err
}

const getMessagesAsc = `-- name: GetMessagesAsc :many
SELECT id, created_at, raw FROM messages ORDER BY id ASC LIMIT $1 OFFSET $2
`

type GetMessagesAscParams struct {
	Limit  int32
	Offset int32
}

type GetMessagesAscRow struct {
	ID        int64
	CreatedAt pgtype.Timestamp
	Raw       pgtype.Text
}

func (q *Queries) GetMessagesAsc(ctx context.Context, arg GetMessagesAscParams) ([]GetMessagesAscRow, error) {
	rows, err := q.db.Query(ctx, getMessagesAsc, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetMessagesAscRow
	for rows.Next() {
		var i GetMessagesAscRow
		if err := rows.Scan(&i.ID, &i.CreatedAt, &i.Raw); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getMessagesDesc = `-- name: GetMessagesDesc :many
SELECT id, created_at, raw FROM messages ORDER BY id DESC LIMIT $1 OFFSET $2
`

type GetMessagesDescParams struct {
	Limit  int32
	Offset int32
}

type GetMessagesDescRow struct {
	ID        int64
	CreatedAt pgtype.Timestamp
	Raw       pgtype.Text
}

func (q *Queries) GetMessagesDesc(ctx context.Context, arg GetMessagesDescParams) ([]GetMessagesDescRow, error) {
	rows, err := q.db.Query(ctx, getMessagesDesc, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetMessagesDescRow
	for rows.Next() {
		var i GetMessagesDescRow
		if err := rows.Scan(&i.ID, &i.CreatedAt, &i.Raw); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
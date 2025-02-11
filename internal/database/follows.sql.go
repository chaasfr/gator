// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: follows.sql

package database

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const createFeedFollow = `-- name: CreateFeedFollow :one
WITH ff AS (
INSERT INTO feed_follows(id, created_at, updated_at, user_id, feed_id)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5
)
RETURNING id, created_at, updated_at, user_id, feed_id)
select ff.id, ff.created_at, ff.updated_at, ff.user_id, ff.feed_id, u.name as username, f.name as feedname
from ff
inner join users as u on ff.user_id = u.id
inner join feeds as f on ff.feed_id = f.id
`

type CreateFeedFollowParams struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	UserID    uuid.UUID
	FeedID    uuid.UUID
}

type CreateFeedFollowRow struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	UserID    uuid.UUID
	FeedID    uuid.UUID
	Username  string
	Feedname  string
}

func (q *Queries) CreateFeedFollow(ctx context.Context, arg CreateFeedFollowParams) (CreateFeedFollowRow, error) {
	row := q.db.QueryRowContext(ctx, createFeedFollow,
		arg.ID,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.UserID,
		arg.FeedID,
	)
	var i CreateFeedFollowRow
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.UserID,
		&i.FeedID,
		&i.Username,
		&i.Feedname,
	)
	return i, err
}

const deleteFeedFollow = `-- name: DeleteFeedFollow :one
WITH feed_found AS (
    SELECT id as feed_id
    FROM feeds 
    WHERE url = $2
), 
deleted AS (
    DELETE FROM feed_follows as ff
    USING feed_found
    WHERE ff.user_id = $1 
      AND ff.feed_id = feed_found.feed_id
    RETURNING ff.id, ff.created_at, ff.updated_at, ff.user_id, ff.feed_id
)
SELECT COUNT(*) 
FROM deleted
`

type DeleteFeedFollowParams struct {
	UserID uuid.UUID
	Url    string
}

func (q *Queries) DeleteFeedFollow(ctx context.Context, arg DeleteFeedFollowParams) (int64, error) {
	row := q.db.QueryRowContext(ctx, deleteFeedFollow, arg.UserID, arg.Url)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const getFeedFollowsForUser = `-- name: GetFeedFollowsForUser :many
select u.name as username, f.name as feedname
from feed_follows as ff
inner join users as u on ff.user_id = u.id
inner join feeds as f on ff.feed_id = f.id
where ff.user_id = $1
`

type GetFeedFollowsForUserRow struct {
	Username string
	Feedname string
}

func (q *Queries) GetFeedFollowsForUser(ctx context.Context, userID uuid.UUID) ([]GetFeedFollowsForUserRow, error) {
	rows, err := q.db.QueryContext(ctx, getFeedFollowsForUser, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetFeedFollowsForUserRow
	for rows.Next() {
		var i GetFeedFollowsForUserRow
		if err := rows.Scan(&i.Username, &i.Feedname); err != nil {
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

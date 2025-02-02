-- name: CreateFeed :one
INSERT INTO feeds(id, created_at, updated_at, name, url, user_id)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6
)
RETURNING *;

-- name: GetFeeds :many
select feeds.name as name, feeds.url as url, users.name as username from feeds
left join users on feeds.user_id = users.id;

-- name: GetFeedIdFromUrl :one
select * from feeds where url= $1 order by created_at asc limit 1;

-- name: MarkFeedFetched :exec
UPDATE feeds
SET updated_at = $2, last_fetched_at = $2
WHERE id = $1;

-- name: GetNextFeedToFetch :one
SELECT * from feeds
order by last_fetched_at asc NULLS FIRST;
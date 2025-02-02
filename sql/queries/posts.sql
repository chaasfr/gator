-- name: CreatePost :one
INSERT INTO posts(id, created_at, updated_at, title, url, description, published_at,feed_id)
VALUES ($1, $2, $2, $3, $4, $5, $6, $7)
RETURNING *;

-- name: GetPostsForUser :many
SELECT * from posts
inner join feed_follows as ff on posts.feed_id = ff.feed_id
where ff.user_id = $1
order by published_at DESC
limit $2;
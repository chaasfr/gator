-- name: CreateFeedFollow :one
WITH ff AS (
INSERT INTO feed_follows(id, created_at, updated_at, user_id, feed_id)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5
)
RETURNING *)
select ff.*, u.name as username, f.name as feedname
from ff
inner join users as u on ff.user_id = u.id
inner join feeds as f on ff.feed_id = f.id;

-- name: GetFeedFollowsForUser :many
select u.name as username, f.name as feedname
from feed_follows as ff
inner join users as u on ff.user_id = u.id
inner join feeds as f on ff.feed_id = f.id
where ff.user_id = $1;


-- name: DeleteFeedFollow :one
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
    RETURNING ff.*
)
SELECT COUNT(*) 
FROM deleted;

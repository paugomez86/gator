-- name: CreateFeed :exec
INSERT INTO feeds (id, created_at, updated_at, name, url, user_id)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6
);

-- name: GetFeeds :many
SELECT feeds.*, users.name AS username 
FROM feeds
INNER JOIN users ON feeds.user_id = users.id;
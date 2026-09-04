-- name: CreateFeedFollow :one
WITH insert_query AS (
    INSERT INTO feed_follows (id, created_at, updated_at, user_id, feed_id)
    VALUES (
        $1, --id
        $2, --created_at
        $3, --updated_at
        $4, --user_id
        $5  --feed_id
    )
    RETURNING *
) 
SELECT 
    insert_query.*,
    users.name AS user_name,
    feeds.name AS feed_name
FROM insert_query
INNER JOIN users ON insert_query.user_id = users.id
INNER JOIN feeds ON insert_query.feed_id = feeds.id;
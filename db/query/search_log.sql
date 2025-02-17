-- name: CreateSearchLog :one
INSERT INTO search_log (
    search_id,
    user_id,
    search_query,
    search_filters,
    results_count
) VALUES (
    $1,
    $2,
    $3,
    $4,
    $5
) RETURNING *;

-- name: GetSearchLogList :many
SELECT 
    * 
FROM  
    search_log
ORDER BY 
    searched_at DESC;

-- name: GetSearchLogByUserId :many
SELECT 
    * 
FROM 
    search_log
WHERE
    user_id = $1;

-- name: GetSearchLogByTime :many
SELECT 
    * 
FROM 
    search_log
WHERE
    searched_at BETWEEN $1 AND $2;

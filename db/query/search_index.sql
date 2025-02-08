-- name: CreateSearchIndex :one
INSERT INTO search_index(
    product_id,
    product_name,
    product_desc,
    product_attributes
) VALUES (
    $1,
    $2,
    $3,
    $4
) RETURNING *;

-- name: GetSearchIndexByProductId :many
SELECT 
    * 
FROM 
    search_index 
WHERE 
    product_id = $1;
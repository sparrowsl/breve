-- name: GetLink :one
SELECT * 
FROM links
WHERE id = ? 
LIMIT 1;

-- name: ListAllLinks :many
SELECT * 
FROM links
ORDER BY created_at;

-- name: CreateLink :execresult
INSERT INTO links (redirect, link, random) 
VALUES (?, ?, ?);

-- name: DeleteLink :exec
DELETE FROM links
WHERE id = ?;

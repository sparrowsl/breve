-- name: GetLink :one
SELECT * 
FROM links
WHERE id = $1 
LIMIT 1;

-- name: GetLinkByURL :one
SELECT * 
FROM links
WHERE url = $1 
LIMIT 1;

-- name: ListAllLinks :many
SELECT * 
FROM links
ORDER BY created_at;

-- name: CreateLink :one
INSERT INTO links (redirect, url, random) 
VALUES ($1, $2, $3)
RETURNING *;

-- name: UpdateLink :one
UPDATE links 
SET 
  url = $1,
  redirect = $2,
  random = $3
WHERE id = $4
RETURNING *;

-- name: DeleteLink :exec
DELETE FROM links
WHERE id = $1;

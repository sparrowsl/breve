-- name: GetLink :one
SELECT * 
FROM links
WHERE id = $1 
LIMIT 1;

-- name: ListAllLinks :many
SELECT * 
FROM links
ORDER BY created_at;

-- name: CreateLink :exec
INSERT INTO links (redirect, url, random) 
VALUES ($1, $2, $3);


-- name: UpdateLink :exec
UPDATE links 
SET 
  url = $1,
  redirect = $2,
  random = $3
WHERE id = $4;

-- name: DeleteLink :exec
DELETE FROM links
WHERE id = $1;

-- name: GetBreve :one
SELECT * FROM breve
WHERE id = ?
LIMIT 1;

-- name: GetBreves :many
SELECT * FROM breve
ORDER BY created;

-- name: CreateBreve :execresult
INSERT INTO breve(redirect, link) VALUES(?, ?);

-- name: DeleteAuthor :exec
DELETE FROM breve
WHERE id = ?;

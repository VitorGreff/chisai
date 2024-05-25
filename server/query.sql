-- name: GetLink :one
SELECT * FROM urls
WHERE long_url = $1 LIMIT 1;

-- name: GetLongURL :one
SELECT long_url FROM urls
WHERE short_url = $1 LIMIT 1;

-- name: ListLinks :many
SELECT * FROM urls
ORDER BY long_url;

-- name: CreateLink :one
INSERT INTO urls (
  long_url, short_url
) VALUES (
  $1, $2
)
RETURNING *;

-- name: DeleteLink :exec
DELETE FROM urls
WHERE short_url = $1;

-- name: DeleteAllLinks :exec
DELETE FROM urls;

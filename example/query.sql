-- name: GetNums :many
SELECT * FROM uint_table;

-- name: InsertNum :one
INSERT INTO uint_table (num) VALUES ($1) RETURNING *;
-- name: CreateTransfer :one
INSERT INTO transfers (from_account_id, to_account_id, amount) VALUES ($1, $2, $3) RETURNING *;

-- name: GetTransfer :one
SELECT * FROM transfers WHERE id = $1;

-- name: ListTransfers :many
SELECT * FROM transfers WHERE from_account_id = $1 OR to_account_id = $1 ORDER BY id LIMIT $2 OFFSET $3;

-- name: UpdateTransfer :one
UPDATE transfers SET amount = $2 WHERE id = $1 RETURNING *;

-- name: DeleteTransfer :exec
DELETE FROM transfers WHERE id = $1;

-- name: ListTransfersByAccount :many
SELECT * FROM transfers WHERE from_account_id = $1 OR to_account_id = $1 ORDER BY id LIMIT $2 OFFSET $3;

-- name: ListTransfersByAccountAndDate :many
SELECT * FROM transfers WHERE (from_account_id = $1 OR to_account_id = $1) AND created_at >= $2 AND created_at <= $3 ORDER BY id LIMIT $4 OFFSET $5;

-- name: ListTransfersByDate :many
SELECT * FROM transfers WHERE created_at >= $1 AND created_at <= $2 ORDER BY id LIMIT $3 OFFSET $4;

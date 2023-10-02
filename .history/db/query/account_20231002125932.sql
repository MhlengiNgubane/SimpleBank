-- name: CreateAccount :one
INSERT INTO accounts (
    owner,
    balance,
    currency,
) VA
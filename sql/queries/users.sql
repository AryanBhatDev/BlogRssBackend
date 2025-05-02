-- name: CreateUser :one
INSERT INTO users(id,created_at,updated_at,name,email,password,api_key)
VALUES ($1,$2,$3,$4,$5,$6,
    encode(sha256(random()::text::bytea),'hex')
)
RETURNING *;

-- name: GetUserByApiKey :one
SELECT * FROM users WHERE api_key =$1;
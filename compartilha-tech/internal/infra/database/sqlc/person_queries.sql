-- name: InsertPerson :exec
INSERT INTO person(id,name,age,active,created_at,updated_at)
VALUES($1,$2,$3,$4,$5,$6);

-- name: GetPersons :many
SELECT * FROM person WHERE deleted_at IS NULL;

-- name: GetPersonById :one
SELECT * FROM person WHERE id = $1 AND deleted_at IS NULL;

-- name: UpdatePerson :exec
UPDATE person SET
name = COALESCE(sqlc.narg('name'), name),
age = COALESCE(sqlc.narg('age'), age),
active = COALESCE(sqlc.narg('active'), active),
updated_at = sqlc.arg('updated_at')
WHERE id = $1;

-- name: DeletePerson :exec
UPDATE person SET
deleted_at = NOW()
WHERE id = $1;
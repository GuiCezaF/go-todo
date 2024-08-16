-- name: GetToDos :many 
SELECT "id", "title", "description" ,"completed"
FROM to_dos;

-- name: CreateToDo :one
INSERT INTO to_dos
("title", "description", "completed") VALUES
($1, $2, $3)
RETURNING "id";

-- name: UpdateToDo :exec
UPDATE to_dos
SET
    "title" = $1,
    "description" = $2,
    "completed" = $3
WHERE 
    "id" = $4;

-- name: MarkCompleted :one
UPDATE to_dos
SET
    completed = TRUE
WHERE   
  id = $1
RETURNING completed;

-- name: RemoveMarkCompleted :one
UPDATE to_dos
SET
    completed = FALSE
WHERE   
  id = $1
RETURNING completed;

-- name: DeleteToDo :exec
DELETE FROM to_dos
WHERE
  id = $1;
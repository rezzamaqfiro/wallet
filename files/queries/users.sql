
-- name: GetBalanceByUserID :one
select * from users
where user_id = $1 and deleted is false;

-- name: UpdateBalanceByUserID :one
UPDATE users set balance = balance + $1
where user_id = $2
RETURNING balance;
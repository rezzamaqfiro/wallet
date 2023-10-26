
-- name: GetOrdersByUserID :many
select * from orders
where user_id = $1 and deleted is false;

-- name: GetOrderByInvoiceID :one
select * from orders
where invoice = $1 and deleted is false;

-- name: UpdateOrderStatusByInvoiceID :one
UPDATE orders set status = $1 
where invoice = $2
RETURNING status;

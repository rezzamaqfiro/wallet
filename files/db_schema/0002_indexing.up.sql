-- orders indexing
CREATE INDEX IF NOT EXISTS orders_user_id_idx
    ON orders (user_id)
    where deleted IS FALSE;
CREATE INDEX IF NOT EXISTS orders_status_idx
    ON orders(status)
    where deleted IS FALSE;

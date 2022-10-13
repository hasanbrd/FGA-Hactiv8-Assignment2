CREATE TABLE IF NOT EXISTS orders (
    order_id SERIAL PRIMARY,
	customer_name varchar(25) NOT NULL,
	ordered_at timestamptz NOT NULL,
)

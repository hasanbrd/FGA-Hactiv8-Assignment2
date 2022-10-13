CREATE TABLE IF NOT EXISTS items (
    item_id SERIAL PRIMARY KEY,
	item_code varchar(20) NOT NULL,
	description varchar(100) NULL,
	quantity INT NOT NULL,
	order_id text NOT NULL,
)

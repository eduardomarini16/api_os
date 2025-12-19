CREATE TABLE services_order_updates (
    id SERIAL PRIMARY KEY,
    service_order_id INT NOT NULL REFERENCES service_orders(id) ON DELETE CASCADE,
    user_id INT NOT NULL REFERENCES users(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
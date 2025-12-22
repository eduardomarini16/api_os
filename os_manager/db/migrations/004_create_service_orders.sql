CREATE TABLE service_orders (
    id SERIAL PRIMARY KEY,
    client_id INT NOT NULL REFERENCES clients(id),
    opened_by INT NOT NULL REFERENCES users(id),
    assigned_to INT REFERENCES users(id),
    status service_order_status NOT NULL DEFAULT 'open',
    priority service_order_priority NOT NULL DEFAULT 'medium',
    description TEXT NOT NULL,
    opened_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    closed_at TIMESTAMP
);
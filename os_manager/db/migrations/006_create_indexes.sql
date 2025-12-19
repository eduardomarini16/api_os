CREATE INDEX idx_so_status ON service_orders(status);
CREATE INDEX idx_so_priority ON service_orders(priority);
CREATE INDEX idx_so_assigned ON service_orders(assigned_to);
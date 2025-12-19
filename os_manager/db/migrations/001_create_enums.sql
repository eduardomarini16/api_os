CREATE TYPE user_role AS ENUM (
    'admin',
    'manager',
    'technician',
    'client'
);

CREATE TYPE service_order_status AS ENUM (
    'open',
    'in_progress',
    'finished',
    'canceled'
);

CREATE TYPE service_order_priority AS ENUM (
    'low',
    'medium',
    'high',
    'critical'
);

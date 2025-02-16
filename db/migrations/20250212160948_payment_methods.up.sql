CREATE TYPE payment_method_status AS ENUM('active', 'inactive');

CREATE TABLE IF NOT EXISTS payment_methods (
    id SERIAL PRIMARY KEY,
    method_name VARCHAR(255) NOT NULL,
    details TEXT,
    status payment_method_status NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT NULL
)
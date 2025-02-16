CREATE TYPE spot_status AS ENUM(
    'available',
    'reserved',
    'occupied'
);

CREATE TABLE IF NOT EXISTS spots (
    id SERIAL PRIMARY KEY,
    zone_id INT NOT NULL,
    spot_number VARCHAR(255) NOT NULL,
    status spot_status NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT NULL,
    CONSTRAINT fk_spot_zone FOREIGN KEY (zone_id) REFERENCES zones (id) ON DELETE CASCADE ON UPDATE CASCADE
)
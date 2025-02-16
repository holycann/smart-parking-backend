CREATE TYPE transaction_status AS ENUM(
    'pending',
    'paid',
    'failed',
    'refunded'
);

CREATE TABLE IF NOT EXISTS transactions (
    id SERIAL PRIMARY KEY,
    reservation_id INT NOT NULL,
    amount INT NOT NULL,
    payment_method_id INT NOT NULL,
    status transaction_status NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT NULL,
    CONSTRAINT fk_transaction_reservation FOREIGN KEY (reservation_id) REFERENCES reservations (id) ON DELETE CASCADE ON UPDATE CASCADE,
    CONSTRAINT fk_transaction_payment_method FOREIGN KEY (payment_method_id) REFERENCES payment_methods (id) ON DELETE CASCADE ON UPDATE CASCADE
)
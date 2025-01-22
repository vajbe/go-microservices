CREATE TABLE
    IF NOT EXISTS orders (
        id UUID PRIMARY KEY DEFAULT gen_random_uuid (), -- Unique identifier for each order
        user_id UUID, -- Reference to the user who placed the order
        order_status VARCHAR(20) NOT NULL DEFAULT 'PENDING', -- Status of the order (e.g., PENDING, CONFIRMED, SHIPPED)
        total_amount NUMERIC(10, 2) NOT NULL, -- Total amount for the order
        payment_status VARCHAR(20) NOT NULL DEFAULT 'PENDING', -- Status of payment (e.g., PENDING, PAID, FAILED)
        products JSONB NOT NULL,
        created_at BIGINT NOT NULL DEFAULT EXTRACT(
            EPOCH
            FROM
                now ()
        ), -- Order creation timestamp
        updated_at BIGINT NOT NULL DEFAULT EXTRACT(
            EPOCH
            FROM
                now ()
        ), -- Last updated timestamp
        FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE -- Reference to the users table
    );
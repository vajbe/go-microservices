CREATE TABLE
    products (
        id UUID PRIMARY KEY DEFAULT gen_random_uuid (),
        name TEXT NOT NULL,
        description TEXT,
        price NUMERIC(10, 2) NOT NULL,
        stock INT NOT NULL,
        category TEXT,
        created_at BIGINT NOT NULL DEFAULT EXTRACT(
            EPOCH
            FROM
                now ()
        )
    );
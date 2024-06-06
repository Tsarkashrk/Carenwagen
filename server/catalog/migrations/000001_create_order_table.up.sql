CREATE TABLE cars IF NOT EXISTS (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP(0) WITH TIME ZONE NOT NULL DEFAULT NOW(),
    title VARCHAR(255) NOT NULL,
    model VARCHAR(255) NOT NULL,
    make VARCHAR(255) NOT NULL,
    year INT NOT NULL,
    color VARCHAR(255) NOT NULL,
    price NUMERIC(10, 2) NOT NULL,
    mileage INT NOT NULL,
    description TEXT,
    transmission VARCHAR(50) NOT NULL,
    fuel_type VARCHAR(50) NOT NULL,
    image_url VARCHAR(255),
    version INT NOT NULL
);

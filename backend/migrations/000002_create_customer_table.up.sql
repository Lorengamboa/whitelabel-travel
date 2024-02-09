-- migrations/000002_create_customer_table.up
-- Add up migration script here

-- Create tables customer
CREATE TABLE IF NOT EXISTS customers(
    id UUID NOT NULL PRIMARY KEY DEFAULT gen_random_uuid(),
    name TEXT NOT NULL,
    email TEXT NOT NULL UNIQUE,
    phone_number TEXT NOT NULL,
    address TEXT NOT NULL,
    date_joined TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    logo TEXT NULL,
    url TEXT NULL
);

-- Create table Client
CREATE TABLE IF NOT EXISTS clients(
    id UUID NOT NULL PRIMARY KEY DEFAULT gen_random_uuid(),
    customer_id UUID NOT NULL,
    name TEXT NOT NULL,
    email TEXT NOT NULL UNIQUE,
    phone_number TEXT NOT NULL,
    address TEXT NOT NULL,
    date_joined TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    logo TEXT NULL,
    url TEXT NULL,
    FOREIGN KEY (customer_id) REFERENCES customers(id) ON DELETE CASCADE
);


-- Insert data into customers
INSERT INTO customers (name, email, phone_number, address, url) VALUES ('Hakuna matata', 'hakunamatata@gmail.com', '+34654456401', 'Madrid, Spain', 'https://www.hakunamatata.com');
INSERT INTO customers (name, email, phone_number, address, url) VALUES ('Kenya experience', 'kenyaexperience@gmail.com', '+254114085558', 'Nairobi, Kenya', 'https://www.kenyaexperience.com');
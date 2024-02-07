-- migrations/000001_create_users_table.up.sql
-- Add up migration script here

CREATE TYPE roles AS ENUM ('super_admin');

-- User table
CREATE TABLE IF NOT EXISTS users(
    id UUID NOT NULL PRIMARY KEY DEFAULT gen_random_uuid(),
    email TEXT NOT NULL UNIQUE,
    password TEXT NOT NULL,
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL,
    is_active BOOLEAN DEFAULT FALSE,
    role roles DEFAULT 'super_admin',
    thumbnail TEXT NULL,
    date_joined TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
CREATE INDEX IF NOT EXISTS users_id_email_is_active_indx ON users (id, email, is_active);
-- Create a domain for phone data type
CREATE DOMAIN phone AS TEXT CHECK(
    octet_length(VALUE) BETWEEN 1
    /*+*/
    + 8 AND 1
    /*+*/
    + 15 + 3
    AND VALUE ~ '^\+\d+$'
);
-- User details table (One-to-one relationship)
CREATE TABLE user_profile (
    id UUID NOT NULL PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL UNIQUE,
    phone_number phone NULL,
    birth_date DATE NULL,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);
CREATE INDEX IF NOT EXISTS users_detail_id_user_id ON user_profile (id, user_id);

-- Create a admin user - Lorenzo
INSERT INTO users (email, password, first_name, last_name, is_active, role) VALUES ('lorenzogamboagarcia@gmail.com', '$argon2id$v=19$m=65536,t=1,p=8$T/lpvVYTETr7q19gosOJWA$4jFhclBeJV/gyglTiyzcgBvTSfKj6Il9xSa7E0bwD8Q', 'Lorenzo', 'Gamboa', TRUE, 'super_admin');
INSERT INTO user_profile (user_id, phone_number, birth_date) VALUES ((SELECT id FROM users WHERE email = 'lorenzogamboagarcia@gmail.com'), '+254114085558', '1993-09-09');

-- Create a admin user - Hadi
INSERT INTO users (email, password, first_name, last_name, is_active, role) VALUES ('abdimohamedhadi@gmail.com', '$argon2id$v=19$m=65536,t=1,p=8$T/lpvVYTETr7q19gosOJWA$4jFhclBeJV/gyglTiyzcgBvTSfKj6Il9xSa7E0bwD8Q', 'Muhammad', 'Hadi', TRUE, 'super_admin');
INSERT INTO user_profile (user_id, phone_number, birth_date) VALUES ((SELECT id FROM users WHERE email = 'abdimohamedhadi@gmail.com'), '+254114085558', '1993-09-27');
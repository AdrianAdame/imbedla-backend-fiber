-- Add UUID extension
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Set timezone
-- For more information, please visit:
-- https://en.wikipedia.org/wiki/List_of_tz_database_time_zones
SET TIMEZONE="Europe/Moscow";

-- Create users table
CREATE TABLE users (
    id UUID DEFAULT uuid_generate_v4 () PRIMARY KEY,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW (),
    updated_at TIMESTAMP NULL,
    email VARCHAR (255) NOT NULL UNIQUE,
    firstname VARCHAR (255) NOT NULL UNIQUE,
    lastname VARCHAR (255) NOT NULL UNIQUE,
    about VARCHAR (255) NOT NULL UNIQUE,
    password_hash VARCHAR (255) NOT NULL,
    user_status INT NOT NULL,
    user_role VARCHAR (25) NOT NULL
);

-- Add indexes
CREATE INDEX active_users ON users (id) WHERE user_status = 1;

--- Tokens Table
CREATE TABLE tokens (
    token_id UUID DEFAULT uuid_generate_v4 () PRIMARY KEY,
    id UUID REFERENCES users(id),
    access_token VARCHAR(255) DEFAULT NULL,
    refresh_token VARCHAR(255) DEFAULT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

--- Rooms Table
CREATE TABLE rooms (
    id UUID DEFAULT uuid_generate_v4 () PRIMARY KEY,
    user_id UUID REFERENCES users(id),
    name VARCHAR(255) NOT NULL,
    color VARCHAR(255) NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP NULL,
	type VARCHAR(255) NOT NULL
);

--- Rooms Table
CREATE TABLE plants (
    id UUID DEFAULT uuid_generate_v4 () PRIMARY KEY,
    user_id UUID REFERENCES users(id),
    room_id UUID REFERENCES rooms(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP NULL,
	module_information json NULL,
	module_specs json NULL
);
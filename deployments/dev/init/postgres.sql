-- Database initialization script
-- Created on March 31, 2025

-- Create enum types
DO $$
BEGIN
    -- Create card_type_enum if it doesn't exist
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'card_type_enum') THEN
        CREATE TYPE card_type_enum AS ENUM ('visa', 'mastercard', 'amex');
    END IF;

    -- Create status_enum if it doesn't exist
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'status_enum') THEN
        CREATE TYPE status_enum AS ENUM ('active', 'inactive');
    END IF;

    -- Create gender_enum if it doesn't exist
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'gender_enum') THEN
        CREATE TYPE gender_enum AS ENUM ('male', 'female');
    END IF;

    -- Create role_enum if it doesn't exist
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'role_enum') THEN
        CREATE TYPE role_enum AS ENUM ('hr', 'admin', 'flight_planner', 'passenger_services', 'ground_services');
    END IF;

    -- Create flight_status_enum if it doesn't exist
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'flight_status_enum') THEN
        CREATE TYPE flight_status_enum AS ENUM ('scheduled', 'delayed', 'cancelled', 'departed', 'arrived');
    END IF;

END$$;

-- Create tables
-- Credit Cards table
-- CREATE TABLE IF NOT EXISTS credit_cards (
--     id SERIAL PRIMARY KEY,
--     card_number VARCHAR(16) UNIQUE NOT NULL,
--     card_holder_name VARCHAR(100) NOT NULL,
--     card_holder_surname VARCHAR(100) NOT NULL,
--     expiration_month INT NOT NULL,
--     expiration_year INT NOT NULL,
--     cvv VARCHAR(4) NOT NULL,
--     card_type card_type_enum NOT NULL,
--     amount DECIMAL(10,2) NOT NULL,
--     currency VARCHAR(3) NOT NULL,
--     issuer_bank VARCHAR(100),
--     status status_enum DEFAULT 'active' NOT NULL,
--     created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
-- );

-- Employees table
-- CREATE TABLE IF NOT EXISTS employees (
--     id SERIAL PRIMARY KEY,
--     employee_id VARCHAR(50) UNIQUE NOT NULL,
--     name VARCHAR(50) NOT NULL,
--     surname VARCHAR(50) NOT NULL,
--     email VARCHAR(100) UNIQUE NOT NULL,
--     phone VARCHAR(15),
--     address VARCHAR(255),
--     gender gender_enum NOT NULL,
--     birth_date TIMESTAMP NOT NULL,
--     hire_date TIMESTAMP NOT NULL,
--     position VARCHAR(100) NOT NULL,
--     role role_enum NOT NULL,
--     salary DECIMAL(10,2) NOT NULL,
--     status status_enum DEFAULT 'active' NOT NULL,
--     emergency_contact VARCHAR(100),
--     emergency_phone VARCHAR(15),
--     profile_image_url VARCHAR(255),
--     created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
--     updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
--     password_hash TEXT NOT NULL,
--     salt TEXT NOT NULL
-- );

-- Flights table
-- CREATE TABLE IF NOT EXISTS flights (
--     id SERIAL PRIMARY KEY,
--     flight_number VARCHAR(10) NOT NULL,
--     departure_airport VARCHAR(3) NOT NULL,
--     destination_airport VARCHAR(3) NOT NULL,
--     departure_datetime TIMESTAMP NOT NULL,
--     arrival_datetime TIMESTAMP NOT NULL,
--     departure_gate_number VARCHAR(5),
--     destination_gate_number VARCHAR(5),
--     plane_registration VARCHAR(10) NOT NULL,
--     status flight_status_enum NOT NULL,
--     price DECIMAL(10,2) NOT NULL,
--     UNIQUE(flight_number, departure_datetime)
-- );

-- Passengers table
-- CREATE TABLE IF NOT EXISTS passengers (
--     id SERIAL PRIMARY KEY,
--     national_id VARCHAR(11) NOT NULL,
--     pnr_no VARCHAR(6) UNIQUE NOT NULL,
--     flight_id INT NOT NULL,
--     payment_id INT NOT NULL,
--     baggage_allowance INT NOT NULL,
--     baggage_id VARCHAR(12) NOT NULL,
--     fare_type VARCHAR(50) NOT NULL,
--     seat INT DEFAULT NULL,
--     meal VARCHAR(50) NOT NULL,
--     extra_baggage INT NOT NULL,
--     check_in BOOLEAN NOT NULL,
--     name VARCHAR(50) NOT NULL,
--     surname VARCHAR(50) NOT NULL,
--     email VARCHAR(100) NOT NULL,
--     phone VARCHAR(15) NOT NULL,
--     gender gender_enum NOT NULL,
--     birth_date TIMESTAMP NOT NULL,
--     cip_member BOOLEAN NOT NULL,
--     vip_member BOOLEAN NOT NULL,
--     disabled BOOLEAN NOT NULL,
--     child BOOLEAN NOT NULL,
--     created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
--     updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
-- );

-- Create luggage table for PostgreSQL
-- CREATE TABLE IF NOT EXISTS baggages (
--     baggage_id VARCHAR(12) PRIMARY KEY,
--     baggage_allowance FLOAT DEFAULT NULL,
--     weight FLOAT NOT NULL,
--     piece INT NOT NULL
-- );

-- Payments table
-- CREATE TABLE IF NOT EXISTS payments (
--     id SERIAL PRIMARY KEY,
--     payment_id VARCHAR(50) UNIQUE NOT NULL,
--     user_id VARCHAR(50) NOT NULL,
--     card_number VARCHAR(16) NOT NULL,
--     amount DECIMAL(10,2) NOT NULL,
--     currency VARCHAR(3) NOT NULL,
--     payment_method VARCHAR(50) NOT NULL,
--     status status_enum DEFAULT 'active' NOT NULL,
--     created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
--     updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
-- );

-- Planes table
-- CREATE TABLE IF NOT EXISTS planes (
--     registration VARCHAR(10) PRIMARY KEY,
--     model VARCHAR(50) NOT NULL,
--     manufacturer VARCHAR(50) NOT NULL,
--     capacity INT NOT NULL,
--     status status_enum DEFAULT 'active' NOT NULL,
--     created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
--     updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
-- );

-- Refunds table
-- CREATE TABLE IF NOT EXISTS refunds (
--     id SERIAL PRIMARY KEY,
--     refund_id VARCHAR(50) UNIQUE NOT NULL,
--     payment_id VARCHAR(50) NOT NULL,
--     amount DECIMAL(10,2) NOT NULL,
--     currency VARCHAR(3) NOT NULL,
--     reason VARCHAR(255),
--     status status_enum DEFAULT 'active' NOT NULL,
--     created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
--     updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
-- );

-- Users table
-- CREATE TABLE IF NOT EXISTS users (
--     id SERIAL PRIMARY KEY,
--     name VARCHAR(50) NOT NULL,
--     surname VARCHAR(50) NOT NULL,
--     username VARCHAR(50) UNIQUE NOT NULL,
--     email VARCHAR(100) UNIQUE NOT NULL,
--     password_hash TEXT NOT NULL,
--     salt TEXT NOT NULL,
--     phone VARCHAR(15),
--     gender gender_enum NOT NULL,
--     birth_date TIMESTAMP NOT NULL,
--     last_login TIMESTAMP,
--     created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
--     updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
--     last_password_change TIMESTAMP
-- );

-- Create foreign key constraints

-- ALTER TABLE flights
--     ADD CONSTRAINT fk_plane_registration
--     FOREIGN KEY (plane_registration) REFERENCES planes(registration)
--     ON DELETE RESTRICT;

-- ALTER TABLE passengers
--     ADD CONSTRAINT fk_flight_id
--     FOREIGN KEY (flight_id) REFERENCES flights(id)
--     ON DELETE RESTRICT;

-- ALTER TABLE passengers
--     ADD CONSTRAINT fk_payment_id
--     FOREIGN KEY (payment_id) REFERENCES payments(id)
--     ON DELETE RESTRICT;

-- ALTER TABLE passengers
--     ADD CONSTRAINT fk_baggage_id
--     FOREIGN KEY (baggage_id) REFERENCES baggages(baggage_id)
--     ON DELETE CASCADE;

-- ALTER TABLE refunds
--     ADD CONSTRAINT fk_payment_id
--     FOREIGN KEY (payment_id) REFERENCES payments(payment_id)
--     ON DELETE RESTRICT;

-- Add indexes for performance (optional but recommended)
-- CREATE INDEX IF NOT EXISTS idx_flights_status ON flights(status);
-- CREATE INDEX IF NOT EXISTS idx_flights_departure ON flights(departure_datetime);
-- CREATE INDEX IF NOT EXISTS idx_payments_user_id ON payments(user_id);

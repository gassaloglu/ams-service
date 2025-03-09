package postgres

import (
	"ams-service/middlewares"
	"database/sql"
	"fmt"
	"log"
)

var CREATE_TABLE_PREFIX string = "create_tables.go"

func CreateTables(db *sql.DB) {
	queries := map[string]string{
		"credit_cards": `
            CREATE TABLE IF NOT EXISTS credit_cards (
                id SERIAL PRIMARY KEY,
                card_number VARCHAR(16) NOT NULL,
                card_holder_name VARCHAR(100) NOT NULL,
                card_holder_surname VARCHAR(100) NOT NULL,
                expiration_month INT NOT NULL,
                expiration_year INT NOT NULL,
                cvv VARCHAR(4) NOT NULL,
                card_type card_type_enum NOT NULL,
                amount DECIMAL(10,2) NOT NULL,
                currency VARCHAR(3) NOT NULL,
                issuer_bank VARCHAR(100),
                status status_enum DEFAULT 'active' NOT NULL,
                created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
            );
        `,
		"employees": `
            CREATE TABLE IF NOT EXISTS employees (
                id SERIAL PRIMARY KEY,
                employee_id VARCHAR(50) UNIQUE NOT NULL,
                name VARCHAR(50) NOT NULL,
                surname VARCHAR(50) NOT NULL,
                email VARCHAR(100) UNIQUE NOT NULL,
                phone VARCHAR(15),
                address VARCHAR(255),
                gender gender_enum NOT NULL,
                birth_date TIMESTAMP NOT NULL,
                hire_date TIMESTAMP NOT NULL,
                position VARCHAR(100) NOT NULL,
                department department_enum NOT NULL,
                salary DECIMAL(10,2) NOT NULL,
                status status_enum DEFAULT 'active' NOT NULL,
                manager_id INT,
                emergency_contact VARCHAR(100),
                emergency_phone VARCHAR(15),
                profile_image_url VARCHAR(255),
                created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
            );
        `,
		"flights": `
            CREATE TABLE IF NOT EXISTS flights (
                flight_number VARCHAR(10) PRIMARY KEY,
                departure_airport VARCHAR(3) NOT NULL,
                destination_airport VARCHAR(3) NOT NULL,
                departure_datetime TIMESTAMP NOT NULL,
                arrival_datetime TIMESTAMP NOT NULL,
                departure_gate_number VARCHAR(5),
                destination_gate_number VARCHAR(5),
                plane_registration VARCHAR(10) NOT NULL,
                status flight_status_enum NOT NULL,
                price DECIMAL(10,2) NOT NULL
            );
        `,
		"passengers": `
            CREATE TABLE IF NOT EXISTS passengers (
                id SERIAL PRIMARY KEY,
                passenger_id VARCHAR(50) UNIQUE NOT NULL,
                name VARCHAR(50) NOT NULL,
                surname VARCHAR(50) NOT NULL,
                email VARCHAR(100) UNIQUE NOT NULL,
                phone VARCHAR(15),
                address VARCHAR(255),
                gender gender_enum NOT NULL,
                birth_date TIMESTAMP NOT NULL,
                passport_number VARCHAR(20) UNIQUE NOT NULL,
                nationality VARCHAR(50) NOT NULL,
                frequent_flyer_number VARCHAR(50),
                created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
            );
        `,
	}

	for tableName, query := range queries {
		if _, err := db.Exec(query); err != nil {
			log.Fatalf("Failed to create %s table: %v", tableName, err)
		} else {
			middlewares.LogInfo(fmt.Sprintf("%s - %s table created successfully or already exists.", CREATE_TABLE_PREFIX, tableName))
		}
	}
}

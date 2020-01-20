package migration

import (
	"database/sql"

	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(Up20200119130812, Down20200119130812)
}

func createExtensions(tx *sql.Tx) error {
	_, err := tx.Exec(`
		CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
	`)
	return err
}

func createRecordType(tx *sql.Tx) error {
	_, err := tx.Exec(`
		CREATE TYPE "record_type" AS ENUM (
			'Asset',
			'Liability'
		);
	`)
	return err
}

func createUsersTable(tx *sql.Tx) error {
	_, err := tx.Exec(`
		CREATE TABLE users (
			id uuid NOT NULL DEFAULT uuid_generate_v4(),
			username varchar(255) NOT NULL UNIQUE,
			password varchar(255) NOT NULL,
			full_name varchar(255) NOT NULL,
			"created_at" timestamp with time zone NOT NULL DEFAULT NOW(),
			"updated_at" timestamp with time zone NOT NULL DEFAULT NOW(),
			PRIMARY KEY(id)
			);
	`)
	return err
}

func createRecordsTable(tx *sql.Tx) error {
	_, err := tx.Exec(`
		CREATE TABLE records (
			id uuid NOT NULL DEFAULT uuid_generate_v4(),
			type "record_type" NOT NULL,
			name varchar(255) NOT NULL,
			balance decimal(19, 4) NOT NULL,
			"created_at" timestamp with time zone NOT NULL DEFAULT NOW(),
			"updated_at" timestamp with time zone NOT NULL DEFAULT NOW(),
			user_id uuid REFERENCES users (id) ON UPDATE CASCADE ON DELETE CASCADE,
			PRIMARY KEY(id)
			);
	`)
	return err
}

// Up20200119130812 creates the users and records tables and adds the uuid extension to the database
func Up20200119130812(tx *sql.Tx) error {
	if err := createExtensions(tx); err != nil {
		return err
	}

	if err := createRecordType(tx); err != nil {
		return err
	}

	if err := createUsersTable(tx); err != nil {
		return err
	}

	if err := createRecordsTable(tx); err != nil {
		return err
	}
	return nil
}

// Down20200119130812 drops the records and users table and the record_type enum
func Down20200119130812(tx *sql.Tx) error {
	if _, err := tx.Exec("DROP TABLE records;"); err != nil {
		return err
	}

	if _, err := tx.Exec("DROP TABLE users;"); err != nil {
		return err
	}

	if _, err := tx.Exec("DROP TYPE record_type;"); err != nil {
		return err
	}
	return nil
}

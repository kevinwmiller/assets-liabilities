package migration

import (
	"database/sql"

	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(Up20200119130812, Down20200119130812)
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

func createRecordsTable(tx *sql.Tx) error {
	_, err := tx.Exec(`
		CREATE TABLE records (
			id varchar(255) NOT NULL,
			type "record_type" NOT NULL,
			name varchar(255) NOT NULL,
			balance decimal(19, 4) NOT NULL,
			"created_at" timestamp with time zone NOT NULL DEFAULT NOW(),
			"updated_at" timestamp with time zone NOT NULL DEFAULT NOW(),
			PRIMARY KEY(id)
			);
	`)
	return err
}

// Up20200119130812 creates the users and records tables
func Up20200119130812(tx *sql.Tx) error {
	if err := createRecordType(tx); err != nil {
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

	if _, err := tx.Exec("DROP TYPE record_type;"); err != nil {
		return err
	}
	return nil
}

package main

import (
	_ "github.com/mattn/go-sqlite3"
	"database/sql"
	"os"
)

// create the connection to the database
func (driver *sqliteDriver) connect(dsn string) error {
	if nil != driver.connection {
		return nil
	}

	mustCreateSchema := schemaExists(dsn)

	var err error
	driver.connection, err = sql.Open("sqlite3", dsn)
	if nil != err {
		return err
	}

	if mustCreateSchema {
		err = createSchema(driver.connection)
	}

	return err
}

// check if the database schema already exists or if it needs to be created
func schemaExists(dsn string) bool {
	_, err := os.Stat(dsn)
	return os.IsNotExist(err)
}

// setup the schema for a new database
func createSchema(connection *sql.DB) error {
	_, err := connection.Exec(SCHEMA)

	return err
}

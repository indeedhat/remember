package main

import (
	"database/sql"
	"github.com/indeedhat/remember"
	"time"
)

// exports
var Driver = sqliteDriver{}

type scannableRow interface {
	Scan(dest ...interface{}) error
}

type sqliteDriver struct {
	connection *sql.DB
}

func (driver *sqliteDriver) Init(dsn string) error {
	return driver.connect(dsn)
}

func (driver *sqliteDriver) Save(object *remember.DataObject) error {
	if 0 != object.ID {
		return updateObject(object, driver.connection)
	}

	return insertObject(object, driver.connection)
}

func (driver *sqliteDriver) Load(id int64) (*remember.DataObject, error) {
	statement, err := driver.connection.Prepare(FETCH)
	if nil != err {
		return nil, err
	}

	defer statement.Close()

	row := statement.QueryRow(id)

	return fetchObject(row)
}

func (driver *sqliteDriver) List(groupId int) ([]*remember.DataObject, error) {
	var results []*remember.DataObject

	statement, err := driver.connection.Prepare(LIST)
	if nil != err {
		return results, err
	}

	defer statement.Close()

	rows, err := statement.Query(groupId)
	if nil != err {
		return results, err
	}

	defer rows.Close()
	for rows.Next() {
		result, err := fetchObject(rows)
		if nil != err {
			return results, err
		}
		results = append(results, result)
	}

	return results, nil
}

func (driver *sqliteDriver) ListAll() ([]*remember.DataObject, error) {
	var results []*remember.DataObject

	statement, err := driver.connection.Prepare(LIST_ALL)
	if nil != err {
		return results, err
	}

	defer statement.Close()

	rows, err := statement.Query()
	if nil != err {
		return results, err
	}

	defer rows.Close()
	for rows.Next() {
		result, err := fetchObject(rows)
		if nil != err {
			return results, err
		}
		results = append(results, result)
	}

	return results, nil
}

func (driver *sqliteDriver) Delete(object *remember.DataObject) error {
	statement, err := driver.connection.Prepare(DELETE)
	if nil != err {
		return err
	}

	defer statement.Close()

	_, err = statement.Exec(object.ID)
	return err
}

func (driver *sqliteDriver) Close() error {
	return driver.connection.Close()
}

// fetch a row from the cursor and load it into a struct
func fetchObject(row scannableRow) (*remember.DataObject, error) {
	object := &remember.DataObject{}

	var created int64
	var updated int64

	err := row.Scan(
		&object.ID,
		&object.Title,
		&object.GroupId,
		&object.Payload,
		created,
		updated,
	)

	object.CreatedAt = time.Unix(created, 0)
	object.UpdatedAt = time.Unix(updated, 0)

	if nil != err {
		return nil, err
	}

	return object, nil
}

// insert a new data object
// manage time stamps and ID
func insertObject(object *remember.DataObject, connection *sql.DB) error {
	statement, err := connection.Prepare(CREATE)
	if nil != err {
		return err
	}
	defer statement.Close()

	object.CreatedAt = time.Now()
	object.UpdatedAt = time.Now()

	response, err := statement.Exec(
		object.Title,
		object.GroupId,
		object.Payload,
		object.CreatedAt.Unix(),
		object.UpdatedAt.Unix(),
	)
	if nil != err {
		return err
	}

	object.ID, err = response.LastInsertId()
	return err
}

// Update an existing object
// manage updated time
func updateObject(object *remember.DataObject, connection *sql.DB) error {
	statement, err := connection.Prepare(UPDATE)
	if nil != err {
		return err
	}
	defer statement.Close()

	object.UpdatedAt = time.Now()

	_, err = statement.Exec(
		object.Title,
		object.GroupId,
		object.Payload,
		object.UpdatedAt.Unix(),
		object.ID,
	)

	return err
}

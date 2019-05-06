package main

const SCHEMA = `
CREATE TABLE data_object (
	title text,
	group_id integer
    payload text,
	created_at integer
    updated_at integer
);

CREATE INDEX group_id ON data_object(group_id);

CREATE INDEX created_at ON data_object(created_at);

CREATE INDEX updated_at ON data_object(updated_at);
`

const FETCH = `
SELECT
	id,
	title,
	group_id,
	payload,
	created_at,
	updated_at
FROM 
	data_object 
WHERE 
	rowid = ?
`

const LIST = `
SELECT
	id,
	title,
	group_id,
	payload,
	created_at,
	updated_at
FROM 
	data_object 
WHERE 
	group_id = ?
`

const LIST_ALL = `
SELECT
	id,
	title,
	group_id,
	payload,
	created_at,
	updated_at
FROM 
	data_object 
`

const DELETE = `
DELETE FROM
	data_object
WHERE
	rowid = ?
`

const CREATE = `
INSERT INTO data_object (
	title,
	group_id,
	payload,
	create_at,
	updated_at
) VALUES (
	?,
	?,
	?,
	?,
	?
)
`

const UPDATE = `
UPDATE
	data_object
SET
	title = ?,
	group_id = ?,
	payload = ?,
	updated_at = ?
WHERE
	rowid = ?
`

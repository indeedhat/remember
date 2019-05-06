package remember

import "time"

type StorageDriver interface {
	// This method will be called on load, it should setup any environment requirements required for the driver to operate
	// Note: this will be called on each load so should only setup requirements if not already done so
	Init(dsn string) error

	// Create or update a data object to the storage engine
	// Should apply the ID value of the struct if the object is newly created
	Save(ob *DataObject) error

	// Load a data object by its id from the storage engine
	Load(id int64) (*DataObject, error)

	// List all data objects by their group id
	List(groupId int)([]*DataObject, error)

	// List all saved data objects
	ListAll() ([]*DataObject, error)

	// Remove a data object from the store
	Delete(ob *DataObject) error

	// Will be called as a shutdown method in case the driver needs to do any cleanup
	Close() error
}

type DataObject struct {
	// This will be the objects unique id
	ID        int64
	Title     string
	GroupId   int
	Payload   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

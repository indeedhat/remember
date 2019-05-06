package remember

import (
	"errors"
	"plugin"
)


// Load a storage driver by path from is shared object
func LoadDriver(path, dsn string) (StorageDriver, error){
	p, err := plugin.Open(path)
	if nil != err {
		return nil, err
	}

	symbol, err := p.Lookup("Driver")
	if nil != err {
		return nil, err
	}

	driver, ok := symbol.(StorageDriver)
	if !ok {
		return nil, errors.New("symbol is not a valid StorageDriver")
	}

	err = driver.Init(dsn)
	if nil != err {
		return nil, err
	}

	return driver, nil
}

// Migrate all data objects from one storage engine to another
func MigrateDataObjects(from StorageDriver, to StorageDriver) error {
	objects, err := from.ListAll()
	if nil != err {
		return err
	}

	for _,object := range objects {
		err := to.Save(object)
		if nil != err {
			return err
		}
	}

	return nil
}

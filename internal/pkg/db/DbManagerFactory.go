package db

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/guiyomh/charlatan/internal/pkg/db/contracts"
	"github.com/guiyomh/charlatan/internal/pkg/db/drivers"
)

var (
	// ErrDbDriverIsNotSupported is returned when the db driver is not supported
	ErrDbDriverIsNotSupported = errors.New("This db driver is not supported")

	// ErrWrongCastNotAMap is returned when a map is not a map[interface{}]interface{}
	ErrWrongCastNotAMap = errors.New("Could not cast record: not a map[interface{}]interface{}")

	// ErrKeyIsNotString is returned when a record is not of type string
	ErrKeyIsNotString = errors.New("Record map key is not string")
)

type DbManagerFactory struct{}

func (dm DbManagerFactory) NewDbManager(
	driverName string,
	host string,
	port int16,
	username string,
	password string,
) (contracts.DbManager, error) {
	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%d)/?charset=utf8", username, password, host, port)
	myDb, err := sql.Open("mysql", dataSource)
	if err != nil {
		panic(err.Error())
	}
	switch driverName {
	case "mysql":
		return drivers.NewMySQL(myDb), nil
	}
	return nil, ErrDbDriverIsNotSupported
}

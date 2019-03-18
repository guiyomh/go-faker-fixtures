package contracts

import "fmt"

// InsertError will be returned if any error happens on database while
// inserting the record
type InsertError struct {
	Err    error
	SQL    string
	Params []interface{}
}

func (i InsertError) Error() string {
	return fmt.Sprintf(
		"charlatan: error inserting record: %v, sql: %s, params: %v",
		i.Err,
		i.SQL,
		i.Params,
	)
}

package contracts

import "database/sql"

type DbManager interface {
	BuildInsertSQL(schema string, table string, fields map[string]interface{}) (string, []interface{}, error)
	TruncateTable(schema string, table string) (sql.Result, error)
	Exec(sqlStr string, params []interface{}) (sql.Result, error)
}

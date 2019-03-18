package drivers

import (
	"database/sql"
	"fmt"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/guiyomh/charlatan/internal/pkg/db/contracts"
	"github.com/guiyomh/charlatan/internal/pkg/db/utils"
)

//MySQL Driver for MySQL and MariaDB
type mySQL struct {
	connection *sql.DB
}

func NewMySQL(conn *sql.DB) contracts.DbManager {
	return &mySQL{connection: conn}
}

func (m *mySQL) escape(str string) string {
	return fmt.Sprintf("`%s`", str)
}

//BuildInsertSQL Convert a record map to a insert string
func (m *mySQL) BuildInsertSQL(schema string, table string, record map[string]interface{}) (sqlStr string, values []interface{}, err error) {
	var (
		sqlColumns []string
		sqlValues  []string
		i          = 1
	)
	for key, value := range record {
		// keyStr, ok := key.(string)
		// if !ok {
		// 	err = ErrKeyIsNotString
		// 	return
		// }
		sqlColumns = append(sqlColumns, m.escape(key))

		// if string, try convert to SQL or time
		// if map or array, convert to json
		switch v := value.(type) {
		case string:
			if strings.HasPrefix(v, "RAW=") {
				sqlValues = append(sqlValues, strings.TrimPrefix(v, "RAW="))
				continue
			}

			if t, err := utils.TryStrToDate(v); err == nil {
				value = t
			}
		case []interface{}, map[interface{}]interface{}:
			value = utils.RecursiveToJSON(v)
		}

		sqlValues = append(sqlValues, "?")

		values = append(values, value)
		i++
	}

	sqlStr = fmt.Sprintf(
		"INSERT INTO %s.%s (%s) VALUES (%s)",
		m.escape(schema),
		m.escape(table),
		strings.Join(sqlColumns, ", "),
		strings.Join(sqlValues, ", "),
	)
	return
}

func (m *mySQL) Exec(sqlStr string, params []interface{}) (sql.Result, error) {
	tx, err := m.connection.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	if _, err = tx.Exec("SET FOREIGN_KEY_CHECKS = 0"); err != nil {
		return nil, err
	}
	result, err := tx.Exec(sqlStr, params...)

	if err != nil {
		return nil, &contracts.InsertError{
			Err:    err,
			SQL:    sqlStr,
			Params: params,
		}
	}
	_, err2 := tx.Exec("SET FOREIGN_KEY_CHECKS = 1")

	if err2 != nil {
		return nil, err2
	}

	if err = tx.Commit(); err != nil {
		return nil, err
	}
	return result, nil
}

func (m *mySQL) TruncateTable(schema string, table string) (sql.Result, error) {
	sqlStr := fmt.Sprintf(
		"TRUNCATE TABLE %s.%s",
		m.escape(schema),
		m.escape(table),
	)

	return m.connection.Exec(sqlStr)
}

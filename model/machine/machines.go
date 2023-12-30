// Package machine provides access to the machines' table in DB.
package machine

import (
	"database/sql"
	"fmt"
)

// Item describes the model of the machine' table.
type Item struct {
	ID     uint32 `db:"id"`
	Name   string `db:"name"`
	OS     string `db:"os"`
	IP     string `db:"ip"`
	Status string `db:"status"`
}

var table = "machines"

// Connection is an interface for making queries.
type Connection interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	Get(dest interface{}, query string, args ...interface{}) error
	Select(dest interface{}, query string, args ...interface{}) error
}

// ByID gets an item by its ID.
//func ByID(db Connection, ID string) (Item, bool, error)  {

//}

// Create adds a new item.
func Create(db Connection, name, os, ip, status string) (sql.Result, error) {
	result, err := db.Exec(fmt.Sprintf(`INSERT INTO %v (name, os, ip, status) VALUES (?,?,?,?)`, table), name, os, ip, status)

	return result, err
}

// Update changes an existing item.
func Update(db Connection, ID, status string) (sql.Result, error) {
	result, err := db.Exec(fmt.Sprintf(`UPDATE %v SET status = ? WHERE ID = ?`, table), status, ID)

	return result, err

}

// SelectAll gets all data from the machine table.
func SelectAll(db Connection) ([]Item, error){
	var result []Item
	err := db.Select(&result, fmt.Sprintf(`SELECT * FROM %v`, table))

	return  result, err
}

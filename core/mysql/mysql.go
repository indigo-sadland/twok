package mysql

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// Values holds parameters for MySQL connection
type Values struct {
	Username  string `json:"Username"`
	Password  string `json:"Password"`
	Database  string `json:"Database"`
	Charset   string `json:"Charset"`
	Collation string `json:"Collation"`
	Hostname  string `json:"Hostname"`
	Parameter string `json:"Parameter"`
	Port      int    `json:"Port"`
}

func (v Values) Connect() (*sqlx.DB, error) {
	dbExists := true
	var dbCon *sqlx.DB

	//Connect to existing DB.
	dbCon, err := sqlx.Connect("mysql", v.dsn(dbExists))
	// If the database doesn't exist or can't connect
	if err != nil {
		// Close the open connection (since 'unknown database' is still an
		// active connection)
		if dbCon != nil {
			dbCon.Close()
		}

		dbExists = false
		// Connect to MySQL without DB.
		dbCon, err = sqlx.Connect("mysql", v.dsn(dbExists))
		if err != nil {
			return nil, err
		}
		err = v.Create(dbCon)
		if err != nil {
			return nil, err
		}
	}

	return dbCon, nil
}

// Create a new database.
func (v Values) Create(sql *sqlx.DB) error {

	// Create the database
	_, err := sql.Exec(fmt.Sprintf(`CREATE DATABASE %v
				DEFAULT CHARSET = %v
				COLLATE = %v
				;`, v.Database, v.Charset, v.Collation))
	return err
}

// dsn returns Data Source Name string.
func (v Values) dsn(dbExists bool) string {
	var dsnStr string

	// Example: charset=utf8mb4&collation=utf8mb4_unicode_ci&parseTime=true
	params := fmt.Sprintf("charset=%v&collation=%v&%v", v.Charset, v.Collation, v.Parameter)

	if dbExists {
		// Example: root:password@tcp(localhost:3306)/test?params
		dsnStr = fmt.Sprintf("%v:%v@tcp(%v:%d)/%v?%s", v.Username, v.Password, v.Hostname, v.Port,
			v.Database, params)
	} else {
		dsnStr = fmt.Sprintf("%v:%v@tcp(%v:%d)/?%s", v.Username, v.Password, v.Hostname, v.Port, params)
	}

	return dsnStr
}

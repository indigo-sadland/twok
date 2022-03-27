// Package holder is a context that can be used by controllers to access application settings.
package holder

import (
	"github.com/indigo-sadland/twok/config"
	"github.com/jmoiron/sqlx"
	"sync"
)

var (
	configValues config.Values
	dbInfo       *sqlx.DB
	mutex        sync.RWMutex
)

// Values structures the application settings.
type Values struct {
	Config config.Values
	//Sess   *sessions.Session
	DB *sqlx.DB
}

// StoreConfig stores the application settings so controller functions can
//access them safely.
func StoreConfig(cv config.Values) {
	mutex.Lock()
	configValues = cv
	mutex.Unlock()
}

// StoreDB stores the database connection settings so controller functions can
// access them safely.
func StoreDB(db *sqlx.DB) {
	mutex.Lock()
	dbInfo = db
	mutex.Unlock()
}

// Context returns the application settings.
func Context() Values {
	//var id string

	//// Get the session
	//sess, err := configValues.Session.Instance(r)
	//
	//// If the session is valid
	//if err == nil {
	//	// Get the user id
	//	id = fmt.Sprintf("%v", sess.Values["id"])
	//}

	mutex.RLock()
	i := Values{
		Config: configValues,
		//Sess:   sess,
		DB: dbInfo,
	}
	mutex.RUnlock()

	return i
}

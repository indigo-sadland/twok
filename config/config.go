// Package config holds functions that handle parsing of the config file.
package config

import (
	"encoding/json"
	"io/ioutil"

	"github.com/indigo-sadland/twok/core/assets"
	"github.com/indigo-sadland/twok/core/mysql"
	"github.com/indigo-sadland/twok/core/server"
)

// Values structures the application settings.
type Values struct {
	Assets assets.Values `json:"Assets"`
	MySQL  mysql.Values  `json:"MySQL"`
	Server server.Values `json:"Server"`
	//	Session     session.Info  `json:"Session"`
}

// Parser must implement ParseJSON.
type Parser interface {
	ParseJSON([]byte) error
}

// ParseJSON unmarshals bytes to structs.
func (v *Values) ParseJSON(b []byte) error {
	return json.Unmarshal(b, &v)
}

// Load the JSON config file.
func Load(configFile string) (*Values, error) {
	// Create a new instance of Values.
	v := &Values{}

	err := parse(configFile, v)

	return v, err

}

func parse(configFile string, p Parser) error {
	// Read the config file
	jsonBytes, err := ioutil.ReadFile(configFile)
	if err != nil {
		return err
	}

	// Parse the config
	if err := p.ParseJSON(jsonBytes); err != nil {
		return err
	}

	return nil
}

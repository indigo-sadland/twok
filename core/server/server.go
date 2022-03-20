package server

import "strconv"

// Values boot's settings.
type Values struct {
	Hostname string
	HTTPPort int
}

func (v Values) GetAddress() string {
	address := v.Hostname + ":" + strconv.Itoa(v.HTTPPort)
	return address
}

// Package controllers handles requests for each of the controllers.
package controllers

import (
	"github.com/indigo-sadland/twok/controllers/home"
	"github.com/indigo-sadland/twok/controllers/machines"
)

// LoadRoutes loads the routes for each of the controllers.
func LoadRoutes() {
	home.Load()
	machines.Load()
}

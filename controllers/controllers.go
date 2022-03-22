package controllers

import (
	"github.com/indigo-sadland/twok/controllers/home"
	"github.com/indigo-sadland/twok/controllers/machines"
)

func Load() {
	home.Load()
	machines.Load()

}

package main

import (
	"github.com/indigo-sadland/twok/boot"
	"github.com/indigo-sadland/twok/config"
)

func main() {
	conf, _ := config.Load("env.json")

	boot.Initialize(conf)
}

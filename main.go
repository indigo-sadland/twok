package main

import (
	"twok/boot"
	"twok/config"
)

func main() {

	conf, _ := config.Load("envt.json")

	boot.Initialize(conf)

}

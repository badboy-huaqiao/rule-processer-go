package main

import (
	"log"
)

var MyPlu MyPlugin

type MyPlugin struct {
}

func (m MyPlugin) Hello() {
	log.Println("Plugin_1.....my plugin")
}

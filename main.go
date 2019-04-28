package main

import (
	"plugin"
	processer "rule-processer-go/processer"
)

func loadPlugin(plugin_path string) *plugin.Plugin {
	plug, _ := plugin.Open(plugin_path)
	return plug
}

func main() {
	plug := loadPlugin("./plugins/plugin_1.so")

	symHello, _ := plug.Lookup("Processer")
	myPlu, _ := symHello.(processer.Processer)
	myPlu.Process()
}

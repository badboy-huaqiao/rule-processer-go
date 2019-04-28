package main

import (
	"plugin"
	message "rule-processer-go/message"
	processor "rule-processer-go/processer"
)

func loadPlugin(plugin_path string) *plugin.Plugin {
	plug, _ := plugin.Open(plugin_path)
	return plug
}

func parseInputs(inputs []model.Input,
	dataCh <-chan []model.Reading,
	filterCh chan []model.Reading) {

	for readings := range dataCh {
		
	}
}

func handleOutputs(outputs []model.Output) {

}

func startupPluin(path string, name string,
	dataCh <-chan []model.Reading) {

	plug := loadPlugin("./plugins/examples/myplugin.so")

	symPlugin, _ := plug.Lookup("MyPlu")
	processorPlugin, _ := symPlu.(processor.Processor)

	filterCh := make(chan []model.Reading, 50)

	inputs := processorPlugin.SetInput()

	parseInputs(inputs, dataCh, filterCh)

	processorPlugin.Process(filterCh)

	outputs := processorPlugin.SetOutput()

	handleOutputs(outputs)
}

func main() {
	dataCh := make(chan []model.Reading, 500)
	message.ChanRegister("",dataCh)
	go startupPluin("", "", dataCh)
	go message.SubscribeFromZMQ()
}

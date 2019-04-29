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

// func isExistPorperty(reading model.Reading) bool {
// 	for _,property := range input.DevicePropertiesName {
//
// 	}
// }

func isExistDevice(inputs []model.Input,reading model.Reading) bool {
	for _,input := range inputs {
		if input.DeviceName == reading.Device &&
			 input.Severity == "normal" {
			return true
		}
	}
	return false
}

func parseInputs(inputs []model.Input,
	dataCh <-chan []model.Reading,
	filterCh chan []model.Reading) {

	go func(){
		
	}()

	go func(){
		for readings := range dataCh {
			for _,reading := range readings {
				if isExistDevice(inputs,reading) {
					filterCh <- readings
					break
				}
			}
		}
	}()
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

	var once bool = true
	for once {
		parseInputs(inputs, dataCh, filterCh)

		once = processorPlugin.Process(filterCh)

		outputs := processorPlugin.SetOutput()

		handleOutputs(outputs)
	}

}

func main() {
	dataCh := make(chan []model.Reading, 500)
	message.ChanRegister("",dataCh)
	go startupPluin("", "", dataCh)
	go message.SubscribeFromZMQ()
}

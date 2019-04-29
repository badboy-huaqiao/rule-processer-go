package main

import (
	"log"
	"rule-processer-go/model"
)

// 必须定义导出的全局变量，唯一，传递给rule-processer，实现接口所有方法
var MyPlu MyPlugin

var outputs = make([]model.Output,0)

type MyPlugin struct {
}

func (m MyPlugin) Process(readingCh <-chan []model.Reading) bool {

	//process data from readingCh

	var output model.Output

	output.Type = "local"
	output.Action.DeviceName = "device-01"
	output.Action.CommandName = "cmdName"
	output.Action.Method = "set"
	output.Action.Value = "1234"
	outputs = append(outputs,output)
	return false
}

func (m Myplugin) SetInput() []model.Input {
	inputs := make([]model.Input,0)
	return inputs
}

func (m Myplugin) SetOutput() []model.Output {
	return outputs
}

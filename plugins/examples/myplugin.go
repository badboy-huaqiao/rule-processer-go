package main

import (
	"log"
	"rule-processer-go/model"
)

var MyPlu MyPlugin

type MyPlugin struct {
}

func (m MyPlugin) Process(reading <-chan []model.Reading) {
	

}

func (m Myplugin) SetInput() []model.Input {
	inputs := make([]model.Input,0)
	return inputs
}

func (m Myplugin) SetOutput() []model.Output {
	outputs := make([]model.Output,0)
	return outputs
}

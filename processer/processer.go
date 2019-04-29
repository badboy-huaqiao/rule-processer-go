package processor

import (
	"rule-processer-go/model"
)

type Processor interface {
	//once return true
	Process(reading <-chan []model.Reading) bool
	SetInput() []model.Input
	SetOutput() []model.Output
}

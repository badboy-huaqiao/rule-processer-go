package processor

import (
	"rule-processer-go/model"
)

type Processor interface {
	Process(reading <-chan []model.Reading)
	SetInput() []model.Input
	SetOutput() []model.Output
}

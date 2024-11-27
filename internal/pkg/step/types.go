package step

import (
	"errors"

	"github.com/domahidizoltan/go-steps/types"
)

type (
	StepType uint8

	TempSteps struct {
		Error        error
		StepWrappers []types.StepWrapper
		Aggregator   types.ReducerFn
		Steps        []types.StepFn
		Validated    StepType
	}

	Transformator struct {
		Error      error
		Aggregator types.ReducerFn
		Steps      []types.StepFn
		Validated  StepType
	}
)

const (
	StepTypeSteps StepType = iota
	StepTypeAggregator
)

var ErrInvalidInputType = errors.New("Invalid input type")

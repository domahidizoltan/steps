package step

import (
	"errors"
	"reflect"
)

const maxArgs = 4

type ArgTypes [maxArgs]reflect.Type

type (
	StepInput struct {
		Args    [maxArgs]any
		ArgsLen uint8
	}

	StepOutput struct {
		Error   error
		Args    [maxArgs]any
		ArgsLen uint8
		Skip    bool
	}

	StepWrapper struct {
		Name     string
		StepFn   StepFn
		Validate func(prevStepArgTypes ArgTypes) (ArgTypes, error)
	}

	StepFn func(StepInput) StepOutput

	ReducerWrapper struct {
		Name      string
		ReducerFn ReducerFn
		Validate  func(prevStepArgTypes ArgTypes) (ArgTypes, error)
	}
	ReducerFn func(StepInput) StepOutput

	stepType uint8

	StepsBranch struct {
		Error             error
		StepWrappers      []StepWrapper
		AggregatorWrapper *ReducerWrapper
		Aggregator        ReducerFn
		Steps             []StepFn
		Validated         stepType
	}

	Transformator struct {
		Error               error
		Aggregator          ReducerFn
		LastAggregatedValue *StepOutput
		Steps               []StepFn
		Validated           stepType
	}
)

var (
	ErrEmptyTransformInputType = errors.New("first step in type is empty")
	ErrStepValidationFailed    = errors.New("step validation failed")
	ErrIncompatibleInArgType   = errors.New("incompatible input argument type")
)

func Steps(s ...StepWrapper) StepsBranch {
	return StepsBranch{
		StepWrappers: s,
	}
}

func (t StepsBranch) Aggregate(fn ReducerWrapper) StepsBranch { //?
	return StepsBranch{
		StepWrappers: t.StepWrappers,
		Aggregator:   fn.ReducerFn,
	}
}

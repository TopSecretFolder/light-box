package framegen

import (
	"fmt"
	"time"

	lua "github.com/yuin/gopher-lua"
)

// framegen provides an api to generate the current state
// of frames that can then be consumed by an led driver

type Pixel struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
	Z float64 `json:"z"`
}

type Frame struct {
	Pixels []Pixel `json:"pixels"`
}

type FrameContext struct {
	CurrentTime time.Time
	DeltaTime   time.Duration
}

// String implements lua.LValue.
func (f FrameContext) String() string {
	return fmt.Sprintf("{current_time = %d, delta_time = %d}", f.CurrentTime.Unix(), f.DeltaTime)
}

// Type implements lua.LValue.
func (f FrameContext) Type() lua.LValueType {
	return lua.LTTable
}

func GenFrame(c FrameContext) (*Frame, error) {

	// set the context into globals

	GlobalScriptRunner.lState.G.Global.Insert(0, c)

	err := GlobalScriptRunner.Execute()
	if err != nil {
		return nil, fmt.Errorf(
			"error generating frame: %w",
			err,
		)
	}

	resultTable, err := GlobalScriptRunner.Result()
	if err != nil {
		return nil, fmt.Errorf(
			"error getting result from script: %w",
			err,
		)
	}

	f, err := MapTable[Frame](resultTable)
	if err != nil {
		return nil, fmt.Errorf(
			"error converting resulting table to frame type: %w",
			err,
		)
	}

	return f, nil
}

package framegen

import (
	"fmt"
	"sync"

	lua "github.com/yuin/gopher-lua"
)

var GlobalScriptRunner ScriptRunner = ScriptRunner{}

type ScriptRunner struct {
	lState *lua.LState
	mu     sync.Mutex
	script string
}

func (s *ScriptRunner) SetScript(script string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.script = script
}

func (s *ScriptRunner) Execute() error {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.lState.DoString(s.script)
}

func (s *ScriptRunner) Result() (*lua.LTable, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	luaTable, ok := s.lState.Get(-1).(*lua.LTable)
	if !ok {
		return luaTable, fmt.Errorf("error getting lua table result from script")
	}

	return luaTable, nil
}

package animation

import (
	"fmt"

	"github.com/yuin/gluamapper"
	"github.com/yuin/gopher-lua"
)

func ResolveLua(script string) (Animation, error) {
	// 1. Create a new Lua state
	L := lua.NewState()
	defer L.Close()

	// 2. Execute the Lua script
	if err := L.DoString(string(script)); err != nil {
		return Animation{}, fmt.Errorf("error executing lua script: %w", err)
	}

	luaTable, ok := L.Get(-1).(*lua.LTable)
	if !ok {
		return Animation{}, fmt.Errorf("error getting resulting animation from lua")
	}

	var ani Animation
	m := gluamapper.NewMapper(gluamapper.Option{
		TagName:     "json",
		ErrorUnused: true,
		NameFunc:    func(s string) string { return s },
	})
	if err := m.Map(luaTable, &ani); err != nil {
		return Animation{}, fmt.Errorf("error mapping lua table to go struct: %w", err)
	}

	return ani, nil
}

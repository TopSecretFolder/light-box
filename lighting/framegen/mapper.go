package framegen

import (
	"github.com/yuin/gluamapper"
	"github.com/yuin/gopher-lua"
)

var mapper = gluamapper.NewMapper(gluamapper.Option{
	TagName:     "json",
	ErrorUnused: true,
	NameFunc:    func(s string) string { return s },
})

func MapTable[T any](table *lua.LTable) (*T, error) {
	t := new(T)
	err := mapper.Map(table, t)
	if err != nil {
		return nil, err
	}
	return t, nil
}

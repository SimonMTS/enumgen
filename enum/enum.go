package enum

import (
	"fmt"
)

type (
	enumValue interface {
		~int
		fmt.Stringer
	}

	Enum[T enumValue] struct {
		values     []T
		strToValue map[string]T
		intToValue map[int]T
	}
)

func Init[T enumValue](values ...T,
) *Enum[T] {
	fromInt := make(map[int]T)
	fromStr := make(map[string]T)

	for _, v := range values {
		fromInt[int(v)] = v
		fromStr[v.String()] = v
	}

	return &Enum[T]{
		values:     values,
		strToValue: fromStr,
		intToValue: fromInt,
	}
}

func (e *Enum[T]) FromInt(n int) (T, bool) {
	// TODO: does it make sense to use a map here
	v, ok := e.intToValue[n]
	if !ok {
		v = -1
	}
	return v, ok
}

func (e *Enum[T]) FromString(s string) (T, bool) {
	v, ok := e.strToValue[s]
	if !ok {
		v = -1
	}
	return v, ok
}

func (e *Enum[T]) List() []T {
	return e.values
}

package factory

import (
	"fmt"
	"math/rand"
	"reflect"
	"time"
)

type Factory[T any] struct {
	object      T
	customTypes map[string]any
}

var DefaultTypeConstructors = map[string]func() any{
	"int":        func() any { return rand.Int() },
	"int8":       func() any { return int8(rand.Intn(1 << 8)) },
	"int16":      func() any { return int16(rand.Intn(1 << 16)) },
	"int32":      func() any { return rand.Int31() },
	"int64":      func() any { return rand.Int63() },
	"uint":       func() any { return uint(rand.Uint32()) },
	"uint8":      func() any { return uint8(rand.Intn(1 << 8)) },
	"uint16":     func() any { return uint16(rand.Intn(1 << 16)) },
	"uint32":     func() any { return rand.Uint32() },
	"uint64":     func() any { return rand.Uint64() },
	"float32":    func() any { return rand.Float32() },
	"float64":    func() any { return rand.Float64() },
	"string":     func() any { return fmt.Sprintf("random-%d", rand.Intn(1000000)) },
	"bool":       func() any { return rand.Intn(2) == 1 },
	"byte":       func() any { return byte(rand.Intn(1 << 8)) },
	"rune":       func() any { return rune(rand.Intn(0x10FFFF)) },
	"complex64":  func() any { return complex(rand.Float32(), rand.Float32()) },
	"complex128": func() any { return complex(rand.Float64(), rand.Float64()) },
	"time.Time":  func() any { return time.Now() },
}

func New[T any](object T) *Factory[T] {
	return &Factory[T]{
		object:      object,
		customTypes: make(map[string]any),
	}
}

func (f *Factory[T]) Build(obj any) error {
}

// Field generates a random string of the given length.
func Field[T any]() (T, error) {
	var t T
	k := reflect.TypeOf(t).Elem().Name()

	constructor, ok := DefaultTypes[k]
	if !ok {
		return t, fmt.Errorf("no constructor found for type %s", k)
	}
	return constructor().(T), nil
}

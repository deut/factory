package factory

// Factory is a struct representing a generic factory.
type Factory struct {
	// Add fields as needed
}

// New creates and returns a new Factory instance.
func New() *Factory {
	return &Factory{}
}

// TypeConstructors maps Go type names to functions returning zero values of those types.
var TypeConstructors = map[string]func() any{
	"int":        func() any { return int(0) },
	"int8":       func() any { return int8(0) },
	"int16":      func() any { return int16(0) },
	"int32":      func() any { return int32(0) },
	"int64":      func() any { return int64(0) },
	"uint":       func() any { return uint(0) },
	"uint8":      func() any { return uint8(0) },
	"uint16":     func() any { return uint16(0) },
	"uint32":     func() any { return uint32(0) },
	"uint64":     func() any { return uint64(0) },
	"float32":    func() any { return float32(0) },
	"float64":    func() any { return float64(0) },
	"string":     func() any { return "" },
	"bool":       func() any { return false },
	"byte":       func() any { return byte(0) },
	"rune":       func() any { return rune(0) },
	"complex64":  func() any { return complex64(0) },
	"complex128": func() any { return complex128(0) },
}

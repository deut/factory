package factory

import (
	"fmt"
	"math/rand"
	"reflect"
	"time"
)

// Factory is a struct that contains the object to be built and the custom types to be used
type Factory[T any] struct {
	object      T
	customTypes map[string]func() any
}

// TypeFn is a function that returns the name of the field, the function to get the value, and the database query wrapper
// fname is the name of the field of struct
// fvalue is the function to get the value
// dbWrapper is the database query wrapper. e.g. `"` will wrap the value in quotes.
type TypeFn func() (fname string, fvalue func() any, dbWrapper rune)

// DefaultTypes is a map of default types constructors
// DefaultTypesConstructors is a map of default types constructors
var DefaultTypes = map[string]func() any{
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

// New creates a new factory for the given object and custom types
func New[T any](object T, customTypes ...func() (fn string, fv func() any)) *Factory[T] {
	f := &Factory[T]{
		object: object,
	}

	for _, customType := range customTypes {
		fn, fv := customType()
		f.customTypes[fn] = fv
	}

	return f
}

func (f *Factory[T]) Build(overrideFunks ...func() (fn string, fv func() any)) (T, error) {
	var result T

	t := reflect.TypeOf(result)
	if t.Kind() != reflect.Struct {
		return result, fmt.Errorf("type T must be a struct")
	}

	overrydeFields := make(map[string]func() any)
	for _, overrideFunk := range overrideFunks {
		fn, fv := overrideFunk()
		overrydeFields[fn] = fv
	}

	v := reflect.New(t).Elem()
	for i := 0; i < t.NumField(); i++ {
		fieldValue := v.Field(i)

		if !fieldValue.CanSet() {
			continue
		}

		fieldType := fieldValue.Kind().String()
		filedValuerFn, isFieldFound := DefaultTypes[fieldType]
		if !isFieldFound {
			return result, fmt.Errorf("unknown type %s for field %s", fieldType, fieldType)
		}

		fieldValue.Set(reflect.ValueOf(filedValuerFn()))

	}

	f.object = v.Interface().(T)

	return f.object, nil
}

// func (f *Factory[T]) SQLInsert() (string, error) {
// 	var zero T
// 	if reflect.DeepEqual(f.object, zero) {
// 		_, err := f.Build()
// 		if err != nil {
// 			return "", fmt.Errorf("failed to build model: %v", err)
// 		}
// 	}

// 	t := reflect.TypeOf(f.object)
// 	if t.Kind() != reflect.Struct {
// 		return "", fmt.Errorf("model must be a struct")
// 	}

// 	v := reflect.ValueOf(f.object)

// 	tableName := strings.ToLower(t.Name())

// 	var fields []string
// 	var values []string

// 	for i := range t.NumField() {
// 		field := t.Field(i)
// 		fieldValue := v.Field(i)

// 		if !fieldValue.CanInterface() {
// 			continue
// 		}

// 		fieldName := strings.ToLower(field.Name)

// 		var valueStr string
// 		switch fieldValue.Kind() {
// 		case reflect.String:
// 			valueStr = fmt.Sprintf("'%s'", fieldValue.String())
// 		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
// 			valueStr = fmt.Sprintf("%d", fieldValue.Int())
// 		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
// 			valueStr = fmt.Sprintf("%d", fieldValue.Uint())
// 		case reflect.Float32, reflect.Float64:
// 			valueStr = fmt.Sprintf("%f", fieldValue.Float())
// 		case reflect.Bool:
// 			valueStr = fmt.Sprintf("%t", fieldValue.Bool())
// 		case reflect.Slice:
// 			valueStr = fmt.Sprintf("'%s'", fieldValue.String())
// 		default:
// 			return "", fmt.Errorf("unsupported type %s for field %s", fieldValue.Kind().String(), fieldName)
// 		}

// 		fields = append(fields, fieldName)
// 		values = append(values, valueStr)
// 	}

// 	if len(fields) == 0 {
// 		return "", fmt.Errorf("no valid fields found in struct")
// 	}

// 	// Build SQL query
// 	query := fmt.Sprintf(
// 		"INSERT INTO %s (%s) VALUES (%s);",
// 		tableName,
// 		strings.Join(fields, ", "),
// 		strings.Join(values, ", "),
// 	)

// 	return query, nil
// }

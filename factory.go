package factory

import (
	"fmt"
	"math/rand"
	"reflect"
	"strings"
	"time"
)

type Factory[T any] struct {
	model        T
	typesMapping map[string]any
}

type Option func(any)

func (f *Factory[T]) setTypesMapping(m map[string]any) {
	f.typesMapping = m
}

func NewFactory[T any](model T, opts ...Option) *Factory[T] {
	f := &Factory[T]{model: model}
	for _, opt := range opts {
		opt(f)
	}
	return f
}

// Example option function:
func WithTypesMapping(mapping map[string]any) Option {
	return func(f any) {
		if factory, ok := f.(interface{ setTypesMapping(map[string]any) }); ok {
			factory.setTypesMapping(mapping)
		}
	}
}

func (f *Factory[T]) Build() (T, error) {
	var result T

	// Initialize random seed
	rand.Seed(time.Now().UnixNano())

	// Check if T is a struct
	t := reflect.TypeOf(result)
	if t.Kind() != reflect.Struct {
		return result, fmt.Errorf("type T must be a struct")
	}

	// Create a new instance
	v := reflect.New(t).Elem()

	// Iterate through all fields
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fieldValue := v.Field(i)

		// Check if field is settable
		if !fieldValue.CanSet() {
			continue
		}

		// Set random value based on field type
		switch fieldValue.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			fieldValue.SetInt(rand.Int63())
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			fieldValue.SetUint(rand.Uint64())
		case reflect.Float32, reflect.Float64:
			fieldValue.SetFloat(rand.Float64() * 1000)
		case reflect.String:
			fieldValue.SetString(fmt.Sprintf("random_%d", rand.Int63()))
		default:
			return result, fmt.Errorf("unknown type for field %s: %v", field.Name, fieldValue.Kind())
		}
	}

	// Save the result in the model field
	f.model = v.Interface().(T)

	return f.model, nil
}

func (f *Factory[T]) Insert() (string, error) {
	// Check if Build was called previously by checking if model has random values
	// If not, call Build first
	var zero T
	if reflect.DeepEqual(f.model, zero) {
		_, err := f.Build()
		if err != nil {
			return "", fmt.Errorf("failed to build model: %v", err)
		}
	}

	// Check if model is a struct
	t := reflect.TypeOf(f.model)
	if t.Kind() != reflect.Struct {
		return "", fmt.Errorf("model must be a struct")
	}

	v := reflect.ValueOf(f.model)

	// Get struct name for table name (convert to snake_case)
	tableName := strings.ToLower(t.Name())

	var fields []string
	var values []string

	// Iterate through all fields
	for i := range t.NumField() {
		field := t.Field(i)
		fieldValue := v.Field(i)

		// Skip unexported fields
		if !fieldValue.CanInterface() {
			continue
		}

		// Get field name (convert to snake_case)
		fieldName := strings.ToLower(field.Name)

		// Format value based on type
		var valueStr string
		switch fieldValue.Kind() {
		case reflect.String:
			valueStr = fmt.Sprintf("'%s'", fieldValue.String())
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			valueStr = fmt.Sprintf("%d", fieldValue.Int())
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			valueStr = fmt.Sprintf("%d", fieldValue.Uint())
		case reflect.Float32, reflect.Float64:
			valueStr = fmt.Sprintf("%f", fieldValue.Float())
		case reflect.Bool:
			valueStr = fmt.Sprintf("%t", fieldValue.Bool())
		default:
			// Skip unsupported types
			continue
		}

		fields = append(fields, fieldName)
		values = append(values, valueStr)
	}

	if len(fields) == 0 {
		return "", fmt.Errorf("no valid fields found in struct")
	}

	// Build SQL query
	query := fmt.Sprintf(
		"INSERT INTO %s (%s) VALUES (%s);",
		tableName,
		strings.Join(fields, ", "),
		strings.Join(values, ", "),
	)

	return query, nil
}

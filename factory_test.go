package factory

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFactory_Build(t *testing.T) {
	t.Run("with default values", func(t *testing.T) {
		type AllTypes struct {
			Int        int
			Int8       int8
			Int16      int16
			Int32      int32
			Int64      int64
			Uint       uint
			Uint8      uint8
			Uint16     uint16
			Uint32     uint32
			Uint64     uint64
			Float32    float32
			Float64    float64
			String     string
			Bool       bool
			Byte       byte
			Rune       rune
			Complex64  complex64
			Complex128 complex128
			Time       time.Time
		}

		factory := New(AllTypes{})

		obj, err := factory.Build()
		require.NoError(t, err)

		assert.NotZero(t, obj.Int)
		assert.NotZero(t, obj.Int8)
		assert.NotZero(t, obj.Int16)
		assert.NotZero(t, obj.Int32)
		assert.NotZero(t, obj.Int64)
		assert.NotZero(t, obj.Uint)
		assert.NotZero(t, obj.Uint8)
		assert.NotZero(t, obj.Uint16)
		assert.NotZero(t, obj.Uint32)
		assert.NotZero(t, obj.Uint64)
		assert.NotZero(t, obj.Float32)
		assert.NotZero(t, obj.Float64)
		assert.NotZero(t, obj.String)
		assert.NotZero(t, obj.Bool)
		assert.NotZero(t, obj.Byte)
		assert.NotZero(t, obj.Rune)
		assert.NotZero(t, obj.Complex64)
		assert.NotZero(t, obj.Complex128)
		assert.NotZero(t, obj.Time)
	})

	// t.Run("with custom type values", func(t *testing.T) {
	// 	type User struct {
	// 		Name      string
	// 		Email     string
	// 		Age       int
	// 		CreatedAt time.Time
	// 	}

	// 	factory := New(User{})

	// 	user, err := factory.Build(func() (string, func() any) {
	// 		return "CreatedAt", func() any {
	// 			return time.Now()
	// 		}
	// 	})
	// 	require.NoError(t, err)

	// 	assert.NotZero(t, user.Name)
	// 	assert.NotZero(t, user.Email)
	// 	assert.NotZero(t, user.Age)
	// 	assert.NotZero(t, user.CreatedAt)
	// })
}

package ptr

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestString(t *testing.T) {
	val := "foo"
	ptr := String(val)
	assert.Equal(t, val, *ptr)
}

func TestBool(t *testing.T) {
	val := true
	ptr := Bool(val)
	assert.Equal(t, val, *ptr)
}

func TestInt(t *testing.T) {
	val := 9223372036854775807
	ptr := Int(val)
	assert.Equal(t, val, *ptr)
}

func TestInt8(t *testing.T) {
	val := int8(127)
	ptr := Int8(val)
	assert.Equal(t, val, *ptr)
}

func TestInt16(t *testing.T) {
	val := int16(32767)
	ptr := Int16(val)
	assert.Equal(t, val, *ptr)
}

func TestInt32(t *testing.T) {
	val := int32(2147483647)
	ptr := Int32(val)
	assert.Equal(t, val, *ptr)
}

func TestInt64(t *testing.T) {
	val := int64(9223372036854775807)
	ptr := Int64(val)
	assert.Equal(t, val, *ptr)
}

func TestUint(t *testing.T) {
	val := uint(18446744073709551615)
	ptr := Uint(val)
	assert.Equal(t, val, *ptr)
}

func TestUint8(t *testing.T) {
	val := uint8(255)
	ptr := Uint8(val)
	assert.Equal(t, val, *ptr)
}

func TestUint16(t *testing.T) {
	val := uint16(65535)
	ptr := Uint16(val)
	assert.Equal(t, val, *ptr)
}

func TestUint32(t *testing.T) {
	val := uint32(4294967295)
	ptr := Uint32(val)
	assert.Equal(t, val, *ptr)
}

func TestUint64(t *testing.T) {
	val := uint64(18446744073709551615)
	ptr := Uint64(val)
	assert.Equal(t, val, *ptr)
}

func TestFloat32(t *testing.T) {
	val := float32(2.7182)
	ptr := Float32(val)
	assert.Equal(t, val, *ptr)
}

func TestFloat64(t *testing.T) {
	val := 2.7182818284
	ptr := Float64(val)
	assert.Equal(t, val, *ptr)
}

func TestByte(t *testing.T) {
	val := byte(255)
	ptr := Byte(val)
	assert.Equal(t, val, *ptr)
}

func TestRune(t *testing.T) {
	val := rune(2147483647)
	ptr := Rune(val)
	assert.Equal(t, val, *ptr)
}

func TestDuration(t *testing.T) {
	val := time.Minute
	ptr := Duration(val)
	assert.Equal(t, val, *ptr)
}

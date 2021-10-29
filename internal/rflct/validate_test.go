package rflct

import (
	"errors"
	"net/url"
	"reflect"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsStructPtr(t *testing.T) {
	tests := []struct {
		name          string
		s             interface{}
		expectedError error
	}{
		{
			"NonStruct",
			new(string),
			errors.New("non-struct type: you should pass a pointer to a struct type"),
		},
		{
			"NonPointer",
			struct{}{},
			errors.New("non-pointer type: you should pass a pointer to a struct type"),
		},
		{
			"OK",
			new(struct{}),
			nil,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			v, err := IsStructPtr(tc.s)

			if tc.expectedError == nil {
				assert.NotNil(t, v)
				assert.NoError(t, err)
			} else {
				assert.Empty(t, v)
				assert.Equal(t, tc.expectedError, err)
			}
		})
	}
}

func TestIsNestedStruct(t *testing.T) {
	vStruct := reflect.ValueOf(struct {
		Int    int
		URL    url.URL
		Regexp regexp.Regexp
		Nested struct {
			String string
		}
	}{})

	vInt := vStruct.FieldByName("Int")
	assert.False(t, IsNestedStruct(vInt.Type()))

	vURL := vStruct.FieldByName("URL")
	assert.False(t, IsNestedStruct(vURL.Type()))

	vRegexp := vStruct.FieldByName("Regexp")
	assert.False(t, IsNestedStruct(vRegexp.Type()))

	vNested := vStruct.FieldByName("Nested")
	assert.True(t, IsNestedStruct(vNested.Type()))
}

func TestIsStructSupported(t *testing.T) {
	tests := []struct {
		name     string
		s        interface{}
		expected bool
	}{
		{"NotSupported", struct{}{}, false},
		{"URL", url.URL{}, true},
		{"Regexp", regexp.Regexp{}, true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tStruct := reflect.TypeOf(tc.s)

			assert.Equal(t, tc.expected, IsStructSupported(tStruct))
		})
	}
}

func TestIsTypeSupported(t *testing.T) {
	var f Flags

	tests := []struct {
		name              string
		field             interface{}
		expectedSupported bool
	}{
		{"String", f.Value.String, true},
		{"Bool", f.Value.Bool, true},
		{"Int", f.Value.Int, true},
		{"Int8", f.Value.Int8, true},
		{"Int16", f.Value.Int16, true},
		{"Int32", f.Value.Int32, true},
		{"Int64", f.Value.Int64, true},
		{"Uint", f.Value.Uint, true},
		{"Uint8", f.Value.Uint8, true},
		{"Uint16", f.Value.Uint16, true},
		{"Uint32", f.Value.Uint32, true},
		{"Uint64", f.Value.Uint64, true},
		{"Float32", f.Value.Float32, true},
		{"Float64", f.Value.Float64, true},
		{"Byte", f.Value.Byte, true},
		{"Rune", f.Value.Rune, true},
		{"Duration", f.Value.Duration, true},
		{"URL", f.Value.URL, true},
		{"Regexp", f.Value.Regexp, true},
		{"String", f.Pointer.String, true},
		{"Bool", f.Pointer.Bool, true},
		{"IntPointer", f.Pointer.Int, true},
		{"Int8Pointer", f.Pointer.Int8, true},
		{"Int16Pointer", f.Pointer.Int16, true},
		{"Int32Pointer", f.Pointer.Int32, true},
		{"Int64Pointer", f.Pointer.Int64, true},
		{"UintPointer", f.Pointer.Uint, true},
		{"Uint8Pointer", f.Pointer.Uint8, true},
		{"Uint16Pointer", f.Pointer.Uint16, true},
		{"Uint32Pointer", f.Pointer.Uint32, true},
		{"Uint64Pointer", f.Pointer.Uint64, true},
		{"Float32Pointer", f.Pointer.Float32, true},
		{"Float64Pointer", f.Pointer.Float64, true},
		{"BytePointer", f.Pointer.Byte, true},
		{"RunePointer", f.Pointer.Rune, true},
		{"DurationPointer", f.Pointer.Duration, true},
		{"URLPointer", f.Pointer.URL, true},
		{"RegexpPointer", f.Pointer.Regexp, true},
		{"StringSlice", f.Slice.String, true},
		{"BoolSlice", f.Slice.Bool, true},
		{"IntSlice", f.Slice.Int, true},
		{"Int8Slice", f.Slice.Int8, true},
		{"Int16Slice", f.Slice.Int16, true},
		{"Int32Slice", f.Slice.Int32, true},
		{"Int64Slice", f.Slice.Int64, true},
		{"UintSlice", f.Slice.Uint, true},
		{"Uint8Slice", f.Slice.Uint8, true},
		{"Uint16Slice", f.Slice.Uint16, true},
		{"Uint32Slice", f.Slice.Uint32, true},
		{"Uint64Slice", f.Slice.Uint64, true},
		{"Float32Slice", f.Slice.Float32, true},
		{"Float64Slice", f.Slice.Float64, true},
		{"ByteSlice", f.Slice.Byte, true},
		{"RuneSlice", f.Slice.Rune, true},
		{"DurationSlice", f.Slice.Duration, true},
		{"URLSlice", f.Slice.URL, true},
		{"RegexpSlice", f.Slice.Regexp, true},
		{"NotSupported", f.Unsupported, false},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			supported := IsTypeSupported(reflect.TypeOf(tc.field))

			assert.Equal(t, tc.expectedSupported, supported)
		})
	}
}

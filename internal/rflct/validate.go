package rflct

import (
	"errors"
	"net/url"
	"reflect"
	"regexp"
	"time"
)

type (
	// Value is a helper type for testing purposes.
	Value struct {
		String   string        `flag:"string" ask:"email"`
		Bool     bool          `flag:"bool" ask:"any,bool value"`
		Int      int           `flag:"int" ask:"any,int value"`
		Int8     int8          `flag:"int8" ask:"any,int8 value"`
		Int16    int16         `flag:"int16" ask:"any,int16 value"`
		Int32    int32         `flag:"int32" ask:"any,int32 value"`
		Int64    int64         `flag:"int64" ask:"any,int64 value"`
		Uint     uint          `flag:"uint" ask:"any,uint value"`
		Uint8    uint8         `flag:"uint8" ask:"any,uint8 value"`
		Uint16   uint16        `flag:"uint16" ask:"any,uint16 value"`
		Uint32   uint32        `flag:"uint32" ask:"any,uint32 value"`
		Uint64   uint64        `flag:"uint64" ask:"any,uint64 value"`
		Float32  float32       `flag:"float32" ask:"any,float32 value"`
		Float64  float64       `flag:"float64" ask:"any,float64 value"`
		Byte     byte          `flag:"byte" ask:"any,byte value"`
		Rune     rune          `flag:"rune" ask:"any,rune value"`
		Duration time.Duration `flag:"duration" ask:"any,duration value"`
		URL      url.URL       `flag:"url,the help text" ask:"any,url value"`
		Regexp   regexp.Regexp `flag:"regexp,the help text" ask:"any,regexp value"`
	}

	// Pointer is a helper type for testing purposes.
	Pointer struct {
		String   *string        `flag:"string-pointer" ask:"secret"`
		Bool     *bool          `flag:"bool-pointer" ask:"any,bool pointer"`
		Int      *int           `flag:"int-pointer" ask:"any,int pointer"`
		Int8     *int8          `flag:"int8-pointer" ask:"any,int8 pointer"`
		Int16    *int16         `flag:"int16-pointer" ask:"any,int16 pointer"`
		Int32    *int32         `flag:"int32-pointer" ask:"any,int32 pointer"`
		Int64    *int64         `flag:"int64-pointer" ask:"any,int64 pointer"`
		Uint     *uint          `flag:"uint-pointer" ask:"any,uint pointer"`
		Uint8    *uint8         `flag:"uint8-pointer" ask:"any,uint8 pointer"`
		Uint16   *uint16        `flag:"uint16-pointer" ask:"any,uint16 pointer"`
		Uint32   *uint32        `flag:"uint32-pointer" ask:"any,uint32 pointer"`
		Uint64   *uint64        `flag:"uint64-pointer" ask:"any,uint64 pointer"`
		Float32  *float32       `flag:"float32-pointer" ask:"any,float32 pointer"`
		Float64  *float64       `flag:"float64-pointer" ask:"any,float64 pointer"`
		Byte     *byte          `flag:"byte-pointer" ask:"any,byte pointer"`
		Rune     *rune          `flag:"rune-pointer" ask:"any,rune pointer"`
		Duration *time.Duration `flag:"duration-pointer" ask:"any,duration pointer"`
		URL      *url.URL       `flag:"url-pointer,the help text" ask:"any,url pointer"`
		Regexp   *regexp.Regexp `flag:"regexp-pointer,the help text" ask:"any,regexp pointer"`
	}

	// Slice is a helper type for testing purposes.
	Slice struct {
		String   []string        `flag:"string-slice" ask:"any,string slice"`
		Bool     []bool          `flag:"bool-slice" ask:"any,bool slice"`
		Int      []int           `flag:"int-slice" ask:"any,int slice"`
		Int8     []int8          `flag:"int8-slice" ask:"any,int8 slice"`
		Int16    []int16         `flag:"int16-slice" ask:"any,int16 slice"`
		Int32    []int32         `flag:"int32-slice" ask:"any,int32 slice"`
		Int64    []int64         `flag:"int64-slice" ask:"any,int64 slice"`
		Uint     []uint          `flag:"uint-slice" ask:"any,uint slice"`
		Uint8    []uint8         `flag:"uint8-slice" ask:"any,uint8 slice"`
		Uint16   []uint16        `flag:"uint16-slice" ask:"any,uint16 slice"`
		Uint32   []uint32        `flag:"uint32-slice" ask:"any,uint32 slice"`
		Uint64   []uint64        `flag:"uint64-slice" ask:"any,uint64 slice"`
		Float32  []float32       `flag:"float32-slice" ask:"any,float32 slice"`
		Float64  []float64       `flag:"float64-slice" ask:"any,float64 slice"`
		Byte     []byte          `flag:"byte-slice" ask:"any,byte slice"`
		Rune     []rune          `flag:"rune-slice" ask:"any,rune slice"`
		Duration []time.Duration `flag:"duration-slice" ask:"any,duration slice"`
		URL      []url.URL       `flag:"url-slice,the help text" ask:"any,url slice"`
		Regexp   []regexp.Regexp `flag:"regexp-slice,the help text" ask:"any,regexp slice"`
	}

	// Flags is a helper type for testing purposes.
	Flags struct {
		Unsupported    chan int
		WithoutFlagTag string
		Value
		Pointer
		Slice
	}
)

func IsStructPtr(s interface{}) (reflect.Value, error) {
	v := reflect.ValueOf(s) // reflect.Value --> v.Type(), v.Kind(), v.NumField()
	t := reflect.TypeOf(s)  // reflect.Type --> t.Kind(), t.Name(), t.NumField()

	// A pointer to a struct should be passed
	if t.Kind() != reflect.Ptr {
		return reflect.Value{}, errors.New("non-pointer type: you should pass a pointer to a struct type")
	}

	// Navigate to the pointer value
	v = v.Elem()
	t = t.Elem()

	if t.Kind() != reflect.Struct {
		return reflect.Value{}, errors.New("non-struct type: you should pass a pointer to a struct type")
	}

	return v, nil
}

func IsNestedStruct(t reflect.Type) bool {
	if t.Kind() != reflect.Struct {
		return false
	}

	if IsStructSupported(t) {
		return false
	}

	return true
}

func IsStructSupported(t reflect.Type) bool {
	return (t.PkgPath() == "net/url" && t.Name() == "URL") ||
		(t.PkgPath() == "regexp" && t.Name() == "Regexp")
}

func IsTypeSupported(t reflect.Type) bool {
	switch t.Kind() {
	case reflect.String:
		return true
	case reflect.Bool:
		return true
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return true
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return true
	case reflect.Float32, reflect.Float64:
		return true
	case reflect.Ptr, reflect.Slice:
		return IsTypeSupported(t.Elem())
	case reflect.Struct:
		return IsStructSupported(t)
	default:
		return false
	}
}

package flagit

import (
	"errors"
	"flag"
	"net/url"
	"os"
	"reflect"
	"regexp"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type (
	Value struct {
		String   string        `flag:"string"`
		Bool     bool          `flag:"bool"`
		Int      int           `flag:"int"`
		Int8     int8          `flag:"int8"`
		Int16    int16         `flag:"int16"`
		Int32    int32         `flag:"int32"`
		Int64    int64         `flag:"int64"`
		Uint     uint          `flag:"uint"`
		Uint8    uint8         `flag:"uint8"`
		Uint16   uint16        `flag:"uint16"`
		Uint32   uint32        `flag:"uint32"`
		Uint64   uint64        `flag:"uint64"`
		Float32  float32       `flag:"float32"`
		Float64  float64       `flag:"float64"`
		Byte     byte          `flag:"byte"`
		Rune     rune          `flag:"rune"`
		Duration time.Duration `flag:"duration"`
		URL      url.URL       `flag:"url,the help text"`
		Regexp   regexp.Regexp `flag:"regexp,the help text"`
	}

	Pointer struct {
		String   *string        `flag:"string-pointer"`
		Bool     *bool          `flag:"bool-pointer"`
		Int      *int           `flag:"int-pointer"`
		Int8     *int8          `flag:"int8-pointer"`
		Int16    *int16         `flag:"int16-pointer"`
		Int32    *int32         `flag:"int32-pointer"`
		Int64    *int64         `flag:"int64-pointer"`
		Uint     *uint          `flag:"uint-pointer"`
		Uint8    *uint8         `flag:"uint8-pointer"`
		Uint16   *uint16        `flag:"uint16-pointer"`
		Uint32   *uint32        `flag:"uint32-pointer"`
		Uint64   *uint64        `flag:"uint64-pointer"`
		Float32  *float32       `flag:"float32-pointer"`
		Float64  *float64       `flag:"float64-pointer"`
		Byte     *byte          `flag:"byte-pointer"`
		Rune     *rune          `flag:"rune-pointer"`
		Duration *time.Duration `flag:"duration-pointer"`
		URL      *url.URL       `flag:"url-pointer,the help text"`
		Regexp   *regexp.Regexp `flag:"regexp-pointer,the help text"`
	}

	Slice struct {
		String   []string        `flag:"string-slice"`
		Bool     []bool          `flag:"bool-slice"`
		Int      []int           `flag:"int-slice"`
		Int8     []int8          `flag:"int8-slice"`
		Int16    []int16         `flag:"int16-slice"`
		Int32    []int32         `flag:"int32-slice"`
		Int64    []int64         `flag:"int64-slice"`
		Uint     []uint          `flag:"uint-slice"`
		Uint8    []uint8         `flag:"uint8-slice"`
		Uint16   []uint16        `flag:"uint16-slice"`
		Uint32   []uint32        `flag:"uint32-slice"`
		Uint64   []uint64        `flag:"uint64-slice"`
		Float32  []float32       `flag:"float32-slice"`
		Float64  []float64       `flag:"float64-slice"`
		Byte     []byte          `flag:"byte-slice"`
		Rune     []rune          `flag:"rune-slice"`
		Duration []time.Duration `flag:"duration-slice"`
		URL      []url.URL       `flag:"url-slice,the help text"`
		Regexp   []regexp.Regexp `flag:"regexp-slice,the help text"`
	}

	Flags struct {
		Unsupported    chan int
		WithoutFlagTag string
		Value
		Pointer
		Slice
	}
)

func TestFlagValue(t *testing.T) {
	d := time.Second

	tests := []struct {
		name             string
		v                flagValue
		setVal           string
		expectedSetError string
	}{
		{
			name: "OK",
			v: flagValue{
				continueOnError: false,
				value:           reflect.ValueOf(&d).Elem(),
				sep:             ",",
			},
			setVal:           "1m",
			expectedSetError: "",
		},
		{
			name: "Error",
			v: flagValue{
				continueOnError: false,
				value:           reflect.ValueOf(&d).Elem(),
				sep:             ",",
			},
			setVal:           "invalid",
			expectedSetError: `time: invalid duration "invalid"`,
		},
		{
			name: "ContinueOnError",
			v: flagValue{
				continueOnError: true,
				value:           reflect.ValueOf(&d).Elem(),
				sep:             ",",
			},
			setVal:           "invalid",
			expectedSetError: "",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Empty(t, tc.v.String())

			err := tc.v.Set(tc.setVal)
			if tc.expectedSetError == "" {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tc.expectedSetError)
			}
		})
	}
}

func TestValidateStruct(t *testing.T) {
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
			v, err := validateStruct(tc.s)

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
			supported := isTypeSupported(reflect.TypeOf(tc.field))

			assert.Equal(t, tc.expectedSupported, supported)
		})
	}
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

			assert.Equal(t, tc.expected, isStructSupported(tStruct))
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
	assert.False(t, isNestedStruct(vInt.Type()))

	vURL := vStruct.FieldByName("URL")
	assert.False(t, isNestedStruct(vURL.Type()))

	vRegexp := vStruct.FieldByName("Regexp")
	assert.False(t, isNestedStruct(vRegexp.Type()))

	vNested := vStruct.FieldByName("Nested")
	assert.True(t, isNestedStruct(vNested.Type()))
}

func TestGetFlagValue(t *testing.T) {
	tests := []struct {
		args              []string
		flag              string
		expectedFlagValue string
	}{
		{[]string{"app=invalid"}, "invalid", ""},

		{[]string{"app", "-enabled"}, "enabled", "true"},
		{[]string{"app", "--enabled"}, "enabled", "true"},
		{[]string{"app", "-enabled=false"}, "enabled", "false"},
		{[]string{"app", "--enabled=false"}, "enabled", "false"},
		{[]string{"app", "-enabled", "false"}, "enabled", "false"},
		{[]string{"app", "--enabled", "false"}, "enabled", "false"},

		{[]string{"app", "-number=-10"}, "number", "-10"},
		{[]string{"app", "--number=-10"}, "number", "-10"},
		{[]string{"app", "-number", "-10"}, "number", "-10"},
		{[]string{"app", "--number", "-10"}, "number", "-10"},

		{[]string{"app", "-text=content"}, "text", "content"},
		{[]string{"app", "--text=content"}, "text", "content"},
		{[]string{"app", "-text", "content"}, "text", "content"},
		{[]string{"app", "--text", "content"}, "text", "content"},

		{[]string{"app", "-enabled", "-text=content"}, "enabled", "true"},
		{[]string{"app", "--enabled", "--text=content"}, "enabled", "true"},
		{[]string{"app", "-enabled", "-text", "content"}, "enabled", "true"},
		{[]string{"app", "--enabled", "--text", "content"}, "enabled", "true"},

		{[]string{"app", "-name-list=alice,bob"}, "name-list", "alice,bob"},
		{[]string{"app", "--name-list=alice,bob"}, "name-list", "alice,bob"},
		{[]string{"app", "-name-list", "alice,bob"}, "name-list", "alice,bob"},
		{[]string{"app", "--name-list", "alice,bob"}, "name-list", "alice,bob"},
	}

	origArgs := os.Args
	defer func() {
		os.Args = origArgs
	}()

	for _, tc := range tests {
		os.Args = tc.args
		flagValue := getFlagValue(tc.flag)

		assert.Equal(t, tc.expectedFlagValue, flagValue)
	}
}

func TestIterateOnFields(t *testing.T) {
	invalid := struct {
		LogLevel string `flag:"log level"`
	}{}

	tests := []struct {
		name               string
		s                  interface{}
		continueOnError    bool
		expectedError      error
		expectedFieldNames []string
		expectedFlagNames  []string
		expectedListSeps   []string
	}{
		{
			name:               "StopOnError",
			s:                  &invalid,
			continueOnError:    false,
			expectedError:      errors.New("invalid flag name: log level"),
			expectedFieldNames: []string{},
			expectedFlagNames:  []string{},
			expectedListSeps:   []string{},
		},
		{
			name:               "ContinueOnError",
			s:                  &invalid,
			continueOnError:    true,
			expectedError:      nil,
			expectedFieldNames: []string{},
			expectedFlagNames:  []string{},
			expectedListSeps:   []string{},
		},
		{
			name:            "OK",
			s:               &Flags{},
			continueOnError: false,
			expectedError:   nil,
			expectedFieldNames: []string{
				"String",
				"Bool",
				"Int", "Int8", "Int16", "Int32", "Int64",
				"Uint", "Uint8", "Uint16", "Uint32", "Uint64",
				"Float32", "Float64",
				"Byte", "Rune", "Duration",
				"URL", "Regexp",
				"String",
				"Bool",
				"Int", "Int8", "Int16", "Int32", "Int64",
				"Uint", "Uint8", "Uint16", "Uint32", "Uint64",
				"Float32", "Float64",
				"Byte", "Rune", "Duration",
				"URL", "Regexp",
				"String",
				"Bool",
				"Int", "Int8", "Int16", "Int32", "Int64",
				"Uint", "Uint8", "Uint16", "Uint32", "Uint64",
				"Float32", "Float64",
				"Byte", "Rune", "Duration",
				"URL", "Regexp",
			},
			expectedFlagNames: []string{
				"string",
				"bool",
				"int", "int8", "int16", "int32", "int64",
				"uint", "uint8", "uint16", "uint32", "uint64",
				"float32", "float64",
				"byte", "rune", "duration",
				"url", "regexp",
				"string-pointer",
				"bool-pointer",
				"int-pointer", "int8-pointer", "int16-pointer", "int32-pointer", "int64-pointer",
				"uint-pointer", "uint8-pointer", "uint16-pointer", "uint32-pointer", "uint64-pointer",
				"float32-pointer", "float64-pointer",
				"byte-pointer", "rune-pointer", "duration-pointer",
				"url-pointer", "regexp-pointer",
				"string-slice",
				"bool-slice",
				"int-slice", "int8-slice", "int16-slice", "int32-slice", "int64-slice",
				"uint-slice", "uint8-slice", "uint16-slice", "uint32-slice", "uint64-slice",
				"float32-slice", "float64-slice",
				"byte-slice", "rune-slice", "duration-slice",
				"url-slice", "regexp-slice",
			},
			expectedListSeps: []string{
				",",
				",",
				",", ",", ",", ",", ",",
				",", ",", ",", ",", ",",
				",", ",",
				",", ",", ",",
				",", ",",
				",",
				",",
				",", ",", ",", ",", ",",
				",", ",", ",", ",", ",",
				",", ",",
				",", ",", ",",
				",", ",",
				",",
				",",
				",", ",", ",", ",", ",",
				",", ",", ",", ",", ",",
				",", ",",
				",", ",", ",",
				",", ",",
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			fieldNames := []string{}
			flagNames := []string{}
			listSeps := []string{}

			vStruct, err := validateStruct(tc.s)
			assert.NoError(t, err)

			err = iterateOnFields("", vStruct, tc.continueOnError, func(f fieldInfo) error {
				fieldNames = append(fieldNames, f.name)
				flagNames = append(flagNames, f.flag)
				listSeps = append(listSeps, f.sep)
				return nil
			})

			assert.Equal(t, tc.expectedError, err)
			assert.Equal(t, tc.expectedFieldNames, fieldNames)
			assert.Equal(t, tc.expectedFlagNames, flagNames)
			assert.Equal(t, tc.expectedListSeps, listSeps)
		})
	}
}

func TestParse(t *testing.T) {
	url1, _ := url.Parse("service-1")
	url2, _ := url.Parse("service-2")

	re1 := regexp.MustCompilePOSIX("[:digit:]")
	re2 := regexp.MustCompilePOSIX("[:alpha:]")

	flags := &Flags{
		Value: Value{
			String:   "foo",
			Bool:     false,
			Float32:  3.1415,
			Float64:  3.14159265359,
			Int:      -9223372036854775808,
			Int8:     -128,
			Int16:    -32768,
			Int32:    -2147483648,
			Int64:    -9223372036854775808,
			Uint:     0,
			Uint8:    0,
			Uint16:   0,
			Uint32:   0,
			Uint64:   0,
			Byte:     0,
			Rune:     -2147483648,
			Duration: time.Second,
			URL:      *url1,
			Regexp:   *re1,
		},
		Pointer: Pointer{
			String:   stringPtr("foo"),
			Bool:     boolPtr(false),
			Float32:  float32Ptr(3.1415),
			Float64:  float64Ptr(3.14159265359),
			Int:      intPtr(-9223372036854775808),
			Int8:     int8Ptr(-128),
			Int16:    int16Ptr(-32768),
			Int32:    int32Ptr(-2147483648),
			Int64:    int64Ptr(-9223372036854775808),
			Uint:     uintPtr(0),
			Uint8:    uint8Ptr(0),
			Uint16:   uint16Ptr(0),
			Uint32:   uint32Ptr(0),
			Uint64:   uint64Ptr(0),
			Byte:     bytePtr(0),
			Rune:     runePtr(-2147483648),
			Duration: durationPtr(time.Second),
			URL:      url1,
			Regexp:   re1,
		},
		Slice: Slice{
			String:   []string{"foo", "bar"},
			Bool:     []bool{false, true},
			Float32:  []float32{3.1415, 2.7182},
			Float64:  []float64{3.14159265359, 2.71828182845},
			Int:      []int{-9223372036854775808, 9223372036854775807},
			Int8:     []int8{-128, 127},
			Int16:    []int16{-32768, 32767},
			Int32:    []int32{-2147483648, 2147483647},
			Int64:    []int64{-9223372036854775808, 9223372036854775807},
			Uint:     []uint{0, 18446744073709551615},
			Uint8:    []uint8{0, 255},
			Uint16:   []uint16{0, 65535},
			Uint32:   []uint32{0, 4294967295},
			Uint64:   []uint64{0, 18446744073709551615},
			Byte:     []byte{0, 255},
			Rune:     []rune{-2147483648, 2147483647},
			Duration: []time.Duration{time.Second, time.Minute},
			URL:      []url.URL{*url1, *url2},
			Regexp:   []regexp.Regexp{*re1, *re2},
		},
	}

	tests := []struct {
		name            string
		args            []string
		s               interface{}
		continueOnError bool
		expectedError   string
		expected        *Flags
	}{
		{
			"NonStruct",
			[]string{"app"},
			new(string),
			false,
			"non-struct type: you should pass a pointer to a struct type",
			&Flags{},
		},
		{
			"NonPointer",
			[]string{"app"},
			Flags{},
			false,
			"non-pointer type: you should pass a pointer to a struct type",
			&Flags{},
		},
		{
			"FromDefaults",
			[]string{"app"},
			flags,
			false,
			"",
			flags,
		},
		{
			"FromFlags",
			[]string{
				"app",
				"-string=foo",
				"-bool=false",
				"-float32=3.1415",
				"-float64=3.14159265359",
				"-int=-9223372036854775808",
				"-int8=-128",
				"-int16=-32768",
				"-int32=-2147483648",
				"-int64=-9223372036854775808",
				"-uint=0",
				"-uint8=0",
				"-uint16=0",
				"-uint32=0",
				"-uint64=0",
				"-byte=0",
				"-rune=-2147483648",
				"-duration=1s",
				"-url=service-1",
				"-regexp=[:digit:]",
				"-string-pointer=foo",
				"-bool-pointer=false",
				"-float32-pointer=3.1415",
				"-float64-pointer=3.14159265359",
				"-int-pointer=-9223372036854775808",
				"-int8-pointer=-128",
				"-int16-pointer=-32768",
				"-int32-pointer=-2147483648",
				"-int64-pointer=-9223372036854775808",
				"-uint-pointer=0",
				"-uint8-pointer=0",
				"-uint16-pointer=0",
				"-uint32-pointer=0",
				"-uint64-pointer=0",
				"-byte-pointer=0",
				"-rune-pointer=-2147483648",
				"-duration-pointer=1s",
				"-url-pointer=service-1",
				"-regexp-pointer=[:digit:]",
				"-string-slice=foo,bar",
				"-bool-slice=false,true",
				"-float32-slice=3.1415,2.7182",
				"-float64-slice=3.14159265359,2.71828182845",
				"-int-slice=-9223372036854775808,9223372036854775807",
				"-int8-slice=-128,127",
				"-int16-slice=-32768,32767",
				"-int32-slice=-2147483648,2147483647",
				"-int64-slice=-9223372036854775808,9223372036854775807",
				"-uint-slice=0,18446744073709551615",
				"-uint8-slice=0,255",
				"-uint16-slice=0,65535",
				"-uint32-slice=0,4294967295",
				"-uint64-slice=0,18446744073709551615",
				"-byte-slice=0,255",
				"-rune-slice=-2147483648,2147483647",
				"-duration-slice=1s,1m",
				"-url-slice=service-1,service-2",
				"-regexp-slice=[:digit:],[:alpha:]",
			},
			&Flags{},
			false,
			"",
			flags,
		},
		{
			"StopOnError",
			[]string{
				"app",
				"-int=invalid",
			},
			&Flags{},
			false,
			`strconv.ParseInt: parsing "invalid": invalid syntax`,
			&Flags{},
		},
		{
			"ContinueOnError",
			[]string{
				"app",
				"-bool=invalid",
				"-float32=invalid",
				"-float64=invalid",
				"-int=invalid",
				"-int8=invalid",
				"-int16=invalid",
				"-int32=invalid",
				"-int64=invalid",
				"-uint=invalid",
				"-uint8=invalid",
				"-uint16=invalid",
				"-uint32=invalid",
				"-uint64=invalid",
				"-url=:",
				"-regexp=[:invalid:",
				"-duration=invalid",
				"-bool-pointer=invalid",
				"-float32-pointer=invalid",
				"-float64-pointer=invalid",
				"-int-pointer=invalid",
				"-int8-pointer=invalid",
				"-int16-pointer=invalid",
				"-int32-pointer=invalid",
				"-int64-pointer=invalid",
				"-uint-pointer=invalid",
				"-uint8-pointer=invalid",
				"-uint16-pointer=invalid",
				"-uint32-pointer=invalid",
				"-uint64-pointer=invalid",
				"-url-pointer=:",
				"-regexp-pointer=[:invalid:",
				"-duration-pointer=invalid",
				"-bool-slice=invalid",
				"-float32-slice=invalid",
				"-float64-slice=invalid",
				"-int-slice=invalid",
				"-int8-slice=invalid",
				"-int16-slice=invalid",
				"-int32-slice=invalid",
				"-int64-slice=invalid",
				"-uint-slice=invalid",
				"-uint8-slice=invalid",
				"-uint16-slice=invalid",
				"-uint32-slice=invalid",
				"-uint64-slice=invalid",
				"-url-slice=:",
				"-regexp-slice=[:invalid:",
				"-duration-slice=invalid",
			},
			&Flags{},
			true,
			"",
			&Flags{},
		},
	}

	origArgs := os.Args
	defer func() {
		os.Args = origArgs
	}()

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			os.Args = tc.args

			err := Parse(tc.s, tc.continueOnError)

			if tc.expectedError == "" {
				assert.NoError(t, err)
				assert.Equal(t, tc.expected, tc.s)
			} else {
				assert.Error(t, err)
				assert.EqualError(t, err, tc.expectedError)
			}
		})
	}
}

func TestRegister(t *testing.T) {
	fs := flag.NewFlagSet("app", flag.ContinueOnError)
	fs.String("string", "", "")

	url1, _ := url.Parse("service-1")
	url2, _ := url.Parse("service-2")

	re1 := regexp.MustCompilePOSIX("[:digit:]")
	re2 := regexp.MustCompilePOSIX("[:alpha:]")

	flags := &Flags{
		Value: Value{
			String:   "foo",
			Bool:     false,
			Float32:  3.1415,
			Float64:  3.14159265359,
			Int:      -9223372036854775808,
			Int8:     -128,
			Int16:    -32768,
			Int32:    -2147483648,
			Int64:    -9223372036854775808,
			Uint:     0,
			Uint8:    0,
			Uint16:   0,
			Uint32:   0,
			Uint64:   0,
			Byte:     0,
			Rune:     -2147483648,
			Duration: time.Second,
			URL:      *url1,
			Regexp:   *re1,
		},
		Pointer: Pointer{
			String:   stringPtr("foo"),
			Bool:     boolPtr(false),
			Float32:  float32Ptr(3.1415),
			Float64:  float64Ptr(3.14159265359),
			Int:      intPtr(-9223372036854775808),
			Int8:     int8Ptr(-128),
			Int16:    int16Ptr(-32768),
			Int32:    int32Ptr(-2147483648),
			Int64:    int64Ptr(-9223372036854775808),
			Uint:     uintPtr(0),
			Uint8:    uint8Ptr(0),
			Uint16:   uint16Ptr(0),
			Uint32:   uint32Ptr(0),
			Uint64:   uint64Ptr(0),
			Byte:     bytePtr(0),
			Rune:     runePtr(-2147483648),
			Duration: durationPtr(time.Second),
			URL:      url1,
			Regexp:   re1,
		},
		Slice: Slice{
			String:   []string{"foo", "bar"},
			Bool:     []bool{false, true},
			Float32:  []float32{3.1415, 2.7182},
			Float64:  []float64{3.14159265359, 2.71828182845},
			Int:      []int{-9223372036854775808, 9223372036854775807},
			Int8:     []int8{-128, 127},
			Int16:    []int16{-32768, 32767},
			Int32:    []int32{-2147483648, 2147483647},
			Int64:    []int64{-9223372036854775808, 9223372036854775807},
			Uint:     []uint{0, 18446744073709551615},
			Uint8:    []uint8{0, 255},
			Uint16:   []uint16{0, 65535},
			Uint32:   []uint32{0, 4294967295},
			Uint64:   []uint64{0, 18446744073709551615},
			Byte:     []byte{0, 255},
			Rune:     []rune{-2147483648, 2147483647},
			Duration: []time.Duration{time.Second, time.Minute},
			URL:      []url.URL{*url1, *url2},
			Regexp:   []regexp.Regexp{*re1, *re2},
		},
	}

	tests := []struct {
		name               string
		args               []string
		fs                 *flag.FlagSet
		s                  interface{}
		continueOnError    bool
		expectedError      error
		expectedParseError string
		expected           *Flags
	}{
		{
			"NonStruct",
			[]string{"app"},
			new(flag.FlagSet),
			new(string),
			false,
			errors.New("non-struct type: you should pass a pointer to a struct type"), "",
			&Flags{},
		},
		{
			"NonPointer",
			[]string{"app"},
			new(flag.FlagSet),
			Flags{},
			false,
			errors.New("non-pointer type: you should pass a pointer to a struct type"), "",
			&Flags{},
		},
		{
			"FlagRegistered_StopOnError",
			[]string{"app"},
			fs,
			&Flags{},
			false,
			errors.New("flag already registered: string"), "",
			&Flags{},
		},
		{
			"FlagRegistered_ContinueOnError",
			[]string{"app"},
			fs,
			&Flags{},
			true,
			nil, "",
			&Flags{},
		},
		{
			"FromDefaults",
			[]string{"app"},
			new(flag.FlagSet),
			flags,
			false,
			nil, "",
			flags,
		},
		{
			"FromFlags",
			[]string{
				"app",
				"-string=foo",
				"-bool=false",
				"-float32=3.1415",
				"-float64=3.14159265359",
				"-int=-9223372036854775808",
				"-int8=-128",
				"-int16=-32768",
				"-int32=-2147483648",
				"-int64=-9223372036854775808",
				"-uint=0",
				"-uint8=0",
				"-uint16=0",
				"-uint32=0",
				"-uint64=0",
				"-byte=0",
				"-rune=-2147483648",
				"-duration=1s",
				"-url=service-1",
				"-regexp=[:digit:]",
				"-string-pointer=foo",
				"-bool-pointer=false",
				"-float32-pointer=3.1415",
				"-float64-pointer=3.14159265359",
				"-int-pointer=-9223372036854775808",
				"-int8-pointer=-128",
				"-int16-pointer=-32768",
				"-int32-pointer=-2147483648",
				"-int64-pointer=-9223372036854775808",
				"-uint-pointer=0",
				"-uint8-pointer=0",
				"-uint16-pointer=0",
				"-uint32-pointer=0",
				"-uint64-pointer=0",
				"-byte-pointer=0",
				"-rune-pointer=-2147483648",
				"-duration-pointer=1s",
				"-url-pointer=service-1",
				"-regexp-pointer=[:digit:]",
				"-string-slice=foo,bar",
				"-bool-slice=false,true",
				"-float32-slice=3.1415,2.7182",
				"-float64-slice=3.14159265359,2.71828182845",
				"-int-slice=-9223372036854775808,9223372036854775807",
				"-int8-slice=-128,127",
				"-int16-slice=-32768,32767",
				"-int32-slice=-2147483648,2147483647",
				"-int64-slice=-9223372036854775808,9223372036854775807",
				"-uint-slice=0,18446744073709551615",
				"-uint8-slice=0,255",
				"-uint16-slice=0,65535",
				"-uint32-slice=0,4294967295",
				"-uint64-slice=0,18446744073709551615",
				"-byte-slice=0,255",
				"-rune-slice=-2147483648,2147483647",
				"-duration-slice=1s,1m",
				"-url-slice=service-1,service-2",
				"-regexp-slice=[:digit:],[:alpha:]",
			},
			new(flag.FlagSet),
			&Flags{},
			false,
			nil, "",
			flags,
		},
		{
			"StopOnError",
			[]string{
				"app",
				"-int=invalid",
			},
			new(flag.FlagSet),
			&Flags{},
			false,
			nil, `invalid value "invalid" for flag -int: strconv.ParseInt: parsing "invalid": invalid syntax`,
			&Flags{},
		},
		{
			"ContinueOnError",
			[]string{
				"app",
				"-bool=invalid",
				"-float32=invalid",
				"-float64=invalid",
				"-int=invalid",
				"-int8=invalid",
				"-int16=invalid",
				"-int32=invalid",
				"-int64=invalid",
				"-uint=invalid",
				"-uint8=invalid",
				"-uint16=invalid",
				"-uint32=invalid",
				"-uint64=invalid",
				"-url=:",
				"-regexp=[:invalid:",
				"-duration=invalid",
				"-bool-pointer=invalid",
				"-float32-pointer=invalid",
				"-float64-pointer=invalid",
				"-int-pointer=invalid",
				"-int8-pointer=invalid",
				"-int16-pointer=invalid",
				"-int32-pointer=invalid",
				"-int64-pointer=invalid",
				"-uint-pointer=invalid",
				"-uint8-pointer=invalid",
				"-uint16-pointer=invalid",
				"-uint32-pointer=invalid",
				"-uint64-pointer=invalid",
				"-url-pointer=:",
				"-regexp-pointer=[:invalid:",
				"-duration-pointer=invalid",
				"-bool-slice=invalid",
				"-float32-slice=invalid",
				"-float64-slice=invalid",
				"-int-slice=invalid",
				"-int8-slice=invalid",
				"-int16-slice=invalid",
				"-int32-slice=invalid",
				"-int64-slice=invalid",
				"-uint-slice=invalid",
				"-uint8-slice=invalid",
				"-uint16-slice=invalid",
				"-uint32-slice=invalid",
				"-uint64-slice=invalid",
				"-url-slice=:",
				"-regexp-slice=[:invalid:",
				"-duration-slice=invalid",
			},
			new(flag.FlagSet),
			&Flags{},
			true,
			nil, `invalid boolean value "invalid" for -bool: parse error`,
			&Flags{},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := Register(tc.fs, tc.s, tc.continueOnError)
			assert.Equal(t, tc.expectedError, err)

			if tc.expectedError == nil {
				err := tc.fs.Parse(tc.args[1:])

				if tc.expectedParseError == "" {
					assert.NoError(t, err)
					assert.Equal(t, tc.expected, tc.s)
				} else {
					assert.Error(t, err)
					assert.EqualError(t, err, tc.expectedParseError)
				}
			}
		})
	}
}

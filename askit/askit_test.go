package askit

import (
	"errors"
	"net/url"
	"reflect"
	"regexp"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/gardenbed/charm/internal/ptr"
	"github.com/gardenbed/charm/internal/rflct"
)

func TestAsk(t *testing.T) {
	url1, _ := url.Parse("service-1.example.com")
	url2, _ := url.Parse("service-2.example.com")

	re1 := regexp.MustCompilePOSIX("[:digit:]")
	re2 := regexp.MustCompilePOSIX("[:alpha:]")

	flags := &rflct.Flags{
		Value: rflct.Value{
			String:   "jane.doe@example.com",
			Bool:     false,
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
			Float32:  2.7182,
			Float64:  2.71828182845,
			Byte:     0,
			Rune:     -2147483648,
			Duration: time.Second,
			URL:      *url1,
			Regexp:   *re1,
		},
		Pointer: rflct.Pointer{
			String:   ptr.String("access_token"),
			Bool:     ptr.Bool(true),
			Int:      ptr.Int(9223372036854775807),
			Int8:     ptr.Int8(127),
			Int16:    ptr.Int16(32767),
			Int32:    ptr.Int32(2147483647),
			Int64:    ptr.Int64(9223372036854775807),
			Uint:     ptr.Uint(18446744073709551615),
			Uint8:    ptr.Uint8(255),
			Uint16:   ptr.Uint16(65535),
			Uint32:   ptr.Uint32(4294967295),
			Uint64:   ptr.Uint64(18446744073709551615),
			Float32:  ptr.Float32(3.1415),
			Float64:  ptr.Float64(3.14159265359),
			Byte:     ptr.Byte(255),
			Rune:     ptr.Rune(2147483647),
			Duration: ptr.Duration(time.Minute),
			URL:      url2,
			Regexp:   re2,
		},
		Slice: rflct.Slice{
			String:   []string{"jane.doe@example.com", "john.doe@example.com"},
			Bool:     []bool{false, true},
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
			Float32:  []float32{2.7182, 3.1415},
			Float64:  []float64{2.71828182845, 3.14159265359},
			Byte:     []byte{0, 255},
			Rune:     []rune{-2147483648, 2147483647},
			Duration: []time.Duration{time.Second, time.Minute},
			URL:      []url.URL{*url1, *url2},
			Regexp:   []regexp.Regexp{*re1, *re2},
		},
	}

	tests := []struct {
		name          string
		s             interface{}
		asker         Asker
		expectedError string
		expected      *rflct.Flags
	}{
		{
			name:          "NonStruct",
			s:             new(string),
			asker:         nil,
			expectedError: "non-struct type: you should pass a pointer to a struct type",
		},
		{
			name:          "NonPointer",
			s:             rflct.Flags{},
			asker:         nil,
			expectedError: "non-pointer type: you should pass a pointer to a struct type",
		},
		{
			name: "Success_WithOldValues",
			s:    flags,
			asker: &MockAsker{
				AskMocks: []AskMock{
					{OutString: "N"}, // Value.String
					{OutString: "N"}, // Value.Bool
					{OutString: "N"}, // Value.Int
					{OutString: "N"}, // Value.Int8
					{OutString: "N"}, // Value.Int16
					{OutString: "N"}, // Value.Int32
					{OutString: "N"}, // Value.Int64
					{OutString: "N"}, // Value.Uint
					{OutString: "N"}, // Value.Uint8
					{OutString: "N"}, // Value.Uint16
					{OutString: "N"}, // Value.Uint32
					{OutString: "N"}, // Value.Uint64
					{OutString: "N"}, // Value.Float32
					{OutString: "N"}, // Value.Float64
					{OutString: "N"}, // Value.Byte
					{OutString: "N"}, // Value.Rune
					{OutString: "N"}, // Value.Duration
					{OutString: "N"}, // Value.URL
					{OutString: "N"}, // Value.Regexp
					{OutString: "N"}, // Pointer.String
					{OutString: "N"}, // Pointer.Bool
					{OutString: "N"}, // Pointer.Int
					{OutString: "N"}, // Pointer.Int8
					{OutString: "N"}, // Pointer.Int16
					{OutString: "N"}, // Pointer.Int32
					{OutString: "N"}, // Pointer.Int64
					{OutString: "N"}, // Pointer.Uint
					{OutString: "N"}, // Pointer.Uint8
					{OutString: "N"}, // Pointer.Uint16
					{OutString: "N"}, // Pointer.Uint32
					{OutString: "N"}, // Pointer.Uint64
					{OutString: "N"}, // Pointer.Float32
					{OutString: "N"}, // Pointer.Float64
					{OutString: "N"}, // Pointer.Byte
					{OutString: "N"}, // Pointer.Rune
					{OutString: "N"}, // Pointer.Duration
					{OutString: "N"}, // Pointer.URL
					{OutString: "N"}, // Pointer.Regexp
					{OutString: "N"}, // Slice.String
					{OutString: "N"}, // Slice.Bool
					{OutString: "N"}, // Slice.Int
					{OutString: "N"}, // Slice.Int8
					{OutString: "N"}, // Slice.Int16
					{OutString: "N"}, // Slice.Int32
					{OutString: "N"}, // Slice.Int64
					{OutString: "N"}, // Slice.Uint
					{OutString: "N"}, // Slice.Uint8
					{OutString: "N"}, // Slice.Uint16
					{OutString: "N"}, // Slice.Uint32
					{OutString: "N"}, // Slice.Uint64
					{OutString: "N"}, // Slice.Float32
					{OutString: "N"}, // Slice.Float64
					{OutString: "N"}, // Slice.Byte
					{OutString: "N"}, // Slice.Rune
					{OutString: "N"}, // Slice.Duration
					{OutString: "N"}, // Slice.URL
					{OutString: "N"}, // Slice.Regexp
				},
				AskSecretMocks: []AskSecretMock{
					{OutString: "access_token"}, // Pointer.String
				},
			},
			expected: flags,
		},
		{
			name: "Success_WithNewValues",
			s:    &rflct.Flags{},
			asker: &MockAsker{
				AskMocks: []AskMock{
					{OutString: "Y"}, {OutString: "jane.doe@example.com"}, // Value.String
					{OutString: "Y"}, {OutString: "false"}, // Value.Bool
					{OutString: "Y"}, {OutString: "-9223372036854775808"}, // Value.Int
					{OutString: "Y"}, {OutString: "-128"}, // Value.Int8
					{OutString: "Y"}, {OutString: "-32768"}, // Value.Int16
					{OutString: "Y"}, {OutString: "-2147483648"}, // Value.Int32
					{OutString: "Y"}, {OutString: "-9223372036854775808"}, // Value.Int64
					{OutString: "Y"}, {OutString: "0"}, // Value.Uint
					{OutString: "Y"}, {OutString: "0"}, // Value.Uint8
					{OutString: "Y"}, {OutString: "0"}, // Value.Uint16
					{OutString: "Y"}, {OutString: "0"}, // Value.Uint32
					{OutString: "Y"}, {OutString: "0"}, // Value.Uint64
					{OutString: "Y"}, {OutString: "2.7182"}, // Value.Float32
					{OutString: "Y"}, {OutString: "2.71828182845"}, // Value.Float64
					{OutString: "Y"}, {OutString: "0"}, // Value.Byte
					{OutString: "Y"}, {OutString: "-2147483648"}, // Value.Rune
					{OutString: "Y"}, {OutString: "1s"}, // Value.Duration
					{OutString: "Y"}, {OutString: "service-1.example.com"}, // Value.URL
					{OutString: "Y"}, {OutString: "[:digit:]"}, // Value.Regexp
					{OutString: "Y"},                      // Pointer.String
					{OutString: "Y"}, {OutString: "true"}, // Pointer.Bool
					{OutString: "Y"}, {OutString: "9223372036854775807"}, // Pointer.Int
					{OutString: "Y"}, {OutString: "127"}, // Pointer.Int8
					{OutString: "Y"}, {OutString: "32767"}, // Pointer.Int16
					{OutString: "Y"}, {OutString: "2147483647"}, // Pointer.Int32
					{OutString: "Y"}, {OutString: "9223372036854775807"}, // Pointer.Int64
					{OutString: "Y"}, {OutString: "18446744073709551615"}, // Pointer.Uint
					{OutString: "Y"}, {OutString: "255"}, // Pointer.Uint8
					{OutString: "Y"}, {OutString: "65535"}, // Pointer.Uint16
					{OutString: "Y"}, {OutString: "4294967295"}, // Pointer.Uint32
					{OutString: "Y"}, {OutString: "18446744073709551615"}, // Pointer.Uint64
					{OutString: "Y"}, {OutString: "3.1415"}, // Pointer.Float32
					{OutString: "Y"}, {OutString: "3.14159265359"}, // Pointer.Float64
					{OutString: "Y"}, {OutString: "255"}, // Pointer.Byte
					{OutString: "Y"}, {OutString: "2147483647"}, // Pointer.Rune
					{OutString: "Y"}, {OutString: "1m"}, // Pointer.Duration
					{OutString: "Y"}, {OutString: "service-2.example.com"}, // Pointer.URL
					{OutString: "Y"}, {OutString: "[:alpha:]"}, // Pointer.Regexp
					{OutString: "Y"}, {OutString: "jane.doe@example.com,john.doe@example.com"}, // Slice.String
					{OutString: "Y"}, {OutString: "false,true"}, // Slice.Bool
					{OutString: "Y"}, {OutString: "-9223372036854775808,9223372036854775807"}, // Slice.Int
					{OutString: "Y"}, {OutString: "-128,127"}, // Slice.Int8
					{OutString: "Y"}, {OutString: "-32768,32767"}, // Slice.Int16
					{OutString: "Y"}, {OutString: "-2147483648,2147483647"}, // Slice.Int32
					{OutString: "Y"}, {OutString: "-9223372036854775808,9223372036854775807"}, // Slice.Int64
					{OutString: "Y"}, {OutString: "0,18446744073709551615"}, // Slice.Uint
					{OutString: "Y"}, {OutString: "0,255"}, // Slice.Uint8
					{OutString: "Y"}, {OutString: "0,65535"}, // Slice.Uint16
					{OutString: "Y"}, {OutString: "0,4294967295"}, // Slice.Uint32
					{OutString: "Y"}, {OutString: "0,18446744073709551615"}, // Slice.Uint64
					{OutString: "Y"}, {OutString: "2.7182,3.1415"}, // Slice.Float32
					{OutString: "Y"}, {OutString: "2.71828182845,3.14159265359"}, // Slice.Float64
					{OutString: "Y"}, {OutString: "0,255"}, // Slice.Byte
					{OutString: "Y"}, {OutString: "-2147483648,2147483647"}, // Slice.Rune
					{OutString: "Y"}, {OutString: "1s,1m"}, // Slice.Duration
					{OutString: "Y"}, {OutString: "service-1.example.com,service-2.example.com"}, // Slice.URL
					{OutString: "Y"}, {OutString: "[:digit:],[:alpha:]"}, // Slice.Regexp
				},
				AskSecretMocks: []AskSecretMock{
					{OutString: "access_token"}, // Pointer.String
				},
			},
			expected: flags,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := Ask(tc.s, tc.asker)

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

func TestIterateOnFields(t *testing.T) {
	t.Run("HandleFails", func(t *testing.T) {
		v := reflect.ValueOf(&rflct.Flags{}).Elem()
		err := iterateOnFields("", v, func(f fieldInfo) error {
			if f.Name == "Value.String" {
				return errors.New("handle error")
			}
			return nil
		})

		assert.EqualError(t, err, "handle error")
	})

	tests := []struct {
		name                 string
		s                    interface{}
		expectedError        error
		expectedNames        []string
		expectedKinds        []Kind
		expectedDescriptions []string
		expectedListSeps     []string
	}{
		{
			name:          "OK",
			s:             &rflct.Flags{},
			expectedError: nil,
			expectedNames: []string{
				"Value.String",
				"Value.Bool",
				"Value.Int", "Value.Int8", "Value.Int16", "Value.Int32", "Value.Int64",
				"Value.Uint", "Value.Uint8", "Value.Uint16", "Value.Uint32", "Value.Uint64",
				"Value.Float32", "Value.Float64",
				"Value.Byte", "Value.Rune", "Value.Duration",
				"Value.URL", "Value.Regexp",
				"Pointer.String",
				"Pointer.Bool",
				"Pointer.Int", "Pointer.Int8", "Pointer.Int16", "Pointer.Int32", "Pointer.Int64",
				"Pointer.Uint", "Pointer.Uint8", "Pointer.Uint16", "Pointer.Uint32", "Pointer.Uint64",
				"Pointer.Float32", "Pointer.Float64",
				"Pointer.Byte", "Pointer.Rune", "Pointer.Duration",
				"Pointer.URL", "Pointer.Regexp",
				"Slice.String",
				"Slice.Bool",
				"Slice.Int", "Slice.Int8", "Slice.Int16", "Slice.Int32", "Slice.Int64",
				"Slice.Uint", "Slice.Uint8", "Slice.Uint16", "Slice.Uint32", "Slice.Uint64",
				"Slice.Float32", "Slice.Float64",
				"Slice.Byte", "Slice.Rune", "Slice.Duration",
				"Slice.URL", "Slice.Regexp",
			},
			expectedKinds: []Kind{
				KindEmail,
				KindAny,
				KindAny, KindAny, KindAny, KindAny, KindAny,
				KindAny, KindAny, KindAny, KindAny, KindAny,
				KindAny, KindAny,
				KindAny, KindAny, KindAny,
				KindAny, KindAny,
				KindSecret,
				KindAny,
				KindAny, KindAny, KindAny, KindAny, KindAny,
				KindAny, KindAny, KindAny, KindAny, KindAny,
				KindAny, KindAny,
				KindAny, KindAny, KindAny,
				KindAny, KindAny,
				KindAny,
				KindAny,
				KindAny, KindAny, KindAny, KindAny, KindAny,
				KindAny, KindAny, KindAny, KindAny, KindAny,
				KindAny, KindAny,
				KindAny, KindAny, KindAny,
				KindAny, KindAny,
			},
			expectedDescriptions: []string{
				"",
				"bool value",
				"int value", "int8 value", "int16 value", "int32 value", "int64 value",
				"uint value", "uint8 value", "uint16 value", "uint32 value", "uint64 value",
				"float32 value", "float64 value",
				"byte value", "rune value", "duration value",
				"url value", "regexp value",
				"",
				"bool pointer",
				"int pointer", "int8 pointer", "int16 pointer", "int32 pointer", "int64 pointer",
				"uint pointer", "uint8 pointer", "uint16 pointer", "uint32 pointer", "uint64 pointer",
				"float32 pointer", "float64 pointer",
				"byte pointer", "rune pointer", "duration pointer",
				"url pointer", "regexp pointer",
				"string slice",
				"bool slice",
				"int slice", "int8 slice", "int16 slice", "int32 slice", "int64 slice",
				"uint slice", "uint8 slice", "uint16 slice", "uint32 slice", "uint64 slice",
				"float32 slice", "float64 slice",
				"byte slice", "rune slice", "duration slice",
				"url slice", "regexp slice",
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
			names := []string{}
			kinds := []Kind{}
			descriptions := []string{}
			listSeps := []string{}

			vStruct, err := rflct.IsStructPtr(tc.s)
			assert.NoError(t, err)

			err = iterateOnFields("", vStruct, func(f fieldInfo) error {
				names = append(names, f.Name)
				kinds = append(kinds, f.Kind)
				descriptions = append(descriptions, f.Description)
				listSeps = append(listSeps, f.Sep)
				return nil
			})

			assert.Equal(t, tc.expectedError, err)
			assert.Equal(t, tc.expectedNames, names)
			assert.Equal(t, tc.expectedKinds, kinds)
			assert.Equal(t, tc.expectedDescriptions, descriptions)
			assert.Equal(t, tc.expectedListSeps, listSeps)
		})
	}
}

func TestIsKindSupported(t *testing.T) {
	tests := []struct {
		name         string
		kind         string
		expectedBool bool
	}{
		{
			name:         "Any",
			kind:         "any",
			expectedBool: true,
		},
		{
			name:         "Email",
			kind:         "email",
			expectedBool: true,
		},
		{
			name:         "Secret",
			kind:         "secret",
			expectedBool: true,
		},
		{
			name:         "Unsupported",
			kind:         "unsupported",
			expectedBool: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedBool, isKindSupported(tc.kind))
		})
	}
}

func TestAskForField(t *testing.T) {
	id := 69
	token := "secret"
	email := "john.doe@example.com"

	tests := []struct {
		name          string
		f             fieldInfo
		asker         Asker
		expectedError string
	}{
		{
			name: "ConfirmAskFails",
			f: fieldInfo{
				Value:       reflect.ValueOf(&id).Elem(),
				Name:        "ID",
				Kind:        KindAny,
				Description: "Your id",
				Sep:         ",",
			},
			asker: &MockAsker{
				AskMocks: []AskMock{
					{OutError: errors.New("io error")},
				},
			},
			expectedError: "io error",
		},
		{
			name: "ConfirmAskNo",
			f: fieldInfo{
				Value:       reflect.ValueOf(&id).Elem(),
				Name:        "ID",
				Kind:        KindAny,
				Description: "Your id",
				Sep:         ",",
			},
			asker: &MockAsker{
				AskMocks: []AskMock{
					{OutString: "n"},
				},
			},
			expectedError: "",
		},
		{
			name: "InputAskFails",
			f: fieldInfo{
				Value:       reflect.ValueOf(&id).Elem(),
				Name:        "ID",
				Kind:        KindAny,
				Description: "Your id",
				Sep:         ",",
			},
			asker: &MockAsker{
				AskMocks: []AskMock{
					{OutString: "Y"},
					{OutError: errors.New("io error")},
				},
			},
			expectedError: "io error",
		},
		{
			name: "InputAskSecretFails",
			f: fieldInfo{
				Value:       reflect.ValueOf(&token).Elem(),
				Name:        "Token",
				Kind:        KindSecret,
				Description: "Your access token",
				Sep:         ",",
			},
			asker: &MockAsker{
				AskMocks: []AskMock{
					{OutString: "Y"},
				},
				AskSecretMocks: []AskSecretMock{
					{OutError: errors.New("io error")},
				},
			},
			expectedError: "io error",
		},
		{
			name: "SetValueFails",
			f: fieldInfo{
				Value:       reflect.ValueOf(&id).Elem(),
				Name:        "ID",
				Kind:        KindAny,
				Description: "Your id",
				Sep:         ",",
			},
			asker: &MockAsker{
				AskMocks: []AskMock{
					{OutString: "Y"},
					{OutString: "invalid_integer"},
				},
			},
			expectedError: `invalid value entered for ID: strconv.ParseInt: parsing "invalid_integer": invalid syntax`,
		},
		{
			name: "InvalidEmail",
			f: fieldInfo{
				Value:       reflect.ValueOf(&email).Elem(),
				Name:        "Email",
				Kind:        KindEmail,
				Description: "Your email address",
				Sep:         ",",
			},
			asker: &MockAsker{
				AskMocks: []AskMock{
					{OutString: "Y"},
					{OutString: "invalid_email"},
				},
			},
			expectedError: "invalid email address entered for Email: mail: missing '@' or angle-addr",
		},
		{
			name: "Success",
			f: fieldInfo{
				Value:       reflect.ValueOf(&email).Elem(),
				Name:        "Email",
				Kind:        KindEmail,
				Description: "Your email address",
				Sep:         ",",
			},
			asker: &MockAsker{
				AskMocks: []AskMock{
					{OutString: "Y"},
					{OutString: "jane.doe@example.com"},
				},
			},
			expectedError: "",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := askForField(tc.f, tc.asker)

			if tc.expectedError == "" {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tc.expectedError)
			}
		})
	}
}

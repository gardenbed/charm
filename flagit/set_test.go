package flagit

import (
	"net/url"
	"reflect"
	"regexp"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSetString(t *testing.T) {
	tests := []struct {
		name            string
		s               string
		val             string
		expectedUpdated bool
		expectedError   string
		expectedResult  string
	}{
		{
			"NewValue",
			"old", "new",
			true, "",
			"new",
		},
		{
			"NoNewValue",
			"same", "same",
			false, "",
			"same",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			v := reflect.ValueOf(&tc.s).Elem()
			updated, err := setString(v, tc.val)

			if tc.expectedError == "" {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tc.expectedError)
			}

			assert.Equal(t, tc.expectedUpdated, updated)
			assert.Equal(t, tc.expectedResult, tc.s)
		})
	}
}

func TestSetBool(t *testing.T) {
	tests := []struct {
		name            string
		b               bool
		val             string
		expectedUpdated bool
		expectedError   string
		expectedResult  bool
	}{
		{
			"NewValue",
			false, "true",
			true, "",
			true,
		},
		{
			"NoNewValue",
			true, "true",
			false, "",
			true,
		},
		{
			"InvalidValue",
			false, "invalid",
			false, `strconv.ParseBool: parsing "invalid": invalid syntax`,
			false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			v := reflect.ValueOf(&tc.b).Elem()
			updated, err := setBool(v, tc.val)

			if tc.expectedError == "" {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tc.expectedError)
			}

			assert.Equal(t, tc.expectedUpdated, updated)
			assert.Equal(t, tc.expectedResult, tc.b)
		})
	}
}

func TestSetInt(t *testing.T) {
	tests := []struct {
		name            string
		i               int
		val             string
		expectedUpdated bool
		expectedError   string
		expectedResult  int
	}{
		{
			"NewValue",
			-9223372036854775808, "9223372036854775807",
			true, "",
			9223372036854775807,
		},
		{
			"NoNewValue",
			9223372036854775807, "9223372036854775807",
			false, "",
			9223372036854775807,
		},
		{
			"InvalidValue",
			-9223372036854775808, "invalid",
			false, `strconv.ParseInt: parsing "invalid": invalid syntax`,
			-9223372036854775808,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			v := reflect.ValueOf(&tc.i).Elem()
			updated, err := setInt(v, tc.val)

			if tc.expectedError == "" {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tc.expectedError)
			}

			assert.Equal(t, tc.expectedUpdated, updated)
			assert.Equal(t, tc.expectedResult, tc.i)
		})
	}
}

func TestSetInt8(t *testing.T) {
	tests := []struct {
		name            string
		i               int8
		val             string
		expectedUpdated bool
		expectedError   string
		expectedResult  int8
	}{
		{
			"NewValue",
			-128, "127",
			true, "",
			127,
		},
		{
			"NoNewValue",
			127, "127",
			false, "",
			127,
		},
		{
			"InvalidValue",
			-128, "invalid",
			false, `strconv.ParseInt: parsing "invalid": invalid syntax`,
			-128,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			v := reflect.ValueOf(&tc.i).Elem()
			updated, err := setInt8(v, tc.val)

			if tc.expectedError == "" {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tc.expectedError)
			}

			assert.Equal(t, tc.expectedUpdated, updated)
			assert.Equal(t, tc.expectedResult, tc.i)
		})
	}
}

func TestSetInt16(t *testing.T) {
	tests := []struct {
		name            string
		i               int16
		val             string
		expectedUpdated bool
		expectedError   string
		expectedResult  int16
	}{
		{
			"NewValue",
			-32768, "32767",
			true, "",
			32767,
		},
		{
			"NoNewValue",
			32767, "32767",
			false, "",
			32767,
		},
		{
			"InvalidValue",
			-32768, "invalid",
			false, `strconv.ParseInt: parsing "invalid": invalid syntax`,
			-32768,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			v := reflect.ValueOf(&tc.i).Elem()
			updated, err := setInt16(v, tc.val)

			if tc.expectedError == "" {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tc.expectedError)
			}

			assert.Equal(t, tc.expectedUpdated, updated)
			assert.Equal(t, tc.expectedResult, tc.i)
		})
	}
}

func TestSetInt32(t *testing.T) {
	tests := []struct {
		name            string
		i               int32
		val             string
		expectedUpdated bool
		expectedError   string
		expectedResult  int32
	}{
		{
			"NewValue",
			-2147483648, "2147483647",
			true, "",
			2147483647,
		},
		{
			"NoNewValue",
			2147483647, "2147483647",
			false, "",
			2147483647,
		},
		{
			"InvalidValue",
			-2147483648, "invalid",
			false, `strconv.ParseInt: parsing "invalid": invalid syntax`,
			-2147483648,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			v := reflect.ValueOf(&tc.i).Elem()
			updated, err := setInt32(v, tc.val)

			if tc.expectedError == "" {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tc.expectedError)
			}

			assert.Equal(t, tc.expectedUpdated, updated)
			assert.Equal(t, tc.expectedResult, tc.i)
		})
	}
}

func TestSetInt64(t *testing.T) {
	tests := []struct {
		name            string
		i               int64
		val             string
		expectedUpdated bool
		expectedError   string
		expectedResult  int64
	}{
		{
			"NewValue",
			-9223372036854775808, "9223372036854775807",
			true, "",
			9223372036854775807,
		},
		{
			"NoNewValue",
			9223372036854775807, "9223372036854775807",
			false, "",
			9223372036854775807,
		},
		{
			"InvalidValue",
			-9223372036854775808, "invalid",
			false, `strconv.ParseInt: parsing "invalid": invalid syntax`,
			-9223372036854775808,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			v := reflect.ValueOf(&tc.i).Elem()
			updated, err := setInt64(v, tc.val)

			if tc.expectedError == "" {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tc.expectedError)
			}

			assert.Equal(t, tc.expectedUpdated, updated)
			assert.Equal(t, tc.expectedResult, tc.i)
		})
	}
}

func TestSetInt64Duration(t *testing.T) {
	tests := []struct {
		name            string
		i               time.Duration
		val             string
		expectedUpdated bool
		expectedError   string
		expectedResult  time.Duration
	}{
		{
			"NewValue",
			time.Second, "1m",
			true, "",
			time.Minute,
		},
		{
			"NoNewValue",
			time.Minute, "1m",
			false, "",
			time.Minute,
		},
		{
			"InvalidValue",
			time.Second, "invalid",
			false, `time: invalid duration "invalid"`,
			time.Second,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			v := reflect.ValueOf(&tc.i).Elem()
			updated, err := setInt64(v, tc.val)

			if tc.expectedError == "" {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tc.expectedError)
			}

			assert.Equal(t, tc.expectedUpdated, updated)
			assert.Equal(t, tc.expectedResult, tc.i)
		})
	}
}

func TestSetUint(t *testing.T) {
	tests := []struct {
		name            string
		u               uint
		val             string
		expectedUpdated bool
		expectedError   string
		expectedResult  uint
	}{
		{
			"NewValue",
			0, "18446744073709551615",
			true, "",
			18446744073709551615,
		},
		{
			"NoNewValue",
			18446744073709551615, "18446744073709551615",
			false, "",
			18446744073709551615,
		},
		{
			"InvalidValue",
			0, "invalid",
			false, `strconv.ParseUint: parsing "invalid": invalid syntax`,
			0,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			v := reflect.ValueOf(&tc.u).Elem()
			updated, err := setUint(v, tc.val)

			if tc.expectedError == "" {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tc.expectedError)
			}

			assert.Equal(t, tc.expectedUpdated, updated)
			assert.Equal(t, tc.expectedResult, tc.u)
		})
	}
}

func TestSetUint8(t *testing.T) {
	tests := []struct {
		name            string
		u               uint8
		val             string
		expectedUpdated bool
		expectedError   string
		expectedResult  uint8
	}{
		{
			"NewValue",
			0, "255",
			true, "",
			255,
		},
		{
			"NoNewValue",
			255, "255",
			false, "",
			255,
		},
		{
			"InvalidValue",
			0, "invalid",
			false, `strconv.ParseUint: parsing "invalid": invalid syntax`,
			0,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			v := reflect.ValueOf(&tc.u).Elem()
			updated, err := setUint8(v, tc.val)

			if tc.expectedError == "" {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tc.expectedError)
			}

			assert.Equal(t, tc.expectedUpdated, updated)
			assert.Equal(t, tc.expectedResult, tc.u)
		})
	}
}

func TestSetUint16(t *testing.T) {
	tests := []struct {
		name            string
		u               uint16
		val             string
		expectedUpdated bool
		expectedError   string
		expectedResult  uint16
	}{
		{
			"NewValue",
			0, "65535",
			true, "",
			65535,
		},
		{
			"NoNewValue",
			65535, "65535",
			false, "",
			65535,
		},
		{
			"InvalidValue",
			0, "invalid",
			false, `strconv.ParseUint: parsing "invalid": invalid syntax`,
			0,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			v := reflect.ValueOf(&tc.u).Elem()
			updated, err := setUint16(v, tc.val)

			if tc.expectedError == "" {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tc.expectedError)
			}

			assert.Equal(t, tc.expectedUpdated, updated)
			assert.Equal(t, tc.expectedResult, tc.u)
		})
	}
}

func TestSetUint32(t *testing.T) {
	tests := []struct {
		name            string
		u               uint32
		val             string
		expectedUpdated bool
		expectedError   string
		expectedResult  uint32
	}{
		{
			"NewValue",
			0, "4294967295",
			true, "",
			4294967295,
		},
		{
			"NoNewValue",
			4294967295, "4294967295",
			false, "",
			4294967295,
		},
		{
			"InvalidValue",
			0, "invalid",
			false, `strconv.ParseUint: parsing "invalid": invalid syntax`,
			0,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			v := reflect.ValueOf(&tc.u).Elem()
			updated, err := setUint32(v, tc.val)

			if tc.expectedError == "" {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tc.expectedError)
			}

			assert.Equal(t, tc.expectedUpdated, updated)
			assert.Equal(t, tc.expectedResult, tc.u)
		})
	}
}

func TestSetUint64(t *testing.T) {
	tests := []struct {
		name            string
		u               uint64
		val             string
		expectedUpdated bool
		expectedError   string
		expectedResult  uint64
	}{
		{
			"NewValue",
			0, "18446744073709551615",
			true, "",
			18446744073709551615,
		},
		{
			"NoNewValue",
			18446744073709551615, "18446744073709551615",
			false, "",
			18446744073709551615,
		},
		{
			"InvalidValue",
			0, "invalid",
			false, `strconv.ParseUint: parsing "invalid": invalid syntax`,
			0,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			v := reflect.ValueOf(&tc.u).Elem()
			updated, err := setUint64(v, tc.val)

			if tc.expectedError == "" {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tc.expectedError)
			}

			assert.Equal(t, tc.expectedUpdated, updated)
			assert.Equal(t, tc.expectedResult, tc.u)
		})
	}
}

func TestSetFloat32(t *testing.T) {
	tests := []struct {
		name            string
		f               float32
		val             string
		expectedUpdated bool
		expectedError   string
		expectedResult  float32
	}{
		{
			"NewValue",
			3.1415, "2.7182",
			true, "",
			2.7182,
		},
		{
			"NoNewValue",
			2.7182, "2.7182",
			false, "",
			2.7182,
		},
		{
			"InvalidValue",
			3.1415, "invalid",
			false, `strconv.ParseFloat: parsing "invalid": invalid syntax`,
			3.1415,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			v := reflect.ValueOf(&tc.f).Elem()
			updated, err := setFloat32(v, tc.val)

			if tc.expectedError == "" {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tc.expectedError)
			}

			assert.Equal(t, tc.expectedUpdated, updated)
			assert.Equal(t, tc.expectedResult, tc.f)
		})
	}
}

func TestSetFloat64(t *testing.T) {
	tests := []struct {
		name            string
		f               float64
		val             string
		expectedUpdated bool
		expectedError   string
		expectedResult  float64
	}{
		{
			"NewValue",
			3.14159265359, "2.7182818284",
			true, "",
			2.7182818284,
		},
		{
			"NoNewValue",
			2.7182818284, "2.7182818284",
			false, "",
			2.7182818284,
		},
		{
			"InvalidValue",
			3.14159265359, "invalid",
			false, `strconv.ParseFloat: parsing "invalid": invalid syntax`,
			3.14159265359,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			v := reflect.ValueOf(&tc.f).Elem()
			updated, err := setFloat64(v, tc.val)

			if tc.expectedError == "" {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tc.expectedError)
			}

			assert.Equal(t, tc.expectedUpdated, updated)
			assert.Equal(t, tc.expectedResult, tc.f)
		})
	}
}

func TestSetStructURL(t *testing.T) {
	url1, _ := url.Parse("service-1")
	url2, _ := url.Parse("service-2")

	tests := []struct {
		name            string
		s               url.URL
		val             string
		expectedUpdated bool
		expectedError   string
		expectedResult  url.URL
	}{
		{
			"URLNewValue",
			*url1, "service-2",
			true, "",
			*url2,
		},
		{
			"URLNoNewValue",
			*url2, "service-2",
			false, "",
			*url2,
		},
		{
			"URLInvalidValue",
			*url1, ":",
			false, `parse ":": missing protocol scheme`,
			*url1,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			v := reflect.ValueOf(&tc.s).Elem()
			updated, err := setStruct(v, tc.val)

			if tc.expectedError == "" {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tc.expectedError)
			}

			assert.Equal(t, tc.expectedUpdated, updated)
			assert.Equal(t, tc.expectedResult, tc.s)
		})
	}
}

func TestSetStructRegexp(t *testing.T) {
	re1 := regexp.MustCompilePOSIX("[:digit:]")
	re2 := regexp.MustCompilePOSIX("[:alpha:]")

	tests := []struct {
		name            string
		s               regexp.Regexp
		val             string
		expectedUpdated bool
		expectedError   string
		expectedResult  regexp.Regexp
	}{
		{
			"RegexpNewValue",
			*re1, "[:alpha:]",
			true, "",
			*re2,
		},
		{
			"RegexpNoNewValue",
			*re2, "[:alpha:]",
			false, "",
			*re2,
		},
		{
			"RegexpInvalidValue",
			*re1, "[:invalid:",
			false, "error parsing regexp: missing closing ]: `[:invalid:`",
			*re1,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			v := reflect.ValueOf(&tc.s).Elem()
			updated, err := setStruct(v, tc.val)

			if tc.expectedError == "" {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tc.expectedError)
			}

			assert.Equal(t, tc.expectedUpdated, updated)
			assert.Equal(t, tc.expectedResult, tc.s)
		})
	}
}

func TestSetStringPtr(t *testing.T) {
	tests := []struct {
		name            string
		s               *string
		val             string
		expectedUpdated bool
		expectedError   string
		expectedResult  *string
	}{
		{
			"Nil",
			nil, "new",
			true, "",
			stringPtr("new"),
		},
		{
			"NewValue",
			stringPtr("old"), "new",
			true, "",
			stringPtr("new"),
		},
		{
			"NoNewValue",
			stringPtr("same"), "same",
			false, "",
			stringPtr("same"),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			v := reflect.ValueOf(&tc.s).Elem()
			updated, err := setStringPtr(v, tc.val)

			if tc.expectedError == "" {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tc.expectedError)
			}

			assert.Equal(t, tc.expectedUpdated, updated)
			assert.Equal(t, tc.expectedResult, tc.s)
		})
	}
}

func TestSetBoolPtr(t *testing.T) {
	tests := []struct {
		name            string
		b               *bool
		val             string
		expectedUpdated bool
		expectedError   string
		expectedResult  *bool
	}{
		{
			"Nil",
			nil, "true",
			true, "",
			boolPtr(true),
		},
		{
			"NewValue",
			boolPtr(false), "true",
			true, "",
			boolPtr(true),
		},
		{
			"NoNewValue",
			boolPtr(true), "true",
			false, "",
			boolPtr(true),
		},
		{
			"InvalidValue",
			boolPtr(false), "invalid",
			false, `strconv.ParseBool: parsing "invalid": invalid syntax`,
			boolPtr(false),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			v := reflect.ValueOf(&tc.b).Elem()
			updated, err := setBoolPtr(v, tc.val)

			if tc.expectedError == "" {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tc.expectedError)
			}

			assert.Equal(t, tc.expectedUpdated, updated)
			assert.Equal(t, tc.expectedResult, tc.b)
		})
	}
}

func TestSetIntPtr(t *testing.T) {
	tests := []struct {
		name            string
		i               *int
		val             string
		expectedUpdated bool
		expectedError   string
		expectedResult  *int
	}{
		{
			"Nil",
			nil, "9223372036854775807",
			true, "",
			intPtr(9223372036854775807),
		},
		{
			"NewValue",
			intPtr(-9223372036854775808), "9223372036854775807",
			true, "",
			intPtr(9223372036854775807),
		},
		{
			"NoNewValue",
			intPtr(9223372036854775807), "9223372036854775807",
			false, "",
			intPtr(9223372036854775807),
		},
		{
			"InvalidValue",
			intPtr(-9223372036854775808), "invalid",
			false, `strconv.ParseInt: parsing "invalid": invalid syntax`,
			intPtr(-9223372036854775808),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			v := reflect.ValueOf(&tc.i).Elem()
			updated, err := setIntPtr(v, tc.val)

			if tc.expectedError == "" {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tc.expectedError)
			}

			assert.Equal(t, tc.expectedUpdated, updated)
			assert.Equal(t, tc.expectedResult, tc.i)
		})
	}
}

func TestSetInt8Ptr(t *testing.T) {
	tests := []struct {
		name            string
		i               *int8
		val             string
		expectedUpdated bool
		expectedError   string
		expectedResult  *int8
	}{
		{
			"Nil",
			nil, "127",
			true, "",
			int8Ptr(127),
		},
		{
			"NewValue",
			int8Ptr(-128), "127",
			true, "",
			int8Ptr(127),
		},
		{
			"NoNewValue",
			int8Ptr(127), "127",
			false, "",
			int8Ptr(127),
		},
		{
			"InvalidValue",
			int8Ptr(-128), "invalid",
			false, `strconv.ParseInt: parsing "invalid": invalid syntax`,
			int8Ptr(-128),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			v := reflect.ValueOf(&tc.i).Elem()
			updated, err := setInt8Ptr(v, tc.val)

			if tc.expectedError == "" {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tc.expectedError)
			}

			assert.Equal(t, tc.expectedUpdated, updated)
			assert.Equal(t, tc.expectedResult, tc.i)
		})
	}
}

func TestSetInt16Ptr(t *testing.T) {
	tests := []struct {
		name            string
		i               *int16
		val             string
		expectedUpdated bool
		expectedError   string
		expectedResult  *int16
	}{
		{
			"Nil",
			nil, "32767",
			true, "",
			int16Ptr(32767),
		},
		{
			"NewValue",
			int16Ptr(-32768), "32767",
			true, "",
			int16Ptr(32767),
		},
		{
			"NoNewValue",
			int16Ptr(32767), "32767",
			false, "",
			int16Ptr(32767),
		},
		{
			"InvalidValue",
			int16Ptr(-32768), "invalid",
			false, `strconv.ParseInt: parsing "invalid": invalid syntax`,
			int16Ptr(-32768),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			v := reflect.ValueOf(&tc.i).Elem()
			updated, err := setInt16Ptr(v, tc.val)

			if tc.expectedError == "" {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tc.expectedError)
			}

			assert.Equal(t, tc.expectedUpdated, updated)
			assert.Equal(t, tc.expectedResult, tc.i)
		})
	}
}

func TestSetInt32Ptr(t *testing.T) {
	tests := []struct {
		name            string
		i               *int32
		val             string
		expectedUpdated bool
		expectedError   string
		expectedResult  *int32
	}{
		{
			"Nil",
			nil, "2147483647",
			true, "",
			int32Ptr(2147483647),
		},
		{
			"NewValue",
			int32Ptr(-2147483648), "2147483647",
			true, "",
			int32Ptr(2147483647),
		},
		{
			"NoNewValue",
			int32Ptr(2147483647), "2147483647",
			false, "",
			int32Ptr(2147483647),
		},
		{
			"InvalidValue",
			int32Ptr(-2147483648), "invalid",
			false, `strconv.ParseInt: parsing "invalid": invalid syntax`,
			int32Ptr(-2147483648),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			v := reflect.ValueOf(&tc.i).Elem()
			updated, err := setInt32Ptr(v, tc.val)

			if tc.expectedError == "" {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tc.expectedError)
			}

			assert.Equal(t, tc.expectedUpdated, updated)
			assert.Equal(t, tc.expectedResult, tc.i)
		})
	}
}

func TestSetInt64Ptr(t *testing.T) {
	tests := []struct {
		name            string
		i               *int64
		val             string
		expectedUpdated bool
		expectedError   string
		expectedResult  *int64
	}{
		{
			"Nil",
			nil, "9223372036854775807",
			true, "",
			int64Ptr(9223372036854775807),
		},
		{
			"NewValue",
			int64Ptr(-9223372036854775808), "9223372036854775807",
			true, "",
			int64Ptr(9223372036854775807),
		},
		{
			"NoNewValue",
			int64Ptr(9223372036854775807), "9223372036854775807",
			false, "",
			int64Ptr(9223372036854775807),
		},
		{
			"InvalidValue",
			int64Ptr(-9223372036854775808), "invalid",
			false, `strconv.ParseInt: parsing "invalid": invalid syntax`,
			int64Ptr(-9223372036854775808),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			v := reflect.ValueOf(&tc.i).Elem()
			updated, err := setInt64Ptr(v, tc.val)

			if tc.expectedError == "" {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tc.expectedError)
			}

			assert.Equal(t, tc.expectedUpdated, updated)
			assert.Equal(t, tc.expectedResult, tc.i)
		})
	}
}

func TestSetInt64DurationPtr(t *testing.T) {
	tests := []struct {
		name            string
		i               *time.Duration
		val             string
		expectedUpdated bool
		expectedError   string
		expectedResult  *time.Duration
	}{
		{
			"Nil",
			nil, "1m",
			true, "",
			durationPtr(time.Minute),
		},
		{
			"NewValue",
			durationPtr(time.Second), "1m",
			true, "",
			durationPtr(time.Minute),
		},
		{
			"NoNewValue",
			durationPtr(time.Minute), "1m",
			false, "",
			durationPtr(time.Minute),
		},
		{
			"InvalidValue",
			durationPtr(time.Second), "invalid",
			false, `time: invalid duration "invalid"`,
			durationPtr(time.Second),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			v := reflect.ValueOf(&tc.i).Elem()
			updated, err := setInt64Ptr(v, tc.val)

			if tc.expectedError == "" {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tc.expectedError)
			}

			assert.Equal(t, tc.expectedUpdated, updated)
			assert.Equal(t, tc.expectedResult, tc.i)
		})
	}
}

func TestSetUintPtr(t *testing.T) {
	tests := []struct {
		name            string
		u               *uint
		val             string
		expectedUpdated bool
		expectedError   string
		expectedResult  *uint
	}{
		{
			"Nil",
			nil, "18446744073709551615",
			true, "",
			uintPtr(18446744073709551615),
		},
		{
			"NewValue",
			uintPtr(0), "18446744073709551615",
			true, "",
			uintPtr(18446744073709551615),
		},
		{
			"NoNewValue",
			uintPtr(18446744073709551615), "18446744073709551615",
			false, "",
			uintPtr(18446744073709551615),
		},
		{
			"InvalidValue",
			uintPtr(0), "invalid",
			false, `strconv.ParseUint: parsing "invalid": invalid syntax`,
			uintPtr(0),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			v := reflect.ValueOf(&tc.u).Elem()
			updated, err := setUintPtr(v, tc.val)

			if tc.expectedError == "" {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tc.expectedError)
			}

			assert.Equal(t, tc.expectedUpdated, updated)
			assert.Equal(t, tc.expectedResult, tc.u)
		})
	}
}

func TestSetUint8Ptr(t *testing.T) {
	tests := []struct {
		name            string
		u               *uint8
		val             string
		expectedUpdated bool
		expectedError   string
		expectedResult  *uint8
	}{
		{
			"Nil",
			nil, "255",
			true, "",
			uint8Ptr(255),
		},
		{
			"NewValue",
			uint8Ptr(0), "255",
			true, "",
			uint8Ptr(255),
		},
		{
			"NoNewValue",
			uint8Ptr(255), "255",
			false, "",
			uint8Ptr(255),
		},
		{
			"InvalidValue",
			uint8Ptr(0), "invalid",
			false, `strconv.ParseUint: parsing "invalid": invalid syntax`,
			uint8Ptr(0),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			v := reflect.ValueOf(&tc.u).Elem()
			updated, err := setUint8Ptr(v, tc.val)

			if tc.expectedError == "" {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tc.expectedError)
			}

			assert.Equal(t, tc.expectedUpdated, updated)
			assert.Equal(t, tc.expectedResult, tc.u)
		})
	}
}

func TestSetUint16Ptr(t *testing.T) {
	tests := []struct {
		name            string
		u               *uint16
		val             string
		expectedUpdated bool
		expectedError   string
		expectedResult  *uint16
	}{
		{
			"Nil",
			nil, "65535",
			true, "",
			uint16Ptr(65535),
		},
		{
			"NewValue",
			uint16Ptr(0), "65535",
			true, "",
			uint16Ptr(65535),
		},
		{
			"NoNewValue",
			uint16Ptr(65535), "65535",
			false, "",
			uint16Ptr(65535),
		},
		{
			"InvalidValue",
			uint16Ptr(0), "invalid",
			false, `strconv.ParseUint: parsing "invalid": invalid syntax`,
			uint16Ptr(0),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			v := reflect.ValueOf(&tc.u).Elem()
			updated, err := setUint16Ptr(v, tc.val)

			if tc.expectedError == "" {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tc.expectedError)
			}

			assert.Equal(t, tc.expectedUpdated, updated)
			assert.Equal(t, tc.expectedResult, tc.u)
		})
	}
}

func TestSetUint32Ptr(t *testing.T) {
	tests := []struct {
		name            string
		u               *uint32
		val             string
		expectedUpdated bool
		expectedError   string
		expectedResult  *uint32
	}{
		{
			"Nil",
			nil, "4294967295",
			true, "",
			uint32Ptr(4294967295),
		},
		{
			"NewValue",
			uint32Ptr(0), "4294967295",
			true, "",
			uint32Ptr(4294967295),
		},
		{
			"NoNewValue",
			uint32Ptr(4294967295), "4294967295",
			false, "",
			uint32Ptr(4294967295),
		},
		{
			"InvalidValue",
			uint32Ptr(0), "invalid",
			false, `strconv.ParseUint: parsing "invalid": invalid syntax`,
			uint32Ptr(0),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			v := reflect.ValueOf(&tc.u).Elem()
			updated, err := setUint32Ptr(v, tc.val)

			if tc.expectedError == "" {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tc.expectedError)
			}

			assert.Equal(t, tc.expectedUpdated, updated)
			assert.Equal(t, tc.expectedResult, tc.u)
		})
	}
}

func TestSetUint64Ptr(t *testing.T) {
	tests := []struct {
		name            string
		u               *uint64
		val             string
		expectedUpdated bool
		expectedError   string
		expectedResult  *uint64
	}{
		{
			"Nil",
			nil, "18446744073709551615",
			true, "",
			uint64Ptr(18446744073709551615),
		},
		{
			"NewValue",
			uint64Ptr(0), "18446744073709551615",
			true, "",
			uint64Ptr(18446744073709551615),
		},
		{
			"NoNewValue",
			uint64Ptr(18446744073709551615), "18446744073709551615",
			false, "",
			uint64Ptr(18446744073709551615),
		},
		{
			"InvalidValue",
			uint64Ptr(0), "invalid",
			false, `strconv.ParseUint: parsing "invalid": invalid syntax`,
			uint64Ptr(0),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			v := reflect.ValueOf(&tc.u).Elem()
			updated, err := setUint64Ptr(v, tc.val)

			if tc.expectedError == "" {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tc.expectedError)
			}

			assert.Equal(t, tc.expectedUpdated, updated)
			assert.Equal(t, tc.expectedResult, tc.u)
		})
	}
}

func TestSetFloat32Ptr(t *testing.T) {
	tests := []struct {
		name            string
		f               *float32
		val             string
		expectedUpdated bool
		expectedError   string
		expectedResult  *float32
	}{
		{
			"Nil",
			nil, "2.7182",
			true, "",
			float32Ptr(2.7182),
		},
		{
			"NewValue",
			float32Ptr(3.1415), "2.7182",
			true, "",
			float32Ptr(2.7182),
		},
		{
			"NoNewValue",
			float32Ptr(2.7182), "2.7182",
			false, "",
			float32Ptr(2.7182),
		},
		{
			"InvalidValue",
			float32Ptr(3.1415), "invalid",
			false, `strconv.ParseFloat: parsing "invalid": invalid syntax`,
			float32Ptr(3.1415),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			v := reflect.ValueOf(&tc.f).Elem()
			updated, err := setFloat32Ptr(v, tc.val)

			if tc.expectedError == "" {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tc.expectedError)
			}

			assert.Equal(t, tc.expectedUpdated, updated)
			assert.Equal(t, tc.expectedResult, tc.f)
		})
	}
}

func TestSetFloat64Ptr(t *testing.T) {
	tests := []struct {
		name            string
		f               *float64
		val             string
		expectedUpdated bool
		expectedError   string
		expectedResult  *float64
	}{
		{
			"Nil",
			nil, "2.7182818284",
			true, "",
			float64Ptr(2.7182818284),
		},
		{
			"NewValue",
			float64Ptr(3.14159265359), "2.7182818284",
			true, "",
			float64Ptr(2.7182818284),
		},
		{
			"NoNewValue",
			float64Ptr(2.7182818284), "2.7182818284",
			false, "",
			float64Ptr(2.7182818284),
		},
		{
			"InvalidValue",
			float64Ptr(3.14159265359), "invalid",
			false, `strconv.ParseFloat: parsing "invalid": invalid syntax`,
			float64Ptr(3.14159265359),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			v := reflect.ValueOf(&tc.f).Elem()
			updated, err := setFloat64Ptr(v, tc.val)

			if tc.expectedError == "" {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tc.expectedError)
			}

			assert.Equal(t, tc.expectedUpdated, updated)
			assert.Equal(t, tc.expectedResult, tc.f)
		})
	}
}

func TestSetStructPtrURL(t *testing.T) {
	url1, _ := url.Parse("service-1")
	url2, _ := url.Parse("service-2")

	tests := []struct {
		name            string
		s               *url.URL
		val             string
		expectedUpdated bool
		expectedError   string
		expectedResult  *url.URL
	}{
		{
			"URLNil",
			nil, "service-2",
			true, "",
			url2,
		},
		{
			"URLNewValue",
			url1, "service-2",
			true, "",
			url2,
		},
		{
			"URLNoNewValue",
			url2, "service-2",
			false, "",
			url2,
		},
		{
			"URLInvalidValue",
			url1, ":",
			false, `parse ":": missing protocol scheme`,
			url1,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			v := reflect.ValueOf(&tc.s).Elem()
			updated, err := setStructPtr(v, tc.val)

			if tc.expectedError == "" {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tc.expectedError)
			}

			assert.Equal(t, tc.expectedUpdated, updated)
			assert.Equal(t, tc.expectedResult, tc.s)
		})
	}
}

func TestSetStructPtrRegexp(t *testing.T) {
	re1 := regexp.MustCompilePOSIX("[:digit:]")
	re2 := regexp.MustCompilePOSIX("[:alpha:]")

	tests := []struct {
		name            string
		s               *regexp.Regexp
		val             string
		expectedUpdated bool
		expectedError   string
		expectedResult  *regexp.Regexp
	}{
		{
			"RegexpNil",
			nil, "[:alpha:]",
			true, "",
			re2,
		},
		{
			"RegexpNewValue",
			re1, "[:alpha:]",
			true, "",
			re2,
		},
		{
			"RegexpNoNewValue",
			re2, "[:alpha:]",
			false, "",
			re2,
		},
		{
			"RegexpInvalidValue",
			re1, "[:invalid:",
			false, "error parsing regexp: missing closing ]: `[:invalid:`",
			re1,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			v := reflect.ValueOf(&tc.s).Elem()
			updated, err := setStructPtr(v, tc.val)

			if tc.expectedError == "" {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tc.expectedError)
			}

			assert.Equal(t, tc.expectedUpdated, updated)
			assert.Equal(t, tc.expectedResult, tc.s)
		})
	}
}

func TestSetStringSlice(t *testing.T) {
	tests := []struct {
		name            string
		s               []string
		vals            []string
		expectedUpdated bool
		expectedError   string
		expectedResult  []string
	}{
		{
			"Nil",
			nil, []string{"new"},
			true, "",
			[]string{"new"},
		},
		{
			"NewValue",
			[]string{"old"}, []string{"new"},
			true, "",
			[]string{"new"},
		},
		{
			"NoNewValue",
			[]string{"same"}, []string{"same"},
			false, "",
			[]string{"same"},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			v := reflect.ValueOf(&tc.s).Elem()
			updated, err := setStringSlice(v, tc.vals)

			if tc.expectedError == "" {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tc.expectedError)
			}

			assert.Equal(t, tc.expectedUpdated, updated)
			assert.Equal(t, tc.expectedResult, tc.s)
		})
	}
}

func TestSetBoolSlice(t *testing.T) {
	tests := []struct {
		name            string
		b               []bool
		vals            []string
		expectedUpdated bool
		expectedError   string
		expectedResult  []bool
	}{
		{
			"Nil",
			nil, []string{"true"},
			true, "",
			[]bool{true},
		},
		{
			"NewValue",
			[]bool{false}, []string{"true"},
			true, "",
			[]bool{true},
		},
		{
			"NoNewValue",
			[]bool{true}, []string{"true"},
			false, "",
			[]bool{true},
		},
		{
			"InvalidValue",
			[]bool{false}, []string{"invalid"},
			false, `strconv.ParseBool: parsing "invalid": invalid syntax`,
			[]bool{false},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			v := reflect.ValueOf(&tc.b).Elem()
			updated, err := setBoolSlice(v, tc.vals)

			if tc.expectedError == "" {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tc.expectedError)
			}

			assert.Equal(t, tc.expectedUpdated, updated)
			assert.Equal(t, tc.expectedResult, tc.b)
		})
	}
}

func TestSetIntSlice(t *testing.T) {
	tests := []struct {
		name            string
		i               []int
		vals            []string
		expectedUpdated bool
		expectedError   string
		expectedResult  []int
	}{
		{
			"Nil",
			nil, []string{"9223372036854775807"},
			true, "",
			[]int{9223372036854775807},
		},
		{
			"NewValue",
			[]int{-9223372036854775808}, []string{"9223372036854775807"},
			true, "",
			[]int{9223372036854775807},
		},
		{
			"NoNewValue",
			[]int{9223372036854775807}, []string{"9223372036854775807"},
			false, "",
			[]int{9223372036854775807},
		},
		{
			"InvalidValue",
			[]int{-9223372036854775808}, []string{"invalid"},
			false, `strconv.ParseInt: parsing "invalid": invalid syntax`,
			[]int{-9223372036854775808},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			v := reflect.ValueOf(&tc.i).Elem()
			updated, err := setIntSlice(v, tc.vals)

			if tc.expectedError == "" {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tc.expectedError)
			}

			assert.Equal(t, tc.expectedUpdated, updated)
			assert.Equal(t, tc.expectedResult, tc.i)
		})
	}
}

func TestSetInt8Slice(t *testing.T) {
	tests := []struct {
		name            string
		i               []int8
		vals            []string
		expectedUpdated bool
		expectedError   string
		expectedResult  []int8
	}{
		{
			"Nil",
			nil, []string{"127"},
			true, "",
			[]int8{127},
		},
		{
			"NewValue",
			[]int8{-128}, []string{"127"},
			true, "",
			[]int8{127},
		},
		{
			"NoNewValue",
			[]int8{127}, []string{"127"},
			false, "",
			[]int8{127},
		},
		{
			"InvalidValue",
			[]int8{-128}, []string{"invalid"},
			false, `strconv.ParseInt: parsing "invalid": invalid syntax`,
			[]int8{-128},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			v := reflect.ValueOf(&tc.i).Elem()
			updated, err := setInt8Slice(v, tc.vals)

			if tc.expectedError == "" {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tc.expectedError)
			}

			assert.Equal(t, tc.expectedUpdated, updated)
			assert.Equal(t, tc.expectedResult, tc.i)
		})
	}
}

func TestSetInt16Slice(t *testing.T) {
	tests := []struct {
		name            string
		i               []int16
		vals            []string
		expectedUpdated bool
		expectedError   string
		expectedResult  []int16
	}{
		{
			"Nil",
			nil, []string{"32767"},
			true, "",
			[]int16{32767},
		},
		{
			"NewValue",
			[]int16{-32768}, []string{"32767"},
			true, "",
			[]int16{32767},
		},
		{
			"NoNewValue",
			[]int16{32767}, []string{"32767"},
			false, "",
			[]int16{32767},
		},
		{
			"InvalidValue",
			[]int16{-32768}, []string{"invalid"},
			false, `strconv.ParseInt: parsing "invalid": invalid syntax`,
			[]int16{-32768},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			v := reflect.ValueOf(&tc.i).Elem()
			updated, err := setInt16Slice(v, tc.vals)

			if tc.expectedError == "" {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tc.expectedError)
			}

			assert.Equal(t, tc.expectedUpdated, updated)
			assert.Equal(t, tc.expectedResult, tc.i)
		})
	}
}

func TestSetInt32Slice(t *testing.T) {
	tests := []struct {
		name            string
		i               []int32
		vals            []string
		expectedUpdated bool
		expectedError   string
		expectedResult  []int32
	}{
		{
			"Nil",
			nil, []string{"2147483647"},
			true, "",
			[]int32{2147483647},
		},
		{
			"NewValue",
			[]int32{-2147483648}, []string{"2147483647"},
			true, "",
			[]int32{2147483647},
		},
		{
			"NoNewValue",
			[]int32{2147483647}, []string{"2147483647"},
			false, "",
			[]int32{2147483647},
		},
		{
			"InvalidValue",
			[]int32{-2147483648}, []string{"invalid"},
			false, `strconv.ParseInt: parsing "invalid": invalid syntax`,
			[]int32{-2147483648},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			v := reflect.ValueOf(&tc.i).Elem()
			updated, err := setInt32Slice(v, tc.vals)

			if tc.expectedError == "" {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tc.expectedError)
			}

			assert.Equal(t, tc.expectedUpdated, updated)
			assert.Equal(t, tc.expectedResult, tc.i)
		})
	}
}

func TestSetInt64Slice(t *testing.T) {
	tests := []struct {
		name            string
		i               []int64
		vals            []string
		expectedUpdated bool
		expectedError   string
		expectedResult  []int64
	}{
		{
			"Nil",
			nil, []string{"9223372036854775807"},
			true, "",
			[]int64{9223372036854775807},
		},
		{
			"NewValue",
			[]int64{-9223372036854775808}, []string{"9223372036854775807"},
			true, "",
			[]int64{9223372036854775807},
		},
		{
			"NoNewValue",
			[]int64{9223372036854775807}, []string{"9223372036854775807"},
			false, "",
			[]int64{9223372036854775807},
		},
		{
			"InvalidValue",
			[]int64{-9223372036854775808}, []string{"invalid"},
			false, `strconv.ParseInt: parsing "invalid": invalid syntax`,
			[]int64{-9223372036854775808},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			v := reflect.ValueOf(&tc.i).Elem()
			updated, err := setInt64Slice(v, tc.vals)

			if tc.expectedError == "" {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tc.expectedError)
			}

			assert.Equal(t, tc.expectedUpdated, updated)
			assert.Equal(t, tc.expectedResult, tc.i)
		})
	}
}

func TestSetInt64DurationSlice(t *testing.T) {
	tests := []struct {
		name            string
		i               []time.Duration
		vals            []string
		expectedUpdated bool
		expectedError   string
		expectedResult  []time.Duration
	}{
		{
			"Nil",
			nil, []string{"1m"},
			true, "",
			[]time.Duration{time.Minute},
		},
		{
			"NewValue",
			[]time.Duration{time.Second}, []string{"1m"},
			true, "",
			[]time.Duration{time.Minute},
		},
		{
			"NoNewValue",
			[]time.Duration{time.Minute}, []string{"1m"},
			false, "",
			[]time.Duration{time.Minute},
		},
		{
			"InvalidValue",
			[]time.Duration{time.Second}, []string{"invalid"},
			false, `time: invalid duration "invalid"`,
			[]time.Duration{time.Second},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			v := reflect.ValueOf(&tc.i).Elem()
			updated, err := setInt64Slice(v, tc.vals)

			if tc.expectedError == "" {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tc.expectedError)
			}

			assert.Equal(t, tc.expectedUpdated, updated)
			assert.Equal(t, tc.expectedResult, tc.i)
		})
	}
}

func TestSetUintSlice(t *testing.T) {
	tests := []struct {
		name            string
		u               []uint
		vals            []string
		expectedUpdated bool
		expectedError   string
		expectedResult  []uint
	}{
		{
			"Nil",
			nil, []string{"18446744073709551615"},
			true, "",
			[]uint{18446744073709551615},
		},
		{
			"NewValue",
			[]uint{0}, []string{"18446744073709551615"},
			true, "",
			[]uint{18446744073709551615},
		},
		{
			"NoNewValue",
			[]uint{18446744073709551615}, []string{"18446744073709551615"},
			false, "",
			[]uint{18446744073709551615},
		},
		{
			"InvalidValue",
			[]uint{0}, []string{"invalid"},
			false, `strconv.ParseUint: parsing "invalid": invalid syntax`,
			[]uint{0},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			v := reflect.ValueOf(&tc.u).Elem()
			updated, err := setUintSlice(v, tc.vals)

			if tc.expectedError == "" {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tc.expectedError)
			}

			assert.Equal(t, tc.expectedUpdated, updated)
			assert.Equal(t, tc.expectedResult, tc.u)
		})
	}
}

func TestSetUint8Slice(t *testing.T) {
	tests := []struct {
		name            string
		u               []uint8
		vals            []string
		expectedUpdated bool
		expectedError   string
		expectedResult  []uint8
	}{
		{
			"Nil",
			nil, []string{"255"},
			true, "",
			[]uint8{255},
		},
		{
			"NewValue",
			[]uint8{0}, []string{"255"},
			true, "",
			[]uint8{255},
		},
		{
			"NoNewValue",
			[]uint8{255}, []string{"255"},
			false, "",
			[]uint8{255},
		},
		{
			"InvalidValue",
			[]uint8{0}, []string{"invalid"},
			false, `strconv.ParseUint: parsing "invalid": invalid syntax`,
			[]uint8{0},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			v := reflect.ValueOf(&tc.u).Elem()
			updated, err := setUint8Slice(v, tc.vals)

			if tc.expectedError == "" {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tc.expectedError)
			}

			assert.Equal(t, tc.expectedUpdated, updated)
			assert.Equal(t, tc.expectedResult, tc.u)
		})
	}
}

func TestSetUint16Slice(t *testing.T) {
	tests := []struct {
		name            string
		u               []uint16
		vals            []string
		expectedUpdated bool
		expectedError   string
		expectedResult  []uint16
	}{
		{
			"Nil",
			nil, []string{"65535"},
			true, "",
			[]uint16{65535},
		},
		{
			"NewValue",
			[]uint16{0}, []string{"65535"},
			true, "",
			[]uint16{65535},
		},
		{
			"NoNewValue",
			[]uint16{65535}, []string{"65535"},
			false, "",
			[]uint16{65535},
		},
		{
			"InvalidValue",
			[]uint16{0}, []string{"invalid"},
			false, `strconv.ParseUint: parsing "invalid": invalid syntax`,
			[]uint16{0},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			v := reflect.ValueOf(&tc.u).Elem()
			updated, err := setUint16Slice(v, tc.vals)

			if tc.expectedError == "" {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tc.expectedError)
			}

			assert.Equal(t, tc.expectedUpdated, updated)
			assert.Equal(t, tc.expectedResult, tc.u)
		})
	}
}

func TestSetUint32Slice(t *testing.T) {
	tests := []struct {
		name            string
		u               []uint32
		vals            []string
		expectedUpdated bool
		expectedError   string
		expectedResult  []uint32
	}{
		{
			"Nil",
			nil, []string{"4294967295"},
			true, "",
			[]uint32{4294967295},
		},
		{
			"NewValue",
			[]uint32{0}, []string{"4294967295"},
			true, "",
			[]uint32{4294967295},
		},
		{
			"NoNewValue",
			[]uint32{4294967295}, []string{"4294967295"},
			false, "",
			[]uint32{4294967295},
		},
		{
			"InvalidValue",
			[]uint32{0}, []string{"invalid"},
			false, `strconv.ParseUint: parsing "invalid": invalid syntax`,
			[]uint32{0},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			v := reflect.ValueOf(&tc.u).Elem()
			updated, err := setUint32Slice(v, tc.vals)

			if tc.expectedError == "" {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tc.expectedError)
			}

			assert.Equal(t, tc.expectedUpdated, updated)
			assert.Equal(t, tc.expectedResult, tc.u)
		})
	}
}

func TestSetUint64Slice(t *testing.T) {
	tests := []struct {
		name            string
		u               []uint64
		vals            []string
		expectedUpdated bool
		expectedError   string
		expectedResult  []uint64
	}{
		{
			"Nil",
			nil, []string{"18446744073709551615"},
			true, "",
			[]uint64{18446744073709551615},
		},
		{
			"NewValue",
			[]uint64{0}, []string{"18446744073709551615"},
			true, "",
			[]uint64{18446744073709551615},
		},
		{
			"NoNewValue",
			[]uint64{18446744073709551615}, []string{"18446744073709551615"},
			false, "",
			[]uint64{18446744073709551615},
		},
		{
			"InvalidValue",
			[]uint64{0}, []string{"invalid"},
			false, `strconv.ParseUint: parsing "invalid": invalid syntax`,
			[]uint64{0},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			v := reflect.ValueOf(&tc.u).Elem()
			updated, err := setUint64Slice(v, tc.vals)

			if tc.expectedError == "" {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tc.expectedError)
			}

			assert.Equal(t, tc.expectedUpdated, updated)
			assert.Equal(t, tc.expectedResult, tc.u)
		})
	}
}

func TestSetFloat32Slice(t *testing.T) {
	tests := []struct {
		name            string
		f               []float32
		vals            []string
		expectedUpdated bool
		expectedError   string
		expectedResult  []float32
	}{
		{
			"Nil",
			nil, []string{"2.7182"},
			true, "",
			[]float32{2.7182},
		},
		{
			"NewValue",
			[]float32{3.1415}, []string{"2.7182"},
			true, "",
			[]float32{2.7182},
		},
		{
			"NoNewValue",
			[]float32{2.7182}, []string{"2.7182"},
			false, "",
			[]float32{2.7182},
		},
		{
			"InvalidValue",
			[]float32{3.1415}, []string{"invalid"},
			false, `strconv.ParseFloat: parsing "invalid": invalid syntax`,
			[]float32{3.1415},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			v := reflect.ValueOf(&tc.f).Elem()
			updated, err := setFloat32Slice(v, tc.vals)

			if tc.expectedError == "" {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tc.expectedError)
			}

			assert.Equal(t, tc.expectedUpdated, updated)
			assert.Equal(t, tc.expectedResult, tc.f)
		})
	}
}

func TestSetFloat64Slice(t *testing.T) {
	tests := []struct {
		name            string
		f               []float64
		vals            []string
		expectedUpdated bool
		expectedError   string
		expectedResult  []float64
	}{
		{
			"Nil",
			nil, []string{"2.7182818284"},
			true, "",
			[]float64{2.7182818284},
		},
		{
			"NewValue",
			[]float64{3.14159265359}, []string{"2.7182818284"},
			true, "",
			[]float64{2.7182818284},
		},
		{
			"NoNewValue",
			[]float64{2.7182818284}, []string{"2.7182818284"},
			false, "",
			[]float64{2.7182818284},
		},
		{
			"InvalidValue",
			[]float64{3.14159265359}, []string{"invalid"},
			false, `strconv.ParseFloat: parsing "invalid": invalid syntax`,
			[]float64{3.14159265359},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			v := reflect.ValueOf(&tc.f).Elem()
			updated, err := setFloat64Slice(v, tc.vals)

			if tc.expectedError == "" {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tc.expectedError)
			}

			assert.Equal(t, tc.expectedUpdated, updated)
			assert.Equal(t, tc.expectedResult, tc.f)
		})
	}
}

func TestSetStructSliceURL(t *testing.T) {
	url1, _ := url.Parse("service-1")
	url2, _ := url.Parse("service-2")

	tests := []struct {
		name            string
		s               []url.URL
		vals            []string
		expectedUpdated bool
		expectedError   string
		expectedResult  []url.URL
	}{
		{
			"URLNil",
			nil, []string{"service-2"},
			true, "",
			[]url.URL{*url2},
		},
		{
			"URLNewValue",
			[]url.URL{*url1}, []string{"service-2"},
			true, "",
			[]url.URL{*url2},
		},
		{
			"URLNoNewValue",
			[]url.URL{*url2}, []string{"service-2"},
			false, "",
			[]url.URL{*url2},
		},
		{
			"URLInvalidValue",
			[]url.URL{*url1}, []string{":"},
			false, `parse ":": missing protocol scheme`,
			[]url.URL{*url1},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			v := reflect.ValueOf(&tc.s).Elem()
			updated, err := setStructSlice(v, tc.vals)

			if tc.expectedError == "" {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tc.expectedError)
			}

			assert.Equal(t, tc.expectedUpdated, updated)
			assert.Equal(t, tc.expectedResult, tc.s)
		})
	}
}

func TestSetStructSliceRegexp(t *testing.T) {
	re1 := regexp.MustCompilePOSIX("[:digit:]")
	re2 := regexp.MustCompilePOSIX("[:alpha:]")

	tests := []struct {
		name            string
		s               []regexp.Regexp
		vals            []string
		expectedUpdated bool
		expectedError   string
		expectedResult  []regexp.Regexp
	}{
		{
			"RegexpNil",
			nil, []string{"[:alpha:]"},
			true, "",
			[]regexp.Regexp{*re2},
		},
		{
			"RegexpNewValue",
			[]regexp.Regexp{*re1}, []string{"[:alpha:]"},
			true, "",
			[]regexp.Regexp{*re2},
		},
		{
			"RegexpNoNewValue",
			[]regexp.Regexp{*re2}, []string{"[:alpha:]"},
			false, "",
			[]regexp.Regexp{*re2},
		},
		{
			"RegexpInvalidValue",
			[]regexp.Regexp{*re1}, []string{"[:invalid:"},
			false, "error parsing regexp: missing closing ]: `[:invalid:`",
			[]regexp.Regexp{*re1},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			v := reflect.ValueOf(&tc.s).Elem()
			updated, err := setStructSlice(v, tc.vals)

			if tc.expectedError == "" {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tc.expectedError)
			}

			assert.Equal(t, tc.expectedUpdated, updated)
			assert.Equal(t, tc.expectedResult, tc.s)
		})
	}
}

func TestSet(t *testing.T) {
	type Struct struct {
		String        string
		Bool          bool
		Int           int
		Int8          int8
		Int16         int16
		Int32         int32
		Int64         int64
		Uint          uint
		Uint8         uint8
		Uint16        uint16
		Uint32        uint32
		Uint64        uint64
		Float32       float32
		Float64       float64
		URL           url.URL
		Regexp        regexp.Regexp
		Byte          byte
		Rune          rune
		Duration      time.Duration
		StringPtr     *string
		BoolPtr       *bool
		IntPtr        *int
		Int8Ptr       *int8
		Int16Ptr      *int16
		Int32Ptr      *int32
		Int64Ptr      *int64
		UintPtr       *uint
		Uint8Ptr      *uint8
		Uint16Ptr     *uint16
		Uint32Ptr     *uint32
		Uint64Ptr     *uint64
		Float32Ptr    *float32
		Float64Ptr    *float64
		URLPtr        *url.URL
		RegexpPtr     *regexp.Regexp
		BytePtr       *byte
		RunePtr       *rune
		DurationPtr   *time.Duration
		StringSlice   []string
		BoolSlice     []bool
		IntSlice      []int
		Int8Slice     []int8
		Int16Slice    []int16
		Int32Slice    []int32
		Int64Slice    []int64
		UintSlice     []uint
		Uint8Slice    []uint8
		Uint16Slice   []uint16
		Uint32Slice   []uint32
		Uint64Slice   []uint64
		Float32Slice  []float32
		Float64Slice  []float64
		URLSlice      []url.URL
		RegexpSlice   []regexp.Regexp
		ByteSlice     []byte
		RuneSlice     []rune
		DurationSlice []time.Duration
	}

	url1, _ := url.Parse("service-1")
	url2, _ := url.Parse("service-2")

	re1 := regexp.MustCompilePOSIX("[:digit:]")
	re2 := regexp.MustCompilePOSIX("[:alpha:]")

	s1 := Struct{
		String:        "old",
		Bool:          false,
		Int:           -9223372036854775808,
		Int8:          -128,
		Int16:         -32768,
		Int32:         -2147483648,
		Int64:         -9223372036854775808,
		Uint:          0,
		Uint8:         0,
		Uint16:        0,
		Uint32:        0,
		Uint64:        0,
		Float32:       3.1415,
		Float64:       3.14159265359,
		URL:           *url1,
		Regexp:        *re1,
		Byte:          0,
		Rune:          -2147483648,
		Duration:      time.Second,
		StringPtr:     stringPtr("old"),
		BoolPtr:       boolPtr(false),
		IntPtr:        intPtr(-9223372036854775808),
		Int8Ptr:       int8Ptr(-128),
		Int16Ptr:      int16Ptr(-32768),
		Int32Ptr:      int32Ptr(-2147483648),
		Int64Ptr:      int64Ptr(-9223372036854775808),
		UintPtr:       uintPtr(0),
		Uint8Ptr:      uint8Ptr(0),
		Uint16Ptr:     uint16Ptr(0),
		Uint32Ptr:     uint32Ptr(0),
		Uint64Ptr:     uint64Ptr(0),
		Float32Ptr:    float32Ptr(3.1415),
		Float64Ptr:    float64Ptr(3.14159265359),
		URLPtr:        url1,
		RegexpPtr:     re1,
		BytePtr:       bytePtr(0),
		RunePtr:       runePtr(-2147483648),
		DurationPtr:   durationPtr(time.Second),
		StringSlice:   []string{"old"},
		BoolSlice:     []bool{false},
		IntSlice:      []int{-2147483648},
		Int8Slice:     []int8{-128},
		Int16Slice:    []int16{-32768},
		Int32Slice:    []int32{-2147483648},
		Int64Slice:    []int64{-9223372036854775808},
		UintSlice:     []uint{0},
		Uint8Slice:    []uint8{0},
		Uint16Slice:   []uint16{0},
		Uint32Slice:   []uint32{0},
		Uint64Slice:   []uint64{0},
		Float32Slice:  []float32{3.1415},
		Float64Slice:  []float64{3.14159265359},
		URLSlice:      []url.URL{*url1, *url2},
		RegexpSlice:   []regexp.Regexp{*re1, *re2},
		ByteSlice:     []byte{0},
		RuneSlice:     []rune{-2147483648},
		DurationSlice: []time.Duration{time.Second},
	}

	s2 := Struct{
		String:        "new",
		Bool:          true,
		Int:           9223372036854775807,
		Int8:          127,
		Int16:         32767,
		Int32:         2147483647,
		Int64:         9223372036854775807,
		Uint:          18446744073709551615,
		Uint8:         255,
		Uint16:        65535,
		Uint32:        4294967295,
		Uint64:        18446744073709551615,
		Float32:       2.7182,
		Float64:       2.7182818284,
		URL:           *url2,
		Regexp:        *re2,
		Byte:          255,
		Rune:          2147483647,
		Duration:      time.Minute,
		StringPtr:     stringPtr("new"),
		BoolPtr:       boolPtr(true),
		IntPtr:        intPtr(9223372036854775807),
		Int8Ptr:       int8Ptr(127),
		Int16Ptr:      int16Ptr(32767),
		Int32Ptr:      int32Ptr(2147483647),
		Int64Ptr:      int64Ptr(9223372036854775807),
		UintPtr:       uintPtr(18446744073709551615),
		Uint8Ptr:      uint8Ptr(255),
		Uint16Ptr:     uint16Ptr(65535),
		Uint32Ptr:     uint32Ptr(4294967295),
		Uint64Ptr:     uint64Ptr(18446744073709551615),
		Float32Ptr:    float32Ptr(2.7182),
		Float64Ptr:    float64Ptr(2.7182818284),
		URLPtr:        url2,
		RegexpPtr:     re2,
		BytePtr:       bytePtr(255),
		RunePtr:       runePtr(2147483647),
		DurationPtr:   durationPtr(time.Minute),
		StringSlice:   []string{"new"},
		BoolSlice:     []bool{true},
		IntSlice:      []int{9223372036854775807},
		Int8Slice:     []int8{127},
		Int16Slice:    []int16{32767},
		Int32Slice:    []int32{2147483647},
		Int64Slice:    []int64{9223372036854775807},
		UintSlice:     []uint{18446744073709551615},
		Uint8Slice:    []uint8{255},
		Uint16Slice:   []uint16{65535},
		Uint32Slice:   []uint32{4294967295},
		Uint64Slice:   []uint64{18446744073709551615},
		Float32Slice:  []float32{2.7182},
		Float64Slice:  []float64{2.7182818284},
		URLSlice:      []url.URL{*url2},
		RegexpSlice:   []regexp.Regexp{*re2},
		ByteSlice:     []byte{255},
		RuneSlice:     []rune{2147483647},
		DurationSlice: []time.Duration{time.Minute},
	}

	values := map[string]string{
		"String":        "new",
		"Bool":          "true",
		"Int":           "9223372036854775807",
		"Int8":          "127",
		"Int16":         "32767",
		"Int32":         "2147483647",
		"Int64":         "9223372036854775807",
		"Uint":          "18446744073709551615",
		"Uint8":         "255",
		"Uint16":        "65535",
		"Uint32":        "4294967295",
		"Uint64":        "18446744073709551615",
		"Float32":       "2.7182",
		"Float64":       "2.7182818284",
		"URL":           "service-2",
		"Regexp":        "[:alpha:]",
		"Byte":          "255",
		"Rune":          "2147483647",
		"Duration":      "1m",
		"StringPtr":     "new",
		"BoolPtr":       "true",
		"IntPtr":        "9223372036854775807",
		"Int8Ptr":       "127",
		"Int16Ptr":      "32767",
		"Int32Ptr":      "2147483647",
		"Int64Ptr":      "9223372036854775807",
		"UintPtr":       "18446744073709551615",
		"Uint8Ptr":      "255",
		"Uint16Ptr":     "65535",
		"Uint32Ptr":     "4294967295",
		"Uint64Ptr":     "18446744073709551615",
		"Float32Ptr":    "2.7182",
		"Float64Ptr":    "2.7182818284",
		"URLPtr":        "service-2",
		"RegexpPtr":     "[:alpha:]",
		"BytePtr":       "255",
		"RunePtr":       "2147483647",
		"DurationPtr":   "1m",
		"StringSlice":   "new",
		"BoolSlice":     "true",
		"IntSlice":      "9223372036854775807",
		"Int8Slice":     "127",
		"Int16Slice":    "32767",
		"Int32Slice":    "2147483647",
		"Int64Slice":    "9223372036854775807",
		"UintSlice":     "18446744073709551615",
		"Uint8Slice":    "255",
		"Uint16Slice":   "65535",
		"Uint32Slice":   "4294967295",
		"Uint64Slice":   "18446744073709551615",
		"Float32Slice":  "2.7182",
		"Float64Slice":  "2.7182818284",
		"URLSlice":      "service-2",
		"RegexpSlice":   "[:alpha:]",
		"ByteSlice":     "255",
		"RuneSlice":     "2147483647",
		"DurationSlice": "1m",
	}

	tests := []struct {
		name            string
		s               Struct
		values          map[string]string
		expectedUpdated bool
		expectedError   string
		expectedResult  Struct
	}{
		{
			"NewValues",
			s1,
			values,
			true, "",
			s2,
		},
		{
			"NoNewValues",
			s2,
			values,
			false, "",
			s2,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			vStruct := reflect.ValueOf(&tc.s).Elem()
			for i := 0; i < vStruct.NumField(); i++ {
				v := vStruct.Field(i)
				f := vStruct.Type().Field(i)

				updated, err := set(v, ",", tc.values[f.Name])

				if tc.expectedError == "" {
					assert.NoError(t, err)
				} else {
					assert.EqualError(t, err, tc.expectedError)
				}

				assert.Equal(t, tc.expectedUpdated, updated)
			}

			assert.Equal(t, tc.expectedResult, tc.s)
		})
	}
}

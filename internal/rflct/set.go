package rflct

import (
	"fmt"
	"net/url"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func setString(v reflect.Value, val string) (bool, error) {
	if v.String() == val {
		return false, nil
	}

	v.SetString(val)
	return true, nil
}

func setBool(v reflect.Value, val string) (bool, error) {
	b, err := strconv.ParseBool(val)
	if err != nil {
		return false, err
	}

	if v.Bool() == b {
		return false, nil
	}

	v.SetBool(b)
	return true, nil
}

func setInt(v reflect.Value, val string) (bool, error) {
	// int size and range are platform-dependent
	i, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		return false, err
	}

	if v.Int() == i {
		return false, nil
	}

	v.SetInt(i)
	return true, nil
}

func setInt8(v reflect.Value, val string) (bool, error) {
	i, err := strconv.ParseInt(val, 10, 8)
	if err != nil {
		return false, err
	}

	if v.Int() == i {
		return false, nil
	}

	v.SetInt(i)
	return true, nil
}

func setInt16(v reflect.Value, val string) (bool, error) {
	i, err := strconv.ParseInt(val, 10, 16)
	if err != nil {
		return false, err
	}

	if v.Int() == i {
		return false, nil
	}

	v.SetInt(i)
	return true, nil
}

func setInt32(v reflect.Value, val string) (bool, error) {
	i, err := strconv.ParseInt(val, 10, 32)
	if err != nil {
		return false, err
	}

	if v.Int() == i {
		return false, nil
	}

	v.SetInt(i)
	return true, nil
}

func setInt64(v reflect.Value, val string) (bool, error) {
	if t := v.Type(); t.PkgPath() == "time" && t.Name() == "Duration" {
		d, err := time.ParseDuration(val)
		if err != nil {
			return false, err
		}

		if v.Interface() == d {
			return false, nil
		}

		v.Set(reflect.ValueOf(d))
		return true, nil
	}

	i, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		return false, err
	}

	if v.Int() == i {
		return false, nil
	}

	v.SetInt(i)
	return true, nil
}

func setUint(v reflect.Value, val string) (bool, error) {
	// uint size and range are platform-dependent
	u, err := strconv.ParseUint(val, 10, 64)
	if err != nil {
		return false, err
	}

	if v.Uint() == u {
		return false, nil
	}

	v.SetUint(u)
	return true, nil
}

func setUint8(v reflect.Value, val string) (bool, error) {
	u, err := strconv.ParseUint(val, 10, 8)
	if err != nil {
		return false, err
	}

	if v.Uint() == u {
		return false, nil
	}

	v.SetUint(u)
	return true, nil
}

func setUint16(v reflect.Value, val string) (bool, error) {
	u, err := strconv.ParseUint(val, 10, 16)
	if err != nil {
		return false, err
	}

	if v.Uint() == u {
		return false, nil
	}

	v.SetUint(u)
	return true, nil
}

func setUint32(v reflect.Value, val string) (bool, error) {
	u, err := strconv.ParseUint(val, 10, 32)
	if err != nil {
		return false, err
	}

	if v.Uint() == u {
		return false, nil
	}

	v.SetUint(u)
	return true, nil
}

func setUint64(v reflect.Value, val string) (bool, error) {
	u, err := strconv.ParseUint(val, 10, 64)
	if err != nil {
		return false, err
	}

	if v.Uint() == u {
		return false, nil
	}

	v.SetUint(u)
	return true, nil
}

func setFloat32(v reflect.Value, val string) (bool, error) {
	f, err := strconv.ParseFloat(val, 32)
	if err != nil {
		return false, err
	}

	if v.Float() == f {
		return false, nil
	}

	v.SetFloat(f)
	return true, nil
}

func setFloat64(v reflect.Value, val string) (bool, error) {
	f, err := strconv.ParseFloat(val, 64)
	if err != nil {
		return false, err
	}

	if v.Float() == f {
		return false, nil
	}

	v.SetFloat(f)
	return true, nil
}

func setStruct(v reflect.Value, val string) (bool, error) {
	t := v.Type()

	if t.PkgPath() == "net/url" && t.Name() == "URL" {
		u, err := url.Parse(val)
		if err != nil {
			return false, err
		}

		// u is a pointer
		if reflect.DeepEqual(v.Interface(), *u) {
			return false, nil
		}

		// u is a pointer
		v.Set(reflect.ValueOf(u).Elem())
		return true, nil
	} else if t.PkgPath() == "regexp" && t.Name() == "Regexp" {
		r, err := regexp.CompilePOSIX(val)
		if err != nil {
			return false, err
		}

		// r is a pointer
		if reflect.DeepEqual(v.Interface(), *r) {
			return false, nil
		}

		// r is a pointer
		v.Set(reflect.ValueOf(r).Elem())
		return true, nil
	}

	return false, fmt.Errorf("unsupported type: %s.%s", t.PkgPath(), t.Name())
}

func setStringPtr(v reflect.Value, val string) (bool, error) {
	if !v.IsZero() && v.Elem().String() == val {
		return false, nil
	}

	v.Set(reflect.ValueOf(&val))
	return true, nil
}

func setBoolPtr(v reflect.Value, val string) (bool, error) {
	b, err := strconv.ParseBool(val)
	if err != nil {
		return false, err
	}

	if !v.IsZero() && v.Elem().Bool() == b {
		return false, nil
	}

	v.Set(reflect.ValueOf(&b))
	return true, nil
}

func setIntPtr(v reflect.Value, val string) (bool, error) {
	// int size and range are platform-dependent
	i64, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		return false, err
	}

	if !v.IsZero() && v.Elem().Int() == i64 {
		return false, nil
	}

	i := int(i64)
	v.Set(reflect.ValueOf(&i))
	return true, nil
}

func setInt8Ptr(v reflect.Value, val string) (bool, error) {
	i64, err := strconv.ParseInt(val, 10, 8)
	if err != nil {
		return false, err
	}

	if !v.IsZero() && v.Elem().Int() == i64 {
		return false, nil
	}

	i8 := int8(i64)
	v.Set(reflect.ValueOf(&i8))
	return true, nil
}

func setInt16Ptr(v reflect.Value, val string) (bool, error) {
	i64, err := strconv.ParseInt(val, 10, 16)
	if err != nil {
		return false, err
	}

	if !v.IsZero() && v.Elem().Int() == i64 {
		return false, nil
	}

	i16 := int16(i64)
	v.Set(reflect.ValueOf(&i16))
	return true, nil
}

func setInt32Ptr(v reflect.Value, val string) (bool, error) {
	i64, err := strconv.ParseInt(val, 10, 32)
	if err != nil {
		return false, err
	}

	if !v.IsZero() && v.Elem().Int() == i64 {
		return false, nil
	}

	i32 := int32(i64)
	v.Set(reflect.ValueOf(&i32))
	return true, nil
}

func setInt64Ptr(v reflect.Value, val string) (bool, error) {
	t := reflect.TypeOf(v.Interface()).Elem()

	if t.PkgPath() == "time" && t.Name() == "Duration" {
		d, err := time.ParseDuration(val)
		if err != nil {
			return false, err
		}

		if !v.IsZero() && v.Elem().Interface() == d {
			return false, nil
		}

		v.Set(reflect.ValueOf(&d))
		return true, nil
	}

	i64, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		return false, err
	}

	if !v.IsZero() && v.Elem().Int() == i64 {
		return false, nil
	}

	v.Set(reflect.ValueOf(&i64))
	return true, nil
}

func setUintPtr(v reflect.Value, val string) (bool, error) {
	// uint size and range are platform-dependent
	u64, err := strconv.ParseUint(val, 10, 64)
	if err != nil {
		return false, err
	}

	if !v.IsZero() && v.Elem().Uint() == u64 {
		return false, nil
	}

	u := uint(u64)
	v.Set(reflect.ValueOf(&u))
	return true, nil
}

func setUint8Ptr(v reflect.Value, val string) (bool, error) {
	u64, err := strconv.ParseUint(val, 10, 8)
	if err != nil {
		return false, err
	}

	if !v.IsZero() && v.Elem().Uint() == u64 {
		return false, nil
	}

	u8 := uint8(u64)
	v.Set(reflect.ValueOf(&u8))
	return true, nil
}

func setUint16Ptr(v reflect.Value, val string) (bool, error) {
	u64, err := strconv.ParseUint(val, 10, 16)
	if err != nil {
		return false, err
	}

	if !v.IsZero() && v.Elem().Uint() == u64 {
		return false, nil
	}

	u16 := uint16(u64)
	v.Set(reflect.ValueOf(&u16))
	return true, nil
}

func setUint32Ptr(v reflect.Value, val string) (bool, error) {
	u64, err := strconv.ParseUint(val, 10, 32)
	if err != nil {
		return false, err
	}

	if !v.IsZero() && v.Elem().Uint() == u64 {
		return false, nil
	}

	u32 := uint32(u64)
	v.Set(reflect.ValueOf(&u32))
	return true, nil
}

func setUint64Ptr(v reflect.Value, val string) (bool, error) {
	u64, err := strconv.ParseUint(val, 10, 64)
	if err != nil {
		return false, err
	}

	if !v.IsZero() && v.Elem().Uint() == u64 {
		return false, nil
	}

	v.Set(reflect.ValueOf(&u64))
	return true, nil
}

func setFloat32Ptr(v reflect.Value, val string) (bool, error) {
	f64, err := strconv.ParseFloat(val, 32)
	if err != nil {
		return false, err
	}

	if !v.IsZero() && v.Elem().Float() == f64 {
		return false, nil
	}

	f32 := float32(f64)
	v.Set(reflect.ValueOf(&f32))
	return true, nil
}

func setFloat64Ptr(v reflect.Value, val string) (bool, error) {
	f64, err := strconv.ParseFloat(val, 64)
	if err != nil {
		return false, err
	}

	if !v.IsZero() && v.Elem().Float() == f64 {
		return false, nil
	}

	v.Set(reflect.ValueOf(&f64))
	return true, nil
}

func setStructPtr(v reflect.Value, val string) (bool, error) {
	t := reflect.TypeOf(v.Interface()).Elem()

	if t.PkgPath() == "net/url" && t.Name() == "URL" {
		u, err := url.Parse(val)
		if err != nil {
			return false, err
		}

		if !v.IsZero() && reflect.DeepEqual(v.Elem().Interface(), *u) {
			return false, nil
		}

		// u is a pointer
		v.Set(reflect.ValueOf(u))
		return true, nil
	} else if t.PkgPath() == "regexp" && t.Name() == "Regexp" {
		r, err := regexp.CompilePOSIX(val)
		if err != nil {
			return false, err
		}

		if !v.IsZero() && reflect.DeepEqual(v.Elem().Interface(), *r) {
			return false, nil
		}

		// r is a pointer
		v.Set(reflect.ValueOf(r))
		return true, nil
	}

	return false, fmt.Errorf("unsupported type: %s.%s", t.PkgPath(), t.Name())
}

func setStringSlice(v reflect.Value, vals []string) (bool, error) {
	if reflect.DeepEqual(v.Interface(), vals) {
		return false, nil
	}

	v.Set(reflect.ValueOf(vals))
	return true, nil
}

func setBoolSlice(v reflect.Value, vals []string) (bool, error) {
	bools := []bool{}
	for _, val := range vals {
		b, err := strconv.ParseBool(val)
		if err != nil {
			return false, err
		}

		bools = append(bools, b)
	}

	if reflect.DeepEqual(v.Interface(), bools) {
		return false, nil
	}

	v.Set(reflect.ValueOf(bools))
	return true, nil
}

func setIntSlice(v reflect.Value, vals []string) (bool, error) {
	// int size and range are platform-dependent
	ints := []int{}
	for _, val := range vals {
		i, err := strconv.ParseInt(val, 10, 64)
		if err != nil {
			return false, err
		}

		ints = append(ints, int(i))
	}

	if reflect.DeepEqual(v.Interface(), ints) {
		return false, nil
	}

	v.Set(reflect.ValueOf(ints))
	return true, nil
}

func setInt8Slice(v reflect.Value, vals []string) (bool, error) {
	ints := []int8{}
	for _, val := range vals {
		i, err := strconv.ParseInt(val, 10, 8)
		if err != nil {
			return false, err
		}

		ints = append(ints, int8(i))
	}

	if reflect.DeepEqual(v.Interface(), ints) {
		return false, nil
	}

	v.Set(reflect.ValueOf(ints))
	return true, nil
}

func setInt16Slice(v reflect.Value, vals []string) (bool, error) {
	ints := []int16{}
	for _, val := range vals {
		i, err := strconv.ParseInt(val, 10, 16)
		if err != nil {
			return false, err
		}

		ints = append(ints, int16(i))
	}

	if reflect.DeepEqual(v.Interface(), ints) {
		return false, nil
	}

	v.Set(reflect.ValueOf(ints))
	return true, nil
}

func setInt32Slice(v reflect.Value, vals []string) (bool, error) {
	ints := []int32{}
	for _, val := range vals {
		i, err := strconv.ParseInt(val, 10, 32)
		if err != nil {
			return false, err
		}

		ints = append(ints, int32(i))
	}

	if reflect.DeepEqual(v.Interface(), ints) {
		return false, nil
	}

	v.Set(reflect.ValueOf(ints))
	return true, nil
}

func setInt64Slice(v reflect.Value, vals []string) (bool, error) {
	t := reflect.TypeOf(v.Interface()).Elem()

	if t.PkgPath() == "time" && t.Name() == "Duration" {
		durations := []time.Duration{}
		for _, val := range vals {
			d, err := time.ParseDuration(val)
			if err != nil {
				return false, err
			}

			durations = append(durations, d)
		}

		if reflect.DeepEqual(v.Interface(), durations) {
			return false, nil
		}

		v.Set(reflect.ValueOf(durations))
		return true, nil
	}

	ints := []int64{}
	for _, val := range vals {
		i, err := strconv.ParseInt(val, 10, 64)
		if err != nil {
			return false, err
		}

		ints = append(ints, i)
	}

	if reflect.DeepEqual(v.Interface(), ints) {
		return false, nil
	}

	v.Set(reflect.ValueOf(ints))
	return true, nil
}

func setUintSlice(v reflect.Value, vals []string) (bool, error) {
	// uint size and range are platform-dependent
	uints := []uint{}
	for _, val := range vals {
		u, err := strconv.ParseUint(val, 10, 64)
		if err != nil {
			return false, err
		}

		uints = append(uints, uint(u))
	}

	if reflect.DeepEqual(v.Interface(), uints) {
		return false, nil
	}

	v.Set(reflect.ValueOf(uints))
	return true, nil
}

func setUint8Slice(v reflect.Value, vals []string) (bool, error) {
	uints := []uint8{}
	for _, val := range vals {
		u, err := strconv.ParseUint(val, 10, 8)
		if err != nil {
			return false, err
		}

		uints = append(uints, uint8(u))
	}

	if reflect.DeepEqual(v.Interface(), uints) {
		return false, nil
	}

	v.Set(reflect.ValueOf(uints))
	return true, nil
}

func setUint16Slice(v reflect.Value, vals []string) (bool, error) {
	uints := []uint16{}
	for _, val := range vals {
		u, err := strconv.ParseUint(val, 10, 16)
		if err != nil {
			return false, err
		}

		uints = append(uints, uint16(u))
	}

	if reflect.DeepEqual(v.Interface(), uints) {
		return false, nil
	}

	v.Set(reflect.ValueOf(uints))
	return true, nil
}

func setUint32Slice(v reflect.Value, vals []string) (bool, error) {
	uints := []uint32{}
	for _, val := range vals {
		u, err := strconv.ParseUint(val, 10, 32)
		if err != nil {
			return false, err
		}

		uints = append(uints, uint32(u))
	}

	if reflect.DeepEqual(v.Interface(), uints) {
		return false, nil
	}

	v.Set(reflect.ValueOf(uints))
	return true, nil
}

func setUint64Slice(v reflect.Value, vals []string) (bool, error) {
	uints := []uint64{}
	for _, val := range vals {
		u, err := strconv.ParseUint(val, 10, 64)
		if err != nil {
			return false, err
		}

		uints = append(uints, u)
	}

	if reflect.DeepEqual(v.Interface(), uints) {
		return false, nil
	}

	v.Set(reflect.ValueOf(uints))
	return true, nil
}

func setFloat32Slice(v reflect.Value, vals []string) (bool, error) {
	floats := []float32{}
	for _, val := range vals {
		f, err := strconv.ParseFloat(val, 32)
		if err != nil {
			return false, err
		}

		floats = append(floats, float32(f))
	}

	if reflect.DeepEqual(v.Interface(), floats) {
		return false, nil
	}

	v.Set(reflect.ValueOf(floats))
	return true, nil
}

func setFloat64Slice(v reflect.Value, vals []string) (bool, error) {
	floats := []float64{}
	for _, val := range vals {
		f, err := strconv.ParseFloat(val, 64)
		if err != nil {
			return false, err
		}

		floats = append(floats, f)
	}

	if reflect.DeepEqual(v.Interface(), floats) {
		return false, nil
	}

	v.Set(reflect.ValueOf(floats))
	return true, nil
}

func setStructSlice(v reflect.Value, vals []string) (bool, error) {
	t := reflect.TypeOf(v.Interface()).Elem()

	if t.PkgPath() == "net/url" && t.Name() == "URL" {
		urls := []url.URL{}
		for _, val := range vals {
			u, err := url.Parse(val)
			if err != nil {
				return false, err
			}

			urls = append(urls, *u)
		}

		// []url.URL
		if reflect.DeepEqual(v.Interface(), urls) {
			return false, nil
		}

		v.Set(reflect.ValueOf(urls))
		return true, nil
	} else if t.PkgPath() == "regexp" && t.Name() == "Regexp" {
		regexps := []regexp.Regexp{}
		for _, val := range vals {
			r, err := regexp.CompilePOSIX(val)
			if err != nil {
				return false, err
			}

			regexps = append(regexps, *r)
		}

		// []regexp.Regexp
		if reflect.DeepEqual(v.Interface(), regexps) {
			return false, nil
		}

		v.Set(reflect.ValueOf(regexps))
		return true, nil
	}

	return false, fmt.Errorf("unsupported type: %s.%s", t.PkgPath(), t.Name())
}

func SetValue(v reflect.Value, sep, val string) (bool, error) {
	switch v.Kind() {
	case reflect.String:
		return setString(v, val)
	case reflect.Bool:
		return setBool(v, val)
	case reflect.Int:
		return setInt(v, val)
	case reflect.Int8:
		return setInt8(v, val)
	case reflect.Int16:
		return setInt16(v, val)
	case reflect.Int32:
		return setInt32(v, val)
	case reflect.Int64:
		return setInt64(v, val)
	case reflect.Uint:
		return setUint(v, val)
	case reflect.Uint8:
		return setUint8(v, val)
	case reflect.Uint16:
		return setUint16(v, val)
	case reflect.Uint32:
		return setUint32(v, val)
	case reflect.Uint64:
		return setUint64(v, val)
	case reflect.Float32:
		return setFloat32(v, val)
	case reflect.Float64:
		return setFloat64(v, val)
	case reflect.Struct:
		return setStruct(v, val)

	case reflect.Ptr:
		tPtr := reflect.TypeOf(v.Interface()).Elem()

		switch tPtr.Kind() {
		case reflect.String:
			return setStringPtr(v, val)
		case reflect.Bool:
			return setBoolPtr(v, val)
		case reflect.Int:
			return setIntPtr(v, val)
		case reflect.Int8:
			return setInt8Ptr(v, val)
		case reflect.Int16:
			return setInt16Ptr(v, val)
		case reflect.Int32:
			return setInt32Ptr(v, val)
		case reflect.Int64:
			return setInt64Ptr(v, val)
		case reflect.Uint:
			return setUintPtr(v, val)
		case reflect.Uint8:
			return setUint8Ptr(v, val)
		case reflect.Uint16:
			return setUint16Ptr(v, val)
		case reflect.Uint32:
			return setUint32Ptr(v, val)
		case reflect.Uint64:
			return setUint64Ptr(v, val)
		case reflect.Float32:
			return setFloat32Ptr(v, val)
		case reflect.Float64:
			return setFloat64Ptr(v, val)
		case reflect.Struct:
			return setStructPtr(v, val)
		}

	case reflect.Slice:
		tSlice := reflect.TypeOf(v.Interface()).Elem()
		vals := strings.Split(val, sep)

		switch tSlice.Kind() {
		case reflect.String:
			return setStringSlice(v, vals)
		case reflect.Bool:
			return setBoolSlice(v, vals)
		case reflect.Int:
			return setIntSlice(v, vals)
		case reflect.Int8:
			return setInt8Slice(v, vals)
		case reflect.Int16:
			return setInt16Slice(v, vals)
		case reflect.Int32:
			return setInt32Slice(v, vals)
		case reflect.Int64:
			return setInt64Slice(v, vals)
		case reflect.Uint:
			return setUintSlice(v, vals)
		case reflect.Uint8:
			return setUint8Slice(v, vals)
		case reflect.Uint16:
			return setUint16Slice(v, vals)
		case reflect.Uint32:
			return setUint32Slice(v, vals)
		case reflect.Uint64:
			return setUint64Slice(v, vals)
		case reflect.Float32:
			return setFloat32Slice(v, vals)
		case reflect.Float64:
			return setFloat64Slice(v, vals)
		case reflect.Struct:
			return setStructSlice(v, vals)
		}
	}

	return false, fmt.Errorf("unsupported type: %s", v.Kind())
}

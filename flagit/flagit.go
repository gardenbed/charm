// Package flagit adds support for a new struct tag: flag.
// You can tag your struct fields with the flag tag and parse command-line arguments into your struct fields.
// Nested structs are also supported. You can either parse the command-line arguments using this package or the built-in flag package.
package flagit

import (
	"flag"
	"fmt"
	"os"
	"reflect"
	"regexp"
	"strings"

	"github.com/gardenbed/charm/internal/rflct"
)

const (
	flagTag = "flag"
	sepTag  = "sep"
)

var (
	flagNameRE = regexp.MustCompile(`^[A-Za-z]([0-9A-Za-z-.]*[0-9A-Za-z])?$`)
	flagArgRE  = regexp.MustCompile("^-{1,2}[A-Za-z]([0-9A-Za-z-.]*[0-9A-Za-z])?")
)

// flagValue implements the flag.Value interface.
type flagValue struct {
	continueOnError bool
	value           reflect.Value
	sep             string
}

// String is called for getting and printing the default value.
// Default value is already included in the usage string.
func (v flagValue) String() string {
	return ""
}

func (v flagValue) Set(val string) error {
	if _, err := rflct.SetValue(v.value, v.sep, val); err != nil {
		if v.continueOnError {
			return nil
		}
		return err
	}

	return nil
}

// Register accepts a flag set and the pointer to a struct type.
// For those struct fields that have the flag tag, it will register a flag on the given flag set.
// The current values of the struct fields will be used as default values for the registered flags.
// Once the Parse method on the flag set is called, the values will be read, parsed to the appropriate types, and assigned to the corresponding struct fields.
func Register(fs *flag.FlagSet, s interface{}, continueOnError bool) error {
	v, err := rflct.IsStructPtr(s)
	if err != nil {
		return err
	}

	return iterateOnFields("", v, continueOnError, func(f fieldInfo) error {
		if fs.Lookup(f.flag) != nil {
			if continueOnError {
				return nil
			}
			return fmt.Errorf("flag already registered: %s", f.flag)
		}

		// Create usage string
		var usage string

		if f.help != "" {
			usage = f.help + "\n"
		}

		switch f.value.Kind() {
		case reflect.Slice:
			usage += fmt.Sprintf("%-15s []%s\n%-15s %v\n%-15s %s",
				"data type:", reflect.TypeOf(f.value.Interface()).Elem(),
				"default value:", f.value.Interface(),
				"separator:", f.sep,
			)
		case reflect.Struct:
			usage += fmt.Sprintf("%-15s %s\n%-15s %+v",
				"data type:", f.value.Type(),
				"default value:", f.value.Interface(),
			)
		default:
			usage += fmt.Sprintf("%-15s %s\n%-15s %v",
				"data type:", f.value.Type(),
				"default value:", f.value.Interface(),
			)
		}

		// Register the flag
		switch f.value.Kind() {
		case reflect.Bool:
			// f.value.CanAddr() expected to be true
			// f.value.Addr().Interface().(*bool) expected to be ok
			ptr := f.value.Addr().Interface().(*bool)
			fs.BoolVar(ptr, f.flag, f.value.Bool(), usage)
		default:
			fv := &flagValue{continueOnError, f.value, f.sep}
			fs.Var(fv, f.flag, usage)
		}

		return nil
	})
}

// Parse accepts the pointer to a struct type.
// For those struct fields that have the flag tag, it will read values from command-line flags and parse them to the appropriate types.
// This method does not use the built-in flag package for parsing and reading the flags.
func Parse(s interface{}, continueOnError bool) error {
	v, err := rflct.IsStructPtr(s)
	if err != nil {
		return err
	}

	return iterateOnFields("", v, continueOnError, func(f fieldInfo) error {
		if val := getFlagValue(f.flag); val != "" {
			if _, err := rflct.SetValue(f.value, f.sep, val); err != nil {
				if continueOnError {
					return nil
				}
				return err
			}
		}

		return nil
	})
}

type fieldInfo struct {
	value reflect.Value
	name  string
	flag  string
	help  string
	sep   string
}

func iterateOnFields(prefix string, vStruct reflect.Value, continueOnError bool, handle func(fieldInfo) error) error {
	// Iterate over struct fields
	for i := 0; i < vStruct.NumField(); i++ {
		v := vStruct.Field(i)        // reflect.Value       --> vField.Kind(), vField.Type().Name(), vField.Type().Kind(), vField.Interface()
		t := v.Type()                // reflect.Type        --> t.Kind(), t.PkgPath(), t.Name(), t.NumField()
		f := vStruct.Type().Field(i) // reflect.StructField --> f.Name, f.Type.Name(), f.Type.Kind(), f.Tag.Get(tag)

		// Recursively, iterate on nested structs
		// Nested structs do not need to have the `flag` tag and can be not settable.
		if rflct.IsNestedStruct(t) {
			newPrefix := prefix + f.Tag.Get(flagTag)
			if err := iterateOnFields(newPrefix, v, continueOnError, handle); err != nil {
				return err
			}
			continue
		}

		// Skip unexported and unsupported fields
		if !v.CanSet() || !rflct.IsTypeSupported(t) {
			continue
		}

		// `flag:"..."`
		val := f.Tag.Get(flagTag)
		if val == "" {
			continue
		}

		var flagName, flagHelp string
		if strings.Contains(val, ",") {
			subs := strings.Split(val, ",")
			flagName, flagHelp = subs[0], subs[1]
		} else {
			flagName = val
		}

		// Apply prefix
		flagName = prefix + flagName

		// Sanitize the flag name
		if !flagNameRE.MatchString(flagName) {
			if continueOnError {
				continue
			}
			return fmt.Errorf("invalid flag name: %s", flagName)
		}

		// `sep:"..."`
		sep := f.Tag.Get(sepTag)
		if sep == "" {
			sep = ","
		}

		fi := fieldInfo{
			value: v,
			name:  f.Name,
			flag:  flagName,
			help:  flagHelp,
			sep:   sep,
		}

		if err := handle(fi); err != nil {
			return err
		}
	}

	return nil
}

func getFlagValue(flag string) string {
	flagRegex := regexp.MustCompile("-{1,2}" + flag)

	for i, arg := range os.Args {
		if flagRegex.MatchString(arg) {
			if s := strings.Index(arg, "="); s > 0 {
				return arg[s+1:]
			}

			if i+1 < len(os.Args) {
				if val := os.Args[i+1]; !flagArgRE.MatchString(val) {
					return val
				}
			}

			// For boolean flags
			return "true"
		}
	}

	return ""
}

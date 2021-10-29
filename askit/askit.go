package askit

import (
	"fmt"
	"net/mail"
	"reflect"
	"strings"

	"github.com/gardenbed/charm/internal/rflct"
)

const (
	askTag = "ask"
	sepTag = "sep"
)

// Kind determines the kind of an input.
type Kind string

const (
	// KindAny denotes any kind of input.
	KindAny Kind = "any"
	// KindEmail denotes an email address input.
	KindEmail Kind = "email"
	// KindSecret denotes a secret input (password, token, etc.).
	KindSecret Kind = "secret"
)

// Asker is the interface for getting inputs.
type Asker interface {
	Output(string)
	Ask(string) (string, error)
	AskSecret(string) (string, error)
}

func Ask(s interface{}, asker Asker) error {
	v, err := rflct.IsStructPtr(s)
	if err != nil {
		return err
	}

	return iterateOnFields("", v, func(f fieldInfo) error {
		return askForField(f, asker)
	})
}

type fieldInfo struct {
	Value       reflect.Value
	Name        string
	Kind        Kind
	Description string
	Sep         string
}

func iterateOnFields(prefix string, vStruct reflect.Value, handle func(fieldInfo) error) error {
	// Iterate over struct fields
	for i := 0; i < vStruct.NumField(); i++ {
		v := vStruct.Field(i)        // reflect.Value       --> vField.Kind(), vField.Type().Name(), vField.Type().Kind(), vField.Interface()
		t := v.Type()                // reflect.Type        --> t.Kind(), t.PkgPath(), t.Name(), t.NumField()
		f := vStruct.Type().Field(i) // reflect.StructField --> f.Name, f.Type.Name(), f.Type.Kind(), f.Tag.Get(tag)

		// Recursively, iterate on nested structs
		if rflct.IsNestedStruct(t) {
			newPrefix := prefix + f.Name
			if err := iterateOnFields(newPrefix, v, handle); err != nil {
				return err
			}
			continue
		}

		// Skip unexported and unsupported fields
		if !v.CanSet() || !rflct.IsTypeSupported(t) {
			continue
		}

		// `ask:"..."`
		val := f.Tag.Get(askTag)
		if val == "" {
			continue
		}
		subs := strings.Split(val, ",")

		// `sep:"..."`
		sep := f.Tag.Get(sepTag)
		if sep == "" {
			sep = ","
		}

		name := f.Name
		if prefix != "" {
			name = prefix + "." + name
		}

		fi := fieldInfo{
			Value: v,
			Name:  name,
			Sep:   sep,
		}

		if isKindSupported(subs[0]) {
			fi.Kind = Kind(subs[0])
		}

		if len(subs) > 1 {
			fi.Description = strings.TrimSpace(subs[1])
		}

		if err := handle(fi); err != nil {
			return err
		}
	}

	return nil
}

func isKindSupported(kind string) bool {
	switch Kind(kind) {
	case KindAny, KindEmail, KindSecret:
		return true
	default:
		return false
	}
}

func askForField(f fieldInfo, asker Asker) error {
	// Print in bold style
	asker.Output(fmt.Sprintf("\033[1m%s\033[0m", f.Name))

	if !f.Value.IsZero() {
		if f.Kind == KindSecret {
			asker.Output("  • Current value: *******")
		} else {
			asker.Output(fmt.Sprintf("  • Current value: %v", f.Value.Interface()))
		}
	}

	ans, err := asker.Ask("  • Would you like to enter a value [Y]?")
	if err != nil {
		return err
	}

	if ans == "" || strings.ToUpper(ans[:1]) != "Y" {
		return nil
	}

	// Create the user prompt
	var prompt string
	if f.Description == "" {
		prompt = "  Enter a new value:"
	} else {
		prompt = fmt.Sprintf("  • Enter a new value (%s):", f.Description)
	}

	// Determine which ask function to use
	askFunc := asker.Ask
	if f.Kind == KindSecret {
		askFunc = asker.AskSecret
	}

	val, err := askFunc(prompt)
	if err != nil {
		return err
	}

	if _, err := rflct.SetValue(f.Value, f.Sep, val); err != nil {
		return fmt.Errorf("invalid value entered for %s: %s", f.Name, err)
	}

	switch f.Kind {
	case KindEmail:
		if _, err := mail.ParseAddress(val); err != nil {
			return fmt.Errorf("invalid email address entered for %s: %s", f.Name, err)
		}
	}

	return nil
}

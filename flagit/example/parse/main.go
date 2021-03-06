package main

import (
	"fmt"
	"net/url"
	"time"

	"github.com/gardenbed/charm/flagit"
)

// Spec is a struct for mapping its fields to command-line flags.
type Spec struct {
	// Flag fields
	Verbose bool `flag:"verbose"`

	// Nested fields
	Options struct {
		Port     uint16 `flag:"port"`
		LogLevel string `flag:"log-level"`
	}

	// Nested fields with prefix
	Config struct {
		Timeout   time.Duration `flag:"timeout"`
		Endpoints []url.URL     `flag:"endpoints"`
	} `flag:"config-"`
}

func main() {
	spec := new(Spec)

	if err := flagit.Parse(spec, false); err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", spec)
}

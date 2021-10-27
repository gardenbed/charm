# flagit

This package allows you to use the `flag` struct tag on your Go struct fields.
You can then read the values for those fields from the command-line arguments.

## Quick Start

```go
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
```

## Examples

You can find more examples [here](./example).

## Documentation

### Supported Types

  - `string`, `*string`, `[]string`
  - `bool`, `*bool`, `[]bool`
  - `int`, `int8`, `int16`, `int32`, `int64`
  - `*int`, `*int8`, `*int16`, `*int32`, `*int64`
  - `[]int`, `[]int8`, `[]int16`, `[]int32`, `[]int64`
  - `uint`, `uint8`, `uint16`, `uint32`, `uint64`
  - `*uint`, `*uint8`, `*uint16`, `*uint32`, `*uint64`
  - `[]uint`, `[]uint8`, `[]uint16`, `[]uint32`, `[]uint64`
  - `float32`, `float64`
  - `*float32`, `*float64`
  - `[]float32`, `[]float64`
  - `url.URL`, `*url.URL`, `[]url.URL`
  - `regexp.Regexp`, `*regexp.Regexp`, `[]regexp.Regexp`
  - `byte`, `*byte`, `[]byte`
  - `rune`, `*rune`, `[]rune`
  - `time.Duration`, `*time.Duration`, `[]time.Duration`

The supported syntax for Regexp is [POSIX Regular Expressions](https://en.wikibooks.org/wiki/Regular_Expressions/POSIX_Basic_Regular_Expressions).
Nested structs are also supported.

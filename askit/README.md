# askit

This package allows you to use the `ask` struct tag on your Go struct fields.
You can then read the values for those fields using an `Asker` interface.

An `Asker` implementation can prompt user and read values from the standard input or any other source
([cli](github.com/mitchellh/cli) provides an implementation for the `Asker` interface).

## Quick Start

```go
package main

import (
  "fmt"
  "os"

  "github.com/gardenbed/charm/askit"
  "github.com/mitchellh/cli"
)

func main() {
  asker := &cli.BasicUi{
    Reader:      os.Stdin,
    Writer:      os.Stdout,
    ErrorWriter: os.Stderr,
  }

  info := struct {
    Name  string `ask:"any, your full name"`
    Email string `ask:"email, your email address"`
    Token string `ask:"secret, your access token"`
  }{
    Name: "Jane Doe",
  }

  err := askit.Ask(&info, asker)
  if err != nil {
    panic(err)
  }

  fmt.Printf("%+v\n", info)
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

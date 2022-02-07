# UI

This package provides a simple interface with implementations for printing information in command-line applications.
It supports various **verbosity** levels as well as formatting **styles**.

## Quick Start

```go
package main

import "github.com/gardenbed/charm/ui"

func main() {
  u := ui.New(ui.Info)
  u.Infof(ui.Green, "Hello, %s!", "World")
}
```

# shell

This package allows you to run shell commands easily and concisely.
You can specify arguments, environment variables, and working directry.
You can also get the exit code, stdout, and any error.

The biggest advantage of using this package over the built-in `os/exec` is **testability** and **brevity**.

## Quick Start

**Running a command:**

```go
package main

import (
  "context"
  "fmt"
  "github.com/gardenbed/charm/shell"
)

func main() {
  exitcode, stdout, err := shell.Run(context.Background(), "date", "-u")
  if err != nil {
    panic(err)
  }
  fmt.Printf("[%d] %s\n", exitcode, stdout)
}
```

**Running a command with environment variables:**

```go
package main

import (
  "context"
  "fmt"
  "github.com/gardenbed/charm/shell"
)

func main() {
  opts := shell.RunOptions{
    Environment: map[string]string{
      "GREETING": "Hello, World!",
    },
  }
  exitcode, stdout, err := shell.RunWith(context.Background(), opts, "printenv", "GREETING")
  if err != nil {
    panic(err)
  }
  fmt.Printf("[%d] %s\n", exitcode, stdout)
}
```

**Building a command and unit testing:**

```go
package main

import (
  "context"
  "fmt"

  "github.com/gardenbed/charm/shell"
)

type service struct {
  funcs struct {
    ls shell.RunnerWithFunc
  }
}

func newService() *service {
  s := new(service)
  s.funcs.ls = shell.RunnerWith("ls")
  s.funcs.ls = s.funcs.ls.WithArgs("-a")
  return s
}

func (s *service) list(path string) (string, error) {
  opts := shell.RunOptions{WorkingDir: path}
  _, stdout, err := s.funcs.ls(context.Background(), opts)
  return stdout, err
}

func main() {
  s := newService()
  out, err := s.list("/opt")
  if err != nil {
    panic(err)
  }

  fmt.Println(out)
}
```

```go
package main

import (
  "context"
  "testing"

  "github.com/gardenbed/charm/shell"
)

func TestService_List(t *testing.T) {
  t.Run("Success", func(t *testing.T) {
    s := new(service)
    s.funcs.ls = func(context.Context, shell.RunOptions, ...string) (int, string, error) {
      return 0, "foo bar", nil
    }
    out, err := s.list("/test")
    if out != "foo bar" || err != nil {
      t.Fail()
    }
  })
}

```

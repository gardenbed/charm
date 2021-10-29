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
		ID      int    `ask:"any, your identification number"`
		Name    string `ask:"any, your full name"`
		Contact struct {
			Email string `ask:"email, your email address"`
		}
		Secrets struct {
			Token string `ask:"secret, your access token"`
		}
	}{
		ID:   1,
		Name: "Jane Doe",
	}

	err := askit.Ask(&info, asker)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", info)
}

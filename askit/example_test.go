package askit_test

import (
	"os"

	"github.com/gardenbed/charm/askit"
	"github.com/mitchellh/cli"
)

func ExampleAsk() {
	asker := &cli.BasicUi{
		Reader:      os.Stdin,
		Writer:      os.Stdout,
		ErrorWriter: os.Stderr,
	}

	info := struct {
		Name  string `ask:"any, your full name"`
		Email string `ask:"email, your email address"`
		Token string `ask:"secret, your access token"`
	}{}

	err := askit.Ask(&info, asker)
	if err != nil {
		panic(err)
	}
}

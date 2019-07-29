package main

import (
	"log"
	"os"

	"github.com/hashicorp/tf-sdk-migrator/cmd/check"
	"github.com/hashicorp/tf-sdk-migrator/cmd/migrate"
	"github.com/mitchellh/cli"
)

func main() {
	c := cli.NewCLI("tf-sdk-migrator", "0.1.0")
	c.Args = os.Args[1:]
	c.Commands = map[string]cli.CommandFactory{
		"check":   check.CommandFactory,
		"migrate": migrate.CommandFactory,
	}

	exitStatus, err := c.Run()
	if err != nil {
		log.Println(err)
	}

	os.Exit(exitStatus)
}

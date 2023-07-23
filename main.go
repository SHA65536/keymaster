package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "Keymaster",
		Usage: "A CLI tools for managing AWS keypairs",
		Commands: []*cli.Command{
			listCommand,
			copyCommand,
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

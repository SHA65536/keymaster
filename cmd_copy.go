package main

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

var copyCommand = &cli.Command{
	Name:  "copy",
	Usage: "copy a keypair from one region to another",
	Action: func(cCtx *cli.Context) error {
		fmt.Println("added task: ", cCtx.Args().First())
		return nil
	},
}

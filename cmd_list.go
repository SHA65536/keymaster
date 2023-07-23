package main

import (
	"fmt"
	"log"

	"github.com/sha65536/keymaster/keymaster"
	"github.com/urfave/cli/v2"
)

var listCommand = &cli.Command{
	Name:  "list",
	Usage: "list aws keypairs",
	Flags: []cli.Flag{
		&cli.StringSliceFlag{
			Name:    "regions",
			Aliases: []string{"r"},
			Value:   cli.NewStringSlice(),
			Usage:   "regions to show keys in",
		},
	},
	Action: func(cCtx *cli.Context) error {
		// Create main keymaster object
		km, err := keymaster.MakeKeymaster()
		if err != nil {
			log.Println("Could not create session with AWS! Check credentials!")
			return err
		}
		regions := cCtx.StringSlice("regions")

		// If no regions were given, take all aws regions
		if len(regions) == 0 {
			if regions, err = km.GetAllRegions(); err != nil {
				log.Println("Could not get AWS regions!")
				return err
			}
		}

		// Getting key list
		keys, err := km.ListKeys(regions...)
		if err != nil {
			log.Println("Could not list keypairs in given regions!")
			return err
		}

		for region, keys := range keys {
			fmt.Printf("%s Region:\n", region)
			for _, key := range keys {
				fmt.Printf("\t- %s\n", key.Name)
			}
		}
		return nil
	},
}

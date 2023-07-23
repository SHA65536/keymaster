package main

import (
	"log"

	"github.com/sha65536/keymaster/keymaster"
	"github.com/urfave/cli/v2"
)

var copyCommand = &cli.Command{
	Name:  "copy",
	Usage: "copy a keypair from one region to another",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "src-key",
			Aliases:  []string{"sk"},
			Usage:    "name of the source key",
			Required: true,
		},
		&cli.StringFlag{
			Name:     "src-region",
			Aliases:  []string{"sr"},
			Usage:    "name of the source region",
			Required: true,
		},
		&cli.StringFlag{
			Name:     "dst-key",
			Aliases:  []string{"dk"},
			Usage:    "name of the destination key",
			Required: false,
		},
		&cli.StringFlag{
			Name:     "dst-region",
			Aliases:  []string{"dr"},
			Usage:    "name of the destination region",
			Required: true,
		},
	},
	Action: func(cCtx *cli.Context) error {
		// Create main keymaster object
		var srcKey *keymaster.Key
		var srcKeyName = cCtx.String("src-key")
		var dstKeyName = cCtx.String("dst-key")
		var srcRegion = cCtx.String("src-region")
		var dstRegion = cCtx.String("dst-region")
		if dstKeyName == "" {
			dstKeyName = srcKeyName
		}

		km, err := keymaster.MakeKeymaster()
		if err != nil {
			log.Println("Could not create session with AWS! Check credentials!")
			return err
		}

		// Getting key list
		keys, err := km.ListKeys(srcRegion)
		if err != nil {
			log.Println("Could not list keypairs in given regions!")
			return err
		}

		// Matching key
		if len(keys) == 0 {
			log.Println("No keys found in the given region!")
			return err
		}
		for _, key := range keys[srcRegion] {
			if key.Name == srcKeyName {
				srcKey = &key
				break
			}
		}
		if srcKey == nil {
			log.Println("Given key not found in the given region!")
			return err
		}

		// Importing key to new region
		newKey := keymaster.Key{
			Name:      dstKeyName,
			PublicKey: srcKey.PublicKey,
		}
		if err := km.CreateKey(dstRegion, newKey); err != nil {
			log.Println("Error creating key in new region!")
			return err
		}

		return nil
	},
}

package cli

import (
	"context"
	"glitchkids/registry_gk/index"
	"log"

	"github.com/urfave/cli/v3"
)

func CreateIndexCommand() *cli.Command {
	return &cli.Command{
		Name:        "index",
		Usage:       "Index current folders based on registry.config.json",
		Description: "Index current folders based on registry.config.json",
		Action: func(ctx context.Context, c *cli.Command) error {
			config, err := index.GetIndexConfig()
			if err != nil {
				log.Fatal(err)
			}

			registry, err := index.CreateRegistry(config)
			if err != nil {
				log.Fatal(err)
			}

			err = index.Save(config, registry)
			if err != nil {
				log.Fatal(err)
			}

			return nil
		},
	}
}

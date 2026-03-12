package cli

import (
	"context"
	"fmt"
	"glitchkids/registry_gk/registry"
	"log"

	"github.com/urfave/cli/v3"
)

func CreateListCommand() *cli.Command {
	return &cli.Command{
		Name:        "ls",
		Usage:       "List registry added",
		Description: "List registry added",
		Action: func(ctx context.Context, c *cli.Command) error {
			config, err := registry.GetOrCreateRemoteRegistry()
			if err != nil {
				log.Fatal(err)
			}

			fmt.Print(config)

			return nil
		},
	}
}

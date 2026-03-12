package cli

import (
	"context"
	"fmt"
	"glitchkids/registry_gk/registry"
	"log"

	"github.com/urfave/cli/v3"
)

func CreateRemoveCommand() *cli.Command {
	return &cli.Command{
		Name:        "remove",
		Usage:       "Remove remote registry",
		Description: "Remove remote registry",
		Arguments: []cli.Argument{
			&cli.StringArg{
				Name: "registry_name",
			},
		},
		Action: func(ctx context.Context, c *cli.Command) error {
			config, err := registry.GetOrCreateRemoteRegistry()
			if err != nil {
				log.Fatal(err)
			}

			registryName := c.StringArg("registry_name")
			if registryName == "" {
				log.Fatal("Argument missing")
			}

			removedRegistryItem, err := registry.RemoveItemToRemoteRegistryByName(config, registryName)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Printf("Registry removed : %v", removedRegistryItem)

			return nil
		},
	}
}

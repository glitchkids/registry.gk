package cli

import (
	"context"
	"fmt"
	"glitchkids/registry_gk/registry"
	"log"

	"github.com/urfave/cli/v3"
)

func CreateAddCommand() *cli.Command {
	return &cli.Command{
		Name:        "add",
		Usage:       "Add remote registry",
		Description: "Add remote registry",
		Arguments: []cli.Argument{
			&cli.StringArg{
				Name: "registry_name",
			},
			&cli.StringArg{
				Name: "registry_url",
			},
		},
		Action: func(ctx context.Context, c *cli.Command) error {
			config, err := registry.GetOrCreateRemoteRegistry()
			if err != nil {
				log.Fatal(err)
			}

			registryName := c.StringArg("registry_name")
			registryUrl := c.StringArg("registry_url")

			if registryName == "" || registryUrl == "" {
				log.Fatal("Arguments missing")
			}

			registryItem, err := registry.AddItemToRemoteRegistry(config, registryName, registryUrl)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Printf("New registry added : %v", registryItem)

			return nil
		},
	}
}

package cli

import (
	"context"
	"fmt"
	"glitchkids/registry_gk/registry"
	"log"
	"path/filepath"

	"github.com/urfave/cli/v3"
)

func CreatePullCommand() *cli.Command {
	return &cli.Command{
		Name:        "pull",
		Usage:       "List or pull from remote registry",
		Description: "List or pull from remote registry",
		Arguments: []cli.Argument{
			&cli.StringArg{
				Name: "remote_registry_name",
			},
			&cli.StringArg{
				Name: "registry_name",
			},
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "out",
				Aliases: []string{"o"},
				Usage:   "Output directory",
				Value:   ".",
			},
			&cli.BoolFlag{
				Name:    "force",
				Aliases: []string{"f"},
				Usage:   "Force write files",
				Value:   false,
			},
		},
		Action: func(ctx context.Context, c *cli.Command) error {
			remoteRegistry, err := registry.GetOrCreateRemoteRegistry()
			if err != nil {
				log.Fatal(err)
			}

			remoteRegistryName := c.StringArg("remote_registry_name")
			registryName := c.StringArg("registry_name")

			outDir := c.String("out")
			force := c.Bool("force")

			if filepath.IsLocal(outDir) == false {
				log.Fatal("Output path must be local")

			}

			if remoteRegistryName != "" && registryName == "" {
				registryIndex, err := registry.ListRemoteRegistryIndex(remoteRegistry, remoteRegistryName)
				if err != nil {
					log.Fatal(err)
				}

				fmt.Print(registryIndex)

				return nil
			} else if remoteRegistryName == "" || registryName == "" {
				log.Fatalf("Missing arguments")
			}

			err = registry.PullFromRemoteRegistry(remoteRegistry, remoteRegistryName, registryName, outDir, force)
			if err != nil {
				log.Fatal(err)
			}
			return nil
		},
	}
}

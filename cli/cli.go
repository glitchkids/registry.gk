package cli

import (
	"context"
	"os"

	"github.com/urfave/cli/v3"
)

func RunClI() {
	(&cli.Command{
		Name:        "registry_gk",
		Usage:       "Glitch Kids Registry",
		Description: "Glitch Kids Registry is a tility CLI to easily retrieve folders or files from remote indexer",
		Commands:    []*cli.Command{CreateIndexCommand(), CreatePullCommand(), CreateAddCommand(), CreateListCommand(), CreateRemoveCommand()},
	}).Run(context.Background(), os.Args)
}

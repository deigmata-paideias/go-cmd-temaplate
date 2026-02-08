package root

import (
	"github.com/spf13/cobra"

	"github.com/deigmata-paideias/go-cmd-temaplate/internal/cmd"
)

func GetRootCommand() *cobra.Command {

	c := &cobra.Command{
		Use:   "Go cmd app",
		Short: "Go CLI Tool",
		Long:  "Go cli desc",
	}

	c.AddCommand(cmd.VersionCommand())

	return c
}

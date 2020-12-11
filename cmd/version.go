package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var Version = "v0.0.1"

func init () {
	versionCmd := &cobra.Command{
		Use: "version",
		Short: "Print version then exit",
		Run: func(command *cobra.Command, args []string) {
			fmt.Printf("dcli - Version %s\n", Version)
		},
	}
	Root.AddCommand(versionCmd)
}

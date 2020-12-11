package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var JsonOutput bool
var Token string

var Root = &cobra.Command{
	Use: "dcli",
	Short: "DCLI is a basic Discord CLI client for automating things",
	Version: Version,
}

func init() {
	Root.SetVersionTemplate("{{.Name}} - Version {{.Version}}\n")

	Root.PersistentFlags().BoolVarP(&JsonOutput, "json", "j", false, "Output as JSON")
	Root.PersistentFlags().StringVarP(&Token, "token", "t", "DISCORD_TOKEN", "Discord token")
	if os.Getenv("DISCORD_TOKEN") == "" {
		Root.MarkPersistentFlagRequired("token")
	} else {
		Root.PersistentFlags().Set("token", os.Getenv("DISCORD_TOKEN"))
	}

	cobra.OnInitialize()
}

func Execute() {
	if err := Root.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing command %s\n", err)
		os.Exit(1)
	}
}

// CheckArgs here is copied from github.com/ncw/rclone/cmd
// CheckArgs checks there are enough arguments and prints a message if not
func CheckArgs(MinArgs, MaxArgs int, cmd *cobra.Command, args []string) {
	if len(args) < MinArgs {
		_ = cmd.Usage()
		fmt.Println()
		_, _ = fmt.Fprintf(os.Stderr, "Command %s needs %d arguments minimum: you provided %d non flag arguments: %q\n", cmd.Name(), MinArgs, len(args), args)
		os.Exit(1)
	} else if len(args) > MaxArgs {
		_ = cmd.Usage()
		fmt.Println()
		_, _ = fmt.Fprintf(os.Stderr, "Command %s needs %d arguments maximum: you provided %d non flag arguments: %q\n", cmd.Name(), MaxArgs, len(args), args)
		os.Exit(1)
	}
}

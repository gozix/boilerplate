// Package command contains cli commands.
package command

import (
	"github.com/spf13/cobra"
)

// TagCookieSubCommand is tag marks cookie subcommands.
const TagCookieSubCommand = "cli.cmd.cookie.subcommand"

// NewCookie is command constructor.
func NewCookie(subCommands []*cobra.Command) *cobra.Command {
	var cmd = cobra.Command{
		Use:           "cookie <command>",
		Short:         "Cookie command group",
		SilenceUsage:  true,
		SilenceErrors: true,
	}

	cmd.AddCommand(subCommands...)

	return &cmd
}

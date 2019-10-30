// Package command contains cli commands.
package command

import (
	"github.com/sarulabs/di/v2"
	"github.com/spf13/cobra"

	gzGlue "github.com/gozix/glue/v2"
)

// DefCommandCookieName is container name.
const DefCommandCookieName = "cli.command.cookie"

// DefCommandCookie register command in di container.
func DefCommandCookie() di.Def {
	return di.Def{
		Name: DefCommandCookieName,
		Tags: []di.Tag{{
			Name: gzGlue.TagCliCommand,
		}},
		Build: func(ctn di.Container) (_ interface{}, err error) {
			var addCmd *cobra.Command
			if err = ctn.Fill(DefCommandCookieAddName, &addCmd); err != nil {
				return nil, err
			}

			var fetchCmd *cobra.Command
			if err = ctn.Fill(DefCommandCookieFetchName, &fetchCmd); err != nil {
				return nil, err
			}

			var cmd = &cobra.Command{
				Use:           "cookie <command>",
				Short:         "Cookie command group",
				SilenceUsage:  true,
				SilenceErrors: true,
			}

			cmd.AddCommand(addCmd)
			cmd.AddCommand(fetchCmd)

			return cmd, nil
		},
	}
}

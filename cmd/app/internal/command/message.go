// Package command contains cli commands.
package command

import (
	"github.com/gozix/glue"
	"github.com/gozix/viper"
	"github.com/gozix/zap"
	"github.com/sarulabs/di"
	"github.com/spf13/cobra"
)

// DefCommandMessageName is container name.
const DefCommandMessageName = "cli.command.message"

// DefCommandMessage register command in di container.
func DefCommandMessage() di.Def {
	return di.Def{
		Name: DefCommandMessageName,
		Tags: []di.Tag{{
			Name: glue.TagCliCommand,
		}},
		Build: func(ctn di.Container) (_ interface{}, err error) {
			return &cobra.Command{
				Use:           "message",
				Short:         "Write configured message to log",
				SilenceUsage:  true,
				SilenceErrors: true,
				RunE: func(cmd *cobra.Command, args []string) error {
					var cfg *viper.Viper
					if err = ctn.Fill(viper.BundleName, &cfg); err != nil {
						return err
					}

					var logger *zap.Logger
					if err = ctn.Fill(zap.BundleName, &logger); err != nil {
						return err
					}

					logger.Info(cfg.GetString("message"))

					return nil
				},
			}, nil
		},
	}
}

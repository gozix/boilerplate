// Package command contains cli commands.
package command

import (
	"github.com/gozix/glue"
	"github.com/gozix/viper"
	"github.com/gozix/zap"
	"github.com/sarulabs/di"
	"github.com/spf13/cobra"
)

// DefCommandConfig is container name.
const DefCommandConfig = "cli.command.config"

// RegisterMessageCommand register command in di container.
func RegisterMessageCommand(builder *di.Builder) {
	builder.Add(di.Def{
		Name: DefCommandConfig,
		Tags: []di.Tag{{
			Name: glue.TagCliCommand,
		}},
		Build: func(ctn di.Container) (_ interface{}, err error) {
			var cfg *viper.Viper
			if err = ctn.Fill(viper.BundleName, &cfg); err != nil {
				return nil, err
			}

			var logger *zap.Logger
			if err = ctn.Fill(zap.BundleName, &logger); err != nil {
				return nil, err
			}

			return NewMessageCommand(cfg, logger), nil
		},
	})
}

// NewMessageCommand is command constructor.
func NewMessageCommand(cfg *viper.Viper, logger *zap.Logger) *cobra.Command {
	return &cobra.Command{
		Use:           "message",
		Short:         "Write configured message to log",
		SilenceUsage:  true,
		SilenceErrors: true,
		Run: func(cmd *cobra.Command, args []string) {
			logger.Info(cfg.GetString("message"))
		},
	}
}

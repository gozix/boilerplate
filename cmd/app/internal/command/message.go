package command

import (
	"github.com/gozix/glue"
	"github.com/gozix/viper"
	"github.com/gozix/zap"
	"github.com/sarulabs/di"
	"github.com/spf13/cobra"
)

func RegisterConfigCommand(builder *di.Builder) {
	builder.AddDefinition(di.Definition{
		Name: "cli.command.config",
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

			return &cobra.Command{
				Use:           "message",
				Short:         "Display configured message",
				SilenceUsage:  true,
				SilenceErrors: true,
				Run: func(cmd *cobra.Command, args []string) {
					logger.Info(cfg.GetString("message"))
				},
			}, nil
		},
	})
}

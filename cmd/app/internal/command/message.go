// Package command contains cli commands.
package command

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

// NewMessage is command constructor.
func NewMessage(cfg *viper.Viper, logger *zap.Logger) *cobra.Command {
	return &cobra.Command{
		Use:           "message",
		Short:         "Write configured message to log",
		SilenceUsage:  true,
		SilenceErrors: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			logger.Info(cfg.GetString("message"))

			return nil
		},
	}
}

// Package command contains cli commands.
package command

import (
	"strconv"

	"github.com/gozix/di"
	"github.com/spf13/cobra"
	"go.uber.org/zap"

	"github.com/gozix/boilerplate/cmd/app/internal/domain"
)

func NewCookieFetch(ctn di.Container) *cobra.Command {
	return &cobra.Command{
		Use:           "fetch <id>",
		Short:         "Fetch cookie by id",
		SilenceUsage:  true,
		SilenceErrors: true,
		Args:          cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return ctn.Call(func(logger *zap.Logger, repository domain.CookieRepository) (err error) {
				var id int64
				if id, err = strconv.ParseInt(args[0], 10, 64); err != nil {
					return err
				}

				var cookie *domain.Cookie
				if cookie, err = repository.FindOneByID(id); err != nil {
					return err
				}

				logger.Info("Fetched cookie", zap.Reflect("cookie", cookie))

				return nil
			})
		},
	}
}

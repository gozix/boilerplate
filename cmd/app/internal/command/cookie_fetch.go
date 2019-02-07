// Package command contains cli commands.
package command

import (
	"strconv"

	"github.com/gozix/glue"
	ut "github.com/gozix/universal-translator"
	validatorBundle "github.com/gozix/validator"
	zapBundle "github.com/gozix/zap"
	"github.com/sarulabs/di"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"gopkg.in/go-playground/validator.v9"

	"github.com/gozix/boilerplate/cmd/app/internal/database"
	"github.com/gozix/boilerplate/cmd/app/internal/domain"
)

// DefCommandCookieFetchName is container name.
const DefCommandCookieFetchName = "cli.command.cookie_fetch"

// DefCommandCookieFetch register command in di container.
func DefCommandCookieFetch() di.Def {
	return di.Def{
		Name: DefCommandCookieFetchName,
		Tags: []di.Tag{{
			Name: glue.TagCliCommand,
		}},
		Build: func(ctn di.Container) (_ interface{}, err error) {
			return &cobra.Command{
				Use:           "fetch <id>",
				Short:         "Fetch cookie by id",
				SilenceUsage:  true,
				SilenceErrors: true,
				Args:          cobra.ExactArgs(1),
				RunE: func(cmd *cobra.Command, args []string) (err error) {
					var logger *zap.Logger
					if err = ctn.Fill(zapBundle.BundleName, &logger); err != nil {
						return err
					}

					var repository domain.CookieRepository
					if err = ctn.Fill(database.DefCookieRepositoryName, &repository); err != nil {
						return err
					}

					var translator *ut.UniversalTranslator
					if err = ctn.Fill(ut.BundleName, &translator); err != nil {
						return err
					}

					var validate *validator.Validate
					if err = ctn.Fill(validatorBundle.BundleName, &validate); err != nil {
						return err
					}

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
				},
			}, nil
		},
	}
}

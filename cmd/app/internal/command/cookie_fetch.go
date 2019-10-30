// Package command contains cli commands.
package command

import (
	"strconv"

	ut "github.com/go-playground/universal-translator"
	"github.com/sarulabs/di/v2"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"gopkg.in/go-playground/validator.v9"

	gzGlue "github.com/gozix/glue/v2"
	gzUT "github.com/gozix/universal-translator/v2"
	gzValidator "github.com/gozix/validator/v2"
	gzZap "github.com/gozix/zap/v2"

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
			Name: gzGlue.TagCliCommand,
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
					if err = ctn.Fill(gzZap.BundleName, &logger); err != nil {
						return err
					}

					var repository domain.CookieRepository
					if err = ctn.Fill(database.DefCookieRepositoryName, &repository); err != nil {
						return err
					}

					var translator *ut.UniversalTranslator
					if err = ctn.Fill(gzUT.BundleName, &translator); err != nil {
						return err
					}

					var validate *validator.Validate
					if err = ctn.Fill(gzValidator.BundleName, &validate); err != nil {
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

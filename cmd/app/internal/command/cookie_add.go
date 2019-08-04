// Package command contains cli commands.
package command

import (
	utBundle "github.com/gozix/universal-translator/v2"
	validatorBundle "github.com/gozix/validator/v2"
	zapBundle "github.com/gozix/zap/v2"
	"github.com/pkg/errors"
	"github.com/sarulabs/di/v2"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"gopkg.in/go-playground/validator.v9"

	"github.com/gozix/boilerplate/cmd/app/internal/database"
	"github.com/gozix/boilerplate/cmd/app/internal/domain"
)

// DefCommandCookieAddName is container name.
const DefCommandCookieAddName = "cli.command.cookie_add"

// DefCommandCookieAdd register command in di container.
func DefCommandCookieAdd() di.Def {
	return di.Def{
		Name: DefCommandCookieAddName,
		Build: func(ctn di.Container) (_ interface{}, err error) {
			return &cobra.Command{
				Use:           "add <name>",
				Short:         "Add new cookie",
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

					var translator *utBundle.UniversalTranslator
					if err = ctn.Fill(utBundle.BundleName, &translator); err != nil {
						return err
					}

					var validate *validator.Validate
					if err = ctn.Fill(validatorBundle.BundleName, &validate); err != nil {
						return err
					}

					var cookie = &domain.Cookie{
						Name: args[0],
					}

					if err = validate.Struct(cookie); err != nil {
						switch v := err.(type) {
						case validator.ValidationErrors:
							for _, e := range v {
								return errors.New(e.Translate(translator.GetFallback()))
							}
						}

						return err
					}

					return repository.Save(cookie)
				},
			}, nil
		},
	}
}

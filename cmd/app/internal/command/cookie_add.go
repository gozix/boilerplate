// Package command contains cli commands.
package command

import (
	"errors"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/gozix/di"
	"github.com/spf13/cobra"
	"go.uber.org/zap"

	"github.com/gozix/boilerplate/cmd/app/internal/domain"
)

// NewCookieAdd is cookie subcommand constructor.
func NewCookieAdd(ctn di.Container) *cobra.Command {
	return &cobra.Command{
		Use:           "add <name>",
		Short:         "Add new cookie",
		SilenceUsage:  true,
		SilenceErrors: true,
		Args:          cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return ctn.Call(func(
				logger *zap.Logger,
				repository domain.CookieRepository,
				translator *ut.UniversalTranslator,
				validate *validator.Validate,
			) (err error) {
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
			})
		},
	}
}

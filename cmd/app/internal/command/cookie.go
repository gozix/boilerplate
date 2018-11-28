// Package command contains cli commands.
package command

import (
	"strconv"

	"github.com/gozix/glue"
	"github.com/gozix/universal-translator"
	validatorBundle "github.com/gozix/validator"
	zapBundle "github.com/gozix/zap"
	"github.com/pkg/errors"
	"github.com/sarulabs/di"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"gopkg.in/go-playground/validator.v9"

	"github.com/gozix/boilerplate/cmd/app/internal/database"
	"github.com/gozix/boilerplate/cmd/app/internal/domain"
)

// DefCommandCookieName is container name.
const DefCommandCookieName = "cli.command.cookie"

// DefCommandCookie register command in di container.
func DefCommandCookie() di.Def {
	return di.Def{
		Name: DefCommandCookieName,
		Tags: []di.Tag{{
			Name: glue.TagCliCommand,
		}},
		Build: func(ctn di.Container) (_ interface{}, err error) {
			var logger *zap.Logger
			if err = ctn.Fill(zapBundle.BundleName, &logger); err != nil {
				return nil, err
			}

			var repository domain.CookieRepository
			if err = ctn.Fill(database.DefCookieRepositoryName, &repository); err != nil {
				return nil, err
			}

			var translator *ut.UniversalTranslator
			if err = ctn.Fill(ut.BundleName, &translator); err != nil {
				return nil, err
			}

			var validate *validator.Validate
			if err = ctn.Fill(validatorBundle.BundleName, &validate); err != nil {
				return nil, err
			}

			return NewCookieCommand(logger, repository, translator, validate), nil
		},
	}
}

// NewCookieCommand is command constructor.
func NewCookieCommand(
	logger *zap.Logger,
	repository domain.CookieRepository,
	translator *ut.UniversalTranslator,
	validate *validator.Validate,
) *cobra.Command {
	var root = cobra.Command{
		Use:           "cookie <command>",
		Short:         "Cookie command group",
		SilenceUsage:  true,
		SilenceErrors: true,
	}

	root.AddCommand(&cobra.Command{
		Use:           "add <name>",
		Short:         "Add new cookie",
		SilenceUsage:  true,
		SilenceErrors: true,
		Args:          cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
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
	})

	root.AddCommand(&cobra.Command{
		Use:           "fetch <id>",
		Short:         "Fetch cookie by id",
		SilenceUsage:  true,
		SilenceErrors: true,
		Args:          cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
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
	})

	return &root
}

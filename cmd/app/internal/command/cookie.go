// Package command contains cli commands.
package command

import (
	"strconv"

	"github.com/gozix/glue"
	zapBundle "github.com/gozix/zap"
	"github.com/sarulabs/di"
	"github.com/spf13/cobra"
	"go.uber.org/zap"

	"github.com/gozix/boilerplate/cmd/app/internal/database"
	"github.com/gozix/boilerplate/cmd/app/internal/domain"
)

// DefCommandCookie is container name.
const DefCommandCookie = "cli.command.cookie"

// RegisterCookieCommand register command in di container.
func RegisterCookieCommand(builder *di.Builder) {
	builder.Add(di.Def{
		Name: DefCommandCookie,
		Tags: []di.Tag{{
			Name: glue.TagCliCommand,
		}},
		Build: func(ctn di.Container) (_ interface{}, err error) {
			var logger *zap.Logger
			if err = ctn.Fill(zapBundle.BundleName, &logger); err != nil {
				return nil, err
			}

			var repository domain.CookieRepository
			if err = ctn.Fill(database.DefCookieRepository, &repository); err != nil {
				return nil, err
			}

			return NewCookieCommand(logger, repository), nil
		},
	})
}

// NewCookieCommand is command constructor.
func NewCookieCommand(logger *zap.Logger, repository domain.CookieRepository) *cobra.Command {
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
		RunE: func(cmd *cobra.Command, args []string) error {
			return repository.Save(&domain.Cookie{
				Name: args[0],
			})
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

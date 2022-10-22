// Package internal provide component implementation.
package internal

import (
	"github.com/gozix/di"
	gzGlue "github.com/gozix/glue/v3"
	gzPrometheus "github.com/gozix/prometheus/v2"
	gzSQLMigrate "github.com/gozix/sql-migrate/v3"
	gzSQL "github.com/gozix/sql/v3"
	gzViper "github.com/gozix/viper/v3"
	gzZapGelf "github.com/gozix/zap-gelf/v2"
	gzZap "github.com/gozix/zap/v3"

	"github.com/gozix/boilerplate/cmd/app/internal/command"
	"github.com/gozix/boilerplate/cmd/app/internal/database"
	"github.com/gozix/boilerplate/cmd/app/internal/domain"
)

// Bundle is component bundle.
type Bundle struct{}

// NewBundle is bundle constructor.
func NewBundle() *Bundle {
	return &Bundle{}
}

// Name implements the gzGlue.Bundle interface.
func (*Bundle) Name() string {
	return "app"
}

// Build implements the gzGlue.Bundle interface.
func (*Bundle) Build(builder di.Builder) error {
	var err = builder.Apply(
		di.BuilderOptions( // command
			di.Provide(
				command.NewCookie,
				di.Constraint(0, di.WithTags(command.TagCookieSubCommand)),
				gzGlue.AsCliCommand(),
			),
			di.Provide(command.NewCookieAdd, di.Tags{{
				Name: command.TagCookieSubCommand,
			}}),
			di.Provide(command.NewCookieFetch, di.Tags{{
				Name: command.TagCookieSubCommand,
			}}),
			di.Provide(command.NewMessage, gzGlue.AsCliCommand()),
		),
		di.BuilderOptions( // database
			di.Provide(database.NewCookie, di.As(new(domain.CookieRepository))),
		),
	)

	return err
}

func (*Bundle) DependsOn() []string {
	return []string{
		gzPrometheus.BundleName,
		gzSQL.BundleName,
		gzSQLMigrate.BundleName,
		gzViper.BundleName,
		gzZap.BundleName,
		gzZapGelf.BundleName,
	}
}

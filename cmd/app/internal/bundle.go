// Package internal provide component implementation.
package internal

import (
	sqlBundle "github.com/gozix/sql/v2"
	"github.com/sarulabs/di/v2"

	"github.com/gozix/boilerplate/cmd/app/internal/command"
	"github.com/gozix/boilerplate/cmd/app/internal/database"
)

// Bundle is component bundle.
type Bundle struct{}

// NewBundle is bundle constructor.
func NewBundle() *Bundle {
	return &Bundle{}
}

// Name implements the glue.Bundle interface.
func (*Bundle) Name() string {
	return "app"
}

// Build implements the glue.Bundle interface.
func (*Bundle) Build(builder *di.Builder) error {
	return builder.Add(
		// commands
		command.DefCommandCookie(),
		command.DefCommandCookieAdd(),
		command.DefCommandCookieFetch(),
		command.DefCommandMessage(),

		// database
		database.DefCookieRepository(),
	)
}

// DependsOn implements the glue.BundleDependsOn interface.
func (*Bundle) DependsOn() []string {
	return []string{
		sqlBundle.BundleName,
	}
}

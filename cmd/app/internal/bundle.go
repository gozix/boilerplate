// Package internal provide component implementation.
package internal

import (
	"github.com/sarulabs/di"

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
	// commands
	command.RegisterCookieCommand(builder)
	command.RegisterMessageCommand(builder)

	// database
	database.RegisterCookieRepository(builder)

	return nil
}

// DependsOn implements the glue.BundleDependsOn interface.
func (*Bundle) DependsOn() []string {
	return []string{"sql"}
}

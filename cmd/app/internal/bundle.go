package internal

import (
	"github.com/sarulabs/di"

	"github.com/gozix/boilerplate/cmd/app/internal/command"
)

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
	command.RegisterConfigCommand(builder)

	return nil
}

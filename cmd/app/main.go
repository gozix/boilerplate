// Package main provide component entry point.
package main

import (
	"log"

	"github.com/gozix/glue"
	"github.com/gozix/redigo"
	"github.com/gozix/sql"
	migrate "github.com/gozix/sql-migrate"
	ut "github.com/gozix/universal-translator"
	"github.com/gozix/validator"
	"github.com/gozix/viper"
	"github.com/gozix/zap"
	_ "github.com/lib/pq" // Postgres database/sql driver

	"github.com/gozix/boilerplate/cmd/app/internal"
)

// Version is component version.
const Version = "0.0.1"

func main() {
	var app, err = glue.NewApp(
		glue.Version(Version),
		glue.Bundles(
			internal.NewBundle(),
			viper.NewBundle(),
			zap.NewBundle(),
			sql.NewBundle(),
			redigo.NewBundle(),
			migrate.NewBundle(),
			validator.NewBundle(),
			ut.NewBundle(),
		),
	)

	if err != nil {
		log.Fatalf("Some error occurred during create app. Error: %v\n", err)
	}

	if err = app.Execute(); err != nil {
		log.Fatalf("Some error occurred during execute app. Error: %v\n", err)
	}
}

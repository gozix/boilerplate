// Package main provide component entry point.
package main

import (
	"log"

	glueBundle "github.com/gozix/glue/v2"
	redigoBundle "github.com/gozix/redigo/v2"
	migrateBundle "github.com/gozix/sql-migrate/v2"
	sqlBundle "github.com/gozix/sql/v2"
	utBundle "github.com/gozix/universal-translator/v2"
	validatorBundle "github.com/gozix/validator/v2"
	viperBundle "github.com/gozix/viper/v2"
	zapBundle "github.com/gozix/zap/v2"
	_ "github.com/lib/pq" // Postgres database/sql driver

	internalBundle "github.com/gozix/boilerplate/cmd/app/internal"
)

// Version is component version.
const Version = "0.0.1"

func main() {
	var app, err = glueBundle.NewApp(
		glueBundle.Version(Version),
		glueBundle.Bundles(
			redigoBundle.NewBundle(),
			migrateBundle.NewBundle(),
			sqlBundle.NewBundle(),
			utBundle.NewBundle(),
			validatorBundle.NewBundle(),
			viperBundle.NewBundle(),
			zapBundle.NewBundle(),
			internalBundle.NewBundle(),
		),
	)

	if err != nil {
		log.Fatalf("Some error occurred during create app. Error: %v\n", err)
	}

	if err = app.Execute(); err != nil {
		log.Fatalf("Some error occurred during execute app. Error: %v\n", err)
	}
}

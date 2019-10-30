// Package main provide component entry point.
package main

import (
	"log"

	gzGlue "github.com/gozix/glue/v2"
	gzRedigo "github.com/gozix/redigo/v3"
	gzMigrate "github.com/gozix/sql-migrate/v2"
	gzSQL "github.com/gozix/sql/v2"
	gzUT "github.com/gozix/universal-translator/v2"
	gzValidator "github.com/gozix/validator/v2"
	gzViper "github.com/gozix/viper/v2"
	gzZap "github.com/gozix/zap/v2"
	_ "github.com/lib/pq" // Postgres database/sql driver

	gzInternal "github.com/gozix/boilerplate/cmd/app/internal"
)

// Version is component version.
const Version = "0.0.1"

func main() {
	var app, err = gzGlue.NewApp(
		gzGlue.Version(Version),
		gzGlue.Bundles(
			gzRedigo.NewBundle(),
			gzMigrate.NewBundle(),
			gzSQL.NewBundle(),
			gzUT.NewBundle(),
			gzValidator.NewBundle(),
			gzViper.NewBundle(),
			gzZap.NewBundle(),
			gzInternal.NewBundle(),
		),
	)

	if err != nil {
		log.Fatalf("Some error occurred during create app. Error: %v\n", err)
	}

	if err = app.Execute(); err != nil {
		log.Fatalf("Some error occurred during execute app. Error: %v\n", err)
	}
}

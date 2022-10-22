// Package main provide component entry point.
package main

import (
	"log"
	"runtime/debug"

	gzAerospike "github.com/gozix/aerospike/v2"
	gzClockwork "github.com/gozix/clockwork/v3"
	gzConsul "github.com/gozix/consul/v3"
	gzEcho "github.com/gozix/echo/v3"
	gzGlue "github.com/gozix/glue/v3"
	gzGoredis "github.com/gozix/goredis/v4"
	gzPrometheus "github.com/gozix/prometheus/v2"
	gzRedigo "github.com/gozix/redigo/v4"
	gzSQLMigrate "github.com/gozix/sql-migrate/v3"
	gzSQL "github.com/gozix/sql/v3"
	gzUT "github.com/gozix/universal-translator/v3"
	gzValidator "github.com/gozix/validator/v3"
	gzViper "github.com/gozix/viper/v3"
	gzZapGelf "github.com/gozix/zap-gelf/v2"
	gzZap "github.com/gozix/zap/v3"
	_ "github.com/lib/pq" // Postgres database/sql driver

	gzInternal "github.com/gozix/boilerplate/cmd/app/internal"
)

// Version is component version.
const Version = "0.0.1"

func main() {
	var app, err = gzGlue.NewApp(
		gzGlue.Version(Version),
		gzGlue.Bundles(
			gzInternal.NewBundle(),

			gzAerospike.NewBundle(),
			gzClockwork.NewBundle(),
			gzConsul.NewBundle(),
			gzEcho.NewBundle(),
			gzGoredis.NewBundle(),
			gzPrometheus.NewBundle(),
			gzRedigo.NewBundle(),
			gzSQL.NewBundle(),
			gzSQLMigrate.NewBundle(),
			gzUT.NewBundle(),
			gzValidator.NewBundle(),
			gzViper.NewBundle(),
			gzZap.NewBundle(),
			gzZapGelf.NewBundle(),
		),
	)

	if err != nil {
		log.Fatalf("Some error occurred during create app. Error: %s\n\n%s", err, debug.Stack())
	}

	if err = app.Execute(); err != nil {
		log.Fatalf("Some error occurred during execute app. Error: %s\n\n%s", err, debug.Stack())
	}
}

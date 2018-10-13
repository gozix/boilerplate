package main

import (
	"fmt"
	"os"

	"github.com/gozix/glue"
	"github.com/gozix/viper"
	"github.com/gozix/zap"

	"github.com/gozix/boilerplate/cmd/app/internal"
)

const Version = "0.0.1"

func main() {
	var app, err = glue.NewApp(
		glue.Version(Version),
		glue.Bundles(
			viper.NewBundle(),
			zap.NewBundle(),
			internal.NewBundle(),
		),
	)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Some error occurred during create app. Error: %v\n", err)
		os.Exit(1)
	}

	if err = app.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Some error occurred during execute app. Error: %v\n", err)
		os.Exit(2)
	}
}

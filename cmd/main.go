package main

import (
	_ "github.com/psyb0t/logrus-configurator" //nolint:depguard
	"github.com/psyb0t/sql-migrator/internal/app"
	"github.com/sirupsen/logrus"
)

func main() {
	if err := app.Run(); err != nil {
		logrus.Fatalf("failed to execute: %v", err)
	}
}

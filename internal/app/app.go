package app

import (
	"github.com/pkg/errors"
	"github.com/psyb0t/sql-migrator/internal/app/generate"
	"github.com/psyb0t/sql-migrator/internal/app/migrate"
	"github.com/spf13/cobra"
)

const (
	cmdName = "sql-migrator"
	cmdDesc = "A CLI for database migrations."

	cmdNameMigrate = "migrate"
	cmdDescMigrate = "Migrate the database."

	cmdNameGenerate = "generate"
	cmdDescGenerate = "Generate migration file."
)

func Run() error {
	cmd := &cobra.Command{
		Use:   cmdName,
		Short: cmdDesc,
		CompletionOptions: cobra.CompletionOptions{
			DisableDefaultCmd: true,
		},
	}

	migrateCmd := &cobra.Command{
		Use:   cmdNameMigrate,
		Short: cmdDescMigrate,
		Run:   migrate.Run,
	}

	migrate.HandleFlags(migrateCmd)

	generateCmd := &cobra.Command{
		Use:   cmdNameGenerate,
		Short: cmdDescGenerate,
		Run:   generate.Run,
	}

	generate.HandleFlags(generateCmd)

	cmd.AddCommand(migrateCmd, generateCmd)

	if err := cmd.Execute(); err != nil {
		return errors.Wrap(err, "failed to execute cmd")
	}

	return nil
}

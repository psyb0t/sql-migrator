package migrate

import (
	"database/sql"

	"github.com/pkg/errors"
	"github.com/psyb0t/sql-migrator/internal/pkg/migrate"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

const (
	dbTypeSqlite3 = migrate.DBTypeSqlite3
)

func isSupportedDBType(dbType migrate.DBType) bool {
	supportedDBTypes := []migrate.DBType{
		dbTypeSqlite3,
	}

	for _, supportedDBType := range supportedDBTypes {
		if dbType == supportedDBType {
			return true
		}
	}

	return false
}

func Run(_ *cobra.Command, _ []string) {
	logrus.Debug("started migrate")

	validateFlags()

	logrus.Infof(
		"dbType: %s, dsn: %s, dbName: %s, migrationsDirPath: %s",
		flags.dbType,
		flags.dsn,
		flags.dbName,
		flags.dirPath,
	)

	dbType := migrate.DBType(flags.dbType)

	if !isSupportedDBType(dbType) {
		logrus.Fatalf("unsupported db type: %s", dbType)
	}

	db, err := sql.Open(dbType.String(), flags.dsn)
	if err != nil {
		logrus.Fatalf("could not connect to db: %v", err)
	}

	err = migrate.Migrate(
		db,
		dbType,
		flags.dbName,
		flags.dirPath,
	)

	if err != nil {
		logrus.Fatal(errors.Wrap(err, "failed to migrate"))
	}

	logrus.Info("migrated successfully")
}

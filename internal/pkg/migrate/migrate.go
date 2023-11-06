package migrate

import (
	"database/sql"
	"fmt"

	"github.com/golang-migrate/migrate"                  //nolint:depguard
	"github.com/golang-migrate/migrate/database"         //nolint:depguard
	"github.com/golang-migrate/migrate/database/sqlite3" //nolint:depguard
	_ "github.com/golang-migrate/migrate/source/file"    //nolint:depguard
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type DBType string

func (t DBType) String() string {
	return string(t)
}

const (
	DBTypeSqlite3 DBType = "sqlite3"
)

type Migrator interface {
	Up() error
	Down() error
}

func Migrate(
	db *sql.DB,
	dbType DBType,
	dbName string,
	migrationsDirPath string,
) error {
	logrus.Debug("entered Migrate")

	logrus.Debug("getting driver")

	driver, err := getDriverByDBType(db, dbType)
	if err != nil {
		logrus.Debugf("error getting driver: %v", err)

		return errors.Wrap(err, "failed to get driver")
	}

	logrus.Debug("getting migrator")

	m, err := getMigrator(getMigrationsDirURI(migrationsDirPath), dbName, driver)
	if err != nil {
		logrus.Debugf("error getting migrator: %v", err)

		return errors.Wrap(err, "failed to get migrator")
	}

	logrus.Debug("running migration up")

	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		logrus.Debugf("error running migration up: %v", err)

		return errors.Wrap(err, "failed to migrate up")
	}

	logrus.Debug("completed migration")

	return nil
}

func Rollback(
	db *sql.DB,
	dbType DBType,
	dbName string,
	migrationsDirPath string,
) error {
	logrus.Debug("entered Rollback")

	logrus.Debug("getting driver")

	driver, err := getDriverByDBType(db, dbType)
	if err != nil {
		logrus.Debugf("error getting driver: %v", err)

		return errors.Wrap(err, "failed to get driver")
	}

	logrus.Debug("getting migrator")

	m, err := getMigrator(getMigrationsDirURI(migrationsDirPath), dbName, driver)
	if err != nil {
		logrus.Debugf("error getting migrator: %v", err)

		return errors.Wrap(err, "failed to get migrator")
	}

	logrus.Debug("running migration down")

	if err := m.Down(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		logrus.Debugf("error running migration down: %v", err)

		return errors.Wrap(err, "failed to migrate down")
	}

	logrus.Debug("completed rollback")

	return nil
}

func getDriverByDBType(db *sql.DB, dbType DBType) (database.Driver, error) { //nolint:ireturn
	logrus.Debug("entering getdriverbydbtype")

	logrus.Debug("switching on db type")

	switch dbType {
	case DBTypeSqlite3:
		logrus.Debug("db type is sqlite3")

		driver, err := sqlite3.WithInstance(db, &sqlite3.Config{})
		if err != nil {
			logrus.Debugf("error creating sqlite3 driver: %v", err)

			return nil, errors.Wrap(err, "failed to create sqlite3 driver")
		}

		logrus.Debug("returning sqlite3 driver")

		return driver, nil
	default:
		logrus.Debug("unsupported db type")

		return nil, errors.Wrap(ErrUnsupportedDBType, string(dbType))
	}
}

func getMigrator( //nolint:ireturn
	migrationsURI,
	dbName string,
	driver database.Driver,
) (Migrator, error) {
	logrus.Debug("entered getMigrator")

	logrus.Debug("getting migrator instance")

	m, err := migrate.NewWithDatabaseInstance(migrationsURI, dbName, driver)
	if err != nil {
		logrus.Debugf("error creating migrate instance %v", err)

		return nil, errors.Wrap(err, "failed to create migrate instance")
	}

	logrus.Debug("returning migrator instance")

	return m, nil
}

func getMigrationsDirURI(migrationsDirPath string) string {
	logrus.Debug("entered getMigrationsDirURI")

	logrus.Debugf("migrations directory path: %v", migrationsDirPath)
	uri := fmt.Sprintf("file://%s", migrationsDirPath)
	logrus.Debugf("created migrations directory URI: %v", uri)

	return uri
}

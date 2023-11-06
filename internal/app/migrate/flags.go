package migrate

import (
	"os"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

//nolint:gochecknoglobals
var flags struct {
	dbType  string
	dsn     string
	dbName  string
	dirPath string
}

func HandleFlags(cmd *cobra.Command) {
	var name, value, usage string

	name = "dbtype"
	value = ""
	usage = "Database type (e.g., sqlite3)"
	cmd.Flags().StringVar(&flags.dbType, name, value, usage)

	name = "dsn"
	value = ""
	usage = "Data Source Name (for sqlite it's the path, for others it might be user:pass@/dbname etc.)"
	cmd.Flags().StringVar(&flags.dsn, name, value, usage)

	name = "dbname"
	value = ""
	usage = "Name of the database"
	cmd.Flags().StringVar(&flags.dbName, name, value, usage)

	name = "dir"
	value = ""
	usage = "Path to migrations directory"
	cmd.Flags().StringVar(&flags.dirPath, name, value, usage)
}

func validateFlags() {
	logrus.Debug("entered validateMigrateFlags function")

	if flags.dbType == "" {
		logrus.Fatalf("dbType cannot be empty")
	}

	if flags.dsn == "" {
		logrus.Fatalf("dsn cannot be empty")
	}

	if flags.dbName == "" {
		logrus.Fatalf("dbName cannot be empty")
	}

	if flags.dirPath == "" {
		logrus.Fatalf("directory path cannot be empty")
	}

	if _, err := os.Stat(flags.dirPath); errors.Is(err, os.ErrNotExist) {
		logrus.Fatalf("directory does not exist: %s", flags.dirPath)
	}
}

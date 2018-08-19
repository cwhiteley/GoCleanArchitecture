package console

import (
	"fmt"
	"strings"

	"recipes/appcontext"

	// blank import needed as it's a dependency of the migrate library
	_ "github.com/mattes/migrate/driver/postgres"
	"github.com/mattes/migrate/migrate"
	pipep "github.com/mattes/migrate/pipe"
)

const dbMigrationsPath = "./migrations"

func RunDatabaseMigrations() error {
	allErrors, ok := migrate.UpSync(appcontext.DBConnectionString(), dbMigrationsPath)
	if !ok {
		return joinErrors(allErrors)
	}

	fmt.Println("Migration successful")

	return nil
}

func RollbackLatestMigration() error {
	pipe := pipep.New()

	go migrate.Migrate(pipe, appcontext.DBConnectionString(), dbMigrationsPath, -1)
	return joinErrors(pipep.ReadErrors(pipe))
}

func joinErrors(errors []error) error {
	var errorMsgs []string
	for _, err := range errors {
		errorMsgs = append(errorMsgs, err.Error())
	}

	return fmt.Errorf(strings.Join(errorMsgs, ","))
}

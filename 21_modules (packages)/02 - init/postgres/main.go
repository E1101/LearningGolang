package postgres

// Each package has the ability to provide as many init functions
// as necessary to be invoked at the beginning of execution time.
// All the init functions that are discovered by the compiler are scheduled
// to be executed prior to the main function being executed.

import (
	"database/sql"
)

func init() {
	// When a program imports this package, the init function will be called, causing the
	// database driver to be registered with Go’s sql package as an available driver.
	sql.Register("postgres", new(PostgresDriver))
}

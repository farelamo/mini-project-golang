package database

import (
	"database/sql"
	"fmt"
	"github.com/gobuffalo/packr/v2"
	migrate "github.com/rubenv/sql-migrate"
)

var (
	DbConnection *sql.DB
)

func DbMigrate(dbParam *sql.DB){
	migrations := &migrate.PackrMigrationSource{
		Box: packr.New("migrations", "./migrations"),
	}

	n, errors := migrate.Exec(dbParam, "postgres", migrations, migrate.Up)
	if errors != nil {
		panic(errors)
	}

	DbConnection = dbParam

	fmt.Println("Applied ", n, " migrations!")
}
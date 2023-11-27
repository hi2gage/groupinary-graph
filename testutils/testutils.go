// package testutils

package testutils

import (
	"context"
	"database/sql"
	"groupinary/ent"
	"groupinary/utils"
	"time"

	entsql "entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/schema"

	"github.com/go-testfixtures/testfixtures/v3"
)

func OpenTest(paths ...string) (*ent.Client, error) {
	// Use SQLite in-memory database for testing.
	db, err := sql.Open("sqlite3", "file:test:memory:?cache=shared&_fk=1")
	if err != nil {
		return nil, err
	}
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxLifetime(time.Hour)

	// Create an ent.Driver from `db`.
	drv := entsql.OpenDB("sqlite3", db)

	// Run migrations.
	client := ent.NewClient(ent.Driver(drv))
	if err := runMigrations(client); err != nil {
		return nil, err
	}

	// Load fixtures.
	if err := loadFixtures(db, paths...); err != nil {
		return nil, err
	}

	return client, nil
}

func runMigrations(client *ent.Client) error {
	return client.Schema.Create(context.Background(), schema.WithDropIndex(true))
}

func loadFixtures(db *sql.DB, paths ...string) error {
	fixtures, err := testfixtures.New(
		testfixtures.Database(db),
		testfixtures.Dialect("sqlite3"),
		testfixtures.Paths(paths...),
	)

	if err != nil {
		return err
	}

	if err := fixtures.Load(); err != nil {
		return err
	}

	return nil
}

func TestContext(userID int) context.Context {
	return utils.AddUserIdToContext(context.Background(), userID)
}

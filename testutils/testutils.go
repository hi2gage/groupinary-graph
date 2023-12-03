// package testutils

package testutils

import (
	"context"
	"database/sql"
	"groupinary/ent"
	"groupinary/utils"
	"os"
	"testing"
	"time"

	entsql "entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/schema"

	"github.com/go-testfixtures/testfixtures/v3"
)

// LoadFixtures loads fixtures into the provided database.
func LoadFixtures(db *sql.DB, paths ...string) error {
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

// OpenTest opens a new test environment with an SQLite in-memory database,
// runs migrations, and loads fixtures.
func OpenTest() (*ent.Client, *sql.DB, error) {
	// Use SQLite in-memory database for testing.
	dbPath := "file:test:memory:?_fk=1"
	// dbPath := "database.db?_fk=1"
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, nil, err
	}
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxLifetime(time.Hour)

	// Create an ent.Driver from `db`.
	drv := entsql.OpenDB("sqlite3", db)

	// Run migrations.
	client := ent.NewClient(ent.Driver(drv))
	if err := runMigrations(client); err != nil {
		return nil, nil, err
	}

	return client, db, nil
}

func runMigrations(client *ent.Client) error {
	ctx := context.Background()
	return client.Schema.Create(ctx, schema.WithDropIndex(true), schema.WithGlobalUniqueID(true))
}

func TestContext(userID int) context.Context {
	return utils.AddUserIdToContext(context.Background(), userID)
}

// CleanupTestEnvironment performs cleanup for the test environment, including
// closing the client and removing the SQLite in-memory database file.
func CleanupTestEnvironment(t *testing.T, client *ent.Client) {
	if err := client.Close(); err != nil {
		t.Error(err)
	}
	if err := os.Remove("test:memory:"); err != nil {
		t.Errorf("Error removing database file: %v", err)
	}
}

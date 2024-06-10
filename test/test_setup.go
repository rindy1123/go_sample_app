package test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/api/internal/infra/models"
	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
	pg "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Add the new models to the tables slice
var tables = []interface{}{
	&models.Todo{},
}

func SetupTestDB(t *testing.T) *gorm.DB {
	t.Parallel()
	ctx := context.Background()
	dbname := "test"
	username := "postgres"
	password := "postgres"
	pgContainer, err := postgres.RunContainer(
		ctx,
		testcontainers.WithImage("postgres:16.2-alpine3.19"),
		postgres.WithDatabase(dbname),
		postgres.WithUsername(username),
		postgres.WithPassword(password),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).WithStartupTimeout(5*time.Second)),
	)
	assert.NoError(t, err)

	t.Cleanup(func() {
		if err := pgContainer.Terminate(ctx); err != nil {
			t.Fatalf("failed to terminate pgContainer: %s", err)
		}
	})
	host, err := pgContainer.Host(ctx)
	assert.NoError(t, err)

	port, err := pgContainer.MappedPort(ctx, "5432/tcp")
	assert.NoError(t, err)

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s",
		host,
		username,
		password,
		dbname,
		port.Port(),
	)
	db, err := gorm.Open(pg.Open(dsn), &gorm.Config{})
	assert.NoError(t, err)

	migrator := db.Migrator()
	migrator.AutoMigrate(tables...)
	return db
}

package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/golang-migrate/migrate/v4"
	pgMigrate "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/sirupsen/logrus"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	pgGorm "gorm.io/driver/postgres"
	"gorm.io/gen"
	"gorm.io/gorm"
)

func main() {
	os.Exit(generate())
}

func generate() int {
	ctx := context.Background()

	// Setup Testcontainer
	pgC, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: testcontainers.ContainerRequest{
			Image:        "postgres:16.3-alpine",
			ExposedPorts: []string{"5432/tcp"},
			Env: map[string]string{
				"POSTGRES_USER":     "user",
				"POSTGRES_PASSWORD": "password",
				"POSTGRES_DB":       "testdb",
			},
			WaitingFor: wait.ForListeningPort("5432/tcp"),
		},
		Started: true,
	})
	if err != nil {
		logrus.Errorf("Failed to start container: %v", err)
		return 1
	}
	defer pgC.Terminate(ctx)

	host, err := pgC.Host(ctx)
	if err != nil {
		logrus.Errorf("Failed to get container host: %v", err)
		return 1
	}

	port, err := pgC.MappedPort(ctx, "5432")
	if err != nil {
		logrus.Errorf("Failed to get container port: %v", err)
		return 1
	}

	dsn := fmt.Sprintf("host=%s port=%s user=user password=password dbname=testdb sslmode=disable TimeZone=UTC", host, port.Port())

	// Connect to the database
	db, err := gorm.Open(pgGorm.Open(dsn), &gorm.Config{})
	if err != nil {
		logrus.Errorf("Failed to connect to database: %v", err)
		return 1
	}

	// Run migrations
	sqlDB, err := db.DB()
	if err != nil {
		logrus.Errorf("Failed to get database connection: %v", err)
		return 1
	}
	driver, err := pgMigrate.WithInstance(sqlDB, &pgMigrate.Config{})
	if err != nil {
		logrus.Errorf("Failed to create migrate instance: %v", err)
		return 1
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://"+filepath.Join("migrations"),
		"postgres",
		driver,
	)
	if err != nil {
		logrus.Errorf("Failed to create migrate instance: %v", err)
		return 1
	}

	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		logrus.Errorf("Failed to run migrations: %v", err)
		return 1
	}

	// Initialize GORM Gen
	g := gen.NewGenerator(gen.Config{
		OutPath:       "internal/domain/query",
		ModelPkgPath:  "internal/domain/model",
		Mode:          gen.WithQueryInterface | gen.WithoutContext,
		FieldNullable: true,
	})
	g.UseDB(db)

	// Generate models & query interfaces
	for _, table := range g.GenerateAllTable() {
		g.ApplyBasic(table)
		g.ApplyInterface(func() {}, table)
	}

	g.Execute()
	return 0
}

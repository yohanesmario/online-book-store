package gormdb

import (
	"fmt"

	gorm_logrus "github.com/onrik/gorm-logrus"
	log "github.com/sirupsen/logrus"
	"github.com/yohanesmario/online-book-store/conf"
	"go.uber.org/fx"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var GormDBModule = fx.Module("gormdb", fx.Provide(
	NewGormDB,
	NewDBConnection,
	NewDBTransactionProvider,
))

func NewGormDB() (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s",
		conf.Configuration.DBHost,
		conf.Configuration.DBPort,
		conf.Configuration.DBUser,
		conf.Configuration.DBPassword,
		conf.Configuration.DBName,
	)
	db, err := gorm.Open(
		postgres.New(postgres.Config{
			DSN: dsn,
		}),
		&gorm.Config{
			Logger: gorm_logrus.New(),
		},
	)

	if err != nil {
		return nil, err
	}

	log.Infof("Connected to database: %s", conf.Configuration.DBName)

	return db, nil
}

package postgres

import (
	"context"
	"fmt"

	"github.com/scalarorg/xchains-api/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type PostgresClient struct {
	DbName string
	Db     *gorm.DB
	cfg    config.PostgresDBConfig
}

func New(ctx context.Context, cfg config.PostgresDBConfig) (*PostgresClient, error) {
	dsn := fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s sslmode=disable", cfg.Host, cfg.Port, cfg.DbName, cfg.User, cfg.Password)
	gormLogLevel := logger.Silent

	switch cfg.LogLevel {
	case "info":
		gormLogLevel = logger.Info
	}
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(gormLogLevel)})
	if err != nil {
		fmt.Println("Error while connecting to the database", err)
		return nil, err
	}

	return &PostgresClient{
		DbName: cfg.DbName,
		Db:     db,
		cfg:    cfg,
	}, err
}

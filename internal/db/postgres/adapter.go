package postgres

import (
	"fmt"

	"github.com/scalarorg/xchains-api/internal/config"
	"gorm.io/gorm"
)

var dbAdapter *DbAdapter

type DbAdapter struct {
	XchainsIndexerClient *gorm.DB
}

func NewDatabaseAdapter(config *config.Config) (*DbAdapter, error) {
	if dbAdapter == nil {
		db, err := NewPostgresDbClient(config.IndexerDb)
		if err != nil {
			return nil, fmt.Errorf("failed to create postgres client: %w", err)
		}
		dbAdapter = &DbAdapter{
			XchainsIndexerClient: db,
		}
	}
	return dbAdapter, nil
}

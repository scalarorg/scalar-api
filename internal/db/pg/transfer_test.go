package pg_test

import (
	"context"
	"os"
	"testing"

	"github.com/scalarorg/xchains-api/internal/config"
	"github.com/scalarorg/xchains-api/internal/db/pg"
	"github.com/scalarorg/xchains-api/internal/db/pg/models"
	"github.com/stretchr/testify/assert"
)

var pgClient *pg.PostgresClient

func TestMain(m *testing.M) {
	var err error
	pgClient, err = pg.New(context.Background(), config.PostgresDBConfig{
		Host:     "localhost",
		Port:     5432,
		DbName:   "relayer",
		User:     "postgres",
		Password: "postgres",
	})
	if err != nil {
		panic(err)
	}
	os.Exit(m.Run())
}

func TestTransfer(t *testing.T) {
	options := &models.Options{
		Offset: 0,
		Size:   10,
	}
	transfers, total, err := pgClient.TokenSearchTransfers(context.Background(), options)
	assert.Nil(t, err)
	t.Logf("total: %d", total)
	for _, transfer := range transfers {
		t.Logf("transfer: %+v", transfer)
	}
	assert.NotNil(t, transfers)
	assert.Equal(t, total, len(transfers))
}

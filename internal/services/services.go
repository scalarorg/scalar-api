package services

import (
	"context"

	"github.com/rs/zerolog/log"

	"github.com/scalarorg/xchains-api/internal/config"
	"github.com/scalarorg/xchains-api/internal/db/pg"
)

// Service layer contains the business logic and is used to interact with
// the database and other external clients (if any).
type Services struct {
	Pg  *pg.PostgresClient
	cfg *config.Config
}

func New(
	ctx context.Context,
	cfg *config.Config,
) (*Services, error) {
	pgdb, err := pg.New(ctx, cfg.RelayerDb)
	if err != nil {
		log.Ctx(ctx).Fatal().Err(err).Msg("error while creating Scalar Postgres client")
		return nil, err
	}

	return &Services{
		Pg:  pgdb,
		cfg: cfg,
	}, nil
}

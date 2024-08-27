package services

import (
	"context"
	"net/http"

	"github.com/rs/zerolog/log"

	"github.com/scalarorg/xchains-api/internal/config"
	"github.com/scalarorg/xchains-api/internal/db"
	"github.com/scalarorg/xchains-api/internal/db/gmp"
	"github.com/scalarorg/xchains-api/internal/db/postgres"
	"github.com/scalarorg/xchains-api/internal/types"
)

// Service layer contains the business logic and is used to interact with
// the database and other external clients (if any).
type Services struct {
	DbClient          db.DBClient
	GmpClient         gmp.GmpClient
	cfg               *config.Config
	params            *types.GlobalParams
	finalityProviders []types.FinalityProviderDetails
}

func New(
	ctx context.Context,
	cfg *config.Config,
	globalParams *types.GlobalParams,
	finalityProviders []types.FinalityProviderDetails,
) (*Services, error) {
	dbClient, err := db.New(ctx, cfg.MongoDb)
	if err != nil {
		log.Ctx(ctx).Fatal().Err(err).Msg("error while creating mongodb client")
		return nil, err
	}
	indexer, err := postgres.New(ctx, cfg.IndexerDb)
	if err != nil {
		log.Ctx(ctx).Fatal().Err(err).Msg("error while creating Indexer client")
		return nil, err
	}
	relayer, err := postgres.New(ctx, cfg.RelayerDb)
	if err != nil {
		log.Ctx(ctx).Fatal().Err(err).Msg("error while creating Relayer client")
		return nil, err
	}
	gmpClient := gmp.New(indexer, relayer)

	return &Services{
		DbClient:          dbClient,
		GmpClient:         *gmpClient,
		cfg:               cfg,
		params:            globalParams,
		finalityProviders: finalityProviders,
	}, nil
}

// DoHealthCheck checks the health of the services by ping the database.
func (s *Services) DoHealthCheck(ctx context.Context) error {
	return s.DbClient.Ping(ctx)
}

func (s *Services) SaveUnprocessableMessages(ctx context.Context, messageBody, receipt string) *types.Error {
	err := s.DbClient.SaveUnprocessableMessage(ctx, messageBody, receipt)
	if err != nil {
		log.Ctx(ctx).Error().Err(err).Msg("error while saving unprocessable message")
		return types.NewErrorWithMsg(http.StatusInternalServerError, types.InternalServiceError, "error while saving unprocessable message")
	}
	return nil
}

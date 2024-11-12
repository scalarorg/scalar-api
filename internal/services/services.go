package services

import (
	"context"
	"net/http"

	"github.com/rs/zerolog/log"

	"github.com/scalarorg/xchains-api/internal/config"
	"github.com/scalarorg/xchains-api/internal/db"
	"github.com/scalarorg/xchains-api/internal/db/gmp"
	"github.com/scalarorg/xchains-api/internal/db/postgres"
	"github.com/scalarorg/xchains-api/internal/db/vault"
	"github.com/scalarorg/xchains-api/internal/types"
)

// Service layer contains the business logic and is used to interact with
// the database and other external clients (if any).
type Services struct {
	DbClient  db.DBClient
	GmpClient gmp.GmpClient

	// TODO: fix me
	VaultClient       vault.VaultClient
	ScalarClient      postgres.ScalarClient
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
	scalarPostgresClient, err := postgres.New(ctx, cfg.ScalarDb)
	if err != nil {
		log.Ctx(ctx).Fatal().Err(err).Msg("error while creating Scalar Postgres client")
		return nil, err
	}
	gmpClient := gmp.New(scalarPostgresClient)
	vaultClient := vault.New(scalarPostgresClient)
	scalarClient := postgres.NewScalarClient(scalarPostgresClient)
	// Migrate the tables
	err = scalarClient.MigrateTables()
	if err != nil {
		log.Ctx(ctx).Fatal().Err(err).Msg("error while migrating Scalar tables")
		return nil, err
	}
	// Init dApps
	err = scalarClient.InitDApps(cfg.InitDApps)
	if err != nil {
		log.Ctx(ctx).Warn().Err(err).Msg("warning while initializing dApps")
	}

	return &Services{
		DbClient:          dbClient,
		GmpClient:         *gmpClient,
		VaultClient:       *vaultClient,
		ScalarClient:      *scalarClient,
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

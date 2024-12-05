package services

import (
	"context"

	"github.com/rs/zerolog/log"

	"github.com/scalarorg/xchains-api/internal/config"
	"github.com/scalarorg/xchains-api/internal/db/gmp"
	"github.com/scalarorg/xchains-api/internal/db/postgres"
	"github.com/scalarorg/xchains-api/internal/db/vault"
	"github.com/scalarorg/xchains-api/internal/types"
)

// Service layer contains the business logic and is used to interact with
// the database and other external clients (if any).
type Services struct {
	GmpClient         gmp.GmpClient
	VaultClient       vault.VaultClient
	ScalarClient      postgres.ScalarClient
	DbAdapter         *postgres.DbAdapter
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
	scalarPostgresClient, err := postgres.New(ctx, cfg.ScalarDb)
	if err != nil {
		log.Ctx(ctx).Fatal().Err(err).Msg("error while creating Scalar Postgres client")
		return nil, err
	}
	gmpClient := gmp.New(scalarPostgresClient)
	vaultClient := vault.New(scalarPostgresClient)
	scalarClient := postgres.NewScalarClient(scalarPostgresClient)

	dbAdapter, err := postgres.NewDatabaseAdapter(cfg)
	if err != nil {
		log.Ctx(ctx).Fatal().Err(err).Msg("error while creating database adapter")
		return nil, err
	}

	// Init dApps
	err = scalarClient.InitDApps(cfg.InitDApps)
	if err != nil {
		log.Ctx(ctx).Warn().Err(err).Msg("warning while failed to initialize dApps")
	}

	return &Services{
		GmpClient:         *gmpClient,
		VaultClient:       *vaultClient,
		ScalarClient:      *scalarClient,
		DbAdapter:         dbAdapter,
		cfg:               cfg,
		params:            globalParams,
		finalityProviders: finalityProviders,
	}, nil
}

package main

import (
	"context"
	"fmt"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
	"github.com/scalarorg/xchains-api/cmd/xchains-api/cli"
	"github.com/scalarorg/xchains-api/cmd/xchains-api/scripts"
	"github.com/scalarorg/xchains-api/internal/api"
	"github.com/scalarorg/xchains-api/internal/config"
	"github.com/scalarorg/xchains-api/internal/db/model"
	"github.com/scalarorg/xchains-api/internal/observability/healthcheck"
	"github.com/scalarorg/xchains-api/internal/observability/metrics"
	"github.com/scalarorg/xchains-api/internal/queue"
	"github.com/scalarorg/xchains-api/internal/services"
	"github.com/scalarorg/xchains-api/internal/types"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Debug().Msg("failed to load .env file")
	}
}

func main() {
	ctx := context.Background()

	// setup cli commands and flags
	if err := cli.Setup(); err != nil {
		log.Fatal().Err(err).Msg("error while setting up cli")
	}

	// load config
	cfgPath := cli.GetConfigPath()
	cfg, err := config.New(cfgPath)
	if err != nil {
		log.Fatal().Err(err).Msg(fmt.Sprintf("error while loading config file: %s", cfgPath))
	}

	paramsPath := cli.GetGlobalParamsPath()
	params, err := types.NewGlobalParams(paramsPath)
	if err != nil {
		log.Fatal().Err(err).Msg(fmt.Sprintf("error while loading global params file: %s", paramsPath))
	}

	finalityProvidersPath := cli.GetFinalityProvidersPath()
	finalityProviders, err := types.NewFinalityProviders(finalityProvidersPath)
	if err != nil {
		log.Fatal().Err(err).Msg(fmt.Sprintf("error while loading finality providers file: %s", finalityProvidersPath))
	}

	// initialize metrics with the metrics port from config
	metricsPort := cfg.Metrics.GetMetricsPort()
	metrics.Init(metricsPort)

	err = model.Setup(ctx, cfg)
	if err != nil {
		log.Fatal().Err(err).Msg("error while setting up staking db model")
	}
	services, err := services.New(ctx, cfg, params, finalityProviders)
	if err != nil {
		log.Fatal().Err(err).Msg("error while setting up staking services layer")
	}
	// Start the event queue processing
	queues := queue.New(&cfg.Queue, services)

	// Check if the replay flag is set
	if cli.GetReplayFlag() {
		log.Info().Msg("Replay flag is set. Starting replay of unprocessable messages.")
		err := scripts.ReplayUnprocessableMessages(ctx, cfg, queues, services.DbClient)
		if err != nil {
			log.Fatal().Err(err).Msg("error while replaying unprocessable messages")
		}
		return
	}

	queues.StartReceivingMessages()

	healthcheck.StartHealthCheckCron(ctx, queues, cfg.Server.HealthCheckInterval)

	apiServer, err := api.New(ctx, cfg, services)
	if err != nil {
		log.Fatal().Err(err).Msg("error while setting up staking api service")
	}
	if err = apiServer.Start(); err != nil {
		log.Fatal().Err(err).Msg("error while starting staking api service")
	}
}
